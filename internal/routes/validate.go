package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/asullivan219/newsletter/internal/db"
	"github.com/asullivan219/newsletter/internal/emailer"
	"github.com/asullivan219/newsletter/internal/views"
)

const (
	PAGE_TITLE = "Newsletter - Verify"
)

type ValidateHandler struct {
	Db          db.I_database
	EmailClient emailer.I_Notifier
}

func (v *ValidateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		v.validate(w, r)
		return
	default:
		status := http.StatusNotFound
		w.WriteHeader(status)
		page := generatePage("404")
		page.Render(r.Context(), w)
	}

}

func (v *ValidateHandler) validate(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	code := r.URL.Query().Get("code")
	errorComponent := generatePage("Whoops, something went wrong")
	subscriber, err := v.Db.GetSubscriber(email)

	if err != nil {
		errorComponent.Render(r.Context(), w)
		return
	}

	if subscriber.VerificationCode != code {
		errorComponent.Render(r.Context(), w)
		return
	}

	if subscriber.Verified {
		layout := generatePage("Already subscribed!")
		layout.Render(r.Context(), w)
		return
	}

	_, err = v.Db.VerifySubscriber(subscriber)
	if err != nil {
		errorComponent.Render(r.Context(), w)
		return
	}
	layout := generatePage("Subscription Confirmed!")
	layout.Render(r.Context(), w)
	return
}

func generatePage(content string) templ.Component {
	contentComponent := views.Verification(content)
	return views.Layout(PAGE_TITLE, contentComponent)
}
