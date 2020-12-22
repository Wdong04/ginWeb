package main

import (
	"github.com/Wdong04/ginWeb/dao"
	"github.com/Wdong04/ginWeb/models"
	"github.com/Wdong04/ginWeb/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// connect to db
	dao.InitMySQL()
	// 关联模型
	err := dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
	// 获取默认gin引擎
	r := gin.Default()
	// 加载渲染模板
	routers.LoadTemplates(r)
	// 注册路由
	routers.SetupRouters(r)
	// 启动服务，默认端口8080
	r.Run()
}
