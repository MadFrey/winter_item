package service

import (
	"blog/dao"
	"blog/model"
)

func JudgeUserExist(username string, password string) bool {
	id := model.QueryIdWithUserPwd(dao.DB, username, password)
	if id > 0 {
		return true
	} else {
		return false
	}
}
