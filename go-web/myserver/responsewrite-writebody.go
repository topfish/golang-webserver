package myserver

import (
	"net/http"
)

func WriteResponseBody(w http.ResponseWriter, r *http.Request){
	str := `<html>
	<head><tile>Go Web Programming</title><head>
	<body><h1>Hello World</h1></body>
	</html>`
	w.Write([]byte(str))
}
