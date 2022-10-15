package adv

import (
	"os"
	"v1/src/handler"
)

func main() {

	telegramApiToken, _ := os.LookupEnv("TELEGRAM_APITOKEN")

	handler.MessageCatcher(telegramApiToken)

}
