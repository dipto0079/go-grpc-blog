package postgres

import (
	"context"
	"go-grpc-blog/blog/storage"
	"log"

	// "log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCreateBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Blog
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_BLOG_SUCCESS",
			in: storage.Blog{
				Title:       "This is title ",
				Description: "This is description ",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Create(context.Background(), tt.in)
			log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Storage.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      int64
		want    *storage.Blog
		wantErr bool
	}{
		{
			name: "GET_BLOG_SUCCESS",
			in: 1,
			want: &storage.Blog{
				ID:          1,
				Title:       "This is title",
				Description: "This is Description",
				IsCompleted: false,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			got, err := s.Create(context.Background(), *tt.want)
			log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == 0 {
				t.Errorf("Storage.Get(), not ok ")
			}
			blog,err:= s.Get(context.Background(),got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if cmp.Equal( blog,tt.want) {
				t.Errorf("Storage.Create() = %v, want %v", blog, tt.want)
			}
		})
	}
}
