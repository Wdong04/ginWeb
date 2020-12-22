package routers

import (
	"net/http"

	"github.com/Wdong04/ginWeb/controllers"
	"github.com/gin-gonic/gin"
)

// LoadTemplates 加载模板
func LoadTemplates(engine *gin.Engine) {
	// 加载静态文件
	engine.Static("/static", "static")
	// 加载模板
	engine.LoadHTMLGlob("templates/*")
}

// SetupRouters 注册路由
func SetupRouters(engine *gin.Engine) {
	// 注册路由
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	v1Group := engine.Group("v1")
	{
		v1Group.POST("/todo", controllers.CreateTodo)
		v1Group.GET("/todo", controllers.GetTodo)
		v1Group.PUT("/todo/:id", controllers.UpdateTodoByID)
		v1Group.DELETE("/todo/:id", controllers.DeleteTodoByID)
	}
}
