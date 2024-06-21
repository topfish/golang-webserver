package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webapp/myserver"
	"webapp/mytemplate"
)

func main() {
	defer cleanUp()

	server := http.Server{
		Addr: ":80",
	}
	fmt.Println("Hello I Am WebServer")

	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT)
		<-sig
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			panic(err)
		}
	}()
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

func cleanUp() {
	//index.html缓存清理
	err := os.Remove("index.html")
	//_, err := os.ReadFile("index.html")
	if err != nil {
		log.Println("WARNING: delete cache file index.html failed")
	} else {
		log.Println("clearn cache finshed")
	}
}
