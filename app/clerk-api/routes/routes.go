package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func Ok(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}

func BadRequest(w http.ResponseWriter, err error) {
	http.Error(
		w,
		err.Error(),
		http.StatusBadRequest,
	)
}

func InternalServerError(w http.ResponseWriter, err error) {
	http.Error(
		w,
		err.Error(),
		http.StatusInternalServerError,
	)
}

func getIdFromURL(r *http.Request) string {
	vars := mux.Vars(r)
	return fmt.Sprintf("#%s", vars["id"])
}
