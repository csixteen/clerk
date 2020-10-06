package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/csixteen/clerk/pkg/models"
)

func NotesRoutes(db *sql.DB) []*Route {
	var routes []*Route

	routes = append(routes, listNotes(db))
	routes = append(routes, createNote(db))
	routes = append(routes, appendNote(db))
	routes = append(routes, showNote(db))
	routes = append(routes, deleteNote(db))

	return routes
}

func listNotes(db *sql.DB) *Route {
	// List all the tasks
	return &Route{
		Method: http.MethodGet,
		Path:   "/notes",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			notes, err := models.ListNotes(db)
			if err != nil {
				InternalServerError(w, err)
			} else {
				json.NewEncoder(w).Encode(notes)
			}
		},
	}
}

func createNote(db *sql.DB) *Route {
	// List all the tasks
	return &Route{
		Method: http.MethodPost,
		Path:   "/notes",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				BadRequest(w, err)
				return
			}

			var note map[string]string
			json.Unmarshal(reqBody, &note)

			id, err := models.AddNote(
				db,
				note["name"],
				note["contents"],
				time.Now(),
			)

			if err != nil {
				BadRequest(w, err)
			} else {
				json.NewEncoder(w).Encode(map[string]int64{"Id": id})
			}

		},
	}
}

func appendNote(db *sql.DB) *Route {
	// List all the tasks
	return &Route{
		Method: http.MethodPut,
		Path:   "/notes/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				BadRequest(w, err)
				return
			}

			var c map[string]string
			json.Unmarshal(reqBody, &c)
			note := getIdFromURL(r)

			err = models.AppendNote(
				db,
				note,
				c["contents"],
			)
			if err != nil {
				InternalServerError(w, err)
			} else {
				Ok(w)
			}
		},
	}
}

func showNote(db *sql.DB) *Route {
	// List all the tasks
	return &Route{
		Method: http.MethodGet,
		Path:   "/notes/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			note := getIdFromURL(r)
			n, err := models.GetNote(db, note)
			if err != nil {
				BadRequest(w, err)
			} else {
				json.NewEncoder(w).Encode(n)
			}
		},
	}
}

func deleteNote(db *sql.DB) *Route {
	// List all the tasks
	return &Route{
		Method: http.MethodDelete,
		Path:   "/notes/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			note := getIdFromURL(r)
			err := models.DeleteNote(db, note)
			if err != nil {
				BadRequest(w, err)
			} else {
				Ok(w)
			}
		},
	}
}
