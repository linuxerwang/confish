package confish

import (
	"fmt"
	"reflect"
)

func indirect(a interface{}) interface{} {
	if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
		// Avoid creating a reflect.Value if it's not a pointer.
		return a
	}
	v := reflect.ValueOf(a)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func indirectVal(val reflect.Value) reflect.Value {
	for val.Kind() == reflect.Ptr {
		if val.IsNil() {
			newVal := reflect.New(val.Type().Elem())
			val.Set(newVal)
			return newVal.Elem()
		}
		val = val.Elem()
	}
	return val
}

func indirectType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}

func primitiveString(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%t", v.Interface())
	case reflect.Int, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%d", v.Interface())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Interface())
	case reflect.String:
		return fmt.Sprintf("\"%s\"", v.Interface())
	}

	return ""
}
