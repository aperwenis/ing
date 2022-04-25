package main

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"ing/training"
	"testing"
)

func TestMessageHandler(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := training.NewMockRepository(ctrl)
	activity := training.Training{Id: 321}
	repositoryMock.EXPECT().Create(activity).Return(nil)
	msg, _ := json.Marshal(activity)

	// when
	err := createMessageHandler(repositoryMock)(msg)

	// then
	assert.Nil(t, err)
}
