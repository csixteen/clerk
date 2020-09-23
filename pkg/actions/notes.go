package actions

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type NoteModel struct {
	id        string
	name      string
	contents  []string
	createdAt time.Time
}

func (n *NoteModel) String() string {
	s := "- id: %s | name: %s | created_at: %s%s\n"

	var c strings.Builder
	if len(n.contents) > 0 {
		c.WriteString("\n  Contents: ")
		c.WriteString(strings.Join(n.contents, "; "))
	}

	return fmt.Sprintf(
		s,
		n.id,
		n.name,
		n.createdAt.Format(dateLayout),
		c.String(),
	)
}

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
		err = rows.Scan(&n.id, &n.name, &createdAt)
		if err != nil {
			return nil, err
		}

		cr, _ := time.Parse(dateLayout, createdAt)
		n.createdAt = cr

		res = append(res, n)
	}

	return res, nil
}

func GetNote(db *sql.DB, note string) (*NoteModel, error) {
	field, id := getIdFieldAndValue(note)
	noteMetadataQuery := fmt.Sprintf(`SELECT id, name FROM notes WHERE %s = ?`, field)
	row := db.QueryRow(noteMetadataQuery, id)
	n := new(NoteModel)
	err := row.Scan(&n.id, &n.name)
	if err != nil {
		return nil, err
	}

	noteContentsQuery := `SELECT contents FROM notes_contents WHERE note_id = ?`
	rows, err := db.Query(noteContentsQuery, n.id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contents string

		if err := rows.Scan(&contents); err != nil {
			return nil, err
		}

		n.contents = append(n.contents, contents)
	}

	return n, nil
}

func AddNote(db *sql.DB, name string, contents string, t time.Time) error {
	insertQuery := `INSERT INTO notes(name, created_at) VALUES (?, ?)`
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(name, t.Format(dateLayout))
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	insertContentsQuery := `INSERT INTO notes_contents (note_id, contents) VALUES (?, ?)`
	stmt, err = db.Prepare(insertContentsQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, contents)

	return err
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
