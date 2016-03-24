package handlers

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
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
	err := getHomeTemplate.Execute(w, nil)
	if err != nil {
		logrus.Warn(err.Error())
	}
}
