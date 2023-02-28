package handler

import (
	"github.com/cbotte21/auth-go/internal/datastore"
	"github.com/cbotte21/auth-go/internal/schema"
	"github.com/cbotte21/auth-go/internal/utilities"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
	"time"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
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

	if !utilities.ParseEmail(credentials.Get("email")) || !utilities.ParsePassword(credentials.Get("password")) { //Validate username and password
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email or password does not meet requirements.\n"))
		return
	}

	//Check if an account already exists with email or username
	emailCheckQuery := schema.User{
		Email: credentials.Get("email"),
	}
	_, err = datastore.Find(emailCheckQuery)
	if err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Email is already registered.\n"))
		return
	}

	//Register
	currTime := strconv.FormatInt(time.Now().Unix(), 10)
	candideUser := schema.User{
		Email:            credentials.Get("email"),
		InitialTimestamp: currTime,
		RecentTimestamp:  currTime,
	}

	if candideUser.SetPassword(credentials.Get("password")) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Please try again later.\n"))
		return
	}

	err = datastore.Create(candideUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Please try again later.\n"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{ "status": "account created" }`))
}
