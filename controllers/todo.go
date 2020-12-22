package controllers

import (
	"net/http"

	"github.com/Wdong04/ginWeb/models"
	"github.com/gin-gonic/gin"
)

// CreateTodo 创建待办事项
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	if err := models.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

// GetTodo 获取待办事项
func GetTodo(c *gin.Context) {
	if todoList, err := models.GetTodo(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

// UpdateTodoByID 根据ID更新待办事项
func UpdateTodoByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效ID"})
		return
	}
	todo, err := models.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(&todo)
	if err := models.UpdateTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// DeleteTodoByID 根据ID删除待办事项
func DeleteTodoByID(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效ID"})
		return
	}
	if err := models.DeleteTodoByID(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "删除成功"})
}
