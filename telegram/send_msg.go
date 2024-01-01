package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Telegram) SendMsg(msg string, chatId int64) error {
	tgMessage := tgbotapi.NewMessage(chatId, msg)

	_, err := c.bot.Send(tgMessage)
	if err != nil {
		return err
	}

	return nil
}
