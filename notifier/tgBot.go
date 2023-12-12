package notifier

import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
)

func TgBot() {
	token := viper.GetString("telegram.token")
	chatId := viper.GetInt64("telegram.chatId")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Token Problem")
		log.Panic(err)
	}
	message := tgbotapi.NewMessage(chatId, "Hello There ! ")
	_, err = bot.Send(message)
	if err != nil {
		log.Println(err)
	}
}
