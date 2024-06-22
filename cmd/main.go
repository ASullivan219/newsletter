package main

import (
	"github.com/asullivan219/newsletter/components"

	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

const (
	inProduction = false
)

func main() {

	renamedHandler := components.Hello()
	http.Handle("/", templ.Handler(renamedHandler))
	http.Handle("/asdf", templ.Handler(renamedHandler))

	if !inProduction {
		fmt.Println("Not production serving on port 8080:")
		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println("Production server listening")
		http.ListenAndServeTLS(":8080", "./certs/cert.pem", "./certs/privateKey.key", nil)
	}

}
