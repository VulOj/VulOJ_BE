package main

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/router"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}
