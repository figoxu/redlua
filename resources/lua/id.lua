local key = KEYS[1]
local len = ARGV[1]
local id = redis.call('get',key)
local fmt = "%0"..len.."d"
if not id
then
    redis.call('set',key,1)
    return key..string.format(fmt,1)
else
    redis.call('set',key,id+1)
    return key..string.format(fmt,id + 1)
end