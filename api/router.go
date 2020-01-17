package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhuchen/learngin/app/activity"
	"github.com/zhuchen/learngin/app/follow"
)

// RouterRun 起 服务
func RouterRun() {
	router := gin.Default()
	// Simple group: v1
	v1 := router.Group("/v1")
	{
		_activity := v1.Group("/activity")
		_activity.GET("", activity.GetActivityInfoByID)
		_follow := v1.Group("/follow")
		_follow.GET("/:userHash", follow.UserFollowers)
		_follow.POST("", follow.Follow)
		_follow.DELETE("", follow.UnFollow)
	}
	router.Run(":8080")
}
