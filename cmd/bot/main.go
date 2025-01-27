package main

import (
	"log"
	"os"

	"github.com/aalikorn/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	productService := product.NewService()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehavior(bot, update.Message)
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n/list - list")
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, productService *product.Service) {
	msgText := "All the products: \n\n"
	products := productService.List()
	for _, p := range products {
		msgText += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)
	// msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
