package routes

import (
	"bytes"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/asullivan219/newsletter/internal/db"
	"github.com/asullivan219/newsletter/internal/emailer"
	"github.com/asullivan219/newsletter/internal/views"
)

type SubscriberHandler struct {
	Db          db.I_database
	EmailClient emailer.I_Notifier
}

func (h *SubscriberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	w.Write([]byte("Hello :)"))
}

func (h *SubscriberHandler) postSubscriber(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.PostFormValue("Name")
	email := r.PostFormValue("Email")

	var nameErr string
	var emailErr string

	if name == "" {
		nameErr = "Name blank"
	}

	if !emailer.ValidateEmail(email) {
		emailErr = "Invalid Email"
		if email == "" {
			emailErr = "Email cant be blank"
		}
	}

	if nameErr != "" || emailErr != "" {
		slog.Error("Error in form",
			"Name", name, "NameErr", nameErr,
			"Email", email, "EmailErr", emailErr,
		)
		errComponent := views.SignUpForm(name, nameErr, email, emailErr)
		errComponent.Render(r.Context(), w)
		return
	}

	// Place Subscriber in Database
	err := h.Db.CreateSubscriber(email, name)
	if err != nil {
		slog.Error(
			"Error Creating subscriber",
			"error", err.Error())

		emailErr = "Email Taken!"

		errComponent := views.SignUpForm(name, nameErr, email, emailErr)
		errComponent.Render(r.Context(), w)
		return
	}

	subscriber, _ := h.Db.GetSubscriber(email)

	validationLink := fmt.Sprintf(
		"%s://%s/validate?email=%s&code=%s",
		os.Getenv("PROTOCOL"),
		os.Getenv("DOMAIN"),
		subscriber.Email,
		subscriber.VerificationCode)

	templEmail := views.VerifySignupEmail(subscriber, validationLink)

	buffer := bytes.NewBuffer([]byte(""))
	templEmail.Render(r.Context(), buffer)

	message := fmt.Sprintf(
		"Subject: alex-sullivan.com Newsletter - Verification\n"+
			"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"+
			"%s", buffer.String())

	slog.Info(
		"Sending signup email",
		"email", email)

	go h.EmailClient.NotifyOne(message, email)
	component := views.SignUpResponse()
	component.Render(r.Context(), w)
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
