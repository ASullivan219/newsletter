package routes

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/asullivan219/newsletter/internal/db"
	"github.com/asullivan219/newsletter/internal/views"
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

	fmt.Println(name)
	fmt.Println(email)
	var nameErr string
	var emailErr string

	if name == "" {
		nameErr = "Name blank"
	}

	if email == "" {
		emailErr = "Email blank"
	}

	if nameErr != "" || emailErr != "" {

		errComponent := views.SignUpForm(name, nameErr, email, emailErr)
		errComponent.Render(r.Context(), w)
		return
	}

	// Place Subscriber in Database
	err := h.Db.CreateSubscriber(email, name)
	if err != nil {
		slog.Error("Error creating subscriber in DB")
		emailErr = "Email Used already!"

		errComponent := views.SignUpForm(name, nameErr, email, emailErr)
		errComponent.Render(r.Context(), w)
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
