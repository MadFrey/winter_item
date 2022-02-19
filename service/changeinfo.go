package service

import (
	"blog/dao"
	"blog/model"
)

func UpdatePersonalInfo	(username string,email string,NickName string,AvatarUrl string,Gender string,introduction string,qq int,birthday string,phone int) (int64,error) {
	user:= model.QueryUserInfo(dao.DB,username)
	user.Email=email
	user.NickName=NickName
	user.AvatarUrl=AvatarUrl
	user.Gender=Gender
	user.Phone=phone
	user.Introduction=introduction
	user.Qq=qq
	user.Birthday=birthday
	return model.UpdateUserInfo(dao.DB,user)
}
func UpdatePersonalInfoNoAvatar(username string,email string,NickName string,Gender string,introduction string,qq int,birthday string,phone int) (int64,error) {
	user:= model.QueryUserInfo(dao.DB,username)
	user.Email=email
	user.NickName=NickName
	user.Gender=Gender
	user.Phone=phone
	user.Introduction=introduction
	user.Qq=qq
	user.Birthday=birthday
	return model.UpdateUserInfo(dao.DB,user)
}

func GetUserinfo(username string) model.User {
	return model.QueryUserInfo(dao.DB,username)
}

func UpdateUserPwd(username string,NewPassword string)(int64,error)  {
	return model.ChangeUserPwd(dao.DB,username,NewPassword)
}


