package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
	"time"
)


type Redis struct {
	pool      *redis.Pool
	luaScript map[string]*LuaScript //脚本名=>LuaScript
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

//
func NewRedis(address string) (*Redis, error) {
	pool := newPool(address)
	cache := &Redis{
		pool: pool,
	}

	err := cache.batchLoadLuaScript(scripts)
	if err != nil {
		return nil, errors.WithMessage(err, "batchLoadLuaScript")
	}
	return cache, nil
}

func (d *Redis) GetString(key string) (ret string, err error) {
	d.Wrap(func(conn redis.Conn) {
		ret, err = redis.String(conn.Do("GET", key))
		err = getErr(err)
	})
	return
}

func (d *Redis) Set(key, value string) (err error) {
	d.Wrap(func(conn redis.Conn) {
		_, err = conn.Do("SET", key, value)
		err = getErr(err)
	})
	return
}

func (d *Redis) Wrap(doSomething func(conn redis.Conn)) {
	conn := d.pool.Get()
	defer conn.Close()

	doSomething(conn)
}

func (d *Redis) SetNX(key, value string) (ok bool, err error) {
	d.Wrap(func(conn redis.Conn) {
		ok, err = redis.Bool(conn.Do("SETNX", key, value))
		err = getErr(err)
	})
	return
}

//pair = <score, value>
func (d *Redis) ZAdd(key string, pairs ...interface{}) (newAddNum int, err error) {
	if len(pairs) == 0 || len(pairs)%2 != 0 {
		return 0, errors.New("invalid pairs num")
	}
	var args []interface{}
	args = append(args, key)
	args = append(args, pairs...)
	d.Wrap(func(conn redis.Conn) {
		newAddNum, err = redis.Int(conn.Do("ZADD", args...))
	})
	return
}

func (d *Redis) ZRank(key, member string) (rank int, err error) {
	d.Wrap(func(conn redis.Conn) {
		rank, err = redis.Int(conn.Do("ZRANK", key, member))
		err = getErr(err)
	})
	return
}

func (d *Redis) ZScore(key, member string) (score int64, err error) {
	d.Wrap(func(conn redis.Conn) {
		score, err = redis.Int64(conn.Do("ZSCORE", key, member))
		err = getErr(err)
	})
	return
}

func (d *Redis) ZCard(key string) (num int64, err error) {
	d.Wrap(func(conn redis.Conn) {
		num, err = redis.Int64(conn.Do("ZCARD", key))
		err = getErr(err)
	})
	return
}

func getErr(err error) error {
	//if err == redis.ErrNil {
	//	err = code.ErrNotFound
	//} else if err != nil {
	//	err = status.ErrInternalServerError
	//}
	return err
}

func (d *Redis) ZRem(key string, members ...interface{}) (remNum int, err error) {
	if len(members) == 0 {
		return 0, errors.New("invalid members num")
	}
	var args []interface{}
	args = append(args, key)
	args = append(args, members...)
	d.Wrap(func(conn redis.Conn) {
		remNum, err = redis.Int(conn.Do("ZREM", args...))
		err = getErr(err)
	})
	return
}

func (d *Redis) HSetBytes(key, field string, value []byte) (err error) {
	d.Wrap(func(conn redis.Conn) {
		_, err = conn.Do("HSET", key, field, value)
		err = getErr(err)
	})
	return
}

func (d *Redis) HGetBytes(key, field string) (value []byte, err error) {
	d.Wrap(func(conn redis.Conn) {
		value, err = redis.Bytes(conn.Do("HGET", key, field))
		err = getErr(err)
	})
	return
}

func (d *Redis) HSetString(key, field, value string) (ret int, err error) {
	d.Wrap(func(conn redis.Conn) {
		ret, err = redis.Int(conn.Do("HSET", key, field, value))
		err = getErr(err)
	})
	return
}

func (d *Redis) LPushInt64(key string, value int64) (err error) {
	d.Wrap(func(conn redis.Conn) {
		_, err = conn.Do("LPUSH", key, value)
		err = getErr(err)
	})
	return
}

func (d *Redis) LRangeAllInt64(key string) (ret []int64, err error) {
	d.Wrap(func(conn redis.Conn) {
		ret, err = redis.Int64s(conn.Do("LRANGE", key, 0, -1))
		err = getErr(err)
	})
	return
}

func (d *Redis) HIncrby(key, field string, num int32) (err error) {
	d.Wrap(func(conn redis.Conn) {
		_, err = redis.Int(conn.Do("HINCRBY", key, field, num))
		err = getErr(err)
	})
	return
}

func (d *Redis) HGetInt64(key, field string) (num int64, err error) {
	d.Wrap(func(conn redis.Conn) {
		num, err = redis.Int64(conn.Do("HGET", key, field))
		err = getErr(err)
	})
	return
}

func (d *Redis) HGetString(key, field string) (ret string, err error) {
	d.Wrap(func(conn redis.Conn) {
		ret, err = redis.String(conn.Do("HGET", key, field))
		err = getErr(err)
	})
	return
}

//注意: 即使都不存在会,ret长度不为0,返回空字符串数组
func (d *Redis) HMGetString(key string, field ...interface{}) (ret []string, err error) {
	d.Wrap(func(conn redis.Conn) {
		var args []interface{}
		args = append(args, key)
		args = append(args, field...)
		ret, err = redis.Strings(conn.Do("HMGET", args...))
		err = getErr(err)
	})
	return
}

func (d *Redis) HGetAllBytes(key string) (bytes [][]byte, err error) {
	d.Wrap(func(conn redis.Conn) {
		ret, err := redis.StringMap(conn.Do("HGETALL", key))
		err = getErr(err)
		if err != nil {
			return
		}
		for _, s := range ret {
			if len(s) <= 0 {
				continue
			}
			bytes = append(bytes, []byte(s))
		}
	})
	return
}

func (d *Redis) HDel(hKey string, field ...interface{}) (delNum int, err error) {
	d.Wrap(func(conn redis.Conn) {
		var args []interface{}
		args = append(args, hKey)
		args = append(args, field...)
		delNum, err = redis.Int(conn.Do("HDEL", args...))
		err = getErr(err)
	})
	return
}

func (d *Redis) Del(key string) (err error) {
	d.Wrap(func(conn redis.Conn) {
		_, err = conn.Do("DEL", key)
		err = getErr(err)
	})
	return
}

func (d *Redis) Expire(key string, sec int32) (err error) {
	d.Wrap(func(conn redis.Conn) {
		_, err = conn.Do("Expire", key, sec)
		err = getErr(err)
	})
	return
}

func (d *Redis) Keys(pattern string) (keys []string, err error) {
	d.Wrap(func(conn redis.Conn) {
		keys, err = redis.Strings(conn.Do("KEYS", pattern))
		err = getErr(err)
	})
	return
}
