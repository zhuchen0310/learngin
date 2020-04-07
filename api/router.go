package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuchen/learngin/app/activity"
	"github.com/zhuchen/learngin/app/follow"
	"github.com/zhuchen/learngin/app/todo_list"
)

// RouterRun 起 服务
func RouterRun() {
	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/api/v1")
	{
		_activity := v1.Group("/activity")
		_activity.GET("", activity.GetActivityInfoByID)
		_follow := v1.Group("/follow")
		_follow.GET("/:userHash", follow.UserFollowers)
		_follow.POST("", follow.Follow)
		_follow.DELETE("", follow.UnFollow)
		_todo := v1.Group("/todo")
		_todo.GET("/", todo_list.All)
		_todo.POST("/", todo_list.Add)
		_todo.GET("/:id", todo_list.Take)
		_todo.PUT("/:id", todo_list.Update)
		_todo.DELETE("/:id", todo_list.Del)
	}
	v2 := router.Group("/api/v2")
	{
		_todo := v2.Group("/todo")
		_todo.GET("/", todo_list.All)
		_todo.POST("/", todo_list.Add)
		_todo.GET("/:id", todo_list.Take)
		_todo.PUT("/:id", todo_list.Update)
		_todo.DELETE("/:id", todo_list.Del)
	}
	router.Run(":8080")
}
