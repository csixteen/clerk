package routes

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/csixteen/clerk/pkg/models"
)

func TasksRoutes(db *sql.DB) []*Route {
	var routes []*Route

	routes = append(routes, listTasks(db))
	routes = append(routes, createTask(db))
	routes = append(routes, modifyTask(db))
	routes = append(routes, deleteTask(db))
	routes = append(routes, completeTask(db))

	return routes
}

func listTasks(db *sql.DB) *Route {
	return &Route{
		Method: http.MethodGet,
		Path:   "/tasks",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			tasks, err := models.ListTasks(db)
			if err != nil {
				InternalServerError(w, err)
			} else {
				json.NewEncoder(w).Encode(tasks)
			}
		},
	}
}

func createTask(db *sql.DB) *Route {
	return &Route{
		Method: http.MethodPost,
		Path:   "/tasks",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				BadRequest(w, err)
				return
			}

			var task models.TaskModel
			json.Unmarshal(reqBody, &task)

			id, err := models.AddTask(
				db,
				task.Name,
				task.Contents,
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

func modifyTask(db *sql.DB) *Route {
	return &Route{
		Method: http.MethodPut,
		Path:   "/tasks/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			reqBody, err := ioutil.ReadAll(r.Body)
			if err != nil {
				BadRequest(w, err)
				return
			}

			var c map[string]string
			json.Unmarshal(reqBody, &c)

			task := getIdFromURL(r)

			err = models.EditTask(db, task, c["contents"])
			if err != nil {
				BadRequest(w, err)
			} else {
				Ok(w)
			}
		},
	}
}

func deleteTask(db *sql.DB) *Route {
	return &Route{
		Method: http.MethodDelete,
		Path:   "/tasks/{id}",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			task := getIdFromURL(r)

			err := models.DeleteTask(db, task)
			if err != nil {
				BadRequest(w, err)
			} else {
				Ok(w)
			}
		},
	}
}

func completeTask(db *sql.DB) *Route {
	return &Route{
		Method: http.MethodPost,
		Path:   "/tasks/{id}/complete",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			task := getIdFromURL(r)

			err := models.CompleteTask(db, task, time.Now())
			if err != nil {
				BadRequest(w, err)
			} else {
				Ok(w)
			}
		},
	}
}
