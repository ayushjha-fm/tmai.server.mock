package response

import (
	"encoding/json"
	"net/http"

	lorem "github.com/drhodes/golorem"
	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/request"
)

type AutocompleteType struct {
	Suggestions []string `json:"suggestions"`
	Query       string   `json:"query"`
}

func AutocompleteResponse(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	request, _ := request.GetRequestObj(*r)
	SetupHeaders(&w, r)
	if r.Method == "OPTIONS" {
		return
	}
	conversation_header := w.Header().Get(config.HeaderConversationToken)
	for _, ep := range config.Api.Endpoints {
		// check if method matches
		methodMatches := false
		for _, m := range ep.Methods {
			if m == r.Method {
				methodMatches = true
				break
			}
		}
		if r.URL.Path == ep.Path && methodMatches {
			w.Header().Set(config.HeaderContentType, config.MIMEApplicationJSON)
			if conversation_header == "" {
				w.Header().Set(config.HeaderConversationToken, config.CONV_TOKEN)
			} else {
				w.Header().Set(config.HeaderConversationToken, conversation_header)
			}
			w.WriteHeader(ep.Status)
			b := getAutocomplete(ep.Path, request)
			w.Write(b)
		}
		continue
	}
}

func getAutocomplete(path string, request request.Request) []byte {
	response := AutocompleteType{}
	query := request.Body.Query
	if query == "" {
		return nil
	}
	response.Query = query
	for i := 0; i < 5; i++ {
		response.Suggestions = append(response.Suggestions, lorem.Sentence(3, 6))
	}
	js_out, _ := json.Marshal(response)
	return js_out
}
