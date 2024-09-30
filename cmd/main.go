package main

import (
	"net/http"
	"os"

	"github.com/asullivan219/newsletter/internal/db"
	"github.com/asullivan219/newsletter/internal/routes"
	"github.com/asullivan219/newsletter/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database := db.NewDb(os.Getenv("DB_FILE"))
	server := server.Server{Mux: http.NewServeMux(), Db: database, Port: "8080"}
	subscriberHandler := routes.SubscriberHandler{Db: database}
	server.AddRoute("/", routes.Index())
	server.AddRoute("/subscriber", &subscriberHandler)
	server.Serve()
}
