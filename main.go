package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const listenAddr = ":8000"

func main() {
	log.Printf("Staring listening at: %s", listenAddr)

	err := http.ListenAndServe(listenAddr, http.HandlerFunc(echo))
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
