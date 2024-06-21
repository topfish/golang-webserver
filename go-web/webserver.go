package main

import(
	"net/http"
	"webapp/myserver"
	"webapp/mytemplate"
	"fmt"
)

func main() {
	server := http.Server{
		Addr: ":80",
	}
	fmt.Println("Hello I Am WebServer")

	http.HandleFunc("/", mytemplate.MyHome)
	http.HandleFunc("/list", myserver.ServerList)
	http.HandleFunc("/getcookie", myserver.Mysetcookie)
	http.HandleFunc("/mybody", myserver.Mybody)
	http.HandleFunc("/header", myserver.HeaderReturn)
	http.HandleFunc("/returnbody", myserver.Mybody)
	http.HandleFunc("/getresponsebody", myserver.WriteResponseBody)
	http.HandleFunc("/writeexample", myserver.Writeexample)
	http.HandleFunc("/write501", myserver.MyWriteHeader_501)
	http.HandleFunc("/write302", myserver.MyWriteHeader_302)
	http.HandleFunc("/none", mytemplate.MyProcess)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
