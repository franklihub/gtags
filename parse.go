package gtags

import (
	"reflect"
)

func ParseStructTags(obj interface{}) *Field {
	typ := reflect.TypeOf(obj)
	typ = indirectType(typ)

	return ParseStructType(typ)
	// return parseStructType(typ)
}

// func parseStruct(typ reflect.Type) *Field {

// }

func indirectType(typ reflect.Type) reflect.Type {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}
