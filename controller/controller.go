package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/t-bonatti/license-manager/datastore"
	"github.com/t-bonatti/license-manager/model"
)

func Create(ds datastore.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var license model.License
		if err := json.NewDecoder(r.Body).Decode(&license); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		license.CreatedAt = time.Now()
		if err := ds.Create(license); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	}
}

func Get(ds datastore.DataStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		version := mux.Vars(r)["version"]

		license, err := ds.Get(id, version)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if license.ID == "" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(license); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
