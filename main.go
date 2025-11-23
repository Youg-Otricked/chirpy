package main

import (
	"log"
	"net/http"
)

func main() {
	servemux := http.NewServeMux()
	servemux.Handle("/", http.FileServer(http.Dir(".")))
	server := http.Server{
		Handler: servemux,
		Addr:    ":8080",
	}
	log.Printf("Serving on address: localhost%s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}
