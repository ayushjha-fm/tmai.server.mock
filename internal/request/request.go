package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	PreDefinedMessage  string = "predefined"
	UserDefinedMessage string = "userdefined"
)

type RequestBody struct {
	Query string `json:"query"`
	Id    string `json:"id"`
}

type Request struct {
	MessagesType    []string
	SuggestionsType []string
	HttpRequest     http.Request
	Body            RequestBody
}

// GetRequestObj converts http Request to Request
func GetRequestObj(hr http.Request) (r Request, err error) {
	defer hr.Body.Close()
	body, err := ioutil.ReadAll(hr.Body)
	if err != nil {
		return r, err
	}
	json.Unmarshal(body, &r.Body)

	suggestionsHeader := hr.Header.Get("Suggestions")
	r.SuggestionsType = strings.Split(suggestionsHeader, ",")

	messagesHeader := hr.Header.Get("Messages")

	r.MessagesType = strings.Split(messagesHeader, ",")
	r.HttpRequest = hr
	return r, nil
}
