package notifier

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"log"
	"reflect"
)

func TgBot(data interface{}) error {
	token := viper.GetString("telegram.token")
	chatIDList := viper.GetIntSlice("telegram.chatIDList")
	fmt.Println(chatIDList)

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println("Token Problem")
		log.Panic(err)
		return err
	}

	stringified := stringifyStruct(data)
	for _, chatID := range chatIDList {

		message := tgbotapi.NewMessage(int64(chatID), stringified)
		_, err = bot.Send(message)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	//hashed := getHashOfData(stringified)
	// TODO: insert on database
	return nil
}

func stringifyArrayOfStructs(data interface{}) string {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Slice {
		return "Not a slice"
	}

	var str string
	for i := 0; i < val.Len(); i++ {
		structElement := val.Index(i)
		str += stringifyStruct(structElement.Interface()) + "\n"
	}

	return str
}

func stringifyStruct(data interface{}) string {
	val := reflect.ValueOf(data)
	typ := reflect.TypeOf(data)

	if val.Kind() != reflect.Struct {
		return "Not a struct"
	}

	var str string
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)

		str += fmt.Sprintf("%s: %v\n", fieldType.Name, field.Interface())
	}

	return str
}

func getHashOfData(str string) string {
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(str))
	hashBytes := sha1Hash.Sum(nil)
	sha1String := hex.EncodeToString(hashBytes)
	return sha1String
}
