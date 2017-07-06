package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/t-bonatti/license-manager/controller"
	"gopkg.in/mgo.v2"
)

func New(session *mgo.Session) *http.Server {

	var mux = mux.NewRouter()
	mux.Path("/status").Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	mux.Path("/license").Methods(http.MethodPost).HandlerFunc(controller.CreateLicense(session))
	mux.Path("/license/{id}/versions/{version}").Methods(http.MethodGet).HandlerFunc(controller.GetLicense(session))

	var server = &http.Server{
		Handler: mux,
		Addr:    ":3000",
	}

	return server
}
