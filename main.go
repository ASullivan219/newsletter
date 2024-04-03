package main

import (
	"components"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

const (
	inProduction = false
)

func main()  {

	helloHandler := components.Hello()
	http.Handle("/", templ.Handler(helloHandler))

	if (!inProduction){
		fmt.Println("Not production serving on port 8080:")
		http.ListenAndServe(":8080", nil)
	} else {
		fmt.Println("Production server listening")
		http.ListenAndServeTLS(":8080","./certs/cert.pem", "./certs/privateKey.key" , nil)
	}

}
