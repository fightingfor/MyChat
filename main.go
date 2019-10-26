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
	http.HandleFunc("/addr", GetAddr)
	http.HandleFunc("/getLocation", GetLocation)
	http.HandleFunc("/api/chat", Chat)

	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//
	//	addr := request.RemoteAddr
	//	log.Println("addr>>>", addr)
	//	writer.Write([]byte("hello"))
	//})
	err := http.ListenAndServe(":8081", nil)
	//err := http.ListenAndServe("8080", nil)
	if err != nil {
		log.Println("错误错误", err.Error())
	}
}
