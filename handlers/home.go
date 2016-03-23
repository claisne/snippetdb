package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
)

var getHomeTemplates *template.Template

func init() {
	var err error
	getHomeTemplates, err = template.ParseFiles("templates/layout.html", "templates/home.html", "templates/sign-modal.html")
	if err != nil {
		logrus.Fatal("Failed to parse home templates")
	}
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	err := getHomeTemplates.Execute(w, nil)
	if err != nil {
		logrus.Warn(err.Error())
	}
}
