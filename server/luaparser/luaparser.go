package luaparser

import (
	"errors"
	"github.com/yuin/gopher-lua"
	"strconv"
)

func New() *LuaParser {
	p := LuaParser{
		state: lua.NewState(lua.Options{SkipOpenLibs: true}),
	}

	return &p
}

type LuaParser struct {
	state *lua.LState
}

func (this *LuaParser) ParseMap(expr string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	err := this.state.DoString(expr)
	if err != nil {
		return result, err
	}

	lv := this.state.Get(-1)

	tbl, ok := lv.(*lua.LTable)
	if !ok {
		return result, errors.New("parsing failed")
	}

	tbl.ForEach(func(k, v lua.LValue) {
		key, ok := k.(lua.LString)

		if !ok {
			return
		}

		boolValue, ok := v.(lua.LBool)
		if ok {
			result[key.String()] = boolValue == lua.LTrue
		}
		intValue, ok := v.(lua.LNumber)
		if ok {
			result[key.String()], _ = strconv.Atoi(intValue.String())
		}
		strValue, ok := v.(lua.LString)
		if ok {
			result[key.String()] = strValue.String()
		}
	})

	return result, nil
}
