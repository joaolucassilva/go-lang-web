package main

import (
	"github.com/joaolucassilva/go-lang-web/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
