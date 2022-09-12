package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/sirupsen/logrus"
	// docs are generated by Swag CLI, you have to import them.
	_ "example.choreo.dev/docs"

	"example.choreo.dev/api/routes"
	"example.choreo.dev/internal/config"
)

// @title Choreo sample api
// @version 1.0
// @description This is a sample swagger
// @host localhost:8000
// @BasePath /
func main() {
	app := fiber.New(fiber.Config{
		AppName:               "choreo-example-app",
		ReadTimeout:           time.Second * 2,
		Prefork:               false,
		DisableKeepalive:      true,
		DisableStartupMessage: true,
	})
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	routes.Initialize(app)

	go func() {
		logrus.WithFields(logrus.Fields{"port": cfg.Port}).Info("choreo-example-app is running")

		if err := app.Listen(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	sigtermC := make(chan os.Signal, 1)
	signal.Notify(sigtermC, os.Interrupt, syscall.SIGTERM)

	<-sigtermC // block until SIGTERM is received
	logrus.Info("SIGTERM received: gracefully shutting down...")

	if err := app.Shutdown(); err != nil {
		logrus.Errorf("server shutdown error: %v", err)
	}

}
