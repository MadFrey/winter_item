package model

import (
	"database/sql"
	"fmt"
	"log"
)

type Topic struct {
	TopicId int
	LogoUrl string
	TopicName string
	Introduction string
}

func QueryAllTopicFromDb(DB *sql.DB) []Topic {
	sqlstr:="select * from topic where topicId > ?"
	 rows,err:=DB.Query(sqlstr,0)
	if err != nil {
		log.Println(err)
		return nil
	}
	var topics []Topic
	defer rows.Close()
	 for rows.Next(){
		 var topic Topic
		 err = rows.Scan(&topic.TopicId, &topic.LogoUrl, &topic.TopicName, &topic.Introduction)
		 if err != nil {
			 log.Println(err)
			 return nil
		 }
		 fmt.Println(topic)
		 topics=append(topics,topic)
	 }
	return topics
}
