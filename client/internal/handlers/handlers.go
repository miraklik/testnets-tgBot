package handlers

import (
	"strings"
	"tg-bot/internal/telegram"
)

func HandlerCallBackQuery(callbackID, data string, chatID int64) {
	telegram.AnswerCallbackQuery(callbackID)
}

func createInlineButtons(testnets []string, prefix string) [][]map[string]interface{} {
	var inlineButtons [][]map[string]interface{}
	for _, testnet := range testnets {
		inlineButtons = append(inlineButtons, []map[string]interface{}{
			{
				"text":          testnet,
				"callback_data": prefix + testnet,
			},
		})
	}
	return inlineButtons
}

func HandleCallbackQuery(callbackID, data string, chatID int64) {
	switch {
	case data == "back_to_testnets":
		handleBackToCategories(chatID)
	case strings.HasPrefix(data, "testnet_"):
		categoryName := strings.TrimPrefix(data, "testnet_")
		handleCategorySelection(chatID, categoryName)
	}
	telegram.AnswerCallbackQuery(callbackID)
}
