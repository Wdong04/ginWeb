package models

import (
	"github.com/Wdong04/ginWeb/dao"
)

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// CreateTodo 创建待办事项
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return err
}

// GetTodo 获取待办事项
func GetTodo() (todoList []Todo, err error) {
	err = dao.DB.Find(&todoList).Error
	return todoList, err
}

// GetTodoByID 根据ID获取待办事项
func GetTodoByID(id string) (todo Todo, err error) {
	err = dao.DB.Where("id=?", id).First(&todo).Error
	return todo, err
}

// UpdateTodo 更新待办事项
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return err
}

// DeleteTodoByID 根据ID删除待办事项
func DeleteTodoByID(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return err
}
