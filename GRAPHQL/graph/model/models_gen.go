// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  string `json:"category_id"`
}
