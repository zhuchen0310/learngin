package todo_list

import (
	"net/http"
	"strconv"

	"github.com/zhuchen/learngin/database"

	"github.com/gin-gonic/gin"
)

const (
	JSON_SUCCESS int = 1
	JSON_ERROR   int = 0
)

// Add 创建TODO条目
func Add(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := database.TodoModel{Title: c.PostForm("title"), Completed: completed}
	db := database.NewMySQL()
	db.Save(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":     JSON_SUCCESS,
		"message":    "创建成功",
		"resourceId": todo.ID,
	})
}

// All 获取所有条目
func All(c *gin.Context) {
	var todos []database.TodoModel
	var _todos []database.TransformedTodo
	db := database.NewMySQL()
	db.Find(&todos)

	// 没有数据
	if len(todos) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "没有数据",
		})
		return
	}

	// 格式化
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, database.TransformedTodo{
			ID:        item.ID,
			Title:     item.Title,
			Completed: completed,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "ok",
		"data":    _todos,
	})

}

// Take 根据id获取一个条目
func Take(c *gin.Context) {
	var todo database.TodoModel
	todoID := c.Param("id")
	db := database.NewMySQL()
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  JSON_ERROR,
			"message": "条目不存在",
		})
		return
	}
	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := database.TransformedTodo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: completed,
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "ok",
		"data":    _todo,
	})
}

// Update 更新一个条目
func Update(c *gin.Context) {
	var todo database.TodoModel
	todoID := c.Param("id")
	db := database.NewMySQL()
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  JSON_ERROR,
			"message": "条目不存在",
		})
		return
	}

	db.Model(&todo).Update("title", c.PostForm("title"))
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	db.Model(&todo).Update("completed", completed)
	c.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "更新成功",
	})
}

// Del 删除条目
func Del(c *gin.Context) {
	var todo database.TodoModel
	todoID := c.Param("id")
	db := database.NewMySQL()
	db.First(&todo, todoID)
	if todo.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  JSON_ERROR,
			"message": "条目不存在",
		})
		return
	}
	db.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":  JSON_SUCCESS,
		"message": "删除成功!",
	})
}
