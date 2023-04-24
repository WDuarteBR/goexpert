package repository

import (
	"context"
	"database/sql"

	"github.com/wduartebr/goexpert/uow/internal/db"
	"github.com/wduartebr/goexpert/uow/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	Db      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dtb *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		Db:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx,
		db.CreateCategoryParams{
			Name: category.Name,
		})
}
