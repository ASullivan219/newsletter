package routes

import (
	"github.com/a-h/templ"
	"github.com/asullivan219/newsletter/internal/views"
)

func Index() *templ.ComponentHandler {
	title := "Newsletter - signup"
	mainContent := views.SignUpForm("", "", "", "")
	return templ.Handler(views.Layout(title, mainContent))
}
