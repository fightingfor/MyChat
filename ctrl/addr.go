package ctrl

import (
	"MyChat/service"
	"MyChat/util"
	"net/http"
)

var addrserver service.AddrServer

func GetAddr(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	query := r.URL.Query()
	location := query.Get("location")
	addr := query.Get("addr")
	e := addrserver.GetAddr(location, ip, addr)
	if e != nil {
		util.RespFail(w, e.Error())
	} else {
		util.RespOk(w, nil, "成功")
	}
}
