package model

import (
	"blog/dao"
	"database/sql"
	"log"
	"strconv"
	"time"
)

type Comment struct {
	Id                string
	PostId            string
	PublishTime       time.Time
	Content           string
	Username          string
	Avatar            string
	Nickname          string
	ReplyUserId       string
	ReplyUserNickname string
	IsPraised         bool
	IsFocus           bool
	PraiseCount       int
	Picture           string
	Model             int
}

type RetComment struct {
	Id          string
	PostId      string
	PublishTime time.Time
	Content     string
	Username    string
	Avatar      string
	Nickname    string
	IsPraised   bool
	IsFocus     bool
	PraiseCount int
	Picture     []string
	Model       int
}

type SecondComment struct {
	Id                string
	PostId            string
	PublishTime       time.Time
	Content           string
	Username          string
	Avatar            string
	Nickname          string
	ReplyUserId       string
	ReplyUserNickname string
	IsPraised         bool
	IsFocus           bool
	PraiseCount       int
	Model             int
}

type PraiseComment struct {
	PostId   int
	username string
	IsPraise bool
}

type FocusComment struct {
	PostId   int
	username string
	IsPraise bool
}

type RetQueryComment struct {
	PostId int
	Username string
	Nickname string
}

func QueryIsPraise(postId int, username string) bool {
	sqlstr := "select id from commentPraise where commentId=? and username=?"
	row := dao.DB.QueryRow(sqlstr, postId, username)
	var praise PraiseComment
	err := row.Scan(&praise.PostId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func QueryIsFocus(postId int, username string) bool {
	sqlstr := "select id from commentFocus where postId=? and username=?"
	row := dao.DB.QueryRow(sqlstr, postId, username)
	var f FocusComment
	err := row.Scan(&f.PostId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func QueryAllPraiseComment(postId string) int {
	sqlstr := "select count(*) from commentpraise where commentId =?"
	rows, err := dao.DB.Query(sqlstr, postId)
	if err != nil {
		log.Println(err)
	}
	num := 0
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&num)
		if err != nil {
			log.Println(err)
			return 0
		}
	}
	return num
}

func NoTokenGetComment(DB *sql.DB, postId int, page int, size int) []Comment {
	var commentList []Comment
	count := 0
	max := 0
	sqlstr := "select * from comment where postId=?"
	rows, err := DB.Query(sqlstr, postId)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)
	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.Id, &comment.PostId,&comment.Content,&comment.Picture,&comment.Model,&comment.PublishTime ,&comment.Username, &comment.Avatar,
			&comment.Nickname)
		if err != nil {
			log.Println(err)
			return nil
		}
		comment.PraiseCount = QueryAllPraiseComment(comment.Id)
		comment.IsPraised = false
		comment.IsFocus = false
		comments = append(comments, comment)
		count++
	}
	if count < page*size {
		max = count
	} else {
		max = page * size
	}
	for i := (page - 1) * size; i < max; i++ {
		commentList = append(commentList, comments[i])
	}

	return commentList
}

func GetComment(DB *sql.DB, postId int, page int, size int,username string) []Comment {
	var commentList []Comment
	count := 0
	max := 0
	sqlstr := "select * from comment where postId=?"
	rows, err := DB.Query(sqlstr, postId)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println(err)
		}
	}(rows)
	var comments []Comment
	for rows.Next() {
		var comment Comment
		rows.Scan(&comment.Id, &comment.PostId, &comment.Content, &comment.Picture, &comment.Model, &comment.PublishTime, &comment.Username, &comment.Avatar,
			&comment.Nickname)
		ID, err2 := strconv.Atoi(comment.Id)
		if err2 != nil {
			return nil
		}
		comment.PraiseCount = QueryAllPraiseComment(comment.Id)
		comment.IsPraised = QueryIsPraise(ID, username)
		comment.IsFocus = QueryIsFocus(ID, username)
		comments = append(comments, comment)
		count++
	}
	if count < page*size {
		max = count
	} else {
		max = page * size
	}
	for i := (page - 1) * size; i < max; i++ {
		commentList = append(commentList, comments[i])
	}
	return commentList
}

func InsertComment(DB *sql.DB, postId int, content string, photo string, model int, username string) (int64, error) {
	sqlstr2 := "select AvatarUrl, NickName from user where username=?"
	row := DB.QueryRow(sqlstr2, username)
	avatar := ""
	nickname := ""
	Time := time.Now()
	err := row.Scan(&avatar, &nickname)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	sqlstr := "insert into comment(PostId,Content,Picture,Model,Username,Avatar,Nickname,PublishTime) values(?,?,?,?,?,?,?,?)"
	result, err := DB.Exec(sqlstr, postId, content, photo, model, username, avatar, nickname, Time)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	cid, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return cid, err
}

func UpdateComment(DB *sql.DB, id int, content string, photo string) (int64, error) {
	sqlstr := "update comment set Content=?,Picture=? where id=?"
	result, err := DB.Exec(sqlstr, content, photo, id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, err
}

func DeleteCommentWithId(DB *sql.DB, id int,username string) (int64, bool) {
	sqlstr := "select Model,Username from comment where id = ?"
	row := DB.QueryRow(sqlstr, id)
	var model int
	var UserId string
	err := row.Scan(&model,UserId)
	if err != nil {
		log.Println(err)
		return 0, false
	}
	if UserId!=username {
		return 0, false
	}
	if model == 2 {
		sqlstr := "delete from comment where id=?"
		result, err := DB.Exec(sqlstr, id)
		if err != nil {
			log.Println(err)
			return 0, false
		}
		count, err := result.RowsAffected()
		if err != nil {
			log.Println(err)
			return 0, false
		}
		return count, true
	} else {
		if model == 1 {
			sqlstr := "delete from comment where postId=?"
			result, err := DB.Exec(sqlstr, id)
			if err != nil {
				log.Println(err)
				return 0, false
			}
			count, err := result.RowsAffected()
			if err != nil {
				log.Println(err)
				return 0, false
			}
			sqlstr2 := "delete from comment where id=?"
			_, err = DB.Exec(sqlstr2, id)
			if err != nil {
				log.Println(err)
				return 0, false
			}
			return count, true
		}
	}
	return 0, true
}
func DeleteArticleCommentWithId(DB *sql.DB, PostId int,username string) (int64, bool) {
	var count int64
	var ok bool
	sqlstr := "select id from comment where PostId = ?"
	rows, err := DB.Query(sqlstr, PostId)
	if err != nil {
		log.Println(err)
		return 0, false
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			return 0, false
		}
		_, err = DeleteCommentPraise(id)
		if err != nil {
			log.Println(err)
			return 0, false
		}
		count,ok= DeleteCommentWithId(dao.DB, id,username)
		if ok==false {
			return 0, false
		}
	}
	return count, true
}

func QueryCommentWithId(id int) (error,RetQueryComment) {
	sqlstr := "select PostId,Username,Nickname from comment where id=?"
	row:= dao.DB.QueryRow(sqlstr, id)
	 var retCom RetQueryComment
	err := row.Scan(&retCom.PostId,&retCom.Username,&retCom.Nickname)
	return err,retCom
}

func DeleteCommentPraise(id int) (int64, error) {
	sqlstr := "delete from commentpraise where commentId=?"
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
