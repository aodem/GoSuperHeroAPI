package main

import (
	"net/http"

	routes "github.com/aodem/super_heros/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/marvel", routes.AllCharacters)
	r.HandleFunc("/new", routes.NewCharacter)
	r.HandleFunc("/new/{name}", routes.UpdateCharacter)
	r.Path("/api").Queries("name", "{name}").HandlerFunc(routes.DeleteCharacter)
	http.ListenAndServe(":8080", r)
}
