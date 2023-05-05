package presenter

import (
	"todo-clean/entity"
)

type Todo struct {
	ID     entity.ID `json:"id"`
	Title  string    `json:"title"`
	Status bool      `json:"status"`
}
