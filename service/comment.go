package service

import (
	"blog/dao"
	"blog/model"
	"strconv"
	"strings"
)

func NoTokenQueryComment(PostId int, page int, size int) []model.RetComment {
	comments := model.NoTokenGetComment(dao.DB, PostId, page, size)
	var commentList []model.RetComment
	for _, comment := range comments {
		newPicture := strings.Split(comment.Picture, ",")
		commentList = append(commentList, model.RetComment{Id: comment.Id, PostId: comment.PostId, PublishTime: comment.PublishTime, Content: comment.Content,
			Username: comment.Username, Avatar: comment.Avatar, Nickname: comment.Nickname, IsPraised: comment.IsPraised,
			IsFocus: comment.IsFocus, PraiseCount: comment.PraiseCount, Picture: newPicture, Model: comment.Model})
	}
	return commentList
}

func QueryComment(PostId int, page int, size int, username string) []model.RetComment {
	comments := model.GetComment(dao.DB, PostId, page, size, username)
	var commentList []model.RetComment
	for _, comment := range comments {
		newPicture := strings.Split(comment.Picture, ",")
		commentList = append(commentList, model.RetComment{Id: comment.Id, PostId: comment.PostId, PublishTime: comment.PublishTime, Content: comment.Content,
			Username: comment.Username, Avatar: comment.Avatar, Nickname: comment.Nickname, IsPraised: comment.IsPraised,
			IsFocus: comment.IsFocus, PraiseCount: comment.PraiseCount, Picture: newPicture, Model: comment.Model})
	}
	return commentList
}

func NoTokenSecondQueryComment(PostId int, page int, size int) []model.SecondComment {
	comments := model.NoTokenGetComment(dao.DB, PostId, page, size)
	var commentList []model.SecondComment
	for _, comment := range comments {
		Id,_:=strconv.Atoi(comment.Id)
		_,retCom:=model.QueryCommentWithId(Id)
		_,Com:=model.QueryCommentWithId(retCom.PostId)
		commentList = append(commentList, model.SecondComment{Id: comment.Id, PostId: comment.PostId, PublishTime: comment.PublishTime, Content: comment.Content,
			Username: comment.Username, Avatar: comment.Avatar, Nickname: comment.Nickname, ReplyUserId: Com.Username, ReplyUserNickname: Com.Nickname,
			IsPraised: comment.IsPraised, IsFocus: comment.IsFocus, PraiseCount: comment.PraiseCount, Model: comment.Model})
	}
	return commentList
}

func QuerySecondComment(PostId int, page int, size int, username string) []model.SecondComment {
	comments := model.GetComment(dao.DB, PostId, page, size, username)
	var commentList []model.SecondComment
	for _, comment := range comments {
		Id,_:=strconv.Atoi(comment.Id)
		_,retCom:=model.QueryCommentWithId(Id)
		_,Com:=model.QueryCommentWithId(retCom.PostId)
		commentList = append(commentList, model.SecondComment{Id: comment.Id, PostId: comment.PostId, PublishTime: comment.PublishTime, Content: comment.Content,
			Username: comment.Username, Avatar: comment.Avatar, Nickname: comment.Nickname, ReplyUserId: Com.Username, ReplyUserNickname: Com.Nickname,
			IsPraised: comment.IsPraised, IsFocus: comment.IsFocus, PraiseCount: comment.PraiseCount, Model: comment.Model})
	}
	return commentList
}

func PostComment(postId int, content string, photo string, Model int, username string) (int64, error) {
	id, err := model.InsertComment(dao.DB, postId, content, photo, Model, username)
	return id, err
}

func UpdateComment(id int, content string, photo string) (int64, error) {
	_, err := model.UpdateComment(dao.DB, id, content, photo)
	return 0, err
}

func DeleteComment(id int,username string) (int64, bool) {
	_,  ok:= model.DeleteCommentWithId(dao.DB, id,username)
	return 0, ok
}
