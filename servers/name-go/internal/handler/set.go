package handler

import (
	"fmt"
	"github.com/cbotte21/name-go/internal/datastore"
	"github.com/cbotte21/name-go/internal/schema"
	"net/http"
)

func SetHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //Populate PostForm
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Please try again later.\n"))
		return
	}

	credentials := r.PostForm

	if !credentials.Has("jwt") || !credentials.Has("name") { //HAS email and password
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request must contain a valid jwt and name.\n"))
		return
	}

	query := schema.Name{Id: credentials.Get("jwt"), Name: credentials.Get("name")}

	err = datastore.Update(query, query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Please try again later.\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{ "status": true }`)))
}
