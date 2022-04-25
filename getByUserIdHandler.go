package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (server Server) getByUserIdHandler(c echo.Context) error {
	t, err := server.repository.GetByUserId(c.Param("userId"))
	if err == sql.ErrNoRows {
		return c.NoContent(http.StatusNoContent)
	}
	if err != nil {
		log.Error().Err(err).Send()
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusCreated, t)
}
