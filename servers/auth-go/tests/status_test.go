package tests

import (
	"net/http"
	"testing"
)

func TestStatus(t *testing.T) {
	_, err := http.Get("http://localhost:5000")
	if err != nil {
		t.Fatalf("api is not online")
	}
}
