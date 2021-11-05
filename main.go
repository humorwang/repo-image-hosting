package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"repo-image-hosting/config"
	"repo-image-hosting/routes"
	"runtime"
)

func init() {
	config.Setup()

	if config.Conf.App.Token == "" {
		panic("token 必须！")
	}

}

func main() {
	// 关闭debug模式
	gin.SetMode(gin.ReleaseMode)

	port := config.Conf.App.Port
	router := routes.InitRoute()

	link := "http://127.0.0.1:" + port
	log.Println("监听端口", link, " 请不要关闭终端")
	// 调用浏览器打开网页
	if runtime.GOOS == "windows" {
		//exec.Command("cmd", "/c", "start", link).CombinedOutput()
	} else {
		// 没有测试环境，暂时不实现
		//exec.Command("open", link).CombinedOutput() // macos
		//exec.Command("x-www-browser", link).CombinedOutput() // linux
	}

	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}
