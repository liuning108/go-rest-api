package comment

import (
	"context"
	"errors"
	"fmt"
)

type ContextKey string

var (
	ErrFetchingComment = errors.New("Fail to fetch comment by id ")
	ErrNotImplemented  = errors.New("Not Implemented")
)

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Server struct {
	Store Store
}

func NewServer(store Store) *Server {
	return &Server{
		Store: store,
	}
}

func (s *Server) GetComment(ctx context.Context, id string) (Comment, error) {

	fmt.Println("retrieving  a commiet")
	key := ContextKey("request_id")
	ctx = context.WithValue(ctx, key, "unique-string")
	fmt.Println(ctx.Value("request_id"))
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Server) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented

}

func (s *Server) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Server) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
