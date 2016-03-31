// Package middlewares provides common middleware handlers.
package middlewares

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"

	"github.com/claisne/snippetdb/models"
	"github.com/claisne/snippetdb/store"
)

func SetSessionStore(store sessions.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "sessionStore", store)
			next.ServeHTTP(w, r)
		})
	}
}

func SetStore(store store.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			context.Set(r, "store", store)
			next.ServeHTTP(w, r)
		})
	}
}

// MustLogin is a middleware that checks existence of current user.
func MustLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionStore := context.Get(r, "sessionStore").(sessions.Store)
		session, _ := sessionStore.Get(r, "snippetdb-session")
		user, ok := session.Values["user"].(*models.User)

		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		context.Set(r, "user", user)
		next.ServeHTTP(w, r)
	})
}
