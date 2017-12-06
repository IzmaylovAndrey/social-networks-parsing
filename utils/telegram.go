package utils

import (
	"github.com/IzmaylovAndrey/social-networks-parsing/models"
	"bytes"
	"strings"
	"encoding/json"
	"net/http"
	"fmt"
)

func SendToTelegram (login string, accounts []models.Accounts) error {
	var rows []string
	rows = append(rows, "New account " + login + " was created.")
	for _, value := range accounts {
		var buffer bytes.Buffer
		buffer.WriteString(value.SocialNetwork + ": " + value.Data)
		rows = append(rows, buffer.String())
	}
	str := strings.Join(rows[:],"\n")

	message := map[string]string{"chat_id": "@go_social", "text": str}
	jsonValue, _ := json.Marshal(message)
	if _, err := http.Post("https://api.telegram.org/bot313653847:AAFAMN20ebJXvctcdNw99tLEMweVGsVTEKY/sendMessage", "application/json", bytes.NewBuffer(jsonValue)); err != nil{
		fmt.Printf("Error sending message to Telegram: %s", err)
		return err
	}
	return nil
}