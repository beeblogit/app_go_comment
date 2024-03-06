package comment

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/beeblogit/lib_go_domain/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(ctx context.Context, comment *domain.Comment) error
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.Comment, error)
		Get(ctx context.Context, id string) (*domain.Comment, error)
		Update(ctx context.Context, ID, userID string, name *string, comment *string) error
		Count(ctx context.Context, filters Filters) (int, error)
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func (r *repo) Create(ctx context.Context, comment *domain.Comment) error {

	if err := r.db.WithContext(ctx).Create(course).Error; err != nil {
		r.log.Println(err)
		return err
	}

	return nil
}

func (r *repo) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]domain.Comment, error){
	return nil, nil
}

func (r *repo) Get(ctx context.Context, id string) (*domain.Comment, error){
	return nil, nil
}

func (r *repo) Update(ctx context.Context, ID, userID string, name *string, comment *string) error{
	return nil
}

func (r *repo) Count(ctx context.Context, filters Filters) (int, error){
	return 0, nil
}

