package cache

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets88ad6344964997bfc6fb419be9b4458a51c98efe = "local key = KEYS[1]\nlocal id = redis.call('get',key)\nif(id == false)\nthen\n    redis.call('set',key,1)\n    return key..\"0001\"\nelse\n    redis.call('set',key,id+1)\n    return key..string.format('%04d',id + 1)\nend"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"resources"}, "/resources": []string{"lua"}, "/resources/lua": []string{"id.lua"}}, map[string]*assets.File{
	"/resources/lua/id.lua": &assets.File{
		Path:     "/resources/lua/id.lua",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1550794659, 1550794659170338187),
		Data:     []byte(_Assets88ad6344964997bfc6fb419be9b4458a51c98efe),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1550794908, 1550794908604969176),
		Data:     nil,
	}, "/resources": &assets.File{
		Path:     "/resources",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1550794627, 1550794627722247706),
		Data:     nil,
	}, "/resources/lua": &assets.File{
		Path:     "/resources/lua",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1550794659, 1550794659171961897),
		Data:     nil,
	}}, "")
