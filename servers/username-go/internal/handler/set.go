package handler

import (
	"fmt"
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

	if !credentials.Has("jwt") || !credentials.Has("username") { //HAS email and password
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request must contain a valid jwt and the desired username.\n"))
		return
	}

	//TODO: Validate JWT, get ID

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{ "status": true }`)))
}
