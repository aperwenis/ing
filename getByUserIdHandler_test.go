package main

import (
	"database/sql"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"ing/mocks"
	"ing/training"
	"net/http"
	"testing"
)

func TestGetByUserIdHandler(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	result := []training.Training{training.Training{Id: uint(24)}}
	repositoryMock := training.NewMockRepository(ctrl)
	repositoryMock.EXPECT().GetByUserId("123").Return(result, nil)
	server := Server{
		repository: repositoryMock,
	}
	echoMock := mocks.NewMockContext(ctrl)
	echoMock.EXPECT().Param("userId").Return("123")
	echoMock.EXPECT().JSON(http.StatusCreated, result).Return(nil)

	// when
	err := server.getByUserIdHandler(echoMock)

	// then
	assert.Nil(t, err)
}

func TestGetByUserIdHandler_NotFound(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := training.NewMockRepository(ctrl)
	repositoryMock.EXPECT().GetByUserId("123").Return([]training.Training{}, sql.ErrNoRows)
	server := Server{
		repository: repositoryMock,
	}
	echoMock := mocks.NewMockContext(ctrl)
	echoMock.EXPECT().Param("userId").Return("123")
	echoMock.EXPECT().NoContent(http.StatusNoContent).Return(nil)

	// when
	err := server.getByUserIdHandler(echoMock)

	// then
	assert.Nil(t, err)
}
