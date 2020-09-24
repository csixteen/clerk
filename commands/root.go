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

package commands

import (
	"database/sql"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	database *sql.DB
	dbFile   string

	RootCmd = &cobra.Command{
		Use:   "clerk",
		Short: "clerk is your command-line personal Jarvis.",
	}
)

func init() {
	var err error

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	dbFile = path.Join(homeDir, ".clerk.db")

	database, err = setupDatabase()
	if err != nil {
		panic(err)
	}

	addCommands()
}

func addCommands() {
	RootCmd.AddCommand(Notes())
	RootCmd.AddCommand(Tasks())
	RootCmd.AddCommand(Search())
}

func Execute() {
	defer database.Close()
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

//-------------------------------------------
//         Database related stuff

func setupDatabase() (*sql.DB, error) {
	var err error

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		file, err := os.Create(dbFile)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, _ := sql.Open(
		"sqlite3",
		fmt.Sprintf("%s?_foreign_keys=true", dbFile),
	)
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
