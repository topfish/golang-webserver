package main

import(
	"net/http"
	"html/template"
)

func MyProcess(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("none.html")
	t.Execute(w, "hello yuyang")
}

func main() {
	server := http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/none", MyProcess)
	
	server.ListenAndServe()
}
