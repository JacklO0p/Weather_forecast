package telegram

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TelegramMessage struct {
	CHAT_ID    string `json:"chat_id"`
	Text       string `json:"text"`
	PARSE_MODE string `json:"parse_mode"`
}

func SendTelegramMessage(message, chatId string) {
	token := "bot" + TOKEN
	sendMessageUrl := "https://api.telegram.org/" + token + "/sendMessage"
	message = strings.ReplaceAll(message, "-", "\\-")
	message = strings.ReplaceAll(message, ".", "\\.")
	message = strings.ReplaceAll(message, "_", "\\_")
	message = strings.ReplaceAll(message, "{", "\\{")
	message = strings.ReplaceAll(message, "}", "\\}")
	msg := new(TelegramMessage)
	msg.PARSE_MODE = "MarkdownV2"
	msg.Text = message
	msg.CHAT_ID = chatId

	jBytes, err := json.Marshal(&msg)

	if err != nil {
		log.Println(err)
		return
	}

	request, err := http.NewRequest("POST", sendMessageUrl, bytes.NewBuffer(jBytes))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	} else {
		defer response.Body.Close()
	}

	log.Println("response Status:", response.Status)
	log.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	log.Println("response Body:", string(body))

}
