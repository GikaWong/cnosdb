package tsm1

import (
	"context"

	"github.com/cnosdatabase/db/tsdb"
)

func (e *Engine) CreateCursorIterator(ctx context.Context) (tsdb.CursorIterator, error) {
	return &arrayCursorIterator{e: e}, nil
}
