package redis

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/zhuchen/learngin/config"
)

var conn *redis.Client

// init

func init() {
	conn = redis.NewClient(&redis.Options{
		Addr:         config.RedisDsn,
		Password:     "",              // no password set
		DB:           1,               // use default DB
		PoolSize:     50,              //连接池大小
		MinIdleConns: 10,              //最小空闲连接数
		IdleTimeout:  5 * time.Second, //最小空闲连接数
	})
}

// NewRedis 新建Redis client
func NewRedis() *redis.Client {
	return conn
	// Output: PONG <nil>
}
