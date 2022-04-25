package main

import (
	"context"
	"github.com/go-playground/validator/v10"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"ing/kafka"
	"ing/training"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	setUpLogger(os.Getenv("DEBUG_MODE") == "true")
	db, err := sqlx.Open("pgx", os.Getenv("DB_ADDRESS"))
	if err != nil {
		log.Fatal().Err(err).Send()
		return
	}
	repository := training.CreateRepository(db)
	producer := kafka.CreateProducer(os.Getenv("KAFKA_TOPIC"), os.Getenv("KAFKA_BROKER_ADDRESS"))
	consumer := kafka.CreateConsumer(
		os.Getenv("KAFKA_TOPIC"),
		os.Getenv("KAFKA_BROKER_ADDRESS"),
		os.Getenv("KAFKA_GROUP_ID"),
		createMessageHandler(repository),
	)
	go consumer.StartConsuming()
	server := createServer(producer, repository)
	e := startEcho(server)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info().Msg("Initialised closing")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Info().Msg("Shutting down echo")
	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Send()
	}
	if err := consumer.Close(); err != nil {
		log.Fatal().Err(err).Send()
	}
	if err := producer.Close(); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func setUpLogger(isDebugMode bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if isDebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func startEcho(server Server) *echo.Echo {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.POST("/training", server.addHandler)
	e.GET("/training/:id", server.getByIdHandler)
	e.GET("/training/user/:userId", server.getByUserIdHandler)
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("shutting down the server")
		}
	}()

	return e
}
