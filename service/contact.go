package service

import (
	"MyChat/model"
	"errors"
	"time"
)

type ContactServer struct {
}

/**
添加好友
*/
func (server ContactServer) AddFriends(userid, dstid int64) error {
	//自己不能加自己
	if userid == dstid {
		return errors.New("自己不能加自己")
	}

	temp := model.Contact{}
	//查询 判断是否已经是好友
	_, e := DbEngin.Where("ownerid = ?", userid).
		And("dstobj = ?", dstid).And("cate = ?", model.CONCAT_CATE_USER).Get(&temp)

	if e != nil {
		return e
	}
	if temp.Id > 0 { //已经是好友关系
		return errors.New("已经是好友")
	}
	//查询该好友是否存在
	userTemp := model.User{}
	_, e2 := DbEngin.Where("id = ?", dstid).Get(&userTemp)
	if e2 != nil {
		return e2
	}
	if userTemp.Id <= 0 {
		return errors.New("不存在该用户")
	}
	//插入各自信息到 联系人库
	session := DbEngin.NewSession()
	_ = session.Begin() //开启事务
	//插入到自己一方
	_, e3 := session.InsertOne(model.Contact{
		Ownerid:  userid,
		Dstobj:   dstid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})
	//插入到对方
	_, e4 := session.InsertOne(model.Contact{
		Ownerid:  dstid,
		Dstobj:   userid,
		Cate:     model.CONCAT_CATE_USER,
		Createat: time.Now(),
	})

	if e3 == nil && e4 == nil {
		e := session.Commit()
		if e != nil {
			return e
		} else {
			return nil
		}
	} else {
		e := session.Rollback()
		if e != nil {
			return e
		}
		if e3 != nil {
			return e3
		} else {
			return e4
		}
	}
}

func (server ContactServer) LoadFriends(userid int64) ([]model.User, error) {
	contactList := make([]model.Contact, 0)
	dstobjList := make([]int64, 0)
	userList := make([]model.User, 0)
	e := DbEngin.Where("ownerid = ? and cate = ?", userid, model.CONCAT_CATE_USER).Find(&contactList)
	if e != nil {
		return userList, e
	}
	for _, v := range contactList {
		dstobjList = append(dstobjList, v.Dstobj)
	}

	if len(dstobjList) == 0 {
		return userList, nil
	}

	e = DbEngin.In("id", dstobjList).Find(&userList)
	if e != nil {
		return userList, e
	} else {
		return userList, nil
	}
}
