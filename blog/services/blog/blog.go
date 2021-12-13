package blog

import (
	"context"

	"grpc-todos/todo/storage"
	tpb "grpc-todos/gunk/v1/todo"
)

type todoCoreStore interface {
	Create(context.Context, storage.Todo) (int64, error)
}

type Svc struct{
	tpb.UnimplementedTodoServiceServer
	core todoCoreStore
}

func NewTodoServer(c todoCoreStore) *Svc {
	return &Svc{
		core: c,
	}
}


