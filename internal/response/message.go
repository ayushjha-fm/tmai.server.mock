package response

import (
	"encoding/json"
	"net/http"
	"strings"

	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/logger"
	"tmai.server.mock/internal/messages"
	"tmai.server.mock/internal/request"
	"tmai.server.mock/internal/suggestions"
)

type MessageResponseType struct {
	Query       string          `json:"query"`
	Meta        json.RawMessage `json:"meta"`
	Data        json.RawMessage `json:"data"`
	Suggestions json.RawMessage `json:"suggestions"`
}

func MessageResponse(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	appLogger := logger.CreateLogger()
	request, _ := request.GetRequestObj(*r)

	appLogger.AccessLog(r)
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
			b := getMessageResponse(ep.Path, request)
			w.Write(b)
		}
		continue
	}
}

func getMessageResponse(path string, request request.Request) []byte {
	response := MessageResponseType{}
	response.Query = request.Body.Query
	query_parts := strings.Split(response.Query, ";")
	if len(request.Body.Query) > 0 {
		response.Data = messages.GetMessages(strings.Split(query_parts[0], ","))
	} else {
		response.Data = messages.GetMessages(request.MessagesType)
	}
	if len(query_parts) > 1 {
		response.Suggestions = suggestions.GetSuggestions(strings.Split(query_parts[1], ","))
	} else {
		response.Suggestions = suggestions.GetSuggestions(request.SuggestionsType)
	}

	js_out, _ := json.Marshal(response)
	return js_out
}
