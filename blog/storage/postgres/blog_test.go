package postgres

import (
	"context"
	"go-grpc-blog/blog/storage"
	"log"

	// "log"
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
				Title: "This is title ",
				Description: "This is description ",
			},
			want: 1,
		},

		{
			name: "CREATE_d_SUCCESS",
			in: storage.Blog{
				Title: "This is title ",
				Description: "This is description ",
			},
			wantErr:true,
		},
	}
	for _, tt := range tests {
		tt:=tt
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
