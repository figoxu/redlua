package cache

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets88ad6344964997bfc6fb419be9b4458a51c98efe = "local key = KEYS[1]\nlocal id = redis.call('get',key)\nif not id\nthen\n    redis.call('set',key,1)\n    return key..\"0001\"\nelse\n    redis.call('set',key,id+1)\n    return key..string.format('%04d',id + 1)\nend"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/resources": []string{"lua"}, "/resources/lua": []string{"id.lua"}, "/": []string{"resources"}}, map[string]*assets.File{
	"/resources/lua/id.lua": &assets.File{
		Path:     "/resources/lua/id.lua",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1551004222, 1551004222972951371),
		Data:     []byte(_Assets88ad6344964997bfc6fb419be9b4458a51c98efe),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1551000493, 1551000493131363608),
		Data:     nil,
	}, "/resources": &assets.File{
		Path:     "/resources",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1550795993, 1550795993086363676),
		Data:     nil,
	}, "/resources/lua": &assets.File{
		Path:     "/resources/lua",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1551004222, 1551004222974339896),
		Data:     nil,
	}}, "")
