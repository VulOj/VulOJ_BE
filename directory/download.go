package directory

import (
	"fmt"
	"os"
	"os/exec"
)

func Getwd() (dir string) {
	dir, _ = os.Getwd()
	return dir
}

func DownloadVulhub() {
	fmt.Println(VULHUB_URL)
	fmt.Println(VULHUB_GIT)
	path := Getwd()
	fmt.Println(path)
	mkdir := exec.Command("mkdir", VULHUB_NAME)
	mkdir.Stdout = os.Stdout
	_ = mkdir.Run()
}
