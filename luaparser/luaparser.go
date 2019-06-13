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

func parseMap(t *lua.LTable) map[string]interface{} {
	result := make(map[string]interface{})

	t.ForEach(func(k1, v1 lua.LValue) {

		boolValue, ok := v1.(lua.LBool)
		if ok {
			result[k1.String()] = boolValue == lua.LTrue
		}

		intValue, ok := v1.(lua.LNumber)
		if ok {
			result[k1.String()], _ = strconv.Atoi(intValue.String())
		}

		strValue, ok := v1.(lua.LString)
		if ok {
			result[k1.String()] = strValue.String()
		}

		tblValue, ok := v1.(*lua.LTable)
		if ok {
			result[k1.String()] = parseMap(tblValue)
		}
	})

	return result
}

func (this *LuaParser) ParseList(expr string) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

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
		key, ok := v.(*lua.LTable)

		if !ok {
			return
		}

		mapresult := parseMap(key)

		result = append(result, mapresult)
	})

	return result, nil
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

	return parseMap(tbl), nil
}
