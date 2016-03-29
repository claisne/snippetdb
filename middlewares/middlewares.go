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
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			context.Set(req, "sessionStore", store)
			next.ServeHTTP(res, req)
		})
	}
}

func SetStore(store store.Store) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			context.Set(req, "store", store)
			next.ServeHTTP(res, req)
		})
	}
}

// MustLogin is a middleware that checks existence of current user.
func MustLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		sessionStore := context.Get(req, "sessionStore").(sessions.Store)
		session, _ := sessionStore.Get(req, "snippetdb-session")
		user, ok := session.Values["user"].(*models.User)

		if !ok {
			http.Redirect(res, req, "/login", 302)
			return
		}

		context.Set(req, "user", user)
		next.ServeHTTP(res, req)
	})
}
