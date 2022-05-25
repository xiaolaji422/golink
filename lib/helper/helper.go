package helper

/*
 * @desc         :
 */

import (
	"reflect"

	"git.woa.com/chengdukf/go-link/c/val"
)

func GetMapVal(data map[string]interface{}, key string, def ...interface{}) val.Var {
	if v, ok := data[key]; ok {
		return val.NewVal(v)
	}
	var def_val interface{} = ""
	if len(def) > 0 {
		def_val = def[0]
	}
	return val.NewVal(def_val)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
