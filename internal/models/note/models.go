package models

import (
	"time"
)

type Note struct {
	Id        string    `db:"id" json:"id"`
	UserId    string    `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	Body      string    `db:"body" json:"body"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateNoteParams struct {
	UserId string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Notes struct {
	Notes []Note
}

type ListNotesParams struct {
	UserId string `json:"user_id"`
}

type GetNoteResponse struct {
	ID        string `db:"id"` 
	UserID    string `db:"user_id"` 
	Title     string `db:"title"`   
	Body      string `db:"body"`
}

type UpdateNoteResponse struct {
	Id        string
	UserId    string
	Title     string
	Body      string
	UpdatedAt time.Time
}
