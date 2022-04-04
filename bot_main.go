package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5039188251:AAFOCrkfYHeBrT9ShBtXOUnkoZikRLIZonk")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 15

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/start":
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в QA-Library! Выбери в меню, что будем изучать ⬇"))
		case "/qa_basic":
			qaBasicMessage := tgbotapi.NewMessage(update.Message.Chat.ID, qaBasicLinks)
			qaBasicMessage.ParseMode = "HTML"
			bot.Send(qaBasicMessage)
		case "/qa_manager":
			qaManagerMessage := tgbotapi.NewMessage(update.Message.Chat.ID, qaManagerLinks)
			qaManagerMessage.ParseMode = "HTML"
			bot.Send(qaManagerMessage)
		case "/qa_automatic":
			qaAutomaticMessage := tgbotapi.NewMessage(update.Message.Chat.ID, qaAutomaticLinks)
			qaAutomaticMessage.ParseMode = "HTML"
			bot.Send(qaAutomaticMessage)
		case "/go_basic":
			goBasicMessage := tgbotapi.NewMessage(update.Message.Chat.ID, goBasicLinks)
			goBasicMessage.ParseMode = "HTML"
			bot.Send(goBasicMessage)
		default:
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаком с этой командой. выбери что-нибудь из меню!"))
		}
	}
}
