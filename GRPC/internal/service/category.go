package service

import (
	"context"

	"github.com/wduartebr/goexpert/grpc/internal/database"
	"github.com/wduartebr/goexpert/grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.Id,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil

}

func (c *CategoryService) ListCategory(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var respCategories []*pb.Category
	for _, cat := range categories {
		category := &pb.Category{
			Id:          cat.Id,
			Name:        cat.Name,
			Description: cat.Description,
		}
		respCategories = append(respCategories, category)
	}

	return &pb.CategoryList{Categories: respCategories}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	cat, err := c.CategoryDB.FindById(in.Id)
	if err != nil {
		return nil, err
	}

	category := &pb.Category{
		Id:          cat.Id,
		Name:        cat.Name,
		Description: cat.Description,
	}

	return category, nil
}
