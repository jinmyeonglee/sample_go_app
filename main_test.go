package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Init(t *testing.T) {
	ts := httptest.NewServer(setupServer())
	defer ts.Close()
	response, err := http.Get(fmt.Sprintf("%s/", ts.URL))

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if response.StatusCode != 200 {
		t.Fatalf("Expected status code 200, but got %v", response.StatusCode)
	}
}
