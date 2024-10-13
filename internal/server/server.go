// Wrapper class for the http router
package server

import (
	"fmt"
	"net/http"

	"github.com/asullivan219/newsletter/internal/db"
	"github.com/asullivan219/newsletter/internal/emailer"
)

// HTTP server
type Server struct {
	Mux          *http.ServeMux
	Db           db.I_database
	EmailService emailer.I_Notifier
	Port         string
}

// Add a route for the given path and handler
func (s *Server) AddRoute(path string, handler http.Handler) {
	s.Mux.Handle(path, handler)
}

// Start service on the provided port
func (s *Server) Serve() {
	fmt.Printf("Serving on port %s", s.Port)
	serveString := fmt.Sprintf(":%s", s.Port)
	http.ListenAndServe(serveString, s.Mux)
}
