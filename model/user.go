package model

import (
	"blog/dao"
	"database/sql"
	"log"
	"time"
)

type User struct {
	id int
	Phone int
	Birthday string
	Introduction string
	Qq int
	Username string
	Password     string
	Email string
	CreateTime time.Time  `json:"createTime" form:"createTime"`
	IsAdmin       bool     //是否管理员
	AvatarUrl     string    //头像链接
	NickName      string
	Gender string
}

func QueryUserWithSql(DB *sql.DB, sqlstr string) int {
	rows, _ := dao.QueryRowsDB(DB, sqlstr)
	id := 0
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			continue
		}
		if id > 0 {
			break
		}
	}
	return id
}
func QueryIdWithUsername(DB *sql.DB, username string) int {
	sqlstr := "select id from user where username='" + username + "'"
	return QueryUserWithSql(DB, sqlstr)
}

func QueryIdWithUserPwd(DB *sql.DB, username string, password string) int {
	sqlstr := "select id from user where username='" + username + "' and password='" + password + "'"
	return QueryUserWithSql(DB, sqlstr)
}

func InsertUser(DB *sql.DB, user User) (int64, error) {
	return dao.ModifyDB(DB, "insert into user(username,password,create_time) values (?,?,?)",
		user.Username, user.Password, user.CreateTime)
}

func QueryUserInfo(DB *sql.DB,username string)User {
		id:=QueryIdWithUsername(DB,username)
		sqlstr:="select username,email,create_time,NickName,AvatarUrl,Gender,introduction,qq,birthday,phone from user where id = ?"
	var u1 User
	err := DB.QueryRow(sqlstr, id).Scan(&u1.Username, &u1.Email, &u1.CreateTime, &u1.NickName, &u1.AvatarUrl, &u1.Gender, &u1.Introduction, &u1.Qq, &u1.Birthday, &u1.Phone)
	if err != nil {
		log.Println(err)
		return User{}
	}
		u1.id=id
	return u1
}

func UpdateUserInfo(db *sql.DB,u1 User) (int64, error) {
	id:=QueryIdWithUsername(db,u1.Username)
	sqlstr:="update user set email=?,NickName=?,AvatarUrl=?,Gender=?,qq=?,birthday=?,introduction=?,phone=? where id = ?"
	return dao.ModifyDB(db,sqlstr,u1.Email,u1.NickName,u1.AvatarUrl,u1.Gender,u1.Qq,u1.Birthday,u1.Introduction,u1.Phone,id)
}
func ChangeUserPwd(db *sql.DB, username string, password string) (int64, error) {
	id:=QueryIdWithUsername(db,username)
	sqlstr:="update user set password=? where id=?"
	return dao.ModifyDB(db,sqlstr,password,id)
}

func QueryUserInfoWithId(id int) (string, string, string, string) {
	sqlstr:="select * from user where id=?"
	row:=dao.DB.QueryRow(sqlstr,id)
	var (
		username string
		avatarUrl string
		nickname string
		introduction string
	)
	row.Scan(&username,&avatarUrl,&nickname,&introduction)
	return username,avatarUrl,nickname,introduction
}

func QueryUserPwd(username string) string {
	sqlstr:="select password from user where username=?"
	row:=dao.DB.QueryRow(sqlstr,username)
	pwd:=""
	err := row.Scan(&pwd)
	if err != nil {
		log.Println(err)
		return ""
	}
	return pwd
}


