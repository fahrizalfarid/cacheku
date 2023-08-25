```go
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
```

```bash
$ go test -v -coverprofile cover.out -run=. -bench=. -benchmem
=== RUN   TestCache
--- PASS: TestCache (3.00s)
=== RUN   TestFetchAll
--- PASS: TestFetchAll (0.00s)
goos: linux
goarch: amd64
pkg: github.com/fahrizalfarid/cacheku
cpu: AMD Ryzen 5 4600H with Radeon Graphics
BenchmarkSet
BenchmarkSet-12          1491522               696.2 ns/op           202 B/op          5 allocs/op
BenchmarkGet
BenchmarkGet-12          8527796               140.0 ns/op            23 B/op          1 allocs/op
BenchmarkDelete
BenchmarkDelete-12       9796062               123.3 ns/op            23 B/op          1 allocs/op
PASS
ok      github.com/fahrizalfarid/cacheku        9.654s
```