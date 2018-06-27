package main

import (
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
)

func listdir(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	var files []string
	fileinfo, err := ioutil.ReadDir(viper.GetString("dirname"))
	if err != nil {
		log.Errorf("Error in listdir: %v", err)
	}
	for _, file := range fileinfo {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		files = append(files, file.Name())
	}

	say(bot, update, strings.Join(files, "\n"))
}
