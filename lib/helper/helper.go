package helper

/*
 * @desc         :
 */

import (
	"reflect"

	"github.com/xiaolaji422/golink/c/val"
	"github.com/xiaolaji422/golink/lib/log"
)

func GetMapVal(data map[string]interface{}, key string, def ...interface{}) val.Var {
	var def_val interface{} = ""
	if len(def) > 0 {
		def_val = def[0]
	}
	if v, ok := data[key]; ok {
		def_val = v
	}

	typ := reflect.TypeOf(def_val)
	log.Instance().Info("GetMapVal:", data, key, def_val, typ.Kind())
	return val.NewVal(def_val)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}
