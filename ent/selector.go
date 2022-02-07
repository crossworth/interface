package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
)

type Selectorable interface {
	sqlQuery(ctx context.Context) *sql.Selector
}

func Selector(ctx context.Context, what Selectorable) *sql.Selector {
	return what.sqlQuery(ctx)
}
