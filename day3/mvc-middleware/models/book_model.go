package models

import "time"

type Book struct {
	ID        int64     `json:"id"`
	Tittle    string    `json:"tittle"`
	Author    string    `json:"author"`
	ISBN      string    `json:"isbn"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookInput struct {
	Tittle string `json:"tittle" validate:"required"`
	Author string `json:"author" validate:"required"`
	ISBN   string `json:"isbn" validate:"required"`
}
