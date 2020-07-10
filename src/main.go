package main

import (
	"loja/routes"
	"net/http"
)

//MAIN
func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
