package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	log.Printf("Staring listening at port %s", port)

	err := http.ListenAndServe(":"+port, http.HandlerFunc(echo))
	if err != nil {
		log.Fatalf("Server exited with error: %s", err)
	}
}

func echo(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")
	words := strings.Split(path, "/")
	resp := strings.Join(words, " ")
	fmt.Fprint(w, resp)
}
