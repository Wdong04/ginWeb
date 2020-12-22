package routers

import (
	"gin_demo/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouters 注册路由
func SetupRouters() *gin.Engine {
	// 获取gin默认框架
	r := gin.Default()
	// 加载静态文件
	r.Static("/static", "static")
	// 加载模板
	r.LoadHTMLGlob("templates/*")
	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", todo.CreateTodo)
		v1Group.GET("/todo", todo.GetTodo)
		v1Group.PUT("/todo/:id", todo.UpdateTodoByID)
		v1Group.DELETE("/todo/:id", todo.DeleteTodoByID)
	}
	return r
}
