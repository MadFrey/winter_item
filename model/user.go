package model

import (
	"blog/dao"
	"database/sql"
	"log"
	"time"
)

type User struct {
	Id           int       `gorm:"column:id"`
	Phone        int       `gorm:"column:phone"`
	Birthday     string    `gorm:"column:birthday"`
	Introduction string    `gorm:"column:introduction"`
	Qq           int       `gorm:"column:qq"`
	Username     string    `gorm:"column:username"`
	Password     string    `gorm:"column:password"`
	Email        string    `gorm:"column:email"`
	CreateTime   time.Time `gorm:"create_time" json:"createTime" form:"createTime"`
	IsAdmin      bool      `gorm:"column:IsAdmin"`
	NickName     string    `gorm:"column:NickName"`
	Gender       string    `gorm:"column:Gender"`
	AvatarUrl    string    `gorm:"column:AvatarUrl"`
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
func QueryIdWithUsername(username string) int {
	db, err := dao.OpenGormLink()
	if err != nil {
		log.Println(err)
		return 0
	}
	var user User
	db.Where("username=?", username).First(&user)
	return user.Id
}

func DeleteUser(username string, password string) int64 {
	db, err := dao.OpenGormLink()
	if err != nil {
		log.Println(err)
		return 0
	}
	var user User
	db.Where("username=? AND password=?", username, password).First(&user)
	result := db.Delete(&user) //db.where().delete()
	return result.RowsAffected
}

func InsertUser(DB *sql.DB, user User) (int64, error) {
	return dao.ModifyDB(DB, "insert into user(username,password,create_time) values (?,?,?)",
		user.Username, user.Password, user.CreateTime)
}

func QueryUserInfo(DB *sql.DB, username string) User {
	id := QueryIdWithUsername(username)
	sqlstr := "select username,email,create_time,NickName,AvatarUrl,Gender,introduction,qq,birthday,phone from user where id = ?"
	var u1 User
	err := DB.QueryRow(sqlstr, id).Scan(&u1.Username, &u1.Email, &u1.CreateTime, &u1.NickName, &u1.AvatarUrl, &u1.Gender, &u1.Introduction, &u1.Qq, &u1.Birthday, &u1.Phone)
	if err != nil {
		log.Println(err)
		return User{}
	}
	u1.Id = id
	return u1
}

func UpdateUserInfo(db *sql.DB, u1 User) (int64, error) {
	id := QueryIdWithUsername(u1.Username)
	sqlstr := "update user set email=?,NickName=?,AvatarUrl=?,Gender=?,qq=?,birthday=?,introduction=?,phone=? where id = ?"
	return dao.ModifyDB(db, sqlstr, u1.Email, u1.NickName, u1.AvatarUrl, u1.Gender, u1.Qq, u1.Birthday, u1.Introduction, u1.Phone, id)
}
func ChangeUserPwd(username string, password string) (int64, error) {
	db, err := dao.OpenGormLink()
	if err != nil {
		log.Println(err)
		return 0, nil
	}
	result := db.Where("username=?", username).Update("password", password)
	rows := result.RowsAffected
	return rows, err
}

func QueryUserInfoWithId(id int) (string, string, string, string) {
	sqlstr := "select * from user where id=?"
	row := dao.DB.QueryRow(sqlstr, id)
	var (
		username     string
		avatarUrl    string
		nickname     string
		introduction string
	)
	row.Scan(&username, &avatarUrl, &nickname, &introduction)
	return username, avatarUrl, nickname, introduction
}

func QueryUserPwd(username string) string {
	sqlstr := "select password from user where username=?"
	row := dao.DB.QueryRow(sqlstr, username)
	pwd := ""
	err := row.Scan(&pwd)
	if err != nil {
		log.Println(err)
		return ""
	}
	return pwd
}

func QueryUserGitId(DB *sql.DB, id int) error {
	sqlstr := "select id from user where gitId=?"
	row := DB.QueryRow(sqlstr, id)
	var userId int
	err := row.Scan(&userId)
	if err != nil {
		return err
	}
	return nil
}

func ExecUserInfoWithGit(DB *sql.DB, username string, avatar string, nickname string, gitId int) error {
	sqlstr := "update user set gitId=?,NickName=?,AvatarUrl=? where username=?"
	_, err := DB.Exec(sqlstr, gitId, nickname, avatar, username)
	if err != nil {
		log.Println(err)
	}
	return err
}

func QueryUsernameWithGitId(DB *sql.DB, id int) (string, error) {
	sqlstr := "select username from user where gitId=?"
	var username string
	err := DB.QueryRow(sqlstr, id).Scan(&username)
	return username, err
}
