package main

import (
	"net/http"

	"github.com/abdulhaseeb08/twirp/haberdasher"
	"github.com/abdulhaseeb08/twirp/serverhaberdasher"
)

func main() {
	server := &serverhaberdasher.Server{} //implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server)

	http.ListenAndServe(":8080", twirpHandler)
}
