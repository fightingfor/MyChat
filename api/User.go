package api

import (
	"MyChat/model"
	"MyChat/service"
	"MyChat/util"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"net/http"
)

var userService service.UserService

func UserLogin(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	password := r.PostForm.Get("password")
	log.Println("mobile>>>", mobile, "password>>>", password)
	if mobile == "" || password == "" {
		util.RespFail(w, errors.New("参数错误").Error())
		return
	}
	user, err := userService.Login(mobile, password)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "")
	}
}

func UserRegist(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	mobile := r.PostForm.Get("mobile")
	password := r.PostForm.Get("password")
	log.Println("mobile>>>", mobile, "password>>>", password)

	nickname := fmt.Sprintf("user%06d", rand.Int31())
	avatar := ""
	sex := model.SEX_UNKNOW

	if mobile == "" || password == "" {
		util.RespFail(w, errors.New("参数错误").Error())
		return
	}

	user, err := userService.Regist(mobile, password, nickname, avatar, sex)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		util.RespOk(w, user, "")
	}
}
