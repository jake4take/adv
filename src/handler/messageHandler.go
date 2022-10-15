package handler

import (
	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"v1/src/controller"
)

func MessageCatcher(telegramApiToken string) {
	bot, err := tg.NewBotAPI(telegramApiToken)
	if err != nil {
		panic(err)
	}
	bot.Debug = false
	updateConfig := tg.NewUpdate(0)

	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		go messageHandler(bot, update)
	}
}

func messageHandler(bot *tg.BotAPI, update tg.Update) {

	if update.CallbackQuery != nil {

	}

	if update.Message != nil {

	}

	controller.SendMessage(msg, bot)

}
