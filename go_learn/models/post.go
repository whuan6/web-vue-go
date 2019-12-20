package models

import (
	db "go_learn/database"
	"log"
)

type MPost struct {
	Id         int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Title      string `gorm:"column:title" json:"title"`
	Content    string `gorm:"column:content" json:"content"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	LikeCount  int32  `gorm:"column:like_count" json:"like_count"`
}


func GetPosts() []*MPost {
	var post []*MPost
	err := db.ORM.Order("like_count desc").Find(&post).Error
	if err != nil {
		log.Println("ERROR:", err)
		return nil
	}
	return post
}

