package cache

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/red"
	"github.com/quexer/utee"
	"testing"
)

func TestLuaId(t *testing.T) {
	rp := red.CreatePool(1, "127.0.0.1:6379", "")
	c := rp.Get()
	defer c.Close()

	script := redis.NewScript(1, script_lua_id)
	reply, err := script.Do(c, "idkey")

	utee.Chk(err)
	v, err := redis.String(reply, err)
	utee.Chk(err)
	fmt.Println(v)
}

func TestExampleScript(t *testing.T) {
	rp := red.CreatePool(1, "127.0.0.1:6379", "")
	c := rp.Get()
	defer c.Close()
	var getScript = redis.NewScript(1, `return redis.call('get', KEYS[1])`)
	reply, err := getScript.Do(c, "foo")
	utee.Chk(err)
	v, err := redis.String(reply, err)
	utee.Chk(err)
	fmt.Println(v)
}
