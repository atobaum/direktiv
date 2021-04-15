// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vorteil/direktiv/pkg/secrets/ent/namespacesecret"
)

// NamespaceSecretCreate is the builder for creating a NamespaceSecret entity.
type NamespaceSecretCreate struct {
	config
	mutation *NamespaceSecretMutation
	hooks    []Hook
}

// SetNs sets the "ns" field.
func (nsc *NamespaceSecretCreate) SetNs(s string) *NamespaceSecretCreate {
	nsc.mutation.SetNs(s)
	return nsc
}

// SetName sets the "name" field.
func (nsc *NamespaceSecretCreate) SetName(s string) *NamespaceSecretCreate {
	nsc.mutation.SetName(s)
	return nsc
}

// SetSecret sets the "secret" field.
func (nsc *NamespaceSecretCreate) SetSecret(b []byte) *NamespaceSecretCreate {
	nsc.mutation.SetSecret(b)
	return nsc
}

// Mutation returns the NamespaceSecretMutation object of the builder.
func (nsc *NamespaceSecretCreate) Mutation() *NamespaceSecretMutation {
	return nsc.mutation
}

// Save creates the NamespaceSecret in the database.
func (nsc *NamespaceSecretCreate) Save(ctx context.Context) (*NamespaceSecret, error) {
	var (
		err  error
		node *NamespaceSecret
	)
	if len(nsc.hooks) == 0 {
		if err = nsc.check(); err != nil {
			return nil, err
		}
		node, err = nsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NamespaceSecretMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nsc.check(); err != nil {
				return nil, err
			}
			nsc.mutation = mutation
			node, err = nsc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nsc.hooks) - 1; i >= 0; i-- {
			mut = nsc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nsc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nsc *NamespaceSecretCreate) SaveX(ctx context.Context) *NamespaceSecret {
	v, err := nsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (nsc *NamespaceSecretCreate) check() error {
	if _, ok := nsc.mutation.Ns(); !ok {
		return &ValidationError{Name: "ns", err: errors.New("ent: missing required field \"ns\"")}
	}
	if _, ok := nsc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := nsc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New("ent: missing required field \"secret\"")}
	}
	return nil
}

func (nsc *NamespaceSecretCreate) sqlSave(ctx context.Context) (*NamespaceSecret, error) {
	_node, _spec := nsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nsc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nsc *NamespaceSecretCreate) createSpec() (*NamespaceSecret, *sqlgraph.CreateSpec) {
	var (
		_node = &NamespaceSecret{config: nsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: namespacesecret.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: namespacesecret.FieldID,
			},
		}
	)
	if value, ok := nsc.mutation.Ns(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespacesecret.FieldNs,
		})
		_node.Ns = value
	}
	if value, ok := nsc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: namespacesecret.FieldName,
		})
		_node.Name = value
	}
	if value, ok := nsc.mutation.Secret(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: namespacesecret.FieldSecret,
		})
		_node.Secret = value
	}
	return _node, _spec
}

// NamespaceSecretCreateBulk is the builder for creating many NamespaceSecret entities in bulk.
type NamespaceSecretCreateBulk struct {
	config
	builders []*NamespaceSecretCreate
}

// Save creates the NamespaceSecret entities in the database.
func (nscb *NamespaceSecretCreateBulk) Save(ctx context.Context) ([]*NamespaceSecret, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nscb.builders))
	nodes := make([]*NamespaceSecret, len(nscb.builders))
	mutators := make([]Mutator, len(nscb.builders))
	for i := range nscb.builders {
		func(i int, root context.Context) {
			builder := nscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NamespaceSecretMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, nscb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nscb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, nscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nscb *NamespaceSecretCreateBulk) SaveX(ctx context.Context) []*NamespaceSecret {
	v, err := nscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}