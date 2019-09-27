package main

import (
	. "MyChat/api"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/login", UserLogin)
	http.HandleFunc("/api/registe", UserRegist)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		writer.Write([]byte("hello"))
	})
	err := http.ListenAndServe(":9091", nil)
	//err := http.ListenAndServe("8080", nil)
	if err != nil {
		log.Println("错误错误", err.Error())
	}
}
