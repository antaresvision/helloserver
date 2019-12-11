package main

import (
	"github.com/antaresvision/helloserver/api"
	"github.com/antaresvision/helloserver/db"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	dataStore := db.NewConnection()
	defer dataStore.Close()

	srv := api.NewServer(dataStore)

	r := mux.NewRouter()
	r.Path("/greetings").Methods(http.MethodGet).HandlerFunc(api.GreetingsHandler)
	r.Path("/greetings/{name}").Methods(http.MethodGet).HandlerFunc(api.GreetingsHandler)

	r.Path("/items/{id}").Methods(http.MethodGet).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		api.GetItemById(writer, request, dataStore)
	})
	r.Path("/items/{id}").Methods(http.MethodDelete).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		api.RemoveItemById(writer, request, dataStore)
	})
	r.Path("/items/{id}").Methods(http.MethodPost).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		api.SaveItem(writer, request, dataStore)
	})
	r.Path("/items").Methods(http.MethodGet).HandlerFunc(srv.GetAllItems)


	http.ListenAndServe(":8000", r)
}
