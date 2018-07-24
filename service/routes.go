package service

import (
	"log"
	"net/http"

	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"LookupCharacters",
		"GET",
		"/lookupCharacters/{characters}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			db, err := bolt.Open("dictionary/dict.db", 0600, &bolt.Options{ReadOnly: true})
			if err != nil {
				log.Fatal(err)
			}
			var jsonresponse []byte
			err = db.View(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("dict"))
				jsonresponse = b.Get([]byte(vars["characters"]))
				return nil
			})
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.Write(jsonresponse)
		},
	},
}
