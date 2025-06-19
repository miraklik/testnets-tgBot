package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tg-bot/internal/config"
	"tg-bot/internal/db"
	"tg-bot/internal/handlers"
	"tg-bot/internal/models"
	"tg-bot/internal/telegram"
	"time"
)

func handlerUpdates() {
	offset := 0

	for {
		url := fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", telegram.BotToken, offset)
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("Failed to get updates: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		defer resp.Body.Close()
		if err := json.NewDecoder(resp.Body).Decode(&models.UpdatesResponse); err != nil {
			log.Printf("Failed to decode updates: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		for _, update := range models.UpdatesResponse.Result {
			offset = int(update.UpdateID) + 1

			if update.Message != nil {
				chatID := update.Message.Chat.ID
				if update.Message.Text == "/start" {
					handlers.HandlerStartMessage(chatID)
				}
			} else if update.CallbackQuery != nil {
				callbackID := update.CallbackQuery.ID
				data := update.CallbackQuery.Data
				chatID := update.CallbackQuery.From.ID
				handlers.HandleCallbackQuery(callbackID, data, chatID)
			}
		}
	}
}

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	telegram.InitBotToken(cfg.Tokens.Token)
	db.ConnectMongoDB(cfg.MongoDB.URI)

	go handlerUpdates()
}
