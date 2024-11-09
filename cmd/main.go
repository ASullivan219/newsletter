package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/asullivan219/newsletter/internal/db"
	"github.com/asullivan219/newsletter/internal/emailer"
	"github.com/asullivan219/newsletter/internal/routes"
	"github.com/asullivan219/newsletter/internal/server"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("./resources/.env")

	if err != nil {
		slog.Error("Error loading .env file",
			"error", err.Error(),
		)
	}

	emailClient := emailer.NewEmailNotifier(
		os.Getenv("FROM_EMAIL"),
		os.Getenv("EMAIL_PASSWORD"),
		os.Getenv("EMAIL_HOST"),
		os.Getenv("EMAIL_PORT"))

	database := db.NewDb(
		os.Getenv("DB_FILE"))

	server := server.Server{
		Mux:          http.NewServeMux(),
		Db:           database,
		EmailService: emailClient,
		Port:         os.Getenv("PORT")}

	subscriberHandler := routes.SubscriberHandler{
		Db:          database,
		EmailClient: emailClient,
	}

	validateHandler := routes.ValidateHandler{
		Db:          database,
		EmailClient: emailClient,
	}

	if os.Getenv("DROP_ENABLED") == "true" {
		slog.Info("Drop tables enabled, adding routes")
		dropHandler := routes.DropHandler{
			Db: database,
		}
		server.AddRoute("/drop", &dropHandler)
	}

	server.AddRoute("/", routes.Index())
	server.AddRoute("/subscriber", &subscriberHandler)
	server.AddRoute("/validate", &validateHandler)
	server.Serve()
}
