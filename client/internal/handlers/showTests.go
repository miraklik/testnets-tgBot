package handlers

import (
	"context"
	"log"
	"tg-bot/internal/config"
	"tg-bot/internal/db"
	"tg-bot/internal/telegram"
)

func HandlerStartMessage(chatID int64) {
	testnets, err := db.GetTestnets(context.Background())
	if err != nil {
		log.Fatalf("Failed to get testnets: %v", err)
		return
	}

	inlineButtons := createInlineButtons(testnets, "testnets_")
	replyMerkly := map[string]any{
		"inline_keybord": inlineButtons,
	}

	if err := telegram.SendMessage(chatID, config.HelloMessage, replyMerkly); err != nil {
		log.Fatalf("Failed to send message: %v", err)
		return
	}
}
