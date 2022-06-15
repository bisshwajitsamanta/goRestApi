package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Could not fetch the comment")
	ErrNotImplemented  = errors.New("Not Yet Implemented")
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

//NewService - Structure Function ways to initiate constructor in Golang
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

// UpdateComment - It updates an existing comment
func (s *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

// DeleteComment - It deletes an existing comment
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

// CreateComment - It creates an new comment
func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
