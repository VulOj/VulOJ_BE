package router

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/directory"
	"github.com/gin-gonic/gin"
	"os"
	"os/exec"
)

func Getwd() (dir string) {
	dir, _ = os.Getwd()
	return dir
}

func DownloadVulhub(c gin.Context) {

	//fmt.Println(directory.VULHUB_URL)
	//fmt.Println(directory.VULHUB_GIT)
	//path := Getwd()
	//fmt.Println(path)
	gitClone := exec.Command(directory.GIT_CLONE, directory.VULHUB_GIT)
	gitClone.Stdout = os.Stdout
	err := gitClone.Run()
	if err != nil {
		//Download Vulhub successfully

		//insert Vulhub into database

	} else {

	}
}
