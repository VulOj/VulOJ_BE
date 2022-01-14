package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Run(":8080")
}

func main_test() {
	fmt.Println("Hello World!\n")
}
