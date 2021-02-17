package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestListNotes(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	query := `SELECT
		id, name, created_at FROM notes
		ORDER BY id`
	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"created_at",
	}).AddRow("1", "test", "2020-10-11 19:28")

	mock.ExpectQuery(query).WillReturnRows(rows)

	notes, err := ListNotes(db)
	assert.Equal(t, 1, len(notes))
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}
