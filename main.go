package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/response"
)

func main() {
	if len(os.Args[1:]) == 1 {
		config.BaseDir = os.Args[1]
	}
	log.Println(string("\033[32m"), "TMAI Mock Server running at http://localhost:3000", string("\033[0m"))

	config.CreateApi()

	for _, ep := range config.Api.Endpoints {
		log.Println(ep)
		if len(ep.Folder) > 0 {
			http.Handle(ep.Path+"/", http.StripPrefix(ep.Path+"/", http.FileServer(http.Dir(ep.Folder))))
		} else if ep.Type == "message" {
			http.HandleFunc(ep.Path, response.MessageResponse)
		} else if ep.Type == "autocomplete" {
			http.HandleFunc(ep.Path, response.AutocompleteResponse)
		} else if ep.Type == "trivia" {
			http.HandleFunc(ep.Path, response.TriviaResponse)
		}
	}

	err := http.ListenAndServe(":"+strconv.Itoa(config.Api.Port), nil)

	if err != nil {
		log.Fatal(" ", err)
	}
}
