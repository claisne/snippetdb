package utils

import (
	"bytes"
	"html/template"
	"net/http"
)

var htmlTemplates *template.Template

type HTMLTemplate struct {
	TemplateName string
	Props        map[string]string
}

func InitHTML() error {
	templatesDir, err := FindDir("templates")
	if err != nil {
		return err
	}

	htmlTemplates, err = template.ParseGlob(templatesDir + "*.html")
	if err != nil {
		return err
	}
}

func NewHTMLTemplate(templateName string, locale string) *HTMLTemplate {
	return &HTMLTemplate{
		TemplateName: templateName,
		Props:        make(map[string]string),
	}
}

func (t *HTMLTemplate) Render() (string, error) {
	var text bytes.Buffer

	err := htmlTemplates.ExecuteTemplate(&text, t.TemplateName, t)
	if err != nil {
		return nil, err
	}

	return text.String(), nil
}

func (t *HTMLTemplate) RenderToWriter(w http.ResponseWriter) error {
	return htmlTemplates.ExecuteTemplate(w, t.TemplateName, t)
}
