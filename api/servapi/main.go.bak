package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func server(w http.ResponseWriter, r *http.Request) {
	from := r.URL.String()
	if strings.HasPrefix(r.Host, "api") {
		fmt.Fprintf(w, "API not yet available.")
		return
	}
	var to string
	if strings.HasSuffix(r.URL.String(), "week") {
		to = "http://week.skilstak.io"
	} else {
		r.URL.Scheme = "http"
		r.URL.Host = "skilstak.io"
		to = r.URL.String()
	}
	http.Redirect(w, r, to, 301)
	log.Printf("Redirected: %s to %s\n", from, to)
}

func main() {
	http.HandleFunc("/", server)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
