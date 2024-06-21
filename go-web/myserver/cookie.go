package myserver 

import(
//	"fmt"
	"net/http"
) 

func Mysetcookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: "yuyang",
		Value: "fish",
		HttpOnly: true,
	}
	w.Header().Set("Set-Cookie", c1.String())
}
/*
func Mysetcookie(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%T\n", http.Cookie.Name)
//	http.Cookie.Name = "yuyang"
//	http.Cookie.Value = "hangzhou"
//	fmt.Fprintln(w, "get cookie:", http.Cookie)
}
*/

type Stu struct {
	Name string
	Age int
}

