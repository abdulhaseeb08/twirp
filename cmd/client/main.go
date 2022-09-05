package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/abdulhaseeb08/twirp/haberdasher"
)

func main() {
	client := haberdasher.NewHaberdasherProtobufClient("http://localhost:8080", &http.Client{})

	hat, err := client.MakeHat(context.Background(), &haberdasher.Size{Inches: 12})
	if err != nil {
		fmt.Printf("error sorrryy: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Your new hat!: %+v", hat)
}
