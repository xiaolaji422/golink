package val

import (
	"fmt"
	"reflect"
	"strconv"
)

type Var struct {
	Value interface{}
}

func NewVal(val interface{}) Var {
	return Var{
		Value: val,
	}
}

func (v Var) String() string {
	switch v.Value.(type) {
	case string:
		return v.Value.(string)
	case int, int16, int32, int8, int64:
		return strconv.Itoa(v.Value.(int))
	case float32, float64:
		return fmt.Sprintf("%f", v.Value)
	default:
		return ""
	}
}

func (v Var) Int() int {
	switch v.Value.(type) {
	case string:
		i, err := strconv.Atoi(v.Value.(string))
		if err != nil {
			return 0
		}
		return i
	case int16, int32, int8, int64:
		return int(v.Value.(int64))
	case int:
		return v.Value.(int)
	case float32, float64:
		return int(v.Value.(float64))
	default:
		return 0
	}
}

func (v Var) IsNil() bool {
	vi := reflect.ValueOf(v.Value)
	switch vi.Kind() {
	case reflect.Ptr:
		return vi.IsNil()
	default:
		if v.Value == nil {
			return true
		}
	}
	return false
}
