package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Notebook", Descricao: "Notebook Dell", Preco: 3000.0, Quantidade: 10},
		{Nome: "Mouse", Descricao: "Mouse Gamer", Preco: 100.0, Quantidade: 20},
		{Nome: "Teclado", Descricao: "Teclado Gamer", Preco: 200.0, Quantidade: 30},
		{Nome: "Monitor", Descricao: "Monitor Gamer", Preco: 500.0, Quantidade: 40},
		{Nome: "Headset", Descricao: "Headset Gamer", Preco: 300.0, Quantidade: 50},
	}

	err := templates.ExecuteTemplate(w, "Index", produtos)
	if err != nil {
		return
	}
}
