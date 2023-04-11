package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	flag.StringVar(&addr, "addr", addr, "address to listen on")

	log.Println("Listening on", addr)

	http.ListenAndServe(addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
}
