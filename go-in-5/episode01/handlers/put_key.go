package handlers

import (
	"go-in-5/episode01/storage"
	"io/ioutil"
	"net/http"
)

func PutKey(db storage.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		//val := r.URL.Query().Get("val")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		val, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading the request body", http.StatusInternalServerError)
			return
		}

		if err = db.Set(key, val); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
