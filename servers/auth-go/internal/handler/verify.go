package handler

import (
	"github.com/cbotte21/auth-go/internal/utilities"
	"net/http"
)

func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //Populate PostForm
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Please try again later.\n"))
		return
	}

	payload := r.PostForm

	if !payload.Has("jwt") { //HAS email and password
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request must contain a jwt.\n"))
		return
	}

	//Parse JWT
	err = utilities.ValidateJWT(payload.Get("jwt"))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("{ \"status\": \"account not authorized\" }`"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "status": "account authorized" }`))
}
