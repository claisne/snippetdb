package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/claisne/snippetdb/models"
)

var getHomeTemplate *template.Template

func init() {
	var err error
	getHomeTemplate, err = template.ParseFiles("templates/layout.html",
		"templates/home.html", "templates/sign-modal.html", "templates/sign.html")
	if err != nil {
		logrus.Fatal("Failed to parse home templates")
	}
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	user, loggedIn := getUser(r)

	data := struct {
		User     *models.User
		LoggedIn bool
	}{
		User:     user,
		LoggedIn: loggedIn,
	}

	err := getHomeTemplate.Execute(w, data)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Warn("Unable to execute GetHome templates")
	}
}
