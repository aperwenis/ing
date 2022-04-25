package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

func (server Server) getByIdHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		log.Error().Err(err).Msg("Cannot convert id from path to int")
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	t, err := server.repository.GetById(uint(id))
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNoContent)
	}
	if err != nil {
		log.Error().Err(err).Send()
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, t)
}
