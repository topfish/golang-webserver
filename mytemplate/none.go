package mytemplate

import(
	"net/http"
	"html/template"
	"time"
)

func MyProcess(w http.ResponseWriter, r *http.Request) {
	CreateTemp()
	t, _ := template.ParseFiles("/home/tem.html")
	t.Execute(w, time.Now())
}
