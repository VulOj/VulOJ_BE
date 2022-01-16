package services

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var RedisClient *redis.Client
