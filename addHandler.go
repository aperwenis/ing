package main

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"ing/training"
	"net/http"
)

func (server Server) addHandler(c echo.Context) error {
	t := new(training.Training)
	if err := c.Bind(t); err != nil {
		log.Error().Err(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(t); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	message, err := json.Marshal(t)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := server.producer.Produce(message); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusAccepted)
}
