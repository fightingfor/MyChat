package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type H struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func RespOk(w http.ResponseWriter, data interface{}, msg string) {

	Resp(w, 0, data, msg)
}

func RespFail(w http.ResponseWriter, msg string) {

	Resp(w, -1, nil, msg)
}

func Resp(w http.ResponseWriter, code int, data interface{}, msg string) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	h := H{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	bytes, e := json.Marshal(h)
	if e != nil {
		log.Println(e.Error())
	}
	_, _ = w.Write(bytes)
}
