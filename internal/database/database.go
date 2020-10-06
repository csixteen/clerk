package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDatabase() (*sql.DB, error) {
	var err error
	var dbFile string

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dbFile = path.Join(homeDir, ".clerk.db")

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Println("Database not found. Creating...")
		file, err := os.Create(dbFile)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := sql.Open(
		"sqlite3",
		fmt.Sprintf("%s?_foreign_keys=true", dbFile),
	)
	if err != nil {
		return nil, err
	}

	err = createTables(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	// Tasks table
	createTasksTable := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64),
		contents TEXT,
		created_at VARCHAR(64),
		completed_at VARCHAR(64)
	);`

	stmt, err := db.Prepare(createTasksTable)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	// Notes table
	createNotesTable := `CREATE TABLE IF NOT EXISTS notes (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(64),
		created_at VARCHAR(64)
	);`

	stmt, err = db.Prepare(createNotesTable)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	// Notes-Contents table
	createNotesContentsTable := `CREATE TABLE IF NOT EXISTS notes_contents (
		note_id INTEGER NOT NULL,
		contents TEXT,
		FOREIGN KEY (note_id)
			REFERENCES notes (id)
				ON DELETE CASCADE
	);`

	stmt, err = db.Prepare(createNotesContentsTable)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()

	return err
}
