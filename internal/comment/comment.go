package comment

import (
	"context"
	"fmt"
)

// Comment - A representation of the comment structure of our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - Everybody like add comment update comment will interact with service struct
type Service struct {
	Store Store
}

//NewService - Structure Function
func NewService(store Store) *Service {
	return &Service{
		store,
	}
}

//GetComment - Retrieves a comment from Database
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a Comment ")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}
