package routes

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/asullivan219/newsletter/internal/db"
)

type SubscriberHandler struct {
	Db db.I_database
}

func (h *SubscriberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	slog.Warn("New request to /subscribers",
		"path", r.Method,
	)
	switch r.Method {
	case http.MethodGet:
		h.getSubscriber(w, r)
	case http.MethodPost:
		h.postSubscriber(w, r)
	default:
		w.Write([]byte("Server error"))
	}
}

func (h *SubscriberHandler) getSubscriber(w http.ResponseWriter, r *http.Request) {

	slog.Warn("url Info",
		"url", r.URL.String(),
		"opaque", r.URL.Opaque,
		"scheme", r.URL.Scheme,
	)

	w.Write([]byte("Hello :)"))
}

func (h *SubscriberHandler) postSubscriber(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	err := assertAllNotNil(name, email)

	// Validate information sent is not nil
	if err != nil {
		slog.Error("nil value passed to function")
		w.Write([]byte("Error couldnt subscribe"))
		return
	}

	// Validate form has no empty values
	err = assertAllNotEmpty(name, email)
	if err != nil {
		slog.Error("empty string provided in form")
		w.Write([]byte("Error couldnt subscribe"))
		return
	}

	// Place Subscriber in Database
	err = h.Db.CreateSubscriber(email, name)
	if err != nil {
		slog.Error("Error creating subscriber in DB")
		w.Write([]byte("Error couldnt subscribe"))
		return
	}

	w.Write([]byte("Subscribed!"))
}

func assertAllNotEmpty(T ...string) error {
	for _, curr := range T {
		err := assertNotEmptyString(curr)
		if err != nil {
			return err
		}
	}
	return nil
}

func assertNotEmptyString(T string) error {
	err := assertNotNil(T)
	if err != nil || len(T) <= 0 {
		return errors.New("Empty String error")
	}
	return nil
}

func assertNotNil(T any) error {
	if T != nil {
		return nil
	}
	return errors.New("nil value")
}

func assertAllNotNil(T ...any) error {
	for _, curr := range T {
		err := assertNotNil(curr)
		if err != nil {
			return err
		}
	}
	return nil
}
