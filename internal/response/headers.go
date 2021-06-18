package response

import "net/http"

func SetupHeaders(w *http.ResponseWriter, req *http.Request) {
	if acrh, ok := req.Header["Access-Control-Request-Headers"]; ok {
		(*w).Header().Set("Access-Control-Allow-Headers", acrh[0])
	}
	(*w).Header().Set("Access-Control-Allow-Credentials", "True")
	if acao, ok := req.Header["Access-Control-Allow-Origin"]; ok {
		(*w).Header().Set("Access-Control-Allow-Origin", acao[0])
	} else {
		if _, oko := req.Header["Origin"]; oko {
			(*w).Header().Set("Access-Control-Allow-Origin", req.Header["Origin"][0])
		} else {
			(*w).Header().Set("Access-Control-Allow-Origin", "*")
		}
	}
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	(*w).Header().Set("Connection", "Close")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Messages, Suggestions, Conversation-Token")
	(*w).Header().Set("Content-Type", "application/json")
}
