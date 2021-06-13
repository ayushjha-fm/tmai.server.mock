package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	charsetUTF8             = "charset=UTF-8"
	CONV_TOKEN              = "CONV-0123456789"
	HeaderAccept            = "Accept"
	HeaderContentType       = "Content-Type"
	HeaderConversationToken = "Conversation-Token"
	MIMEApplicationJSON     = "application/json"
)

type Endpoint struct {
	Type    string   `json:"type"`
	Methods []string `json:"methods"`
	Status  int      `json:"status"`
	Path    string   `json:"path"`
	Folder  string   `json:"folder"`
}

type API struct {
	Host      string     `json:"host"`
	Port      int        `json:"port"`
	Endpoints []Endpoint `json:"endpoints"`
}

type RequestBody struct {
	Query string `json:"query"`
}

var BaseDir string = "."

var Api API

// CreateApi this creates the API object - run this before accessing config.Api
func CreateApi() {
	raw, ioerr := ioutil.ReadFile(BaseDir + "/api.json")
	if ioerr != nil {
		fmt.Println(ioerr.Error())
		os.Exit(1)
	}

	err := json.Unmarshal(raw, &Api)
	if err != nil {
		log.Fatal(" ", err)
	}
}
