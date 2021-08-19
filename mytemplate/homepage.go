package mytemplate

import(
	"net/http"
	"html/template"
	"time"
	"io/ioutil"
	"fmt"
)

func MyProcess(w http.ResponseWriter, r *http.Request) {
	filename := "./mytemplate/tem.html"
	CreateTemp()
	t, _ := template.ParseFiles(filename)
	t.Execute(w, time.Now())
	DeleteTemp(filename)
}

func checkPagefile() {
	_, err := ioutil.ReadFile("/etc/webapp/index.html")
	if err != nil {
		fmt.Println("==== please config index.html ====")
		fmt.Println("==== you can execute like this: ====")
		fmt.Println("==== mkdir -p /etc/webapp &&  cp ./mytemplate/index.html /etc/webapp/index.html ====")
	}
}

func MyHome(w http.ResponseWriter, r *http.Request) {
	checkPagefile()
	t, _ := template.ParseFiles("/etc/webapp/index.html")
	t.Execute(w, time.Now())
}
