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
	"strings"
	"time"
)

type NoteModel struct {
	Id        string    `json:"idd"`
	Name      string    `json:"name"`
	Contents  []string  `json:"contents"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *NoteModel) String() string {
	s := "- id: %s | name: %s%s%s\n"

	var createdAtStr string
	if (n.CreatedAt == time.Time{}) {
		createdAtStr = ""
	} else {
		createdAtStr = fmt.Sprintf(
			" | created_at: %s",
			n.CreatedAt.Format(dateLayout),
		)
	}

	var c strings.Builder
	if len(n.Contents) > 0 {
		c.WriteString("\n  Contents: ")
		c.WriteString(strings.Join(n.Contents, "; "))
	}

	return fmt.Sprintf(
		s,
		n.Id,
		n.Name,
		createdAtStr,
		c.String(),
	)
}

func (n *NoteModel) Type() string {
	return "note"
}

// ListNotes lists all the existing notes. The displayed
// fields are the note `id`, the note `name` and `created_at`.
func ListNotes(db *sql.DB) ([]*NoteModel, error) {
	rows, err := db.Query(`SELECT
		id, name, created_at FROM notes
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*NoteModel
	for rows.Next() {
		var createdAt string
		n := &NoteModel{}
		err = rows.Scan(&n.Id, &n.Name, &createdAt)
		if err != nil {
			return nil, err
		}

		cr, _ := time.Parse(dateLayout, createdAt)
		n.CreatedAt = cr

		res = append(res, n)
	}

	return res, nil
}

func GetNote(db *sql.DB, note string) (*NoteModel, error) {
	field, id := getIdFieldAndValue(note)
	noteMetadataQuery := fmt.Sprintf(
		`SELECT id, name, created_at FROM notes WHERE %s = ?`,
		field,
	)
	row := db.QueryRow(noteMetadataQuery, id)
	n := new(NoteModel)
	var createdAt string
	err := row.Scan(&n.Id, &n.Name, &createdAt)
	if err != nil {
		return nil, err
	}

	cr, _ := time.Parse(dateLayout, createdAt)
	n.CreatedAt = cr

	noteContentsQuery := `SELECT contents FROM notes_contents WHERE note_id = ?`
	rows, err := db.Query(noteContentsQuery, n.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contents string

		if err := rows.Scan(&contents); err != nil {
			return nil, err
		}

		n.Contents = append(n.Contents, contents)
	}

	return n, nil
}

func AddNote(db *sql.DB, name string, contents string, t time.Time) (int64, error) {
	insertQuery := `INSERT INTO notes(name, created_at) VALUES (?, ?)`
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(name, t.Format(dateLayout))
	if err != nil {
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	insertContentsQuery := `INSERT INTO notes_contents (note_id, contents) VALUES (?, ?)`
	stmt, err = db.Prepare(insertContentsQuery)
	if err != nil {
		return id, err
	}

	_, err = stmt.Exec(id, contents)

	return id, err
}

func AppendNote(db *sql.DB, note string, contents string) error {
	field, id := getIdFieldAndValue(note)

	var placeHolder string
	if field == "name" {
		placeHolder = "SELECT id FROM notes WHERE name = ?"
	} else {
		placeHolder = "?"
	}

	appendQuery := fmt.Sprintf(
		`INSERT INTO notes_contents (note_id, contents) VALUES ((%s), ?)`,
		placeHolder,
	)

	stmt, err := db.Prepare(appendQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, contents)

	return err
}

func DeleteNote(db *sql.DB, note string) error {
	field, id := getIdFieldAndValue(note)
	deleteQuery := fmt.Sprintf(`DELETE FROM notes WHERE %s = ?`, field)
	stmt, err := db.Prepare(deleteQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	return err
}
