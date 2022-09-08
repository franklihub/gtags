package gtags

import (
	"reflect"
)

func ParseStructTags(obj interface{}) *Structs {
	typ := reflect.TypeOf(obj)
	typ = indirectType(typ)

	return parseStructType(typ)
}

func indirectType(typ reflect.Type) reflect.Type {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}
