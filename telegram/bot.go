package telegram

import (
	"log"
	"strings"

	"../config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func StartBot() {
	bot, err := tgbotapi.NewBotAPI(config.TelegramAPI)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}
		if !strings.HasPrefix(update.Message.Text, "!") {
			continue
		}
		// Declaring reply_msg Text.
		var replyMsg string = ""

		

		reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyMsg)

		bot.Send(reply)
	}
}
