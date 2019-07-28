package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)
	directory := flag.String("d", ".", "")
	http.Handle("/", http.FileServer(http.Dir(*directory)))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
