package response

import (
	"encoding/json"
	"log"
	"net/http"

	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/logger"
	"tmai.server.mock/internal/pdmh"
	"tmai.server.mock/internal/request"
)

func PDMHResponse(w http.ResponseWriter, r *http.Request) {
	appLogger := logger.CreateLogger()
	appLogger.AccessLog(r)
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
			b, err := getGamificationMessage(request.Body.Intent, request)
			if err != nil {
				log.Println(err)
			}
			w.Write(b)
		}
		continue
	}
}

func getGamificationMessage(intent string, request request.Request) ([]byte, error) {
	response := MessageResponseType{}
	response.Query = request.Body.Query
	if request.Body.Intent == "" {
		response.Meta = []byte(`"userdefined"`)
	} else {
		response.Meta = []byte(`"predefined"`)
	}

	response.Data = pdmh.GetGamificationMessages(intent)
	js_out, err := json.Marshal(response)
	return js_out, err
}
