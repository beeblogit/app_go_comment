package comment

import (
	"context"
	"log"
	"time"

	"github.com/beeblogit/lib_go_domain/domain"
)

type (
	Filters struct {
		ID []string
		UserID []string
		PostID []string
	}

	Service interface {
		Create(ctx context.Context, userID, postID, name, comment string) (*domain.Comment, error)
		Get(ctx context.Context, id string) (*domain.Comment, error)
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.Comment, error)
		//Delete(ctx context.Context, id string) error
		//Update(ctx context.Context, id string, name, startDate, endDate *string) error
		Count(ctx context.Context, filters Filters) (int, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, userID, postID, name, comment string) (*domain.Comment, error) {


	comment := &domain.Comment{
		UserID: userID,
		PostID: postID,
		Name:      name,
		Comment: comment,
	}

	if err := s.repo.Create(ctx, comment); err != nil {
		return nil, err
	}

	return comment, nil
}


func (s service) Get(ctx context.Context, id string) (*domain.Comment, error) {
	return nil, nil
}

func (s service) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.Comment, error) {
	return nil, nil
}

func (s service) Count(ctx context.Context, filters Filters) (int, error) {
	return 0, nil
}