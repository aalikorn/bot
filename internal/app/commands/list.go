package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	msgText := "All the products: \n\n"
	products := c.productService.List()
	for _, p := range products {
		msgText += p.Title + "\n"
	}
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, msgText)
	c.bot.Send(msg)
}
