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

func (s *Storage) Get(ctx context.Context, id int64) (*storage.Blog, error) {
	var b storage.Blog

	if err :=s.db.Get(&b,"SELECT * FROM blogs WHERE id=$1",id);err != nil{
		return nil,err
	}
	return &b,nil
}


const updateBlog =`
	UPDATE blogs 
	SET
		title =:title,
		description =:description
	WHERE
	id =:id
	RETURNING *;
`

func (s *Storage) Update(ctx context.Context,t storage.Blog) (*storage.Blog, error) {
	stmt,err :=s.db.PrepareNamed(updateBlog)
	if err!=nil{
		return nil,err
	}
	if err :=stmt.Get(&t,t);err != nil{
		return nil,err
	}
	return &t,nil
}


func (s *Storage) Delete(ctx context.Context,id int64) error {
	var b storage.Blog
	if err := s.db.Get(&b,"DELETE FROM blogs WHERE id=$1 RETURNING * ",id); err != nil {
		return err;
	}
	return nil
}
