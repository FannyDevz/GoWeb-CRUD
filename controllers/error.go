package error

import (
	"net/http"
	"text/template"
)

func ErrorController(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/error.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}
