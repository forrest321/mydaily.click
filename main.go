package main

import (
	"github.com/gorilla/mux"
	"net/http"
"github.com/forrest321/mydaily.click/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)

	http.ListenAndServe(":8000", r)
}
