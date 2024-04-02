package comment

import (
	"context"
	"log"
	//"time"

	blogDomain "github.com/beeblogit/lib_go_domain/domain/blog"
)

type (
	Filters struct {
		ID     []string
		UserID []string
		PostID []string
	}

	Service interface {
		Create(ctx context.Context, userID, postID, name, comment string) (*blogDomain.Comment, error)
		Get(ctx context.Context, id string) (*blogDomain.Comment, error)
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]blogDomain.Comment, error)
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

func (s service) Create(ctx context.Context, userID, postID, name, comment string) (*blogDomain.Comment, error) {

	comObj := &blogDomain.Comment{
		UserID:  userID,
		PostID:  postID,
		Name:    name,
		Comment: comment,
	}

	if err := s.repo.Create(ctx, comObj); err != nil {
		return nil, err
	}

	return comObj, nil
}

func (s service) Get(ctx context.Context, id string) (*blogDomain.Comment, error) {
	return nil, nil
}

func (s service) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]blogDomain.Comment, error) {
	comments, err := s.repo.GetAll(ctx, filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (s service) Count(ctx context.Context, filters Filters) (int, error) {
	return s.repo.Count(ctx, filters)
}
