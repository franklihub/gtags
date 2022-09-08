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
	name2alias map[string]string
	///
	nesteds          map[string]*Structs
	nestedname2alias map[string]string
}

func (a *Structs) Name() string {
	return a.name
}
func (a *Structs) Alise() string {
	return a.alias
}
func (a *Structs) IsAnon() bool {
	//todo: isnano
	return a.isAnon
}

func (a *Structs) NestedNames() []string {
	///
	size := len(a.nesteds)
	names := make([]string, size)
	pos := 0
	for k, _ := range a.nesteds {
		names[pos] = k
		pos += 1
	}

	return names
}
func (a *Structs) FieldNames() []string {
	///
	size := len(a.fields)
	names := make([]string, size)
	pos := 0
	for k, _ := range a.fields {
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

func (a *Structs) NestedByName(nestedfield string) *Structs {
	if a == nil {
		return nil
	}
	return a.nesteds[nestedfield]
}

///
////

///notice: must be elem
func parseStructType(typ reflect.Type) *Structs {
	// typ := reflect.TypeOf(val)
	s := &Structs{
		///
		//contain a struct that has UnmarshalJson method
		fields:     map[string]*Field{},
		name2alias: map[string]string{},

		///
		nesteds:          map[string]*Structs{},
		nestedname2alias: map[string]string{},
	}
	for i := 0; i < typ.NumField(); i++ {
		structfield := typ.Field(i)
		s.addField(structfield)
	}
	return s
}

func (a *Structs) addAnonStruct(structfield reflect.StructField) {
	typ := structfield.Type
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		a.addField(field)
	}
}
func (a *Structs) addStruct(structfield reflect.StructField) {
	tags := parseTags(string(structfield.Tag))
	_, has := structfield.Type.MethodByName("UnmarshalJson")
	///
	s := &Structs{
		name:         structfield.Name,
		alias:        tags.Get(AliasTag).Val(),
		tags:         tags,
		hasUnmarshal: has,
		isAnon:       false,
		///
		//contain a struct that has UnmarshalJson method
		fields:     map[string]*Field{},
		name2alias: map[string]string{},

		///
		nesteds:          map[string]*Structs{},
		nestedname2alias: map[string]string{},
	}
	typ := structfield.Type
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		s.addField(field)
	}
	a.nesteds[structfield.Name] = s
	a.nestedname2alias[s.name] = s.alias
	a.nestedname2alias[s.alias] = s.name
}

func (a *Structs) addField(structfield reflect.StructField) {
	///
	switch structfield.Type.Kind() {
	case reflect.Struct:
		if structfield.Anonymous {
			a.addAnonStruct(structfield)
		} else {
			a.addStruct(structfield)
		}
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
	a.name2alias[f.Alias()] = f.Name()
	a.name2alias[f.Name()] = f.Alias()
	return f
}
