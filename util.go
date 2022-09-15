package gtags

import (
	"reflect"
	"strconv"
	"strings"
)

func mergermap(a, b map[string]any) {
	for k, v := range b {
		if vv, ok := a[k]; !ok {
			a[k] = v
		} else {
			if _, ok := vv.(map[string]any); ok {
				if _, ok := v.(map[string]any); ok {
					mergermap(vv.(map[string]any), v.(map[string]any))
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

func formatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}
