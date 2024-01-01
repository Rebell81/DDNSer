package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	bot *tgbotapi.BotAPI
}

func InitBot(token string) (*Telegram, error) {
	fmt.Println(fmt.Sprintf("try init bot '%s'", token))
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	log.Printf("telegram bot authorized as '%s'", bot.Self.UserName)

	return &Telegram{bot}, nil
}
