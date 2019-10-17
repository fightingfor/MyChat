package ctrl

import (
	"MyChat/service"
	"MyChat/util"
	"net/http"
)

var locationservice service.LocationService

func GetLocation(w http.ResponseWriter, r *http.Request) {
	ip := r.RemoteAddr
	query := r.URL.Query()
	location := query.Get("location")

	e := locationservice.GetLocation(ip, location)
	if e != nil {
		util.RespFail(w, "失败"+e.Error())

	} else {
		util.RespOk(w, nil, "成功")
	}
}
