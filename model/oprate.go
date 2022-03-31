package model

import (
	"blog/dao"
	"database/sql"
	"log"
	"time"
)

type Praised struct {
	Id       int
	Username string
}

type FocusUser struct {
	id       int
	username string
	FocusId  int
}

type ReturnInfo struct {
	Username     string
	AvatarUrl    string
	NickName     string
	Introduction string
}

type ReturnCollect struct {
	PostId      int
	Title       string
	PublishTime time.Time
	Username    string
	Avatar      string
	NickName    string
}

func PraisedProcess(DB *sql.DB, Id int, username string, Model int) error {
	var err error
	_, err = QueryArticleWithId(dao.DB, Id, username)
	if err != nil {
		return err
	}
	err, _ = QueryCommentWithId(Id)
	if err != nil {
		return err
	}
	if Model == 1 {
		sqlstr := "insert into praise(postId,username) value(?,?)"
		_, err := DB.Exec(sqlstr, Id, username)
		if err != nil {
			log.Println(err)
		}
	} else {
		if Model == 2 {
			sqlstr := "insert into commentPraise(commentId,username) value(?,?)"
			_, err := DB.Exec(sqlstr, Id, username)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return err
}

func AddFocusUser(DB *sql.DB, username string, FocusId string) bool {
	sqlstr2 := "select id from focusUser where username=? and focusId=?"
	row := DB.QueryRow(sqlstr2, username, FocusId)
	id := 0
	err := row.Scan(&id)
	if err == nil {
		return false
	}
	sqlstr := "insert into focusUser(username ,focusId) values(?,?)"
	_, err = DB.Exec(sqlstr, username, FocusId)
	if err != nil {
		return false
	}
	return true
}

func QueryFocusUserList(DB *sql.DB, username string) []ReturnInfo {
	sqlstr := "select focusId from focusUser where username = ?"
	rows, err := DB.Query(sqlstr, username)
	if err != nil {
		log.Println(err)
		return nil
	}
	var rets []ReturnInfo
	defer rows.Close()
	for rows.Next() {
		var FocusId string
		rows.Scan(&FocusId)
		var ret ReturnInfo
		u := QueryUserInfo(dao.DB, FocusId)
		ret.Username = u.Username
		ret.AvatarUrl = u.AvatarUrl
		ret.NickName = u.NickName
		ret.Introduction = u.Introduction
		rets = append(rets, ret)
	}
	return rets
}

func CollectPost(DB *sql.DB, postId int, username string) bool {
	ok := QuerySingleCollect(postId, username)
	if ok == true {
		_, err := QueryArticleWithId(dao.DB, postId, username)
		if err != nil {
			return false
		}
		sqlstr := "insert into collect(postId,username) value(?,?)"
		_, err = DB.Exec(sqlstr, postId, username)
		if err != nil {
			return false
		}
	} else {
		return false
	}
	return true
}

func QueryUserCollect(DB *sql.DB, username string) []ReturnCollect {
	sqlstr := "select postId from collect where username = ?"
	rows, err := DB.Query(sqlstr, username)
	if err != nil {
		log.Println(err)
		return nil
	}
	var collectList []ReturnCollect
	defer rows.Close()
	for rows.Next() {
		var collect ReturnCollect
		err = rows.Scan(&collect.PostId)
		if err != nil {
			log.Println(err)
			return nil
		}
		article, err := QueryArticleWithId(dao.DB, collect.PostId, username)
		if err != nil {
			log.Println(err)
			return nil
		}
		collect.Title = article.Title
		collect.PublishTime = article.CreateTime
		collect.Username = article.Username
		collect.Avatar = article.Avatar
		collect.NickName = article.Author
		collectList = append(collectList, ReturnCollect{PostId: collect.PostId, Title: collect.Title, PublishTime: collect.PublishTime,
			Username: collect.Username, Avatar: collect.Avatar, NickName: collect.NickName})
	}
	return collectList
}

func QuerySingleCollect(id int, username string) bool {
	sqlstr := "select id from collect where postId=? and username=?"
	row := dao.DB.QueryRow(sqlstr, id, username)
	Id := 0
	err := row.Scan(&Id)
	if err != nil {
		return true
	}
	return false
}

//后续改成单独的用户操作
func deleteCollect(id int) (int64, error) {
	sqlstr := "delete from collect where postId=?"
	result, err := dao.DB.Exec(sqlstr, id)
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	return affected, err
}
