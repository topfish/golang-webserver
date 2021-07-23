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

func MyNginx(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("/usr/local/nginx/html/index.html")
	t.Execute(w, time.Now())
}
