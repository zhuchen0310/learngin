package follow

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhuchen/learngin/database"
)

// followData 关注Api 请求参数
type followData struct {
	RequestUserHash string `json:"requester_user_hash" form:"requester_user_hash"`
	OwnerUserHash   string `json:"owner_user_hash" form:"owner_user_hash"`
}

// Follow 关注
func Follow(c *gin.Context) {
	data := followData{}
	c.BindJSON(&data)
	if data.RequestUserHash == "" || data.OwnerUserHash == "" || data.RequestUserHash == data.OwnerUserHash {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "param is error",
		})
		return
	}
	db := database.NewMySQL()
	var followRelationships database.FollowRelationships
	db.FirstOrCreate(&followRelationships, database.FollowRelationships{FolllowerUserHash: data.RequestUserHash, FollloweeUserHash: data.OwnerUserHash})
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": map[string]interface{}{
			"requester_user_hash": followRelationships.FolllowerUserHash,
			"owner_user_hash":     followRelationships.FollloweeUserHash,
		},
	})
}

// UnFollow 取消关注
func UnFollow(c *gin.Context) {
	data := followData{}
	c.BindJSON(&data)
	if data.RequestUserHash == "" || data.OwnerUserHash == "" || data.RequestUserHash == data.OwnerUserHash {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "param is error",
		})
		return
	}
	db := database.NewMySQL()
	ret := new([]interface{})
	db.Raw("Delete from follow_relationships Where follower_user_hash = ? and followee_user_hash = ?", data.RequestUserHash, data.OwnerUserHash).Scan(ret)
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}

// UserFollowers 获取 用户粉丝
func UserFollowers(c *gin.Context) {
	userHash := c.Param("userHash")
	db := database.NewMySQL()
	var followRelationshipses []*database.FollowRelationships
	db.Where(&database.FollowRelationships{FollloweeUserHash: userHash}).Select("follower_user_hash").Find(&followRelationshipses)
	var followRelationshipsList []string
	for _, f := range followRelationshipses {
		followRelationshipsList = append(followRelationshipsList, f.FolllowerUserHash)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":       "success",
		"followers": followRelationshipsList,
	})
}
