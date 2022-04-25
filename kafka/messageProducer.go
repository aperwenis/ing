package kafka

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

type MessageProducer interface {
	Produce(message []byte) error
	Close() error
}

func (producer Producer) Produce(message []byte) error {
	err := producer.writer.WriteMessages(context.Background(),
		kafka.Message{
			Value: message,
		},
	)
	if err != nil {
		log.Error().Err(err).Msg("Cannot send message: " + string(message))
	} else {
		log.Debug().Msg("Msg sent: " + string(message))
	}

	return err
}

func (producer Producer) Close() error {
	log.Info().Msg("Closing kafka producer")
	err := producer.writer.Close()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot close message producer")
	} else {
		log.Info().Msg("Producer connection to kafka closed")
	}
	return err
}

func CreateProducer(topic, addr string) MessageProducer {
	w := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return Producer{writer: w}
}
