package notifier

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
	"match_statistics_scrapper/db"
	"match_statistics_scrapper/db/repos"
	"match_statistics_scrapper/models"
	"strconv"
	"strings"
)

func TgBot(stringified string) error {
	ss := repos.SubscriberStore{DB: db.ConnectDB()}

	token := viper.GetString("telegram.token")
	chatIDList := viper.GetIntSlice("telegram.chatIDList")
	fmt.Println(chatIDList)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Token Problem")
		log.Panic(err)
		return err
	}

	// Update Config From TgBOT
	updateConfig := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(updateConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Process updates
	for update := range updates {
		if update.Message != nil {
			// Get the message text and chat ID
			messageText := update.Message.Text
			chatID := update.Message.Chat.ID
			fmt.Println(chatID)

			// Check if the message is "/start"
			if messageText == "/start" {
				// Send a welcome message
				message := tgbotapi.NewMessage(chatID, "Hello! ")
				_, err := bot.Send(message)
				if err != nil {
					log.Println(err)
				}
			} else if messageText == "/subscribe" {
				message := tgbotapi.NewMessage(chatID, "Admin will approve your request! Thanks! :D ")
				_, err := bot.Send(message)
				if err != nil {
					log.Println("Subscription msg rply error", err)
				}

				sub := &models.Subscriber{
					ChatID:   chatID,
					Approved: false,
				}
				err = ss.Save(sub)
				if err != nil {
					log.Println("Cannot Save Subscriber Info: ", err)
				}

				adminId := viper.GetInt64("telegram.adminID")
				message = tgbotapi.NewMessage(adminId,
					fmt.Sprintf("%v wants to join this channel ! To approve type `/approve %v`", chatID, chatID))
				_, err = bot.Send(message)
				if err != nil {
					log.Println("Admin approve request msg: ", err)
				}
			} else if strings.HasPrefix(messageText, "/approve ") {
				parts := strings.Fields(messageText)
				if len(parts) != 2 {
					log.Println("Invalid /approve command format")
					continue
				}

				chatIDToApprove := parts[1]
				chatID, _ := strconv.Atoi(chatIDToApprove)

				res, err := ss.GetSubscriberData(int64(chatID))
				if err != nil {
					log.Println("Could Find ChatID : ", err)
					continue
				}

				approvalMessage := tgbotapi.NewMessage(int64(chatID), "Your request has been approved!")
				_, err = bot.Send(approvalMessage)
				if err != nil {
					log.Println("Approval confirmation message error:", err)
					continue
				}

				adminId := viper.GetInt64("telegram.adminID")
				approvalMessage = tgbotapi.NewMessage(adminId, fmt.Sprintf("ID %s has been approved!", chatIDToApprove))
				_, err = bot.Send(approvalMessage)
				if err != nil {
					log.Println("Approval confirmation message error:", err)
					continue
				}

				res.Approved = true
				err = ss.Save(res)
				if err != nil {
					log.Println("Cannot Save Subscriber Info: ", err)
				}
			}
		}
	}
	//for _, chatID := range chatIDList {
	//	message := tgbotapi.NewMessage(int64(chatID), stringified)
	//	_, err = bot.Send(message)
	//	if err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//}
	return nil
}
