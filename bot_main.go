package main

import (
	"encoding/json"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var emoji string = "◻"
var btnData string = "process"

var cmdMessages = map[string]cmdData{
	"/qa_basic": {
		links: qaBasicLinks,
		text:  "QA основы",
	},
	"/qa_manager": {
		links: qaManagerLinks,
		text:  "Управление",
	},
	"/qa_automatic": {
		links: qaAutomaticLinks,
		text:  "Автоматизация",
	},
	"/protocols_helper": {
		links: protocolLinks,
		text:  "О протоколах",
	},
	"/api_testing": {
		links: apiLinks,
		text:  "API testing",
	},
	"/git_helper": {
		links: gitLinks,
		text:  "GIThub",
	},
	"/go_basic": {
		links: goBasicLinks,
		text:  "GOLANG",
	},
	"/go_tgbot": {
		links: goTgBotLinks,
		text:  "Пишем бота на GO",
	},
	"/coding": {
		links: codingLinks,
		text:  "Чистый код",
	},
	"/start": {
		links: nil,
		text:  "Добро пожаловать в QA-Library! Выбери в меню, что будем изучать⬇",
	},
}

type cmdData struct {
	links []buttonLink
	text  string
}

type Config struct {
	TelegramBotToken string
}

type buttonLink struct {
	Name string
	Link string
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

	for update := range updates {

		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "process" {
				emoji = "▶"
				btnData = "done"
			} else if update.CallbackQuery.Data == "done" {
				emoji = "✔"
				btnData = "empty"
			} else {
				emoji = "◻"
				btnData = "process"
			}
			continue
		}

		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

		var msg tgbotapi.MessageConfig

		if data, ok := cmdMessages[update.Message.Text]; ok {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, data.text)
			if data.links != nil {
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				msg.ReplyMarkup = newKeyboard(data.links)
			}
		} else {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаком с этой командой. выбери что-то из меню⬇")
		}

		bot.Send(msg)
	}
}

func newKeyboard(btns []buttonLink) tgbotapi.InlineKeyboardMarkup {
	var out [][]tgbotapi.InlineKeyboardButton

	for _, btn := range btns {
		out = append(out, []tgbotapi.InlineKeyboardButton{
			tgbotapi.NewInlineKeyboardButtonURL(btn.Name, btn.Link),
			tgbotapi.NewInlineKeyboardButtonData(emoji, btnData),
		})
	}

	kbrd := tgbotapi.NewInlineKeyboardMarkup(out...)

	return kbrd
}
