package suggestions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"tmai.server.mock/internal/config"
)

type Suggestion struct {
	Type      string            `json:"type,omitempty"`
	Questions []config.Question `json:"questions,omitempty"`
}

const (
	ExploreStabbySuggestion    string = "Explore-Stabby"
	ExploreQuestionsSuggestion string = "Explore-Questions"
	SuggestQuestionsSuggestion string = "Suggest-Questions"
	NoSuggestion               string = "No-Suggest"
)

// GetSuggestions returns suggestions as byte array
func GetSuggestions(sugegstionsType []string) []byte {
	var suggestions []Suggestion
	for _, suggestionType := range sugegstionsType {
		suggestion := getSuggestion(suggestionType)
		if suggestion.Type != "" {
			suggestions = append(suggestions, suggestion)
		}
	}
	raw, _ := json.Marshal(suggestions)
	return raw
}

// getSuggestion returns a suggestion
func getSuggestion(suggestionType string) Suggestion {
	var suggestion Suggestion
	if suggestionType == "" {
		return suggestion
	}
	suggestionType = strings.TrimSpace(suggestionType)
	raw, ioerr := ioutil.ReadFile(config.BaseDir + "/suggestions/" + suggestionType + ".json")
	if ioerr != nil {
		fmt.Println(ioerr.Error())
		return suggestion
	}
	json.Unmarshal(raw, &suggestion)
	return suggestion
}
