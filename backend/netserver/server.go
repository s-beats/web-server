package netserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// https://golang.org/doc/articles/wiki/#tmp_1

func Start() {
	middleware := func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTION")
			w.Header().Set("Content-Type", "application/json")
			f(w, r)
		}
	}

	// health check
	http.HandleFunc("/health", middleware(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal("ok")
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(b))
	}))

	// example
	http.HandleFunc("/ping", middleware(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal("pong")
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(b))
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
