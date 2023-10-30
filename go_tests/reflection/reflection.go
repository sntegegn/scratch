package reflection

import "reflect"

func walk(x interface{}, f func(string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	f(field.String())
}
