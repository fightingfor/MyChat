package service

import (
	"MyChat/model"
	"time"
)

type AddrServer struct {
}

func (server AddrServer) GetAddr(location string, ip string, addr string) error {

	temp := model.Address{}
	temp.Ip = ip
	temp.Location = location
	temp.Addr = addr
	temp.Createat = time.Now()
	_, e := DbEngin.Insert(temp)
	if e != nil {
		return e
	}

	return nil
}
