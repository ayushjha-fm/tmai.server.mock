package response

import (
	"encoding/json"
	"net/http"

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
			b := path2Response(ep.Path, request)
			w.Write(b)
		}
		continue
	}
}

func path2Response(path string, request request.Request) []byte {
	response := MessageResponseType{}
	response.Query = request.Body.Query
	response.Data = messages.GetMessages(request)
	response.Suggestions = suggestions.GetSuggestions(request)
	js_out, _ := json.Marshal(response)
	return js_out
}
