package gtags

import (
	"reflect"
)

func ParseStructType(typ reflect.Type) *Field {
	f := newField(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		sf := ParseStructField(field)
		f.subFields = append(f.subFields, sf)
	}
	return f
}

///
func ParseStructField(structfield reflect.StructField) *Field {
	f := newField(structfield.Type)
	f.fieldName = structfield.Name
	f.fieldType = structfield.Type
	f.isAnon = structfield.Anonymous
	f.fieldIndex = structfield.Index
	f.fieldTypeName = structfield.Type.Name()
	///
	f.tags = parseTags(string(structfield.Tag))
	f.alias = f.tags.Get(AliasTag).Val()
	// UnmarshalJSON
	_, has := structfield.Type.MethodByName("UnmarshalJSON")
	f.hasUnmarshal = has
	f.parseType(structfield.Type)
	return f
}

//
func newField(typ reflect.Type) *Field {
	f := &Field{
		fieldType: typ,
	}
	return f
}

func (a *Field) parseType(typ reflect.Type) {
	// func (a *Field) parseType(structfield reflect.StructField) {
	//todo: structfield pointer convto elem
	// typ := indirectType(structfield.Type)
	//

	switch typ.Kind() {
	case reflect.Invalid:
		panic("Invalid Kind")
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.String:
	case reflect.Uintptr:
		panic("not supply Uintptr")
	case reflect.Float32:
	case reflect.Float64:
	case reflect.Complex64:
		panic("not supply Complex64")
	case reflect.Complex128:
		panic("not supply Complex128")
	case reflect.Array:
	case reflect.Chan:
		panic("not supply Chan")
	case reflect.Func:
		panic("not supply Func")
	case reflect.Interface:
		panic("not supply Interface")
	case reflect.Map:
		panic("not supply Map")
	case reflect.Pointer:
		a.parseType(typ.Elem())
	case reflect.Slice:
		a.addSlice(typ)
	case reflect.Struct:
		a.addStruct(typ)
	case reflect.UnsafePointer:
		panic("not supply UnsafePointer")
	}
}

///
