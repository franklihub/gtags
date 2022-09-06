package gtags

import (
	"reflect"
)

type Structs struct {
	name         string
	alias        string
	tags         *Tags
	hasUnmarshal bool
	isAnon       bool
	///
	fields     map[string]*Field
	alias2name map[string]string
	name2alias map[string]string
	///
	nesteds          map[string]*Structs
	alias2nestedname map[string]string
	nestedname2alias map[string]string
}

func (a *Structs) FieldNames() []string {
	///
	size := len(a.fields) + len(a.nesteds)
	names := make([]string, size)
	pos := 0
	for k, _ := range a.fields {
		names[pos] = k
		pos += 1
	}
	for k, _ := range a.nesteds {
		names[pos] = k
		pos += 1
	}

	return names
}

func (a *Structs) FieldByName(fieldname string) *Field {
	if a == nil {
		return nil
	}
	if v, ok := a.fields[fieldname]; ok {
		return v
	}

	return nil
}

///

// func (a *Structs) Field(field string) *Tags {
// 	if a == nil {
// 		return nil
// 	}
// 	if v, ok := a.fieldtags[field]; ok {
// 		return v
// 	}
// 	if v, ok := a.anonestedtags[field]; ok {
// 		return v
// 	}
// 	return nil
// }
func (a *Structs) NestedByName(nestedfield string) *Structs {
	if a == nil {
		return nil
	}
	return a.nesteds[nestedfield]
}

///
////

///notice: must be elem
func parseType(typ reflect.Type) *Structs {
	// typ := reflect.TypeOf(val)
	s := &Structs{
		///
		//contain a struct that has UnmarshalJson method
		fields:     map[string]*Field{},
		alias2name: map[string]string{},
		name2alias: map[string]string{},

		///
		nesteds:          map[string]*Structs{},
		alias2nestedname: map[string]string{},
		nestedname2alias: map[string]string{},
	}
	for i := 0; i < typ.NumField(); i++ {
		structfield := typ.Field(i)
		s.addField(structfield)
	}
	return s
}

func (a *Structs) addStruct(structfield reflect.StructField) *Structs {
	tags := parseTags(string(structfield.Tag))
	_, has := structfield.Type.MethodByName("UnmarshalJson")
	isAnon := structfield.Anonymous
	///
	s := &Structs{
		name:         structfield.Name,
		alias:        tags.Get(AliasTag).Val(),
		tags:         tags,
		hasUnmarshal: has,
		isAnon:       isAnon,
		///
		//contain a struct that has UnmarshalJson method
		fields:     map[string]*Field{},
		alias2name: map[string]string{},
		name2alias: map[string]string{},

		///
		nesteds:          map[string]*Structs{},
		alias2nestedname: map[string]string{},
		nestedname2alias: map[string]string{},
	}

	typ := structfield.Type
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		s.addField(field)
	}
	///
	if isAnon {
		//todo: merge field
		for k, v := range s.fields {
			a.fields[k] = v
		}
		for k, v := range s.alias2name {
			a.alias2name[k] = v
		}
		for k, v := range s.name2alias {
			a.name2alias[k] = v
		}
		///
		for k, v := range s.nesteds {
			a.nesteds[k] = v
		}
		for k, v := range s.nestedname2alias {
			a.nestedname2alias[k] = v
		}
		for k, v := range s.alias2nestedname {
			a.alias2nestedname[k] = v
		}
	} else {
		a.nesteds[structfield.Name] = s
	}
	return s
}

func (a *Structs) addField(structfield reflect.StructField) {
	///
	switch structfield.Type.Kind() {
	case reflect.Struct:
		a.addStruct(structfield)
	case reflect.Pointer:
		a.add(structfield)
	case reflect.Array, reflect.Slice:
		panic("not supply slice/array")
	case reflect.Map:
		panic("not supply Map")
	case reflect.Chan:
		panic("not supply Chan")
	case reflect.Func:
		panic("not supply Func")
	default:
		a.add(structfield)
		//
	}
}

func (a *Structs) add(field reflect.StructField) *Field {
	f := parseField(field)
	a.fields[f.Name()] = f
	///
	a.alias2name[f.Alias()] = f.Name()
	a.alias2name[f.Name()] = f.Alias()
	return f
}
