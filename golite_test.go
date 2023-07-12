package golite

import (
	"fmt"
	"testing"
	"net/http"
)

func Test1(t *testing.T) {
	fmt.Println("test")
}

func Test2(t *testing.T) {
	c := MakeRoute(
		Controller{
			Name:     "Root",
			Path:     "/",
			Template: "{{.message}}",
			Data:     map[string]any{"message": "Welcome To Golite!"},
			Methods:  "GET OPTIONS",
		})
	c.AddData()

	MakeRoute(Controller{
		Name:     "Home",
		Path:     "/Home",
		Template: "<h1>{{.greeting}}</h1>",
		Data:     map[string]any{"greeting": "welcome"},
	})


	fmt.Printf("%T\n",c)
	go http.ListenAndServe(":8082", nil)
	res, _ := http.Get("http://example.com/")

	fmt.Printf(" %T %T\n",res,c)
}


