package usecase

import (
	"context"

	"github.com/wduartebr/goexpert/uow/internal/entity"
	"github.com/wduartebr/goexpert/uow/internal/repository"
	"github.com/wduartebr/goexpert/uow/pkg/uow"
)

type InputUseCaseUow struct {
	CategoryName     string
	CourseName       string
	CourseCategoryId int
}

type AddCourseUsecaseUow struct {
	Uow uow.UowInterface
}

func NewAddCourseUsecaseUow(uow uow.UowInterface) *AddCourseUsecaseUow {
	return &AddCourseUsecaseUow{
		Uow: uow,
	}
}

func (a *AddCourseUsecaseUow) Execute(ctx context.Context, input InputUseCaseUow) error {
	return a.Uow.Do(ctx, func(uow *uow.Uow) error {
		category := entity.Category{
			Name: input.CategoryName,
		}

		repoCategory := a.getCategoryRepository(ctx)

		err := repoCategory.Insert(ctx, category)
		if err != nil {
			return err
		}

		course := entity.Course{
			Name:        input.CourseName,
			Category_id: input.CourseCategoryId,
		}

		repoCourse := a.getCourseRepository(ctx)
		err = repoCourse.Insert(ctx, course)
		if err != nil {
			return err
		}
		return nil
	})
}

func (a *AddCourseUsecaseUow) getCategoryRepository(ctx context.Context) repository.CategoryRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "categoryRepo")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CategoryRepositoryInterface) // <= casting para CategoryRepository
}

func (a *AddCourseUsecaseUow) getCourseRepository(ctx context.Context) repository.CourseRepositoryInterface {
	repo, err := a.Uow.GetRepository(ctx, "courseRepo")
	if err != nil {
		panic(err)
	}

	return repo.(repository.CourseRepositoryInterface)
}
