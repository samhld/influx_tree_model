package main

import "net/http"

func Serve(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
