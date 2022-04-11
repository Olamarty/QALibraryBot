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
	var links []buttonLink
	var text string

	for update := range updates {
		if update.Message == nil {
			continue
		}
		switch update.Message.Text {
		case "/qa_basic":
			links = qaBasicLinks
			text = "QA ОСНОВЫ"
		case "/qa_manager":
			links = qaManagerLinks
			text = "УПРАВЛЕНИЕ"
		case "/qa_automatic":
			links = qaAutomaticLinks
			text = "АВТОМАТИЗАЦИЯ"
		case "/protocols_helper":
			links = protocolLinks
			text = "SSH, SCP, HTTP и пр."
		case "/git_helper":
			links = gitLinks
			text = "GITHUB"
		case "/go_basic":
			links = goBasicLinks
			text = "GOLANG"
		case "/go_tgbot":
			links = goTgBotLinks
			text = "ПИШЕМ БОТА НА GO"
		case "/start":
			links = nil
			text = "Добро пожаловать в QA-Library! Выбери в меню, что будем изучать⬇"
		default:
			links = nil
			text = "Я не знаком с этой командой. выбери что-то из меню⬇"
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
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
