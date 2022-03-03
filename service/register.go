package service

import (
	"blog/dao"
	"blog/model"
	"blog/util"
	"log"
	"time"
)

func AddNewUserProcess(username string,password string) (int64, error) {
	// 用户数据
	hash, err := util.PasswordHash(password)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	user := model.User{Username: username, Password: hash, CreateTime: time.Now()}
	// 返回
	return model.InsertUser(dao.DB, user)
}
