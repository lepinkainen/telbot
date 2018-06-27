package main

import (
	"io/ioutil"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// List directory given in dirname config variable
func (h Handler) listdir() {

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

	h.Say(strings.Join(files, "\n"))
}
