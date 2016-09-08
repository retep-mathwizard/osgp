package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":", http.FileServer(http.Dir("/usr/share/doc"))))
}
