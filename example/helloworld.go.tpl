package main

import (
	"log"
	"net/http"
)

func main() {

	c := MakeRoute(
		Controller{
			Name:     "Root",
			Path:     "/",
			Template: "{{.message}}",
			Data:     map[string]any{"message": "Hello World!"},
			Methods:  "GET OPTIONS",
		})
	c.AddData()

	MakeRoute(Controller{
		Name:     "Home",
		Path:     "/Home",
		Template: "<h1>{{.greeting}}</h1>",
		Data:     map[string]any{"greeting": "welcome"},
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
