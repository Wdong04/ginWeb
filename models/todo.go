package models

import (
	"strconv"

	"github.com/Wdong04/ginWeb/dao"
	"gorm.io/gorm"
)

// Todo Model
type Todo struct {
	gorm.Model
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
func GetTodoByID(id string) (Todo, error) {
	var todo Todo
	ID, err := strconv.Atoi(id)
	if err != nil {
		return todo, err
	}
	err = dao.DB.Where("id=?", ID).First(&todo).Error
	return todo, err
}

// UpdateTodo 更新待办事项
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return err
}

// DeleteTodoByID 根据ID删除待办事项
func DeleteTodoByID(id string) error {
	ID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	err = dao.DB.Where("id=?", ID).Delete(Todo{}).Error
	return err
}
