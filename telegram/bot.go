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

	bot.Debug = false

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

		// Getting command from Message.
		var command string = strings.Split(update.Message.Text, " ")[0]
		var text []string = strings.Split(update.Message.Text, " ")
		// Deletes Command from Text.
		text = append(text[:0], text[0+1:]...)
		msg := strings.Join(text, " ")

		// Command Handling
		switch command {
		case "!add":
			updatePackage := Package{update.Message.From.UserName,
				msg,
				"/"}
			AddToDo(updatePackage)
		}

		// Sending Reply.
		reply := tgbotapi.NewMessage(update.Message.Chat.ID, replyMsg)
		bot.Send(reply)
	}
}

type Package struct {
	userID string
	text   string
	todoID string
}
