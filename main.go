package main

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/router"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	router.DownloadVulhub()
	r.Run(":8080")
}
