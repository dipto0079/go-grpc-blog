package todo

import (
	"context"
	tpb "grpc-todos/gunk/v1/todo"
	"grpc-todos/todo/storage"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Svc) Create(ctx context.Context, req *tpb.CreateTodoRequest) (*tpb.CreateTodoResponse, error){
	log.Printf("Request Todo: %#v\n", req.GetTodo())

	log.Fatal("asdfsdf")
	// Need to Validate request 
	todo := storage.Todo{
		Title: req.GetTodo().Title,
		Description: req.GetTodo().Description,
	}
	id, err := s.core.Create(context.Background(), todo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create todo: %s", err)
	}
	return &tpb.CreateTodoResponse{
		ID: id,
	}, nil
}
