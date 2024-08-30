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
	database := &db.Supabase{ApiUrl: os.Getenv("API_URL"), ServiceKey: os.Getenv("SERVICE_KEY")}
	server := server.Server{Mux: http.NewServeMux(), Db: database, Port: "8070"}
	userHandler := routes.UserHandler{Db: database}

	server.AddRoute("/", routes.Index())
	server.AddRoute("/user", &userHandler)
	server.Serve()
}
