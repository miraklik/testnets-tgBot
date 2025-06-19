package handlers

import (
	"context"
	"log"
	"tg-bot/internal/config"
	"tg-bot/internal/db"
	"tg-bot/internal/telegram"
)

func handleCategorySelection(chatID int64, categoryName string) {
	subcategories, err := db.GetSubTestnets(context.Background(), categoryName)
	if err != nil {
		log.Printf("Failed to get subcategories: %v", err)
		return
	}

	inlineButtons := createInlineButtons(subcategories, "subtestnet_"+categoryName+"_")
	inlineButtons = append(inlineButtons, []map[string]interface{}{
		{
			"text":          "⬅️ Back",
			"callback_data": "back_to_categories",
		},
	})

	replyMarkup := map[string]interface{}{
		"inline_keyboard": inlineButtons,
	}
	if err := telegram.SendMessage(chatID, config.SubcatMessage, replyMarkup); err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}

func handleBackToCategories(chatID int64) {
	HandlerStartMessage(chatID)
}
