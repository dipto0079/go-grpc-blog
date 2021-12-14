package postgres

import (
	"context"
	"go-grpc-blog/blog/storage"
	"testing"


)

func TestCreate(t *testing.T) {
	s:=newTestStorage(t)
	tests := []struct {
		name string
		in storage.Blog
		want    int64
		wantErr bool
	}{
		{
			name: "CREATE_BLOG_SUCCESS",
			in: storage.Blog{
				Title: "This is title 1",
				Description: "This is description 1",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		tt:=tt
		t.Run(tt.name, func(t *testing.T) {
		
			got, err := s.Create(context.TODO(), tt.in)
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
