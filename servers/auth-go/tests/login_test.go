package tests

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	TestStatus(t)

	data := url.Values{}
	data.Set("email", "cbotte21@gmail.com")
	data.Set("password", "Asdfasdf1")

	body := strings.NewReader(data.Encode())

	resp, err := http.Post("http://localhost:5000/login", "application/x-www-form-urlencoded", body)
	//Handle Error
	if err != nil {
		TestStatus(t)
		return
	}
	defer resp.Body.Close()

	//Read the response body
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	if string(res) == "Username and password do not match." {
		t.Fatalf("invalid login")
	}
}
