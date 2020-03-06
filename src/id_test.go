package cache

import (
	"fmt"
	"testing"

	"github.com/quexer/red"
)

func TestNewId(t *testing.T) {
	rp := red.CreatePool(1, "127.0.0.1:6379", "")
	id := NewId("hello", 3, rp)
	for i := 0; i < 10; i++ {
		fmt.Println(id.Next())
	}
}
