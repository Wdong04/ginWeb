package main

import (
	"github.com/Wdong04/ginWeb/dao"
	"github.com/Wdong04/ginWeb/models"
	"github.com/Wdong04/ginWeb/routers"
)

func main() {
	// connect to db
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	// defer closing db
	// defer DB.Close()
	// 关联模型
	err = dao.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		panic(err)
	}
	// 注册路由
	r := routers.SetupRouters()
	// 启动服务，默认端口8080
	r.Run()
}
