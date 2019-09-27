package service

import (
	"MyChat/model"
	"MyChat/util"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type UserService struct {
}

/**
注册
*/
func (service *UserService) Regist(mobile, password, nickname, avatar, sex string) (user model.User, err error) {
	tempUser := model.User{}

	_, err = DbEngin.Where("mobile=?", mobile).Get(&tempUser)
	if err != nil {
		//查询失败
		log.Println("注册查询失败", err.Error())
		return tempUser, err
	}

	if tempUser.Id > 0 {
		//用户已经注册
		return tempUser, errors.New("手机号已存在")
	}
	tempUser.Mobile = mobile
	tempUser.Passwd = util.Md5encode(password)
	tempUser.Nickname = nickname
	tempUser.Avatar = avatar
	tempUser.Sex = sex
	tempUser.Salt = fmt.Sprintf("%06d", rand.Int31n(10000)) //加密的盐
	tempUser.Token = fmt.Sprintf("%08d", rand.Int31())      //随机数充当token
	tempUser.Createat = time.Now()
	_, err = DbEngin.InsertOne(&tempUser)
	if err != nil {
		log.Println("注册数据插入错误 ")
	}
	return tempUser, err
}

/**
登录
*/
func (service *UserService) Login(mobile, password string) (user model.User, err error) {
	tempUser := model.User{}
	//先查询
	_, err = DbEngin.Where("mobile=?", mobile).Get(&tempUser)
	if err != nil {
		//查询失败
		log.Println("登录查询失败", err.Error())
		return tempUser, err
	}
	if tempUser.Id <= 0 {
		return tempUser, errors.New("未注册")
	}
	if tempUser.Passwd != util.Md5encode(password) {
		return tempUser, errors.New("密码错误")
	}

	return tempUser, nil

}
