// MIT License
//
// Copyright (c) 2020 Pedro Rodrigues
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package actions

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error occurred when opening a stub DB connection: %s", err)
	}

	return db, mock
}

func TestListTasks(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	query := `SELECT
		id, name, contents, created_at, COALESCE\(completed_at,''\) FROM tasks
		ORDER BY id`
	rows := sqlmock.NewRows([]string{
		"id",
		"name",
		"contents",
		"created_at",
		"completed_at",
	}).AddRow("1", "test", "test contents", "2020-09-20 15:00", "")

	mock.ExpectQuery(query).WillReturnRows(rows)

	tasks, err := ListTasks(db)
	assert.Equal(t, 1, len(tasks))
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddTask(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	created := time.Now()
	query := "INSERT INTO tasks\\(name, contents, created_at\\) VALUES \\(\\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(
		"test", "test contents", created.Format(dateLayout),
	).WillReturnResult(sqlmock.NewResult(0, 1))

	err := AddTask(db, "test", "test contents", created)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestEditTask(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	query := "UPDATE tasks SET contents = \\? WHERE name = \\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(
		"new contents", "test",
	).WillReturnResult(sqlmock.NewResult(0, 1))

	err := EditTask(db, "test", "new contents")
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteTask(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	query := "DELETE FROM tasks WHERE name = \\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs("test").WillReturnResult(sqlmock.NewResult(0, 1))

	err := DeleteTask(db, "test")
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestCompleteTask(t *testing.T) {
	db, mock := newMockDB(t)
	defer db.Close()

	completed := time.Now()
	query := "UPDATE tasks SET completed_at = \\? WHERE name = \\?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(
		completed.Format(dateLayout), "test",
	).WillReturnResult(sqlmock.NewResult(0, 1))

	err := CompleteTask(db, "test", completed)
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
