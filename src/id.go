package cache

import (
	"github.com/garyburd/redigo/redis"
	"github.com/quexer/utee"
)

type Id struct {
	prefix string
	seqLen int
	rp     *redis.Pool
}

func NewId(prefix string, seqLen int, rp *redis.Pool) *Id {
	return &Id{
		prefix: prefix,
		rp:     rp,
		seqLen: seqLen,
	}
}

func (p *Id) Next() string {
	script := redis.NewScript(1, script_lua_id)
	c := p.rp.Get()
	defer c.Close()
	v, err := redis.String(script.Do(c, p.prefix, p.seqLen, p.seqLen))
	utee.Chk(err)
	return v
}
