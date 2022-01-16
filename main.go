package main

import (
	"fmt"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	r.Run(":8080")
}

func main_test() {
	fmt.Println("Hello World!\n")
}
