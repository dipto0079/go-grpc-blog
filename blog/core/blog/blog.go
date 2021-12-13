package blog

import (
	"context"
	"go-grpc-blog/blog/storage"
)
type CoreSvc struct{

}


func (cs CoreSvc) Create(ctx context.Context, t storage.Blog) (int64, error) {
	return cs.store.Create(ctx, t)
}