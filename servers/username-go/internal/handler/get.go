package handler

import (
	"fmt"
	"github.com/cbotte21/username-go/internal/datastore"
	"github.com/cbotte21/username-go/internal/schema"
	"github.com/gorilla/mux"
	"net/http"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //Populate PostForm
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Please try again later.\n"))
		return
	}

	playerId := mux.Vars(r)["_id"]
	//TODO: Verify playerId is present

	query := schema.Username{
		Id: playerId,
	}

	username, err := datastore.Find(query) //Find user with matching email
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("_id is not linked to a username.\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{ "_id": "%s", "username": "%s" }`, playerId, username.Username)))
}
