package models

import "time"

type ClassToUser struct {
	ID              uint `gorm:"primary_key"`
	ClassID         uint
	UserEmail       string
	JoinAtTimestamp time.Time //用户加入组织的时间
	ClassType       string
}

//这个表主要用于建立组织和用户间的多对多关系
