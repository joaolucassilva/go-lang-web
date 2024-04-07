package controllers

import (
	"github.com/joaolucassilva/go-lang-web/models"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}
