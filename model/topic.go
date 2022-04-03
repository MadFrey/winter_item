package model

import (
	"blog/dao"
	"log"
)

type Topic struct {
	TopicId      int    `gorm:"column:topicId"`
	LogoUrl      string `gorm:"column:logoUrl"`
	TopicName    string `gorm:"column:topicName"`
	Introduction string `gorm:"column:introduction"`
}

func QueryAllTopicFromDb() []Topic {
	db, err := dao.OpenGormLink()
	if err != nil {
		log.Println(err)
		return nil
	}

	var topics []Topic
	db.Find(&topics)
	return topics
}
