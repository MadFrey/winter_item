package service

import (
	"blog/dao"
	"blog/model"
	"time"
)

func AddNewUserProcess(args ...interface{}) (int64, error) {
	// 用户数据
	user := model.User{Username: args[0].(string), Password: args[1].(string), CreateTime: time.Now()}
	// 返回
	return model.InsertUser(dao.DB, user)
}
