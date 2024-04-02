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

func (r *repo) GetAll(ctx context.Context, filters Filters, offset, limit int) ([]blogDomain.Comment, error) {
	var c []blogDomain.Comment

	tx := r.db.WithContext(ctx).Model(&c)
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	result := tx.Order("created_at desc").Find(&c)

	if result.Error != nil {
		r.log.Println(result.Error)
		return nil, result.Error
	}
	return c, nil
}

func (r *repo) Get(ctx context.Context, id string) (*blogDomain.Comment, error) {
	return nil, nil
}

func (r *repo) Update(ctx context.Context, ID, userID string, name *string, comment *string) error {
	return nil
}

func (r *repo) Count(ctx context.Context, filters Filters) (int, error) {
	var count int64
	tx := r.db.WithContext(ctx).Model(blogDomain.Comment{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		r.log.Println(err)
		return 0, err
	}

	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.ID != nil {
		tx = tx.Where("id in (?)", filters.ID)
	}
	if filters.UserID != nil {
		tx = tx.Where("user_id in (?)", filters.UserID)
	}
	if filters.PostID != nil {
		tx = tx.Where("post_id in (?)", filters.PostID)
	}

	return tx
}
