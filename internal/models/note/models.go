package models

import (
	"time"
)

type Note struct {
	Id        string    `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	Body      string    `db:"body" json:"body"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CreateNoteParams struct {
	Title     string
	Body      string
	Completed bool
}

type Notes struct {
	Notes []Note
}

type ListNotesParams struct {
	Completed bool
}

type GetNoteResponse struct {
	Id        string
	Title     string
	Body      string
	Completed bool
}

type UpdateNoteResponse struct {
	Id        string
	Title     string
	Body      string
	Completed bool
	UpdatedAt  time.Time
}
