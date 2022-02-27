package router

import (
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/directory"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/models"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/services"
	"github.com/VulOJ/Vulnerable_Online_Judge_Project/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
)

func Getwd() (dir string) {
	dir, _ = os.Getwd()
	return dir
}

func DownloadDirectory(c *gin.Context) {
	dirName := c.Param("dir_name")
	switch dirName {
	case directory.VULHUB_NAME:
		err := DownloadVulhub()
		if err != nil {
			c.Abort()
			c.JSON(http.StatusBadRequest, gin.H{`msg`: "安装失败"})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "安装成功",
			})
		}
	default:
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "暂无该环境 尽请期待！",
		})
	}
	return
}
func DownloadVulhub() (err error) {

	//fmt.Println(directory.VULHUB_URL)
	//fmt.Println(directory.VULHUB_GIT)
	//path := Getwd()
	//fmt.Println(path)
	gitClone := exec.Command(directory.GIT_CLONE, directory.VULHUB_GIT)
	gitClone.Stdout = os.Stdout
	err = gitClone.Run()
	if err != nil {
		return
	}
	//Download Vulhub successfully

	//insert Vulhub into database
	temp := models.Directories{
		Name:                directory.VULHUB_NAME,
		Path:                Getwd() + "/directory/" + directory.VULHUB_NAME,
		Status:              true,
		DownloadedTimestamp: util.GetTimeStamp(),
	}
	services.CreateDirectory(temp)
	return
}
