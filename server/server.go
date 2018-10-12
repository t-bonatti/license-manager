package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/t-bonatti/license-manager/controller"
	"github.com/t-bonatti/license-manager/datastore"
)

func New(ds datastore.DataStore, url string) *http.Server {

	var mux = mux.NewRouter()
	mux.Path("/status").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	mux.Path("/license").Methods(http.MethodPost).HandlerFunc(controller.Create(ds))
	mux.Path("/license/{id}/versions/{version}").Methods(http.MethodGet).HandlerFunc(controller.Get(ds))

	var server = &http.Server{
		Handler: mux,
		Addr:    url,
	}

	return server
}
