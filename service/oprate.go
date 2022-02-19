package service

import (
	"blog/dao"
	"blog/model"
)

func IsPraiseProcess(commentId int,username string,Model int) error {
	return model.PraisedProcess(dao.DB,commentId,username,Model)
}

func FocusUser(username string, FocusId string) bool {
	ok:=model.AddFocusUser(dao.DB,username,FocusId)
	return ok
}

func QueryFocusList(username string) []model.ReturnInfo {
	list:=model.QueryFocusUserList(dao.DB,username)
	return list
}

func CollectProcess(postId int,username string)bool {
	return model.CollectPost(dao.DB,postId,username)
}

func QueryCollectList(username string) []model.ReturnCollect {
	return model.QueryUserCollect(dao.DB,username)
}
