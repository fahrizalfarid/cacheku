package cacheku

import (
	"fmt"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	c := NewCache(1 * time.Second)
	err := c.Set("hallo", "world", 0)
	if err != nil {
		t.Error("it supposed be nil", err)
	}

	found, data := c.Get("hallo")
	if !found {
		t.Error("supposed to be true", found)
	}
	if data.(string) != "world" {
		t.Error("supposed to be world", data.(string))
	}

	err = c.Set("hallo1", "world", 1*time.Second)
	if err != nil {
		t.Error("supposed to be nil", err)
	}

	time.Sleep(3 * time.Second)

	found, data = c.Get("hallo")
	if found {
		t.Error("supposed to be false", found)
	}
	if data != nil {
		t.Error("supposed to be nil", data)
	}

	c.Delete("hello")
}

func TestFetchAll(t *testing.T) {
	c := NewCache(10 * time.Second)
	for i := 0; i < 10; i++ {
		c.Set(fmt.Sprintf("%d-hallo", i), fmt.Sprintf("%d-world", i), 0)
	}

	f := c.FetchAll()
	if len(f) != 10 {
		t.Error("supposed to be 10", len(f))
	}
}

func BenchmarkSet(b *testing.B) {
	c := NewCache(1 * time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Set(fmt.Sprintf("%d-hallo", i), fmt.Sprintf("%d-world", i), 0)
	}
}

func BenchmarkGet(b *testing.B) {
	c := NewCache(10 * time.Second)

	for i := 0; i < 1e4; i++ {
		c.Set(fmt.Sprintf("%d-hallo", i), fmt.Sprintf("%d-world", i), 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Get(fmt.Sprintf("%d-hallo", i))
	}
}

func BenchmarkDelete(b *testing.B) {
	c := NewCache(10 * time.Second)

	for i := 0; i < 1e4; i++ {
		c.Set(fmt.Sprintf("%d-hallo", i), fmt.Sprintf("%d-world", i), 0)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Delete(fmt.Sprintf("%d-hallo", i))
	}
}
