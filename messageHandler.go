package main

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"ing/kafka"
	"ing/training"
)

func createMessageHandler(repo training.Repository) kafka.HandlerFunc {
	return func(message []byte) error {
		var t training.Training
		if err := json.Unmarshal(message, &t); err != nil {
			log.Error().Err(err).Msg("Cannot unmarshal message: " + string(message))
			return err
		}
		if err := repo.Create(t); err != nil {
			log.Error().Err(err).Msg("Cannot create new training: " + string(message))
			return err
		}

		return nil
	}
}
