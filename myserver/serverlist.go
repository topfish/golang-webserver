package myserver

import (
	"fmt"
	"net/http"
)

func ServerList(w http.ResponseWriter, r *http.Request){
	var list []string
	fmt.Fprintln(w, "可以选择以下path进行访问！！！", "\n")
	list = []string{"/getcookie","/mybody","/header","/returnbody","/getresponsebody","/writeexample","/write501","/write502","/none"}
	for _, k := range list{
        	fmt.Fprintln(w, k, "\n")  
	} 
}

