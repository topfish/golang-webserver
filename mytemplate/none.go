package mytemplate

import(
	"net/http"
	"html/template"
	"time"
)

func MyProcess(w http.ResponseWriter, r *http.Request) {
	filename := "./mytemplate/tem.html"
	CreateTemp()
	t, _ := template.ParseFiles(filename)
	t.Execute(w, time.Now())
	DeleteTemp(filename)
}

func MyHome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./mytemplate/index.html")
	t.Execute(w, time.Now())
}
