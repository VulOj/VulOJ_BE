package models

import (
	"time"
)

type Blog struct {
	ID               uint `gorm:"primary_key"`
	PublishTimestamp time.Time
	AuthEmail        string
	Title            string
	Content          string //文章内容
	OrganizationType string //文章所属组织类型，可以为企业(enterprise)或项目(project)
	OrganizationID   uint
	OrganizationName string
}

type BlogForbidden struct {
	Blog      `gorm:"embedded"`
	IsDeleted bool `gorm:"default:false"`
}
