// Package handlers provides request handlers.
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
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

func (hErr *Error) Render(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	logrus.Info(contentType)

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
