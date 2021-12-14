package postgres

import (
	"context"
	"go-grpc-blog/blog/storage"
)

const insertBlog =`
	INSERT INTO blogs(
		title,
		description
	) VALUES(
		:title,
		:description
	)RETURNING id;
`

func (s *Storage) Create(ctx context.Context, t storage.Blog) (int64, error) {
	stmt,err :=s.db.PrepareNamed(insertBlog)
	if err!=nil{
		return 0,err
	}
	var id int64
	if err :=stmt.Get(&id,t);err != nil{
		return 0,err
	}
	return id,nil
}
