package gtags

import (
	"reflect"
	"strings"
)

func MergerMap(a, b map[string]any) {
	for k, v := range b {
		if vv, ok := a[k]; !ok {
			a[k] = v
		} else {
			if _, ok := vv.(map[string]any); ok {
				if _, ok := v.(map[string]any); ok {
					MergerMap(vv.(map[string]any), v.(map[string]any))
				}
			}
		}
	}
}
func tagDVal(tag *Tag) string {
	d := tag.Val()
	o := tag.Opts()

	if len(o) > 0 {
		d = d + "," + strings.Join(o, ",")
	}
	return d
}
func TypMethod(typ reflect.Type, method string) bool {
	ok := ptrMethod(typ, method)
	if !ok {
		ok = valMethod(typ, method)
	}
	return ok
}

//
func valMethod(typ reflect.Type, method string) bool {
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}
	_, ok := typ.MethodByName(method)
	return ok
}
func ptrMethod(typ reflect.Type, method string) bool {
	if typ.Kind() != reflect.Pointer {
		typ = reflect.PtrTo(typ)
	}
	_, ok := typ.MethodByName(method)
	return ok
}

// func convKind(kind reflect.Kind, str string) (any, error) {
// 	var val any = str
// 	var err error
// 	switch kind {
// 	case reflect.Int:
// 		i, err := strconv.ParseInt(str, 10, 0)
// 		if err != nil {
// 		} else {
// 			val = int(i)
// 		}
// 	case reflect.Int64:
// 		i, err := strconv.ParseInt(str, 10, 0)
// 		if err != nil {
// 		} else {
// 			val = i
// 		}
// 	case reflect.Uint:
// 		i, err := strconv.ParseUint(str, 10, 0)
// 		if err != nil {
// 		} else {
// 			val = i
// 		}
// 	case reflect.Uint64:
// 		i, err := strconv.ParseUint(str, 10, 0)
// 		if err != nil {
// 		} else {
// 			val = uint(i)
// 		}
// 	case reflect.Float32:
// 		val = gconv.Float32(str)
// 	case reflect.Float64:
// 		val = gconv.Float64(str)
// 	case reflect.Bool:
// 		val = gconv.Bool(str)
// 	case reflect.String:
// 		val = str
// 	case reflect.Slice:
// 		v := strings.Split(str, ",")
// 		if v[0] != "" {
// 			val = v
// 		}
// 	case reflect.Struct:
// 	case reflect.Array:
// 	default:
// 		panic(kind.String())
// 	}
// 	return val, err
// }
