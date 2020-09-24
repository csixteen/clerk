package actions

import (
	"database/sql"
	"fmt"
	"strings"

	m "github.com/csixteen/clerk/pkg/models"
)

type Result interface {
	Type() string
	String() string
}

func searchNotes(db *sql.DB, query string) ([]Result, error) {
	searchNotesQuery := `SELECT DISTINCT id, name, GROUP_CONCAT(contents,'|') as contents
		FROM notes
		INNER JOIN notes_contents ON notes.id = notes_contents.note_id
		WHERE name LIKE '%%%s%%' OR contents LIKE '%%%s%%' GROUP BY id`

	rows, err := db.Query(
		fmt.Sprintf(searchNotesQuery, query, query),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Result
	for rows.Next() {
		var contents string
		n := &m.NoteModel{}
		err = rows.Scan(&n.Id, &n.Name, &contents)
		if err != nil {
			return nil, err
		}
		n.Contents = strings.Split(contents, "|")

		res = append(res, n)
	}

	return res, nil
}

func searchTasks(db *sql.DB, query string) ([]Result, error) {
	searchTasksQuery := `SELECT DISTINCT id, name, contents FROM tasks
		WHERE name LIKE '%%%s%%' OR contents LIKE '%%%s%%'`

	rows, err := db.Query(
		fmt.Sprintf(searchTasksQuery, query, query),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []Result
	for rows.Next() {
		t := &m.TaskModel{}
		err = rows.Scan(&t.Id, &t.Name, &t.Contents)
		if err != nil {
			return nil, err
		}

		res = append(res, t)
	}

	return res, nil
}

func Search(db *sql.DB, query string) ([]Result, error) {
	var res []Result

	tasksResult, err := searchTasks(db, query)
	if err != nil {
		return nil, err
	}
	res = append(res, tasksResult...)

	notesResult, err := searchNotes(db, query)
	if err != nil {
		return nil, err
	}
	res = append(res, notesResult...)

	return res, nil
}
