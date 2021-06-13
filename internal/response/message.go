package response

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/logger"
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

	appLogger.AccessLog(r)
	var requestBody config.RequestBody
	body, ioerr := ioutil.ReadAll(r.Body)
	if ioerr != nil {
		log.Fatal(ioerr)
	}
	json.Unmarshal(body, &requestBody)
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
			b := path2Response(ep.Path, requestBody.Query)
			w.Write(b)
		}
		continue
	}
}

func path2Response(path string, query string) []byte {
	file, err := os.Open(config.BaseDir + path + ".json")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	defer file.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	bytes := buf.Bytes()
	if query != "" {
		response := MessageResponseType{}
		jsonErr := json.Unmarshal(buf.Bytes(), &response)
		if jsonErr != nil {
			return bytes
		}
		response.Query = query
		js_out, jsonMarshalError := json.Marshal(response)
		if jsonMarshalError != nil {
			return bytes
		}
		return js_out
	}
	return bytes
}
