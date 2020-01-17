package redis

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/zhuchen/learngin/config"
)

var conn *redis.Client

// NewRedis 新建Redis client
func NewRedis() *redis.Client {
	if conn != nil {
		fmt.Println("单例", conn)
		return conn
	}
	conn = redis.NewClient(&redis.Options{
		Addr:     config.RedisDsn,
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	fmt.Println("init conn", conn)
	pong, err := conn.Ping().Result()
	fmt.Println(pong, err)
	return conn
	// Output: PONG <nil>
}
