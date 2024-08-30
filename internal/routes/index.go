package routes

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/asullivan219/newsletter/internal/db"
	"github.com/asullivan219/newsletter/internal/views"
)

func Index() *templ.ComponentHandler {
	title := "Newsletter - signup"
	return templ.Handler(views.Layout(title))
}

type UserHandler struct {
	Db db.I_database
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getUser(w, r)
	case http.MethodPost:
		h.post(w, r)
	default:
		w.Write([]byte("Server error"))
	}
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {
	component := views.Layout("Get User")
	component.Render(r.Context(), w)
}

func (h *UserHandler) post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		slog.Info(
			"error empty field",
			slog.String("name", name),
			slog.String("email", email),
		)
		w.Write([]byte("server error"))
		return
	}

	err := h.Db.PutSubscriber(email, name)
	if err != nil {
		slog.Info(
			"error writing to database",
			slog.String("error", err.Error()),
		)
		w.Write([]byte("server error"))
		return
	}

	w.Write([]byte("hello world"))
}
