package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Courses struct {
	db          *sql.DB
	Id          string
	Name        string
	Description string
	Category_Id string
}

func NewCourse(db *sql.DB) *Courses {
	return &Courses{db: db}
}

func (c *Courses) Create(name, description, category_id string) (Courses, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("insert into Courses(id, name, description, category_id) values($1, $2, $3, $4)",
		id, name, description, category_id)
	if err != nil {
		return Courses{}, err
	}

	return Courses{
		Id:          id,
		Name:        name,
		Description: description,
		Category_Id: category_id,
	}, nil

}

func (c *Courses) FindAll() ([]Courses, error) {
	rows, err := c.db.Query("select id, name, description, category_id from Courses")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var courses []Courses

	for rows.Next() {
		var id, name, description, category string
		err = rows.Scan(&id, &name, &description, &category)
		if err != nil {
			return nil, err
		}

		courses = append(courses, Courses{
			Id:          id,
			Name:        name,
			Description: description,
			Category_Id: category,
		})
	}

	return courses, nil
}

func (c *Courses) FindByCategoryId(categoryId string) ([]Courses, error) {
	rows, err := c.db.Query("select id, name, description, category_id from Courses where category_id = $1", categoryId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var courses []Courses
	for rows.Next() {
		var id, name, description, category string
		err = rows.Scan(&id, &name, &description, &category)
		if err != nil {
			return nil, err
		}
		courses = append(courses, Courses{
			Id:          id,
			Name:        name,
			Description: description,
			Category_Id: category,
		})
	}
	return courses, nil

}
