package views

import (
	"html/template"
	"net/http"
)

func PageNotFound404(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/error/404.html"))
	t.ExecuteTemplate(w, "layout", "")
}
