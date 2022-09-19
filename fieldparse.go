package gtags

import (
	"reflect"
	"strings"
)

func ParseStructType(typ reflect.Type) *Field {
	f := newField(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		f.ParseStructField(field)
	}
	return f
}

///
func (a *Field) ParseStructField(structfield reflect.StructField) *Field {
	f := newField(structfield.Type)
	f.fieldName = structfield.Name
	f.fieldType = structfield.Type
	f.isAnon = structfield.Anonymous
	f.fieldIndex = append(a.fieldIndex, structfield.Index...)
	f.fieldTypeName = structfield.Type.Name()
	///
	f.tags = parseTags(string(structfield.Tag))
	f.alias = f.tags.Get(AliasTag).Val()
	// UnmarshalJSON
	// _, has := structfield.Type.MethodByName("UnmarshalJSON")
	has := TypMethod(structfield.Type, "UnmarshalJSON")
	f.hasUnmarshal = has
	f.parseType(structfield.Type)
	///addto field list
	a.subFields = append(a.subFields, f)
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

func (a *Field) DMap(dtag string) map[string]any {
	dmap := map[string]any{}

	for _, sf := range a.subFields {
		if sf.IsStruct() {
			if sf.IsAnon() {
				ddmap := sf.DMap(dtag)
				MergerMap(dmap, ddmap)
			} else {
				if sf.Alias() != "" {
					d := sf.DMap(dtag)
					if len(d) > 0 {
						dmap[sf.Alias()] = sf.DMap(dtag)
					} else {
						d := sf.Tags().Get(dtag).Val()
						if d != "" {
							dmap[sf.Alias()] = d
						}
					}
				}
			}
		} else if sf.IsSlice() {
			if sf.fieldType.Elem().Kind() == reflect.String {
				d := sf.Tags().Get(dtag).Val()
				os := sf.Tags().Get(dtag).Opts()
				if d != "" && len(os) > 0 {
					d = d + "," + strings.Join(os, ",")
					s := strings.Split(d, ",")
					if len(s) > 0 {
						dv := []string{}
						for _, v := range s {
							dv = append(dv, v)
						}
						dmap[sf.Alias()] = dv
					}
				}
			} else {

				panic("not supply slice:" + sf.fieldType.Elem().Kind().String())
				// dv := sf.DMap(dtag)
				// if len(dv) > 0 {
				// 	dmap[sf.Alias()] = []any{dv}
				// }
			}
		} else {
			if sf.Alias() != "" {
				d := sf.Tags().Get(dtag).Val()
				if d != "" {
					// v, _ := convKind(sf.fieldType.Kind(), d)
					dmap[sf.Alias()] = d
				}
			}
		}
	}
	return dmap
}
