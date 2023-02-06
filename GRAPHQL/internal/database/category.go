package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	Id          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()

	_, err := c.db.Exec("insert into categories(id, name, description) values (1$, 2$, 3$)",
		id, name, description)

	if err != nil {
		return Category{}, err
	}

	return Category{Id: id, Name: name, Description: description}, nil

}
