package main

import (
	"ing/kafka"
	"ing/training"
)

type Server struct {
	producer   kafka.MessageProducer
	repository training.Repository
}

func createServer(producer kafka.MessageProducer, repository training.Repository) Server {
	return Server{
		producer:   producer,
		repository: repository,
	}
}
