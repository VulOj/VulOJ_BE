package models

import "time"

type Comment struct {
	ID               uint `gorm:"primary_key"`
	BlogID           uint
	Content          string
	AuthEmail        string    //评论者的email
	AuthName         string    //评论者的邮箱
	PublishTimestamp time.Time //评论的时间
}

//帖子下面的评论相关
