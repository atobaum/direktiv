// nolint
package function

import (
	"fmt"
	"strconv"

	"github.com/mitchellh/hashstructure/v2"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
)

type Config struct {
	Namespace string
	Name      string

	ServicePath  string
	WorkflowPath string

	Image string
	CMD   string
	Size  string
	Scale int

	Error error
}

func (c *Config) id() string {
	str := fmt.Sprintf("%s-%s-%s-%s", c.Namespace, c.Name, c.ServicePath, c.WorkflowPath)
	v, err := hashstructure.Hash(str, hashstructure.FormatV2, nil)
	if err != nil {
		panic("unexpected hashstructure.Hash error: " + err.Error())
	}

	return fmt.Sprintf("obj-%d-obj", v)
}

func (c *Config) hash() string {
	str := fmt.Sprintf("%s-%s-%s-%d", c.Image, c.CMD, c.Size, c.Scale)
	v, err := hashstructure.Hash(str, hashstructure.FormatV2, nil)
	if err != nil {
		panic("unexpected hashstructure.Hash error: " + err.Error())
	}

	return strconv.Itoa(int(v))
}

type Status interface {
	data() any
	id() string
	hash() string
}

type ConfigStatus struct {
	Config *Config
	Status Status
}

type K8sFunctionStatus struct {
	*servingv1.Service
}

func (r *K8sFunctionStatus) data() any {
	return r.Service
}

func (r *K8sFunctionStatus) id() string {
	return r.Name
}

func (r *K8sFunctionStatus) hash() string {
	return r.Annotations["direktiv.io/input_hash"]
}

var _ Status = &K8sFunctionStatus{}