package ctrl

import (
	"MyChat/service"
	"MyChat/util"
	"net/http"
	"time"
)

var contactServer service.ContactServer

/**
添加好友
*/
func AddFriend(w http.ResponseWriter, req *http.Request) {

	_ = req.ParseForm()

	userid := req.PostForm.Get("userid")
	mobile := req.PostForm.Get("mobile")
	//todo 添加校验token 逻辑
	_ = req.PostForm.Get("token")
	err := contactServer.AddFriends(util.Str2int64(userid), mobile)
	if err != nil {
		util.RespFail(w, err.Error())
	} else {
		strings := make(map[string]int32)
		strings["server_time"] = int32(time.Now().Unix())
		util.RespOk(w, strings, "添加好友成功")
	}
}

func LoadFriends(w http.ResponseWriter, req *http.Request) {
	_ = req.ParseForm()
	userid := req.PostForm.Get("userid")
	//todo 添加校验token 逻辑
	_ = req.PostForm.Get("token")
	users, e := contactServer.LoadFriends(util.Str2int64(userid))
	if e != nil {
		util.RespFail(w, e.Error())
	} else {
		util.RespOk(w, users, "")
	}

}
