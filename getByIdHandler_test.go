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

func TestGetByIdHandler(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	activity := training.Training{Id: uint(24)}
	repositoryMock := training.NewMockRepository(ctrl)
	repositoryMock.EXPECT().GetById(uint(24)).Return(activity, nil)
	server := Server{
		repository: repositoryMock,
	}
	echoMock := mocks.NewMockContext(ctrl)
	echoMock.EXPECT().Param("id").Return("24")
	echoMock.EXPECT().JSON(http.StatusCreated, activity).Return(nil)

	// when
	err := server.getByIdHandler(echoMock)

	// then
	assert.Nil(t, err)
}

func TestGetByIdHandler_NotFound(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repositoryMock := training.NewMockRepository(ctrl)
	activity := training.Training{Id: uint(24)}
	repositoryMock.EXPECT().GetById(uint(24)).Return(activity, sql.ErrNoRows)
	server := Server{
		repository: repositoryMock,
	}
	echoMock := mocks.NewMockContext(ctrl)
	echoMock.EXPECT().Param("id").Return("24")
	echoMock.EXPECT().NoContent(http.StatusNoContent).Return(nil)

	// when
	err := server.getByIdHandler(echoMock)

	// then
	assert.Nil(t, err)
}
