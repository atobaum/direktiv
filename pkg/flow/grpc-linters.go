package flow

import (
	"context"
	"fmt"

	"github.com/direktiv/direktiv/pkg/flow/bytedata"
	"github.com/direktiv/direktiv/pkg/flow/grpc"
	"github.com/direktiv/direktiv/pkg/refactor/core"
)

func (flow *flow) NamespaceLint(ctx context.Context, req *grpc.NamespaceLintRequest) (*grpc.NamespaceLintResponse, error) {
	flow.sugar.Debugf("Handling gRPC request: %s", this())

	tx, err := flow.beginSqlTx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	ns, err := tx.DataStore().Namespaces().GetByName(ctx, req.GetName())
	if err != nil {
		return nil, err
	}

	rootDir, err := tx.FileStore().ForRootNamespaceID(ns.ID).GetFile(ctx, "/")
	if err != nil {
		return nil, err
	}

	annotations, _ := tx.DataStore().FileAnnotations().Get(ctx, rootDir.ID)

	secretIssues, err := flow.lintSecrets(ctx, tx, ns)
	if err != nil {
		return nil, err
	}

	var resp grpc.NamespaceLintResponse

	resp.Namespace = bytedata.ConvertNamespaceToGrpc(ns, annotations)
	resp.Issues = make([]*grpc.LinterIssue, 0)
	resp.Issues = append(resp.Issues, secretIssues...)

	return &resp, nil
}

func (flow *flow) lintSecrets(ctx context.Context, tx *sqlTx, ns *core.Namespace) ([]*grpc.LinterIssue, error) {
	secrets, err := tx.DataStore().Secrets().GetAll(ctx, ns.ID)
	if err != nil {
		return nil, err
	}

	issues := make([]*grpc.LinterIssue, 0)

	for _, secret := range secrets {
		if secret.Data == nil {
			issues = append(issues, &grpc.LinterIssue{
				Level: "critical",
				Type:  "secret",
				Id:    secret.Name,
				Issue: fmt.Sprintf(`secret '%s' has not been initialized`, secret.Name),
			})
		}
	}

	return issues, nil
}
