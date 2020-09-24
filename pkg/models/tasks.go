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

package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// TaskModel struct representation of a row in `tasks` table
type TaskModel struct {
	Id          string
	Name        string
	Contents    string
	CreatedAt   time.Time
	CompletedAt time.Time
}

// String returns a printable representation of a Task
func (t *TaskModel) String() string {
	var createdAtStr string
	if (t.CreatedAt == time.Time{}) {
		createdAtStr = ""
	} else {
		createdAtStr = fmt.Sprintf(
			" | created_at: %s",
			t.CreatedAt.Format(dateLayout),
		)
	}

	return fmt.Sprintf(
		"- id: %s | name: %s%s\n  Contents: %s\n",
		t.Id,
		t.Name,
		createdAtStr,
		t.Contents,
	)
}

func (t *TaskModel) Type() string {
	return "task"
}

// ListTask returns a slice of TaskModels ordered by `id`
func ListTasks(db *sql.DB) ([]*TaskModel, error) {
	rows, err := db.Query(`SELECT 
		id, name, contents, created_at, COALESCE(completed_at,'') FROM tasks
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*TaskModel
	for rows.Next() {
		var createdAt, completedAt string
		t := &TaskModel{}
		err = rows.Scan(&t.Id, &t.Name, &t.Contents, &createdAt, &completedAt)
		if err != nil {
			return nil, err
		}

		cr, _ := time.Parse(dateLayout, createdAt)
		t.CreatedAt = cr
		co, coErr := time.Parse(dateLayout, completedAt)
		if coErr == nil {
			t.CompletedAt = co
		}

		res = append(res, t)
	}

	return res, nil
}

// AddTask adds a new task given a name, its contents and creation time
func AddTask(db *sql.DB, name string, contents string, t time.Time) error {
	insertQuery := `INSERT INTO tasks(name, contents, created_at) VALUES (?, ?, ?)`
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name, contents, t.Format(dateLayout))

	return err
}

// EditTask sets the contents of a task
func EditTask(db *sql.DB, task string, contents string) error {
	field, id := getIdFieldAndValue(task)
	editQuery := fmt.Sprintf(`UPDATE tasks SET contents = ? WHERE %s = ?`, field)
	stmt, err := db.Prepare(editQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(contents, id)

	return err
}

// DeleteTask deletes a task given its name or id. If `task` starts with a '#',
// then it refers to the task id: #123 refers to id 123.
func DeleteTask(db *sql.DB, task string) error {
	field, id := getIdFieldAndValue(task)
	deleteQuery := fmt.Sprintf(`DELETE FROM tasks WHERE %s = ?`, field)
	stmt, err := db.Prepare(deleteQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	return err
}

// CompleteTask marks a task as completed by setting its `completed_at` field
// to the current time.
func CompleteTask(db *sql.DB, task string, t time.Time) error {
	field, id := getIdFieldAndValue(task)
	completeQuery := fmt.Sprintf(`UPDATE tasks SET completed_at = ? WHERE %s = ?`, field)
	stmt, err := db.Prepare(completeQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(t.Format(dateLayout), id)
	return err
}
