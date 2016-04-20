package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/claisne/snippetdb/models"

	"github.com/gorilla/context"
)

var getAccountTemplate *template.Template

func init() {
	var err error
	getAccountTemplate, err = template.ParseFiles("templates/layout.html",
		"templates/users/account.html", "templates/navbar.html")
	if err != nil {
		logrus.Fatal("Failed to parse account templates")
	}
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "user").(*models.User)

	err := getAccountTemplate.Execute(w, user)
	if err != nil {
		logrus.Warn(err.Error())
	}
}
