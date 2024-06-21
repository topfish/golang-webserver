package myserver

import (
	"net/http"
	"fmt"
)

func HeaderReturn(w http.ResponseWriter, r *http.Request){
	for key,value := range r.Header{
		fmt.Fprintln(w, key, "-->", value)
	}
}
