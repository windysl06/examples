// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"kratos-cqrs/app/logger/service/internal/data/ent/predicate"
	"kratos-cqrs/app/logger/service/internal/data/ent/sensordata"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SensorDataDelete is the builder for deleting a SensorData entity.
type SensorDataDelete struct {
	config
	hooks    []Hook
	mutation *SensorDataMutation
}

// Where appends a list predicates to the SensorDataDelete builder.
func (sdd *SensorDataDelete) Where(ps ...predicate.SensorData) *SensorDataDelete {
	sdd.mutation.Where(ps...)
	return sdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdd *SensorDataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sdd.hooks) == 0 {
		affected, err = sdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SensorDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sdd.mutation = mutation
			affected, err = sdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sdd.hooks) - 1; i >= 0; i-- {
			if sdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sdd *SensorDataDelete) ExecX(ctx context.Context) int {
	n, err := sdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdd *SensorDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: sensordata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: sensordata.FieldID,
			},
		},
	}
	if ps := sdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sdd.driver, _spec)
}

// SensorDataDeleteOne is the builder for deleting a single SensorData entity.
type SensorDataDeleteOne struct {
	sdd *SensorDataDelete
}

// Exec executes the deletion query.
func (sddo *SensorDataDeleteOne) Exec(ctx context.Context) error {
	n, err := sddo.sdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sensordata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sddo *SensorDataDeleteOne) ExecX(ctx context.Context) {
	sddo.sdd.ExecX(ctx)
}
