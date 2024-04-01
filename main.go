package main

import (
	"components"
	"net/http"

	"github.com/a-h/templ"
)

func home(w http.ResponseWriter, r *http.Request)  {
	
}


func main()  {

	helloHandler := components.Hello()
	http.Handle("/", templ.Handler(helloHandler))
	http.ListenAndServe(":8080", nil)

}
