package model

import (
	"blog/dao"
	"blog/util"
	"database/sql"
	"log"
	"strconv"
	"time"
)

type Article struct {
	Id          int
	TopicId     string
	Title       string
	Content     string
	Pictures    string
	Author      string
	Username    string
	CreateTime  time.Time
	Avatar      string
	IsPraised   bool
	IsFocus     bool
	FocusCount  int
	PraiseCount int
}

type RetArticle struct {
	Id          int
	TopicId     string
	Title       string
	Content     string
	Pictures    []string
	Author      string
	Username    string
	CreateTime  time.Time
	Avatar      string
	IsPraised   bool
	IsFocus     bool
	FocusCount  int
	PraiseCount int
}

type Focus struct {
	ArticleId int
	username  string
	IsFocus   bool
}
type Praise struct {
	ArticleId int
	username  string
	IsPraise  bool
}

func QueryFocusCount(id int) int {
	sqlstr := "select count(*) from focus where ArticleId=?"
	num := 0
	totalRow, err := dao.DB.Query(sqlstr, id)
	if err != nil {
		log.Println(err)
	}
	for totalRow.Next() {
		err := totalRow.Scan(&num)
		if err != nil {
			log.Println(err)
			return 0
		}
	}
	return num
}

func QueryPraiseCount(id int) int {
	sqlstr := "select count(*) from praise where postId=?"
	num := 0
	totalRow, err := dao.DB.Query(sqlstr, id)
	if err != nil {
		log.Println(err)
	}
	for totalRow.Next() {
		err := totalRow.Scan(&num)
		if err != nil {
			log.Println(err)
			return 0
		}
	}
	return num
}

func QueryArticleFocus(id int, username string) bool {
	println(id, username)
	sqlstr := "select id from focus where username=? and ArticleId=?"
	var f Focus
	err := dao.DB.QueryRow(sqlstr, username, id).Scan(&f.ArticleId)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func QueryArticlePraise(id int, username string) bool {
	sqlstr := "select postId from praise where username = ? and postId=?"
	var p Praise
	err := dao.DB.QueryRow(sqlstr, username, id).Scan(&p.ArticleId)
	if err != nil {
		return false
	}
	return true
}

func QueryArticleWithId(DB *sql.DB, id int, username string) (Article, error) {
	sqlstr := "select Title,Author,Pictures,Content,CreateTime,Avatar,Username,TopicId from article where id = ?"
	article := Article{}
	row := DB.QueryRow(sqlstr, id)
	err := row.Scan(&article.Title, &article.Author, &article.Pictures, &article.Content, &article.CreateTime,
		&article.Avatar, &article.Username, &article.TopicId)
	if err != nil {
		log.Println(err)
		return Article{}, err
	}
	article.Id = id
	article.IsFocus = QueryArticleFocus(article.Id, username)
	article.IsPraised = QueryArticlePraise(article.Id, username)
	article.FocusCount = QueryFocusCount(article.Id)
	article.PraiseCount = QueryPraiseCount(article.Id)
	article = Article{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Content, Pictures: article.Pictures,
		Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar, PraiseCount: article.PraiseCount,
		FocusCount: article.FocusCount, IsFocus: article.IsFocus, IsPraised: article.IsPraised}
	return article, err
}
func NoTokenQueryArticleWithId(DB *sql.DB, id int) Article {
	sqlstr := "select Title,Author,Pictures,Content,CreateTime,Avatar,Username,TopicId from article where id = ?"
	article := Article{}
	row := DB.QueryRow(sqlstr, id)
	err := row.Scan(&article.Title, &article.Author, &article.Pictures, &article.Content, &article.CreateTime,
		&article.Avatar, &article.Username, &article.TopicId)
	if err != nil {
		log.Println(err)
		return Article{}
	}
	article.Id = id
	article.IsPraised = false
	article.IsFocus = false
	article.FocusCount = QueryFocusCount(article.Id)
	article.PraiseCount = QueryPraiseCount(article.Id)
	article = Article{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Content, Pictures: article.Pictures,
		Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar, PraiseCount: article.PraiseCount,
		FocusCount: article.FocusCount, IsFocus: article.IsFocus, IsPraised: article.IsPraised}
	return article
}

func QueryArticleWithPage(DB *sql.DB, pageNum int, size int, username string) ([]Article, error) {
	if pageNum <= 0 {
		pageNum = 1
	}
	fromIndexStr := strconv.Itoa((pageNum - 1) * size)
	pageNumStr := strconv.Itoa(size)
	sqlstr := "select * from article order by CreateTime desc limit " + fromIndexStr + ", " + pageNumStr
	rows, err := dao.QueryRowsDB(DB, sqlstr)
	var articles []Article
	defer rows.Close()
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.TopicId, &article.Title, &article.Content, &article.Pictures, &article.Author,
			&article.Username, &article.CreateTime, &article.Avatar)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		article.IsFocus = QueryArticleFocus(article.Id, username)
		article.IsPraised = QueryArticlePraise(article.Id, username)
		article.FocusCount = QueryFocusCount(article.Id)
		article.PraiseCount = QueryPraiseCount(article.Id)
		articles = append(articles, Article{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Author, Pictures: article.Pictures,
			Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar, PraiseCount: article.PraiseCount, FocusCount: article.FocusCount, IsFocus: article.IsFocus, IsPraised: article.IsPraised})
	}
	return articles, err
}

func NoTokenQueryArticleWithPage(DB *sql.DB, pageNum int, size int) ([]Article, error) {
	if pageNum <= 0 {
		pageNum = 1
	}
	fromIndexStr := strconv.Itoa((pageNum - 1) * size)
	pageNumStr := strconv.Itoa(size)
	sqlstr := "select * from article order by CreateTime desc limit " + fromIndexStr + ", " + pageNumStr
	rows, err := dao.QueryRowsDB(DB, sqlstr)
	var articles []Article
	defer rows.Close()
	for rows.Next() {
		var article Article
		article.IsPraised = false
		article.IsFocus = false
		rows.Scan(&article.Id, &article.Title, &article.Author, &article.Pictures, &article.Content, &article.CreateTime,
			&article.Avatar, &article.Username, &article.TopicId)
		article.FocusCount = QueryFocusCount(article.Id)
		article.PraiseCount = QueryPraiseCount(article.Id)
		articles = append(articles, Article{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Author, Pictures: article.Pictures,
			Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar, IsFocus: article.IsFocus, IsPraised: article.IsPraised, PraiseCount: article.PraiseCount, FocusCount: article.FocusCount})
	}
	return articles, err
}

func UpdateArticle(DB *sql.DB, article RetArticle) (int64, error) {
	sqlstr := "update article set Title=?,TopicId=?,Pictures=?,Content=? where id=?"
	imageString := util.ArrayToString(article.Pictures)
	return dao.ModifyDB(DB, sqlstr, article.Title, article.TopicId, imageString, article.Content, article.Id)
}

func InsertArticle(DB *sql.DB, article Article) (int64, error) {
	sqlstr := "insert into article(title,content,Author,CreateTime,pictures,Avatar,username,topicId) values(?,?,?,?,?,?,?,?)"
	id, err := dao.ModifyDBID(DB, sqlstr, article.Title, article.Content, article.Author, article.CreateTime, article.Pictures, article.Avatar,
		article.Username, article.TopicId)
	if err != nil {
		log.Fatal(err)
	}
	return id, err
}

//DeleteArticleWithId 改用事务对多条数据进行处理
func DeleteArticleWithId(DB *sql.DB, id int, username string) (int64, bool) {
	ar := NoTokenQueryArticleWithId(DB, id)
	tx, err := dao.OpenTransaction()
	if err != nil {
		return 0, false
	}
	if ar.Username != username {
		return 0, false
	}
	sqlstr := "delete from article where id = ?"
	result, err := tx.Exec(sqlstr, id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Println(err)
			return 0, false
		}
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return 0, false
	}

	sqlstr2 := "delete from focus where ArticleId=?"
	_, err = tx.Exec(sqlstr2, id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Println(err)
			return 0, false
		}
	}
	sqlstr3 := "delete from praise where postId=?"
	_, err = tx.Exec(sqlstr3, id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Println(err)
			return 0, false
		}
	}
	sqlstr4 := "delete from collect where postId=?"
	_, err = tx.Exec(sqlstr4, id)
	if err != nil {
		err := tx.Rollback()
		if err != nil {
			log.Println(err)
			return 0, false
		}
	}
	_, ok := DeleteArticleCommentWithId(DB, id, username)
	if ok == false {
		return 0, false
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return 0, false
	}
	return affected, true
}

func QueryArticleWithKey(DB *sql.DB, key string, page int, size int) []Article {
	sqlstr := "select * from article where Content like ?"
	rows, err := DB.Query(sqlstr, "%"+key+"%")
	if err != nil {
		log.Println(err)
	}
	var (
		count    = 0
		max      int
		articles []Article
	)
	defer rows.Close()
	for rows.Next() {
		var article Article
		rows.Scan(&article.Id, &article.Title, &article.Author, &article.Pictures, &article.Content, &article.CreateTime,
			&article.Avatar, &article.Username, &article.TopicId)
		articles = append(articles, Article{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Author, Pictures: article.Pictures,
			Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar})
		count++
	}
	var searchList []Article
	if count < page*size {
		max = count
	} else {
		max = page * size
	}
	for i := (page - 1) * size; i < max; i++ {
		searchList = append(searchList, articles[i])
	}
	return searchList
}

func DeleteArticleFocus(postId int) (int64, error) { //后续改成单独的用户操作
	sqlstr := "delete from focus where ArticleId=?"
	result, err := dao.DB.Exec(sqlstr, postId)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return affected, err
}

func DeleteArticlePraise(postId int) (int64, error) { //后续改成单独的用户操作
	sqlstr := "delete from praise where postId=?"
	result, err := dao.DB.Exec(sqlstr, postId)
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
