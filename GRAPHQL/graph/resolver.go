package graph

import (
	"github.com/wduartebr/goexpert/graphql/internal/database"
)

type Resolver struct {
	CategoryDB *database.Category
	CourseDB   *database.Courses
}
