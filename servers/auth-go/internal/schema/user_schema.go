package schema

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

//user struct

const DATABASE string = "auth"
const COLLECTION string = "users"

type User struct { //Payload
	Id               string `bson:"_id,omitempty"`
	Email            string `bson:"email,omitempty"`
	Password         string `bson:"password,omitempty"`
	InitialTimestamp string `bson:"intitial_timestamp,omitempty"`
	RecentTimestamp  string `bson:"recent_timestamp,omitempty"`
	Role             int    `bson:"role,omitempty"`
}

func (user User) Database() string {
	return DATABASE
}

func (user User) Collection() string {
	return COLLECTION
}

func (user *User) SetPassword(candidePassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(candidePassword), 10)
	if err != nil {
		return errors.New("could not hash password")
	}
	user.Password = string(hash)
	return nil
}

func (user User) VerifyPassword(candidePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(candidePassword))
}
