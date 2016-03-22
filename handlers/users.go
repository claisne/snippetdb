package handlers

import (
	"net/http"

	"github.com/claisne/snippetdb/libhttp"
	"github.com/claisne/snippetdb/store"

	"github.com/gorilla/context"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := getIdFromPath(r)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	store := context.Get(r, "store").(store.Store)

	user, err := store.User().Get(id)
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	jsonBytes, err := user.ToJson()
	if err != nil {
		libhttp.HandleErrorJson(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/json")
	w.Write(jsonBytes)
	return
}
