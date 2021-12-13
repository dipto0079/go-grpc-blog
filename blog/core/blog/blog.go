package blog

import (
	"context"
	"go-grpc-blog/blog/storage"
)

type blogStore interface {
	Create(context.Context, storage.Blog) (int64, error)
}
type CoreSvc struct {
	store blogStore
}

func NewCoreSve(b blogStore) *CoreSvc {
	return &CoreSvc{
		store: b,
	}
}

func (cs CoreSvc) Create(ctx context.Context, t storage.Blog) (int64, error) {
	return cs.store.Create(ctx, t)
	// return 0, nil
}
