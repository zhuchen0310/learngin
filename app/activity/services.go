package activity

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zhuchen/learngin/common"
	"github.com/zhuchen/learngin/database"
	"github.com/zhuchen/learngin/redis"
)

// activityJSON 活动信息json
type activityJSON struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	SubTitle  string `json:"sub_title"`
	Desc      string `json:"desc"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	HashTag   string `json:"hash_tag"`
}

const ()

// activityKey 获取 缓存key
func activityKey(activityID int) string {
	return fmt.Sprintf("%s%d", "string:activity_", activityID)
}

// GetActivityInfoByID 根据活动id 获取活动详情
func GetActivityInfoByID(c *gin.Context) {

	activityID, err := strconv.Atoi(c.Query("activity_id"))
	if err != nil {
		c.JSON(403, gin.H{"mes": "params error"})
		return
	}
	fmt.Println(activityID)
	var activityjson activityJSON

	activityjson, ok := getActivityInfoFromCache(activityID)
	println(ok)
	if ok {
		fmt.Println("直接从缓存获取")
		c.JSON(200, activityjson)
		return
	}
	db := database.NewMySQL()
	conn := redis.NewRedis()
	var activity database.Activity

	db.Where(&database.Activity{ID: activityID}).First(&activity).Scan(&activityjson)

	key := activityKey(activityID)
	value, _ := common.JSONDumps(activityjson)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_ = conn.Set(key, value, time.Second*60).Err()
	c.JSON(200, activityjson)
}

// GetActivityInfoFromCache 从缓存中获取数据
func getActivityInfoFromCache(activityID int) (activityJSON, bool) {
	var activityjson activityJSON

	conn := redis.NewRedis()
	defer conn.Close()
	key := activityKey(activityID)
	value, err := conn.Get(key).Result()

	if value == "" {
		return activityjson, false
	}

	if err != nil {
		fmt.Println(err.Error())
		return activityjson, false
	}
	common.JSONLoads([]byte(value), &activityjson)
	return activityjson, true
}
