package routes

import (
	"github.com/joaolucassilva/go-lang-web/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
