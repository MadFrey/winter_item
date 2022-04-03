package service

import (
	"blog/dao"
	"blog/model"
	"log"
	"strings"
	"unicode/utf8"
)

func NoTokenMakeArticleList(PageNum int, size int) []model.RetArticle {
	articles, _ := model.NoTokenQueryArticleWithPage(dao.DB, PageNum, size)
	var articleList []model.RetArticle
	for _, article := range articles {
		if utf8.RuneCountInString(article.Content) > 500 {
			article.Content = string([]rune(article.Content)[0:100]) + "..."
		}
		newPicture := strings.Split(article.Pictures, ",")
		articleList = append(articleList, model.RetArticle{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Content, Pictures: newPicture,
			Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar})
	}
	return articleList
}

func MakeArticleList(PageNum int, size int, username string) []model.RetArticle {
	articles, _ := model.QueryArticleWithPage(dao.DB, PageNum, size, username)
	var articleList []model.RetArticle
	for _, article := range articles {
		if utf8.RuneCountInString(article.Content) > 500 {
			article.Content = string([]rune(article.Content)[0:100]) + "..."
		}
		newPicture := strings.Split(article.Pictures, ",")
		articleList = append(articleList, model.RetArticle{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Content, Pictures: newPicture,
			Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar})
	}
	return articleList
}

func AddArticleProcess(article model.Article) (int64, error) {
	return model.InsertArticle(dao.DB, article)
}

func DeleteArticleProcess(id int, username string) (int64, bool) {
	return model.DeleteArticleWithId(dao.DB, id, username)
}

func UpdateArticleProcess(article model.RetArticle) (int64, error) {
	return model.UpdateArticle(dao.DB, article)
}

func QuerySingleArticleProcess(id int, username string) (model.RetArticle, error) {
	article, err := model.QueryArticleWithId(dao.DB, id, username)
	if err != nil {
		log.Println(err)
		return model.RetArticle{}, err
	}
	newPicture := strings.Split(article.Pictures, ",")
	newArticle := model.RetArticle{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Content, Pictures: newPicture,
		Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar, PraiseCount: article.PraiseCount,
		FocusCount: article.FocusCount, IsFocus: article.IsFocus, IsPraised: article.IsPraised}
	return newArticle, err
}

func NoTokenQuerySingleArticleProcess(id int) model.RetArticle {
	article := model.NoTokenQueryArticleWithId(dao.DB, id)
	newPicture := strings.Split(article.Pictures, ",")
	newArticle := model.RetArticle{Id: article.Id, TopicId: article.TopicId, Title: article.Title, Content: article.Content, Pictures: newPicture,
		Author: article.Author, Username: article.Username, CreateTime: article.CreateTime, Avatar: article.Avatar, PraiseCount: article.PraiseCount,
		FocusCount: article.FocusCount, IsFocus: article.IsFocus, IsPraised: article.IsPraised}
	return newArticle
}

func SearchArticlesProcess(key string, page int, size int) []model.Article {
	return model.QueryArticleWithKey(dao.DB, key, page, size)
}

func QueryAllTopicsProcess() []model.Topic {
	return model.QueryAllTopicFromDb()
}
