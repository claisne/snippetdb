package application

import (
	"net/http"

	"github.com/carbocation/interpose"
	gorilla_mux "github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"

	"github.com/claisne/snippetdb/handlers"
	"github.com/claisne/snippetdb/middlewares"
	"github.com/claisne/snippetdb/store"
)

// New is the constructor for Application struct.
func New(config *viper.Viper) (*Application, error) {
	databaseDriver := config.Get("database_driver").(string)
	dsn := config.Get("dsn").(string)
	cookieStoreSecret := config.Get("cookie_secret").(string)

	store, err := store.NewSqlStore(databaseDriver, dsn)
	if err != nil {
		return nil, err
	}

	app := &Application{}
	app.config = config
	app.sessionStore = sessions.NewCookieStore([]byte(cookieStoreSecret))
	app.store = store

	return app, nil
}

// Application is the application object that runs HTTP server.
type Application struct {
	config       *viper.Viper
	sessionStore sessions.Store
	store        store.Store
}

func (app *Application) MiddlewareStruct() (*interpose.Middleware, error) {
	middle := interpose.New()
	middle.Use(middlewares.SetSessionStore(app.sessionStore))
	middle.Use(middlewares.SetStore(app.store))

	middle.UseHandler(app.mux())

	return middle, nil
}

func (app *Application) mux() *gorilla_mux.Router {
	router := gorilla_mux.NewRouter()
	MustLogin := middlewares.MustLogin

	// Home
	router.Handle("/", http.HandlerFunc(handlers.GetHome)).Methods("GET")

	// Auth
	router.Handle("/login", http.HandlerFunc(handlers.GetSign)).Methods("GET")
	router.Handle("/login", http.HandlerFunc(handlers.PostLogin)).Methods("POST")
	router.Handle("/register", http.HandlerFunc(handlers.GetSign)).Methods("GET")
	router.Handle("/register", http.HandlerFunc(handlers.PostRegister)).Methods("POST")
	router.Handle("/logout", http.HandlerFunc(handlers.GetLogout)).Methods("GET")

	// Users
	router.Handle("/account", MustLogin(http.HandlerFunc(handlers.GetAccount))).Methods("GET")

	// Snippets
	snippetsRouter := router.PathPrefix("/snippets").Subrouter()
	snippetsRouter.Handle("/new", MustLogin(http.HandlerFunc(handlers.GetNewSnippet))).Methods("GET")

	// Path of static files must be last!
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return router
}
