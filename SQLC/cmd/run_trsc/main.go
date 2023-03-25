package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/wduartebr/goexpert/sqlc/internal/db"
)

type CousrseDB struct {
	dbCon *sql.DB
	*db.Queries
}

func NewCourseDB(dbCon *sql.DB) *CousrseDB {
	return &CousrseDB{
		dbCon:   dbCon,
		Queries: db.New(dbCon),
	}
}

func (c *CousrseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbCon.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			return fmt.Errorf("error on rollback: %v, original error %w", errTx, err)
		}
		return err
	}
	return tx.Commit()
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CousrseDB) CreateCourseAndCategory(cxt context.Context, argsCourse CourseParams, argsCategory CategoryParams) error {
	err := c.callTx(cxt, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(cxt, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(cxt, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			CategoryID:  argsCategory.ID,
		})

		if err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		return err
	}

	return nil
}
func main() {
	ctx := context.Background()
	dbcnn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbcnn.Close()

	// queries := db.New(dbcnn)

	courseArgs := CourseParams{
		ID:          uuid.New().String(),
		Name:        "Go the best",
		Description: sql.NullString{String: "Go language", Valid: true},
	}

	categoryArgs := CategoryParams{
		ID:          uuid.New().String(),
		Name:        "Program language",
		Description: sql.NullString{String: "Most performance language", Valid: true},
	}

	courseDb := NewCourseDB(dbcnn)
	err = courseDb.CreateCourseAndCategory(ctx, courseArgs, categoryArgs)
	if err != nil {
		panic(err)
	}

}
