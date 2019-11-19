package main

import (
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"kxagency_test/src/controller"
	"log"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	r1 := router.Group("/")
	{
		r1.POST("createinfo", controller.TaskCreateInfo)
		r1.POST("updateinfo", controller.TaskUpdateInfo)
		r1.POST("deleteinfo", controller.TaskDeleteInfo)
		r1.POST("paymentinfo", controller.TaskGrossInfo)
		r1.POST("createrenderframeinfo", controller.RenderFrameCreateInfo)
		r1.POST("updaterenderframeinfo", controller.RenderFrameUpdateInfo)
		r1.POST("useradd",controller.UserRegister)
		r1.POST("userupdate",controller.UserUpdate)
		r1.POST("userdelete",controller.UserDelete)
		r1.POST("taskids",controller.GetTaskidsByCreatetime)
		r1.POST("taskinfo",controller.GetTaskInfo)
		r1.POST("frameinfo",controller.GetFrameInfo)
		r1.POST("download",controller.GetFileUrl)
	}
	log.Println("经销商测试api 地址为:http://127.0.0.1:4456/")
	err := router.Run("127.0.0.1:4456")
	if err != nil {
		log.Println("无法启动经销商测试程序", err)
		panic(err)
	}
}
