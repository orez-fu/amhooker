package cfgtelegram

import (
	"amhooker/amhooker/configs"
	"context"
	"log"
	"sync"
	"time"

	"github.com/go-telegram/bot"
)

var (
	telegramOnce sync.Once
	telegramBot  *bot.Bot
)

func GetTelegramBot(alertConfigPath string) *bot.Bot {
	telegramOnce.Do(func() {
		setup(alertConfigPath)
	})
	return telegramBot
}

func setup(alertConfigPath string) {
	alertConfig := configs.ReadConfig(alertConfigPath)
	log.Println(alertConfig)
	for {
		var err error
		telegramBot, err = bot.New(alertConfig.TelegramConfig.BotToken)
		if err == nil {
			break
		} else {
			log.Printf("Error initializing telegram connection: %s", err)
			time.Sleep(time.Second)
		}
	}
	log.Println("Telegram bot has reborn")
	go telegramBot.Start(context.TODO())
}
