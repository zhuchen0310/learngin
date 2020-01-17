package common

import (
	"encoding/json"
	"fmt"
)

// JSONLoads 将json 转为 map类型
func JSONLoads(jsonData []byte, v interface{}) error {
	err := json.Unmarshal(jsonData, v)
	return err
}

// JSONDumps 将结构换数据 转换为 json
func JSONDumps(data interface{}) ([]byte, error) {
	ret, err := json.Marshal(data)
	return ret, err
}

func main() {
	a := map[string]interface{}{"1": 1, "2": "2"}
	b, _ := JSONDumps(a)
	fmt.Println(a, b)
	fmt.Printf("a, %T, b, %T", a, b)
	var c map[string]interface{}
	err := JSONLoads(b, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println()
	fmt.Println(c)
	fmt.Printf("%T", c)
}
