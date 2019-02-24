package cache

import (
	"fmt"
	"github.com/quexer/red"
	"testing"
)

func TestNewId(t *testing.T) {
	rp := red.CreatePool(1, "127.0.0.1:6379", "")
	id := NewId("201901011121", rp)
	for i := 0; i < 10; i++ {
		fmt.Println(id.Next())
	}
}
