package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Courses struct {
	Db          *sql.DB
	Id          string
	Name        string
	Description string
	Category_Id string
}

func NewCourse(db *sql.DB) *Courses {
	return &Courses{Db: db}
}

func (c *Courses) Create(name string, description string) (Courses, error) {
	id := uuid.New().String()
	_, err := c.Db.Exec("insert into Courses(id, name, description) values($1, $2, $3)", id, name, description)
	if err != nil {
		return Courses{}, err
	}

	return Courses{
		Id:          id,
		Name:        name,
		Description: description,
	}, nil

}

func (c *Courses) FindAll() ([]Courses, error) {
	rows, err := c.Db.Query("select id, name description from Courses")
	if err != nil {
		return nil, err
	}

	var courses []Courses

	for rows.Next() {
		var id, name, description string
		err = rows.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}

		courses = append(courses, Courses{
			Id:          id,
			Name:        name,
			Description: description,
		})
	}

	return courses, nil
}
