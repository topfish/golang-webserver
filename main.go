package main

import(
	"fmt"
	"net/http"
	"webapp/myserver"
	"webapp/mytemplate"
)

func main() {
	fmt.Println("I Am Server")
	s1 := myserver.Stu{"yuyang", 30}
	fmt.Println(s1)
	server := http.Server{
		Addr: ":80",
	}
	http.HandleFunc("/", myserver.ServerList)
	http.HandleFunc("/getcookie", myserver.Mysetcookie)
	http.HandleFunc("/mybody", myserver.Mybody)
	http.HandleFunc("/header", myserver.HeaderReturn)
	http.HandleFunc("/returnbody", myserver.Mybody)
	http.HandleFunc("/getresponsebody", myserver.WriteResponseBody)
	http.HandleFunc("/writeexample", myserver.Writeexample)
	http.HandleFunc("/write501", myserver.MyWriteHeader_501)
	http.HandleFunc("/write302", myserver.MyWriteHeader_302)
	http.HandleFunc("/none", mytemplate.MyProcess)

	server.ListenAndServe()
}
