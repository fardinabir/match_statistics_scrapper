package notifier

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
)

func TgBot(stringified string) error {
	token := viper.GetString("telegram.token")
	chatIDList := viper.GetIntSlice("telegram.chatIDList")
	fmt.Println(chatIDList)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Token Problem")
		log.Panic(err)
		return err
	}

	for _, chatID := range chatIDList {
		message := tgbotapi.NewMessage(int64(chatID), stringified)
		_, err = bot.Send(message)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
