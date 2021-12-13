package blog

import (
	"context"
	tpb "go-grpc-blog/gunk/v1/blog"
	"go-grpc-blog/blog/storage"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (s *Svc) Create(ctx context.Context, req *tpb.CreateBlogRequest) (*tpb.CreateBlogResponse, error){
	log.Printf("Request Blog: %#v\n", req.GetBlog())
	// Need to Validate request 
	blog := storage.Blog{
		Title: req.GetBlog().Title,
		Description: req.GetBlog().Description,
	}
	id, err := s.core.Create(context.Background(), blog)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create todo: %s", err)
	}
	return &tpb.CreateBlogResponse{
		ID: id,
	}, nil
}
