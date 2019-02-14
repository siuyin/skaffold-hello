package web

import (
	"fmt"
	"log"
	"net/http"
)

// Serve launces a web server.
func Serve() {
	go func() {
		log.Println("webServer starting")
		http.HandleFunc("/ans", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "The answer is 42.")
		})
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello, this is from a go webserver running in kubernetes.")
		})
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}
