package response

import (
	"io/ioutil"
	"log"
	"net/http"

	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/logger"
)

func GetTrivia() ([]byte, error) {
	resp, getErr := http.Get("https://catfact.ninja/fact")
	if getErr != nil {
		return nil, getErr
	}
	defer resp.Body.Close()

	content, ioErr := ioutil.ReadAll(resp.Body)
	if ioErr != nil {
		return nil, ioErr
	}
	return content, nil
}

func TriviaResponse(w http.ResponseWriter, r *http.Request) {
	appLogger := logger.CreateLogger()
	appLogger.AccessLog(r)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept, Authorization, Conversation-Token")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
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
