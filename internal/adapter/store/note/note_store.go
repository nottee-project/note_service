package tm

import (
	"context"
	"database/sql"
	"fmt"

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
		Insert(tableNameNote).
		Columns(fieldNameTitle, fieldNameBody).
		Values(note.Title, note.Body, note.Completed).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s, %s",
			fieldNameNoteId,
			fieldNameTitle,
			fieldNameBody,
			fieldNameCreatedAt,
			fieldNameUpdatedAt)).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Note{}, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.Get(&result, sqlStr, args...)

	if err != nil {
		return models.Note{}, errors.Wrap(err, "Get")
	}

	return result, nil
}

func (t *NoteStore) UpdateNote(ctx context.Context, note models.Note) (models.UpdateNoteResponse, error) {
	result := models.UpdateNoteResponse{}

	sqlStr, args, err := squirrel.
		Update(tableNameNote).
		Set(fieldNameTitle, note.Title).
		Set(fieldNameBody, note.Body).
		Set(fieldNameUpdatedAt, squirrel.Expr("NOW()")).
		Where(squirrel.Eq{
			fieldNameNoteId: note.Id,
		}).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s",
			fieldNameNoteId,
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
			&result.Title,
			&result.Body,
			&result.Completed,
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
func (t *NoteStore) GetNote(ctx context.Context, noteId string) (models.GetNoteResponse, error) {
	result := models.GetNoteResponse{}

	query, args, err := squirrel.Select(
		fieldNameNoteId, fieldNameTitle, fieldNameBody,
	).
		From(tableNameNote).
		Where(squirrel.Eq{
			fieldNameNoteId: noteId,
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

func (t *NoteStore) ListNotes(ctx context.Context, listNotesParams models.ListNotesParams) ([]models.GetNoteResponse, error) {
	result := make([]models.GetNoteResponse, 0)

	query := squirrel.Select(
		fieldNameNoteId, fieldNameTitle, fieldNameBody,
	).
		From(tableNameNote).
		PlaceholderFormat(squirrel.Dollar)

	if listNotesParams.Completed {
		query = query.Where(squirrel.Eq{"completed": listNotesParams.Completed})
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "ToSql")
	}

	err = t.Store.DB.SelectContext(ctx, &result, sqlStr, args...)
	if err != nil {
		return nil, errors.Wrap(err, "SelectContext")
	}

	return result, nil
}

func (t *NoteStore) DeleteNote(ctx context.Context, noteID string) error {
	sqlStr, args, err := squirrel.
		Delete(tableNameNote).
		Where(squirrel.Eq{"note_id": noteID}).
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
