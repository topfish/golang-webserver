package myserver

import (
	"fmt"
	"net/http"
)

func Mybody(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Return Request Body:", r.Body)
	len := r.ContentLength
	body := make([]byte , len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
//	fmt.Fprintln(w, body)
}
