package handler

import (
	"fmt"
	"github.com/cbotte21/auth-go/internal/datastore"
	"github.com/cbotte21/auth-go/internal/schema"
	"github.com/cbotte21/auth-go/internal/utilities"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) { //TODO: Update last login
	err := r.ParseForm() //Populate PostForm
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Please try again later.\n"))
		return
	}

	credentials := r.PostForm

	if !credentials.Has("email") || !credentials.Has("password") { //HAS email and password
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Request must contain an email and password.\n"))
		return
	}

	query := schema.User{
		Email: credentials.Get("email"),
	}

	candideUser, err := datastore.Find(query) //Find user with matching email
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email does not exist.\n"))
		return
	}

	if candideUser.VerifyPassword(credentials.Get("password")) != nil { //Validate password
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Username and password do not match."))
		return
	}

	tokenString, err := utilities.GenerateJWT(candideUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Please try again later.\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{ "jwt": "%s" }`, tokenString)))
}
