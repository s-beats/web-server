package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	hostname string
	port     string
	address  string
)

func setenv() {
	hostname = os.Getenv("HOSTNAME")
	port = os.Getenv("PORT")
	address = fmt.Sprintf("%s:%s", hostname, port)
}

// https://golang.org/doc/articles/wiki/#tmp_1
func Start() {
	setenv()

	loggerMiddleware := func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			f(w, r)
			log.Printf("[%v] ", r.Method)
		}
	}

	corsMiddleware := func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// CORS の許可
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTION")
			w.Header().Set("Content-Type", "application/json")
			f(w, r)
		}
	}

	// health check
	http.HandleFunc("/health", loggerMiddleware(corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal("ok")
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(b))
	})))

	// example
	http.HandleFunc("/example", loggerMiddleware(corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(
			struct {
				String string
				Int    int
				Bool   bool
			}{
				String: "example",
				Int:    1,
				Bool:   true,
			},
		)
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(b))
	})))

	// ping
	http.HandleFunc("/ping", loggerMiddleware(corsMiddleware(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal("{ping:pong}")
		if err != nil {
			panic(err)
		}
		io.WriteString(w, string(b))
	})))

	// Only localhost
	fmt.Printf("listen server address: %s", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
