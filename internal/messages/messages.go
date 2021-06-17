package messages

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	lorem "github.com/drhodes/golorem"
	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/request"
)

const (
	ArticleMessage          string = "article"
	TextMessage             string = "text"
	InfographicMessage      string = "infographic"
	SuggestionTopicsMessage string = "suggested-topics"
)

type MessageType struct {
	Type            string          `json:"type"`
	Meta            json.RawMessage `json:"meta"`
	Title           json.RawMessage `json:"title,omitempty"`
	Content         json.RawMessage `json:"content,omitempty"`
	PublicationDate json.RawMessage `json:"pub_date,omitempty"`
	SourceUrl       json.RawMessage `json:"source_url,omitempty"`
	Author          json.RawMessage `json:"author,omitempty"`
	OtherQuestions  json.RawMessage `json:"other_questions,omitempty"`
	Link            json.RawMessage `json:"link,omitempty"`
	Description     json.RawMessage `json:"description,omitempty"`
	Slice           json.RawMessage `json:"slice,omitempty"`
}

// GetMessages returns messages as byte array
func GetMessages(req request.Request) []byte {
	var messages []MessageType
	for _, messageType := range req.MessagesType {
		message := getMessage(messageType)
		messages = append(messages, message)
	}
	raw, err := json.Marshal(messages)
	if err != nil {
		log.Println(err)
	}
	return raw
}

// getMessage returns a message
func getMessage(messageType string) (message MessageType) {
	if messageType == "" {
		return message
	}
	messageType = strings.TrimSpace(messageType)
	switch messageType {
	case ArticleMessage:
		message = getArticleMessage()
	case TextMessage:
		message = getTextMessage()
	case SuggestionTopicsMessage:
		message = getSuggestedTopicsMessage()
	default:
		raw, ioerr := ioutil.ReadFile(config.BaseDir + "/messages/" + messageType + ".json")
		if ioerr != nil {
			fmt.Println(ioerr.Error())
			return message
		}
		json.Unmarshal(raw, &message)
	}
	return message
}

// getTextMessage returns a text message
func getTextMessage() (m MessageType) {
	raw, ioerr := ioutil.ReadFile(config.BaseDir + "/messages/text.json")
	if ioerr != nil {
		fmt.Println(ioerr.Error())
		return m
	}
	m.Type = "text"
	json.Unmarshal(raw, &m)
	m.Content = []byte(`"` + lorem.Sentence(15, 30) + `"`)
	return m
}

// getTextMessage returns a text message
func getArticleMessage() (m MessageType) {
	raw, ioerr := ioutil.ReadFile(config.BaseDir + "/messages/text.json")
	if ioerr != nil {
		fmt.Println(ioerr.Error())
		return m
	}
	json.Unmarshal(raw, &m)
	m.Type = "article"
	m.Link = []byte("\"https://times.com/articlelink\"")
	m.Title = []byte(`"` + lorem.Sentence(5, 15) + `"`)
	m.Content = []byte(`"` + lorem.Paragraph(5, 5) + `"`)
	m.Author = []byte(getAuthor())
	m.SourceUrl = []byte(`"` + "times.com" + `"`)
	m.PublicationDate = []byte(`"2020-12-12"`)
	return m
}

func getAuthor() string {
	return fmt.Sprintf(
		`{"name": "%s", "twitter_username": "%s"}`,
		lorem.Sentence(2, 4),
		lorem.Sentence(1, 1))
}

// getTextMessage returns a text message
func getSuggestedTopicsMessage() (m MessageType) {
	raw, ioerr := ioutil.ReadFile(config.BaseDir + "/messages/text.json")
	if ioerr != nil {
		fmt.Println(ioerr.Error())
		return m
	}
	json.Unmarshal(raw, &m)
	m.Type = "infographic"
	m.Content = []byte(fmt.Sprintf(
		`[
        {
            "text": "%s",
            "icon": "%s"
        },
        {
            "text": "%s",
            "icon": "%s"
        },
        {
            "text": "%s",
            "icon": "%s"
        }
    ]`,
		lorem.Sentence(2, 4),
		lorem.Sentence(2, 4), // image blob
		lorem.Sentence(2, 4),
		lorem.Sentence(2, 4), // image blob
		lorem.Sentence(2, 4),
		lorem.Sentence(2, 4), // image blob
	))
	return m
}
