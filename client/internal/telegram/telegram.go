package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var BotToken string

func InitBotToken(token string) {
	BotToken = token
}

func SendMessage(chatID int64, text string, replyMerkly any) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", BotToken)
	body := map[string]any{
		"chat_id":      chatID,
		"text":         text,
		"reply_markup": replyMerkly,
		"parse_mode":   "HTML",
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		log.Printf("Failed to encode body: %v", err)
		return err
	}

	_, err := http.Post(url, "application/json", buf)
	if err != nil {
		log.Printf("Failed to send message: %v", err)
		return err
	}

	return nil
}

func AnswerCallbackQuery(callbackQueryID string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/answerCallbackQuery", BotToken)
	body := map[string]any{
		"callback_query_id": callbackQueryID,
	}
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		log.Printf("Failed to encode body: %v", err)
		return err
	}
	_, err := http.Post(url, "application/json", buf)
	if err != nil {
		log.Printf("Failed to answer callback query: %v", err)
		return err
	}

	return nil
}
