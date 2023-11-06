package redisLock

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"time"
)

//[基于 go-redis 与 Lua 实现简易且高效的分布式锁](https://www.moemona.com/backend-notebook/1185/)

var Redis *redis.Client

func init() {
	Redis = NewRedis()
}

func NewRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	return rdb
}

type LockInterface interface {
	Lock() bool
	Block(seconds int64) bool // 持续获取锁
	Release() bool
	ForceRelease()
}

const LockPrefix = "redis:lock"

type lock struct {
	context context.Context
	name    string // 锁名称
	owner   string // 锁标识
	seconds int64  // 有效期
}

func NewLock(name string, seconds int64) LockInterface {
	return &lock{
		context.Background(),
		LockPrefix + name,
		uuid.New().String(),
		seconds,
	}
}

func (l *lock) Lock() bool {
	return Redis.SetNX(l.context, l.name, l.owner, time.Duration(l.seconds)*time.Second).Val()
}

func (l *lock) Block(seconds int64) bool {
	starting := time.Now().Unix()
	for {
		if !l.Lock() {
			time.Sleep(time.Duration(1) * time.Second)
			if time.Now().Unix()-seconds >= starting {
				return false
			}
		} else {
			return true
		}
	}
}

const releaseLockLuaScript = `
if redis.call("get", KEYS[1]) == ARGV[1] then
   return redis.call("del", KEYS[1])
else
    return 0
end
`

func (l *lock) Release() bool {
	luaScript := redis.NewScript(releaseLockLuaScript)

	result := luaScript.Run(
		l.context, Redis,
		[]string{l.name}, // []string{l.name}   KEYS[1]
		l.owner,          // l.owner            ARGV[1]
	).Val().(int64)
	return result != 0
}

func (l *lock) ForceRelease() {
	Redis.Del(l.context, l.name).Val()
}
