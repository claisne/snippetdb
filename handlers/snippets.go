package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var getNewSnippetTemplate *template.Template

func init() {
	var err error
	getNewSnippetTemplate, err = template.ParseFiles("templates/layout.html",
		"templates/snippets/new.html", "templates/navbar.html")
	if err != nil {
		logrus.Fatal("Failed to parse snippets templates")
	}
}

func GetNewSnippet(w http.ResponseWriter, r *http.Request) {
	// user, loggedIn := getUser(r)

	err := getNewSnippetTemplate.Execute(w, nil)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Unable to execute GetNewSnippet templates")
	}
}
