package comment

import (
	"context"
	//"fmt"
	"log"
	//"strings"
	//"time"

	blogDomain "github.com/beeblogit/lib_go_domain/domain/blog"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(ctx context.Context, comment *blogDomain.Comment) error
		GetAll(ctx context.Context, filters Filters, offset, limit int) ([]blogDomain.Comment, error)
		Get(ctx context.Context, id string) (*blogDomain.Comment, error)
		Update(ctx context.Context, ID, userID string, name *string, comment *string) error
		Count(ctx context.Context, filters Filters) (int, error)
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(ctx context.Context, comment *blogDomain.Comment) error {

	if err := r.db.WithContext(ctx).Create(comment).Error; err != nil {
		r.log.Println(err)
		return err
	}

	return nil
}

func (r *repo) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]blogDomain.Comment, error){
	return nil, nil
}

func (r *repo) Get(ctx context.Context, id string) (*blogDomain.Comment, error){
	return nil, nil
}

func (r *repo) Update(ctx context.Context, ID, userID string, name *string, comment *string) error{
	return nil
}

func (r *repo) Count(ctx context.Context, filters Filters) (int, error){
	return 0, nil
}

