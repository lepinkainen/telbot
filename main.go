package main

import (
	"os"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
)

// https://godoc.org/github.com/go-telegram-bot-api/telegram-bot-api#Message

func configChange(e fsnotify.Event) {
	log.Infof("Config file changed:", e.Name)
}

type Handler struct {
	bot    *tgbotapi.BotAPI
	update tgbotapi.Update
}

func main() {

	var handler Handler

	viper.SetEnvPrefix("TELBOT")
	viper.BindEnv("KEY")
	viper.BindEnv("DEBUG")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("Fatal error config file: %s \n", err)
		os.Exit(1)
	}
	viper.WatchConfig()
	viper.OnConfigChange(configChange)

	bot, err := tgbotapi.NewBotAPI(viper.GetString("KEY"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = viper.GetBool("DEBUG")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Errorf("Error getting update channel: ")
		return
	}

	handler.bot = bot

	for update := range updates {
		if update.Message == nil {
			continue
		}

		handler.update = update

		log.Infof("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Text {
		case "hello":
			handler.Say("well, hello there")
		case "listdir":
			handler.listdir()
		default:
			handler.Reply(update.Message.Text)
		}
	}
}
