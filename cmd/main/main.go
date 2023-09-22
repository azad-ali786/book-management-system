package main 

import (
	"log"
	"net/http"
	"github.com/azad-ali786/bookmanagementsystem/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreROutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}