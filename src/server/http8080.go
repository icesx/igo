package server

import (
	"fmt"
	"net/http"
)

// Default Request Handler
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", r.URL.Path[1:])
}

func StartServer() {
	http.HandleFunc("/", defaultHandler)
	fmt.Print("dd")
	http.ListenAndServe(":8080", nil)
}
