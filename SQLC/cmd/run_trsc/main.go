package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

func main() {
	ctx := context.Background()
	dbcnn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}

	defer dbcnn.Close()

	queries := db.New(dbcnn)

}
