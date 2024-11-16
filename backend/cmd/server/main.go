package main

import "github.com/torfstack/bitdrifter/pkg/adapter/http"

func main() {
	server := http.NewServer()
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
