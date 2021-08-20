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
//	checkPagefile()
	createIndex()
	t, err := template.ParseFiles("index.html")
	Myerr(err)
/*	t, err := template.ParseFiles("/etc/webapp/index.html")
	if err != nil {
		fmt.Println("==== please config index.html ====")
		fmt.Println("==== you can execute like this: ====")
		fmt.Println("==== mkdir -p /etc/webapp &&  cp ./mytemplate/index.html /etc/webapp/index.html ====")
	}
*/
	t.Execute(w, time.Now())
}

func createIndex() {
	content := `<!DOCTYPE html>
<html>
<head>
<title>Yuyang's persional webServer!</title>
<style>
    body {
        width: 45em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
</head>
<body>
<h1>Welcome to Yuyang's persional webServer!</h1>
<p>If you see this page, the web server is successfully and
working.</p>

<p>For online documentation and support please refer to
<a href="http://zichuandr.xyz/list">zichuandr.xyz</a>.<br/>
</p>

<h3>There is just a space for persional learning and show.</h3>

<p><em>Thank you for your coming.</em></p>

<h4> ==== 分享常用学习网站 ==== </h4>
<p>[substrate.dev] 
<a href="https://substrate.dev/docs/zh-CN/">substrate.com</a></p>
<p>[The Rust Programming Language] 
<a href="https://doc.rust-lang.org/book/">Rust-lang.org</a></p>

<h4>For others</h4>
<p>If you want to search something, you can checkout 
<a href="http://www.baidu.com/">baidu.com</a>.</p>
</body>
</html>
You ComingTime is : {{.}}`

	buff := []byte(content)
	err := ioutil.WriteFile("index.html", buff, 0666)
	Myerr(err)
}

func Myerr(err error) {
	if err != nil {
		panic(err)
	}
}
