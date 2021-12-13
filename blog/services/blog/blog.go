package blog

import (
	"context"

	"go-grpc-blog/blog/storage"
	tpb "go-grpc-blog/gunk/v1/blog"
)

type blogCoreStore interface {
	Create(context.Context, storage.Blog) (int64, error)
}

type Svc struct{
	tpb.UnimplementedBlogServiceServer
	core blogCoreStore
}

func NewTodoServer(c blogCoreStore) *Svc {
	return &Svc{
		core: c,
	}
}


