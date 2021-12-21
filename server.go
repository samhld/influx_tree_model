package main

import "net/http"

// This is an API server that proxies InfluxDB and holds (in-mem) a hierarchical structure of Line Protocl elements

func Serve(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
