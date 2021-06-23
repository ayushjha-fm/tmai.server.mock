package pdmh

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"tmai.server.mock/internal/config"
	"tmai.server.mock/internal/messages"
)

// pdmh - pre-defined message handler

const (
	ExploreStabbyIntent     string = "EXPLORE_STABBY"
	ExploreQuestions        string = "EXPLORE_QUESTIONS"
	AroundWorldIntent       string = "AROUND_WORLD"
	VaccineProgressIntent   string = "PROGRESS_VACCINES"
	SideEffectsIntent       string = "SIDE_EFFECTS"
	TravelDuringCovidIntent string = "TRAVEL_DURING_COVID"
	VaccinationRateIntent   string = "VACCINATION_RATE"
	VaccinePassportIntent   string = "VACCINE_PASSPORT"
)

type GamificationType []messages.MessageType

// GetMessages returns messages as byte array
func GetGamificationMessages(intent string) []byte {
	gamificationMessage := getGamificationMessage(intent)
	raw, err := json.Marshal(gamificationMessage)
	if err != nil {
		log.Println(err)
	}
	return raw
}

// getMessage returns a message
func getGamificationMessage(intent string) (message GamificationType) {
	raw, ioerr := ioutil.ReadFile(config.BaseDir + "/gamification/" + intent + ".json")
	if ioerr != nil {
		fmt.Println(ioerr.Error())
		return message
	}
	json.Unmarshal(raw, &message)

	return message
}
