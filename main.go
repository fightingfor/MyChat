package main

import (
	. "MyChat/ctrl"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/login", UserLogin)         //登录
	http.HandleFunc("/api/registe", UserRegist)      //注册
	http.HandleFunc("/api/addfriend", AddFriend)     //添加好友
	http.HandleFunc("/api/loadfriends", LoadFriends) //加载好友

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

		writer.Write([]byte("hello"))
	})
	err := http.ListenAndServe(":9091", nil)
	//err := http.ListenAndServe("8080", nil)
	if err != nil {
		log.Println("错误错误", err.Error())
	}
}
