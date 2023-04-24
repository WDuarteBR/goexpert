package usecase

import (
	"context"

	"github.com/wduartebr/goexpert/uow/internal/entity"
	"github.com/wduartebr/goexpert/uow/internal/repository"
)

type InputUseCase struct {
	CategoryName     string
	CourseName       string
	CourseCategoryId int
}

type AddCourseUsecase struct {
	CourseRepository   repository.CourseRepositoryInterface
	CategoryRepository repository.CategoryRepositoryInterface
}

func NewAddCourseUsecase(courseRepo repository.CourseRepositoryInterface, categoryRepo repository.CategoryRepositoryInterface) *AddCourseUsecase {
	return &AddCourseUsecase{
		CourseRepository:   courseRepo,
		CategoryRepository: categoryRepo,
	}
}

func (a *AddCourseUsecase) Execute(ctx context.Context, input InputUseCase) error {
	category := entity.Category{
		Name: input.CategoryName,
	}

	err := a.CategoryRepository.Insert(ctx, category)
	if err != nil {
		return err
	}

	course := entity.Course{
		Name:        input.CourseName,
		Category_id: input.CourseCategoryId,
	}

	err = a.CourseRepository.Insert(ctx, course)
	if err != nil {
		return err
	}

	return nil
}
