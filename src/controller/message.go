package controller

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DeleteMessage(message tg.DeleteMessageConfig, bot *tg.BotAPI) {
	if _, err := bot.Request(message); err != nil {
		fmt.Print("Error delete msg: " + err.Error())
	}
}

func SendMessage(message tg.MessageConfig, bot *tg.BotAPI) tg.Message {
	resp, err := bot.Send(message)
	if err != nil {
		fmt.Printf("Error send msg: %s", err.Error())
	}
	return resp
}

func UpdateMessage(updateMsg tg.EditMessageTextConfig, updButton tg.EditMessageReplyMarkupConfig, bot *tg.BotAPI) {
	if _, err := bot.Send(updateMsg); err != nil {
		fmt.Println(fmt.Sprintf("Error update msg: %s, %d", err.Error(), updateMsg.MessageID))
	}
	if _, err := bot.Send(updButton); err != nil {
		fmt.Println(fmt.Sprintf("Error update button: %s, %d", err.Error(), updateMsg.MessageID))
	}
}
