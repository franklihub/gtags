package gtags

import (
	"reflect"
)

var AliaseTag string = "json"

func ParseStructTags(obj interface{}) *StructTags {
	typ := reflect.TypeOf(obj)
	val := reflect.ValueOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = reflect.ValueOf(obj).Elem()
	}
	//  else {
	// 	val = reflect.ValueOf(obj)
	// }
	// for i := 0; i < val.NumField(); i++ {
	// 	field := val.Field(i)
	// 	if field.Type() == reflect.T {
	// 		fmt.Println(field.Type())
	// 	}
	// }

	return parseStructType(val, false)
}

func indirectType(typ reflect.Type) reflect.Type {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}
