package main

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"ing/kafka"
	"ing/mocks"
	"ing/training"
	"net/http"
	"testing"
)

func TestAddHandler(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	producerMock := kafka.NewMockMessageProducer(ctrl)
	msg, _ := json.Marshal(new(training.Training))
	producerMock.EXPECT().Produce(msg).Return(nil)
	echoMock := mocks.NewMockContext(ctrl)
	echoMock.EXPECT().Bind(new(training.Training)).Return(nil)
	echoMock.EXPECT().Validate(new(training.Training)).Return(nil)
	echoMock.EXPECT().NoContent(http.StatusAccepted)
	server := Server{
		producer: producerMock,
	}

	// when
	err := server.addHandler(echoMock)

	// then
	assert.Nil(t, err)
}
