package response

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

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

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept, Authorization, Conversation-Token")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Messages, Suggestions")
	csp := []string{"default-src: 'self'", "font-src: 'fonts.googleapis.com'", "frame-src: 'none'"}
	w.Header().Set("Content-Security-Policy", strings.Join(csp, "; "))

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
