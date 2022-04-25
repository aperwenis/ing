package kafka

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

type HandlerFunc func(message []byte) error

type Consumer struct {
	reader  *kafka.Reader
	handler HandlerFunc
	ctx     context.Context
	cancel  context.CancelFunc
}

type MessageConsumer interface {
	StartConsuming() error
	Close() error
}

func (consumer Consumer) StartConsuming() error {
	log.Debug().Msg("consumer started")
	for {
		m, err := consumer.reader.FetchMessage(consumer.ctx)
		if err != nil && err == consumer.ctx.Err() {
			log.Debug().Msg("Closed consumer goroutine")
			return nil
		}
		if err != nil {
			log.Fatal().Err(err).Send()
			return err
		}
		log.Debug().Msg(fmt.Sprintf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value)))
		if err := consumer.handler(m.Value); err != nil {
			log.Error().Err(err).Msg(string(m.Value))
		} else if err := consumer.reader.CommitMessages(consumer.ctx, m); err != nil {
			log.Error().Err(err).Msg("Failed to commit msg")
		}
	}
}

func (consumer Consumer) Close() error {
	consumer.cancel()
	log.Info().Msg("Closing kafka consumer")
	err := consumer.reader.Close()
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot close message consumer")
	} else {
		log.Info().Msg("Consumer connection to kafka closed")
	}

	return err
}

func CreateConsumer(topic, addr, groupId string, handler HandlerFunc) MessageConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{addr},
		GroupID:  groupId,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	ctx, cancel := context.WithCancel(context.Background())

	return Consumer{
		reader:  r,
		handler: handler,
		ctx:     ctx,
		cancel:  cancel,
	}
}
