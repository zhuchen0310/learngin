package database

import (
	"time"
)

// LocalPush 本地push 结构体
type LocalPush struct {
	ID              int    `gorm:"Column:id"`
	PushType        string `gorm:"Column:push_type"`
	PushExecuteType string `gorm:"Column:push_execute_type"`
	PushExecuteTime string `gorm:"Column:push_execute_time"`
	Title           string `gorm:"Column:title"`
	Body            string `gorm:"Column:body"`
	DeepLink        string `gorm:"Column:deep_link"`
	LargeIcon       string `gorm:"Column:large_icon"`
	Active          int    `gorm:"Column:active"`
	CreatedAt       string `gorm:"Column:created_at"`
	UpdatedAt       string `gorm:"Column:last_edited_at"`
	ImageURL        string `gorm:"Column:image_url"`
}

// TableName 本地push 表名称
func (LocalPush) TableName() string {
	return "local_push_config"
}

// Activity 活动配置表
type Activity struct {
	ID        int    `gorm:"Column:id"`
	Title     string `gorm:"Column:title"`
	SubTitle  string `gorm:"Column:sub_title"`
	Desc      string `gorm:"Column:desc"`
	StartTime string `gorm:"start_time"`
	EndTime   string `gorm:"Column:end_time"`
	HashTag   string `gorm:"Column:hash_tag"`
	Active    int    `gorm:"Column:active"`
	CreatedAt string `gorm:"Column:created_at"`
	UpdatedAt string `gorm:"Column:last_edited_at"`
}

// TableName 活动表名称
func (Activity) TableName() string {
	return "activity"
}

// FollowRelationships 邀请关系表
type FollowRelationships struct {
	FolllowerUserHash string     `gorm:"Column:follower_user_hash"`
	FollloweeUserHash string     `gorm:"Column:followee_user_hash"`
	CreatedAt         *time.Time `gorm:"Column:created_at"`
}

// TableName 邀请关系表名称
func (FollowRelationships) TableName() string {
	return "follow_relationships"
}
