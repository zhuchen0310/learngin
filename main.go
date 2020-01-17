package main

import (
	"fmt"

	"github.com/go-redis/redis/v7"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>

	// key := "test"
	// user := User{Name: "zhuchen", Age: 18}
	// value, err := json.Marshal(user)
	// fmt.Println(value)
	// err = client.Set(key, value, 60*time.Second).Err()

	// err = client.Set("key", "test", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := client.Get(key).Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)
	// var user2 *User
	// json.Unmarshal([]byte(val), &user2)
	// fmt.Println("key", user2.Name, user2.Age)

	// val2, err := client.Get("key2").Result()
	// if err == redis.Nil {
	// 	fmt.Println("key2 does not exist")
	// } else if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("key2", val2)
	// }

	key := "testhash"
	field := "1"
	err = client.HSet(key, field, 1).Err()
	if err != nil {
		panic(err)
	}
	d, err := client.HGet(key, field).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	values := map[string]interface{}{"1": "1", "2": "2"}
	fmt.Println(values)
	client.HMSet(key, values).Err()

	d1, _ := client.HGetAll(key).Result()
	fmt.Println("hgetall", d1)

	d2, _ := client.HKeys(key).Result()
	fmt.Println(d2)

	d3, _ := client.HVals(key).Result()
	fmt.Println(d3)
	fmt.Printf("%T", d3)
	d4 := client.HVals(key).Val()
	fmt.Println(d4)
	fmt.Println("zset: ")
	key = "test_zset"
	m := []*redis.Z{&redis.Z{Score: 10, Member: 1}}
	er := client.ZAdd(key, m...).Err()
	if er != nil {
		panic(er)
	}
	z1, _ := client.ZRangeWithScores(key, 0, -1).Result()
	fmt.Println(z1[0])
	fmt.Printf("%T %s %f", z1, z1[0].Member, z1[0].Score)
}
