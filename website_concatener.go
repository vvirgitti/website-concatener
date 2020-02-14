package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main () {
	http.HandleFunc("/", ServerHandler)
	http.ListenAndServe(":4000", nil)
}

func ServerHandler (w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()["url"]

	if queryParams != nil {
		resp, err := http.Get(queryParams[0])
		if err != nil {
			fmt.Println("error getting the website with status code", resp.StatusCode)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		s := string(body)

		fmt.Fprintf(w, s)
	} else {
		errors.New("a url has not been provided")
	}
}