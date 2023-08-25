package main

import (
	"fmt"
	"time"

	cacheku "github.com/fahrizalfarid/cacheku"
)

func main() {
	c := cacheku.NewCache(10 * time.Minute)

	_ = c.Set("hello", "world", 0)
	_ = c.Set("hello1", "world", 0)
	_ = c.Set("hello2", "world", 0)

	found, data := c.Get("hello")
	if !found {
		fmt.Println("data not found")
		return
	}
	fmt.Println("data found", data.(string))

	items := c.FetchAll()
	for k, v := range items {
		fmt.Printf("key %s, value %s expired %d\n", k, v.Value.(string), v.Expire)
	}

	c.Delete("hello")

	items = c.FetchAll()
	fmt.Println(items)
}
