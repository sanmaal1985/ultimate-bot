package main

import (
	"log"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
)
func main(){
	bot, err := tgbotapi.NewBotAPI("205446592:AAF7JAztEMT_YadcVEXh9PypNVfa8Dml1s0")
	if err != nil{
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)
	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := bot.GetUpdatesChan(ucfg)
	// читаем обновления из канала
	for update := range updates {

		var reply string

		UserName := update.Message.From.UserName
		UserID := update.Message.From.ID

		ChatID := update.Message.Chat.ID

		Text := update.Message.Text

		log.Printf("[%s] %d %d %s", UserName, UserID, ChatID, Text)

		command := update.Message.Command()
		if command !="" {
			log.Printf("получена команда - %s", command)
			switch command {
			case "/help":
				reply = "Бот Ultimate Frisbee сообщества в г.Барнаул"
			}
		}
		if update.Message.NewChatParticipant.UserName != "" {
			// новый пользователь зашел
			reply = fmt.Sprintf(`Привет @%s!`, update.Message.NewChatParticipant.UserName)
		}
		if reply != "" {
			msg := tgbotapi.NewMessage(ChatID, reply)
			bot.Send(msg)
		} else {
			msg := tgbotapi.NewMessage(ChatID, Text)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}