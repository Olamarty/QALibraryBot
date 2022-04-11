package main

import (
	"encoding/json"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Config struct {
	TelegramBotToken string
}

func main() {

	file, _ := os.Open("/root/config.json")
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 15
	updates, _ := bot.GetUpdatesChan(u)
	links := []buttonLink{}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/qa_basic":
			links = qaBasicLinks
		case "/qa_manager":
			links = qaManagerLinks
		case "/qa_automatic":
			links = qaAutomaticLinks
		case "/protocols_helper":
			links = protocolLinks
		case "/git_helper":
			links = gitLinks
		case "/go_basic":
			links = goBasicLinks
		case "/go_tgbot":
			links = goTgBotLinks
		case "/start":
			links = nil
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать в QA-Library! Выбери в меню, что будем изучать⬇"))
		default:
			links = nil
			bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаком с этой командой. выбери что-нибудь из меню⬇"))
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "вот что у меня есть:")
		msg.ReplyMarkup = newKeyboard(links)

		bot.Send(msg)
	}
}

type buttonLink struct {
	Name string
	Link string
}

func newKeyboard(btns []buttonLink) tgbotapi.InlineKeyboardMarkup {
	var out [][]tgbotapi.InlineKeyboardButton
	for _, btn := range btns {
		out = append(out, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonURL(btn.Name, btn.Link)})
	}

	kbrd := tgbotapi.NewInlineKeyboardMarkup(out...)

	return kbrd
}
