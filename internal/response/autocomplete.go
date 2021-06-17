package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	lorem "github.com/drhodes/golorem"
	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/logger"
	"tmai.server.mock/internal/request"
)

type AutocompleteType struct {
	Suggestions []string `json:"suggestions"`
	Query       string   `json:"query"`
}

func AutocompleteResponse(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	appLogger := logger.CreateLogger()
	request, _ := request.GetRequestObj(*r)
	appLogger.AccessLog(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers",
		"Origin, X-Requested-With, Content-Type, Accept, Authorization, Conversation-Token")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Messages, Suggestions")
	csp := []string{"default-src: 'self'", "font-src: 'fonts.googleapis.com'", "frame-src: 'none'"}
	w.Header().Set("Content-Security-Policy", strings.Join(csp, "; "))
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
	fmt.Println(response.Suggestions)
	js_out, _ := json.Marshal(response)
	return js_out
}