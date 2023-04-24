package repository

import (
	"context"
	"database/sql"

	"github.com/wduartebr/goexpert/uow/internal/db"
	"github.com/wduartebr/goexpert/uow/internal/entity"
)

type CourseRepositoryInterface interface {
	Insert(ctx context.Context, course entity.Course) error
}

type CourseRepository struct {
	Db      *sql.DB
	Queries *db.Queries
}

func NewCourseRepository(dtb *sql.DB) *CourseRepository {
	return &CourseRepository{
		Db:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CourseRepository) Insert(ctx context.Context, course entity.Course) error {
	return r.Queries.CreateCourses(ctx, db.CreateCoursesParams{
		Name:       course.Name,
		CategoryID: int32(course.Category_id),
	})
}
