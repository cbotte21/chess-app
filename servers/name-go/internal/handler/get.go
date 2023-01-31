package handler

import (
	"fmt"
	"github.com/cbotte21/name-go/internal/datastore"
	"github.com/cbotte21/name-go/internal/schema"
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

	query := schema.Name{
		Id: playerId,
	}

	name, err := datastore.Find(query) //Find user with matching email
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email does not exist.\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{ "_id": "%s", "name": "%s" }`, playerId, name.Name)))
}
