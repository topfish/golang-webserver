package myserver

import (
	"fmt"
	"net/http"
)

func Writeexample(w http.ResponseWriter, r *http.Request) {
	str := "hello I am Myself"
	w.Write([]byte(str))
}

func MyWriteHeader_501(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(501)
	fmt.Fprintln(w, "No Have This Server")
}

func MyWriteHeader_302(w http.ResponseWriter, r *http.Request){
//	w.WriteHeader(302)
	w.Header().Set("Location", "http://www.baidu.com")
	w.WriteHeader(302)
}
