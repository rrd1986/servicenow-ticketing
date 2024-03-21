package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func LogFlags() {
	for key, value := range viper.GetViper().AllSettings() {
		log.WithFields(log.Fields{
			key: value,
		}).Info("Command Flag")
	}
}
