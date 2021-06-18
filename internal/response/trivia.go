package response

import (
	"encoding/json"
	"log"
	"net/http"

	lorem "github.com/drhodes/golorem"
	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/logger"
)

type TriviaResponseType struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func GetTrivia() ([]byte, error) {
	resp := TriviaResponseType{}
	resp.Fact = lorem.Sentence(4, 5)
	resp.Length = len(resp.Fact)
	data, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func TriviaResponse(w http.ResponseWriter, r *http.Request) {
	appLogger := logger.CreateLogger()
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
			b, err := GetTrivia()
			if err != nil {
				log.Println(err)
			}
			w.Write(b)
		}
		continue
	}
}
