package routes

import (
	"github.com/asullivan219/newsletter/internal/db"
	//"log/slog"
	"net/http"
)

type DropHandler struct {
	Db db.I_database
}

func (d *DropHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		d.drop(w, r)
		return
	default:
		w.Write([]byte("Server error"))
		return
	}
}

func (d *DropHandler) drop(w http.ResponseWriter, r *http.Request) {
	if err := d.Db.DropSubscribers(); err != nil {
		w.Write([]byte("Error Dropping subscriber table"))
	}
	if err := d.Db.InitializeTables(); err != nil {
		w.Write([]byte("Error re initializing tables"))
	}
	w.Write([]byte("Deleted and re initialized tables!"))
}
