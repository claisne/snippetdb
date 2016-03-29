// Package handlers provides request handlers.
package handlers

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/claisne/snippetdb/models"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

const (
	ErrorLevelFail = "fail"
)

type Error struct {
	Level        string `json:"level"`
	Message      string `json:"message"`
	StatusCode   int    `json:"-"`
	RedirectPath string `json:"-"`
}

func init() {
	gob.Register(&Error{})
}

func getIdFromPath(r *http.Request) (int64, error) {
	idString := mux.Vars(r)["id"]
	if idString == "" {
		return -1, errors.New("user id cannot be empty.")
	}

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return -1, err
	}

	return id, nil
}

func getUser(r *http.Request) (*models.User, bool) {
	sessionStore := context.Get(r, "sessionStore").(sessions.Store)
	session, _ := sessionStore.Get(r, "snippetdb-session")
	user, ok := session.Values["user"].(*models.User)
	return user, ok
}

func getFlashesErrors(session *sessions.Session) []*Error {
	flashes := session.Flashes()
	errors := make([]*Error, 0, 5)

	for _, flash := range flashes {
		if hErr, ok := flash.(*Error); ok {
			errors = append(errors, hErr)
		}
	}

	return errors
}

func (hErr *Error) Render(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	if contentType == "application/json" {
		jsonBytes, err := json.Marshal(hErr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(hErr.StatusCode)
		w.Header().Set("Content-Type", "text/json")
		w.Write(jsonBytes)
	} else {
		sessionStore := context.Get(r, "sessionStore").(sessions.Store)
		session, _ := sessionStore.Get(r, "snippetdb-session")

		session.AddFlash(hErr)
		session.Save(r, w)

		http.Redirect(w, r, hErr.RedirectPath, http.StatusFound)
	}
}
