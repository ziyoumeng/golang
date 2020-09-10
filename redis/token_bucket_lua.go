package redis


/*令牌桶算法
特点： rate + 桶当前令牌数, 即为当前请求的峰值
KEYS[1] 接口id
ARGV[1] 令牌桶的容量
ARGV[2] 令牌每秒恢复速度
ARGV[3] 当前时间戳，毫秒
*/

var token_bucket =  &LuaScript{
	scriptName: "token_bucket",
	sha:        "",
	keyCount:   1,
	script:     `
local ratelimit_info = redis.pcall('HMGET',KEYS[1],'last_time','current_token')
local last_time = ratelimit_info[1] -- 上次令牌获取时间
local current_token = tonumber(ratelimit_info[2]) --上次令牌获取后，桶遗留的令牌数
local max_token = tonumber(ARGV[1]) -- 令牌桶的容量
local token_rate = tonumber(ARGV[2]) -- 令牌每秒恢复速度
local current_time = tonumber(ARGV[3]) -- 当前时间戳，毫秒
local reverse_time = 1000/token_rate -- 每生成一个令牌消耗的时间，毫秒
if current_token == nil then -- 第一次运行脚本时默认令牌桶是满的，因此可以将数据的过期时间设置为令牌桶恢复到满所需的时间，及时释放资源
  current_token = max_token
  last_time = current_time
else
  local past_time = current_time-last_time 
  local reverse_token = math.floor(past_time/reverse_time) -- 两次获取的时间段内应该生成多少令牌
  current_token = current_token+reverse_token -- 当前桶的令牌数
  last_time = reverse_time*reverse_token+last_time --更新last_time,考虑math.floor会引入误差，所以不直接使用current_time
  if current_token>max_token then
    current_token = max_token
  end
end
local result = 0
if(current_token>0) then
  result = 1
  current_token = current_token-1
end
redis.call('HMSET',KEYS[1],'last_time',last_time,'current_token',current_token)
-- 设置桶的剩余缓存时间;Redis 的 pexpire 等命令不支持小数，但 Lua 的 Number 类型可以存放小数，因此 Number 类型传递给 Redis 时最好通过 math.ceil() 等方式转换以避免存在小数导致命令失败。
redis.call('pexpire',KEYS[1],math.ceil(reverse_time*(max_token-current_token)+(current_time-last_time))) 
return result
`,
}


