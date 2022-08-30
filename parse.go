package gtags

import (
	"reflect"
)

func ParseStructTags(obj interface{}) *StructTags {
	typ := reflect.TypeOf(obj)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return parseStructType(typ)
}

func indirectType(typ reflect.Type) reflect.Type {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ
}

func parseStructType(typ reflect.Type) *StructTags {
	typ = indirectType(typ)

	stags := &StructTags{
		fieldtags:     map[string]*Tags{},
		anonestedtags: map[string]*Tags{},
		nestedtags:    map[string]*StructTags{},
	}
	///
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		///
		if field.Type.Kind() == reflect.Struct {
			if field.Name == field.Type.Name() {
				//anonymous
				anotags := parseStructType(field.Type)
				for k, v := range anotags.fieldtags {
					stags.anonestedtags[k] = v
				}
			} else {
				//nested
				nesttags := parseStructType(field.Type)
				stags.nestedtags[field.Name] = nesttags
			}
		} else {
			tags := parseTags(string(field.Tag))
			stags.fieldtags[field.Name] = tags
		}
	}

	// typ.Name()
	// typ.PkgPath()
	// typ.Kind().String()
	return stags
}
