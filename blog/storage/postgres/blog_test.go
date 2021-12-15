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
				Title:       "This is title",
				Description: "This is description",
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
				Description: "This is description",
				IsCompleted: false,
			},
		},
		{
			name: "FAILED_BLOG_DOES_NOT_EXIST",
			in: 100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			blog, err := s.Get(context.Background(), tt.in)
			log.Printf("%#v", blog)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !cmp.Equal(blog, tt.want) {
				t.Errorf("Diff: got -, want += %v", cmp.Diff(blog, tt.want))
			}
		})
	}
}


func TestUpdateBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      storage.Blog
		want    *storage.Blog
		wantErr bool
	}{
		{
			name: "UPDATE_BLOG_SUCCESS",
			in: storage.Blog{
				ID: 1,
				Title:       "This is title Update",
				Description: "This is description Update",
			},
			want: &storage.Blog{
				ID: 1,
				Title:       "This is title Update",
				Description: "This is description Update",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Update(context.Background(), tt.in)
			//log.Printf("%#v", got)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(got, tt.want) {
				t.Errorf("Diff: got -, want += %v", cmp.Diff(got, tt.want))
			}
		})
	}
}


func TestDeleteBlog(t *testing.T) {
	s := newTestStorage(t)
	tests := []struct {
		name    string
		in      int64
		wantErr bool
	}{
		{
			name: "DELETE_BLOG_SUCCESS",
			in: 1,
		},
		{
			name: "FAILED_TO_DELETE_BLOG_DOES_NOT_EXIST",
			in: 100,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			err := s.Delete(context.Background(), tt.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

