package tm

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Masterminds/squirrel"
	models "github.com/nottee-project/note_service/internal/models/note"
	"github.com/pkg/errors"

	store "github.com/nottee-project/note_service/internal/adapter/store"
)

type NoteStore struct {
	Store *store.Store
}

func (t *NoteStore) CreateNote(ctx context.Context, note models.CreateNoteParams) (models.Note, error) {
	result := models.Note{}

	sqlStr, args, err := squirrel.
		Insert(tableNameNotes).
		Columns(fieldNameUserId, fieldNameTitle, fieldNameBody).
		Values(note.UserId, note.Title, note.Body).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s, %s",
			fieldNameNoteId,
			fieldNameUserId,
			fieldNameTitle,
			fieldNameBody,
			fieldNameCreatedAt,
			fieldNameUpdatedAt)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Note{}, errors.Wrap(err, "ToSql")
	}

	log.Printf("SQL Query: %s, Args: %v", sqlStr, args)

	err = t.Store.DB.Get(&result, sqlStr, args...)
	if err != nil {
		return models.Note{}, errors.Wrap(err, "Get")
	}

	return result, nil
}

func (t *NoteStore) UpdateNote(ctx context.Context, note models.Note) (models.UpdateNoteResponse, error) {
	result := models.UpdateNoteResponse{}

	sqlStr, args, err := squirrel.
		Update(tableNameNotes).
		Set(fieldNameTitle, note.Title).
		Set(fieldNameBody, note.Body).
		Set(fieldNameUpdatedAt, squirrel.Expr("NOW()")).
		Where(squirrel.Eq{
			fieldNameNoteId: note.Id,
			fieldNameUserId: note.UserId,
		}).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s",
			fieldNameNoteId,
			fieldNameUserId,
			fieldNameTitle,
			fieldNameBody,
			fieldNameUpdatedAt)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.UpdateNoteResponse{}, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.QueryRow(sqlStr, args...).
		Scan(
			&result.Id,
			&result.UserId,
			&result.Title,
			&result.Body,
			&result.UpdatedAt,
		)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.UpdateNoteResponse{}, errors.New("note not found")
		}
		return models.UpdateNoteResponse{}, errors.Wrap(err, "QueryRowContext")
	}

	return result, nil
}

func (t *NoteStore) GetNote(ctx context.Context, noteId, userId string) (models.GetNoteResponse, error) {
	result := models.GetNoteResponse{}

	query, args, err := squirrel.Select(
		fieldNameNoteId, fieldNameUserId, fieldNameTitle, fieldNameBody,
	).
		From(tableNameNotes).
		Where(squirrel.Eq{
			fieldNameNoteId: noteId,
			fieldNameUserId: userId,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.GetNoteResponse{}, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.Get(&result, query, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.GetNoteResponse{}, NoteNotFound
		}
		return models.GetNoteResponse{}, errors.Wrap(err, "Get")
	}

	return result, nil
}

func (t *NoteStore) ListNotes(ctx context.Context, listParams models.ListNotesParams) ([]models.GetNoteResponse, error) {
	result := make([]models.GetNoteResponse, 0)

	query := squirrel.Select(
		fieldNameNoteId, fieldNameUserId, fieldNameTitle, fieldNameBody,
	).
		From(tableNameNotes).
		Where(squirrel.Eq{
			fieldNameUserId: listParams.UserId,
		}).
		PlaceholderFormat(squirrel.Dollar)

	sqlStr, args, err := query.ToSql()
	if err != nil {
		log.Printf("Error generating SQL: %v", err)
		return nil, errors.Wrap(err, "ToSql")
	}

	log.Printf("Generated SQL: %s, Args: %v", sqlStr, args)

	err = t.Store.DB.SelectContext(ctx, &result, sqlStr, args...)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, errors.Wrap(err, "SelectContext")
	}

	return result, nil
}
func (t *NoteStore) DeleteNote(ctx context.Context, noteID, userID string) error {
	sqlStr, args, err := squirrel.
		Delete(tableNameNotes).
		Where(squirrel.Eq{
			fieldNameNoteId: noteID,
			fieldNameUserId: userID,
		}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "ToSql")
	}

	_, err = t.Store.DB.ExecContext(ctx, sqlStr, args...)
	if err != nil {
		return errors.Wrap(err, "ExecContext")
	}

	return nil
}
