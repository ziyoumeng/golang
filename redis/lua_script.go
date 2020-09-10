package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/pkg/errors"
)

var scripts []*LuaScript

func init() {
	scripts = append(scripts, token_bucket)
}

type LuaScript struct {
	scriptName string
	script     string
	sha        string
	keyCount   int
}

func (this *LuaScript) GetScriptName() string {
	return this.scriptName
}

func (this *LuaScript) GetScript() string {
	return this.script
}

func (this *LuaScript) GetKeyCount() int {
	return this.keyCount
}

func (this *LuaScript) SetScriptSha(sha string) {
	this.sha = sha
}

func (this *LuaScript) GetScriptSha() string {
	return this.sha
}

func (this *Redis) EvalSha(scriptName string, values ...interface{}) (ret int64, err error) {
	script, ok := this.luaScript[scriptName]
	if !ok {
		return 0, errors.Errorf("not exist luaScript:%s", scriptName)
	}

	this.Wrap(func(conn redis.Conn) {
		var args []interface{}
		args = append(args, script.GetScriptSha())
		args = append(args, script.GetKeyCount())
		args = append(args, values...)
		ret, err = redis.Int64(conn.Do("EVALSHA", args...))
	})
	return
}

func (this *Redis) batchLoadLuaScript(scripts []*LuaScript) error {
	this.luaScript = make(map[string]*LuaScript)
	for _, script := range scripts {
		err := this.loadLuaScript(script)
		if err != nil {
			return err
		}
		this.luaScript[script.GetScriptName()] = script
	}
	return nil
}

func (this *Redis) loadLuaScript(script *LuaScript) (err error) {
	this.Wrap(func(conn redis.Conn) {
		val := redis.NewScript(script.GetKeyCount(), script.GetScript())
		err = val.Load(conn)
		if err != nil {
			err = errors.WithMessage(err, script.GetScriptName())
			return
		}
		script.SetScriptSha(val.Hash())
	})
	return
}
