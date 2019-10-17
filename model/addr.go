package model

import "time"

type Address struct {
	Id int64 `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	//谁的10000
	Ip string `xorm:"varchar(120)" form:"ip" json:"ip"` // ip地址
	//对端,10001
	Location string `xorm:"varchar(120)" form:"location" json:"location"` // 经纬度
	//
	Addr string `xorm:"varchar(120)" form:"addr" json:"addr"` // 地址
	//
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"` // 创建时间
}
