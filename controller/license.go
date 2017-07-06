package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/t-bonatti/license-manager/model"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CreateLicense(s *mgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var license model.License
		if err := json.NewDecoder(r.Body).Decode(&license); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		license.CreatedAt = time.Now()
		c := session.DB("store").C("licenses")

		if err := c.Insert(license); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetLicense(s *mgo.Session) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		id := mux.Vars(r)["id"]
		version := mux.Vars(r)["version"]

		c := session.DB("store").C("licenses")

		var license model.License
		if err := c.Find(bson.M{"id": id, "version": version}).One(&license); err != nil {
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
