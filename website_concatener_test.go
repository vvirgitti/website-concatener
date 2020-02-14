package main

import (
	"net/http/httptest"
	"testing"
	"net/http"
)

func TestReturnStatusOK(t *testing.T) {
	req, err := http.NewRequest("GET", "?url=http://www.bbc.co.uk", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()

	ServerHandler(resp, req)

	got := resp.Body.String()
	want := resp.Body.String()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}