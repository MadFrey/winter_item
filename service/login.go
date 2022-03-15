package service

import (
	"blog/dao"
	"blog/model"
	"blog/util"
)

func JudgeUserExist(username string, password string) bool {
	pwd := model.QueryUserPwd(username)
	verify := util.PasswordVerify(password, pwd)
	return verify
}

func JudgeUserWithGitId(id int) error {
	err := model.QueryUserGitId(dao.DB, id)
	if err != nil {
		return err
	}
	return nil
}

func QueryUserWithGitId(id int) (string, error) {
	username, err := model.QueryUsernameWithGitId(dao.DB, id)
	return username, err
}
