package cache

import "io/ioutil"

var (
	script_lua_id string
)

func init() {
	b, _ := ioutil.ReadAll(Assets.Files["/resources/lua/id.lua"])
	script_lua_id = string(b)
}
