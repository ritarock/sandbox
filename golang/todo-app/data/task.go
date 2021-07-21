package data

import "time"

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	UserId    int       `json:"user_id"`
	Status    int       `json:"status"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
