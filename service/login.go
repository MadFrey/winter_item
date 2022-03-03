package service

import (
	"blog/model"
	"blog/util"
)

func JudgeUserExist(username string, password string) bool {
	pwd:=model.QueryUserPwd(username)
	verify := util.PasswordVerify(password, pwd)
	return verify
}
