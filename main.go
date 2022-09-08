package main

import (
	"fmt"
	"net/url"
	"os"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := telegram.NewBotAPI(os.Getenv("TOKEN"))
	e(1, err)
	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message != nil {
			if update.Message.Chat.ID != 1765800269 {
				continue
			}
			go messageHandler(bot, update)
		}
	}
}

func messageHandler(bot *telegram.BotAPI, update telegram.Update) {
	if len(update.Message.Text) < 60 {
		u, err := url.ParseRequestURI(update.Message.Text)
		delete := false
		if err != nil {
			delete = true
		} else if u.Host != "github.com" && u.Host != "www.github.com" {
			delete = true
		}
		if delete {
			del := telegram.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			bot.Send(del)
		}
	} else {
		del := telegram.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
		bot.Send(del)
	}

}

func e(id int, err error) {
	if err != nil {
		fmt.Println(id, ":", err)
	}
}
