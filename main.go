package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Printf("Staring listening at port %s", port)

	err := http.ListenAndServe("0.0.0.0:"+port, http.HandlerFunc(echo))
	if err != nil {
		fmt.Printf("Server exited with error: %s", err)
		os.Exit(1)
	}
}

const html = `<!DOCTYPE html>
<html lang="en">
<head>
  <title>Title of the document</title>
</head>
<body>

<u>%s</u>

</body>
</html>`

func echo(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimLeft(r.URL.Path, "/")
	words := strings.Split(path, "/")
	resp := strings.Join(words, " ")
	
	if r.URL.Query().Get("format") == "upper" {
		resp = strings.ToUpper(resp)
	}
	fmt.Fprintf(w, html, resp)
}
