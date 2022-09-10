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
	index        []int
	///
	fields     map[string]*Field
	name2alias map[string]string
	///
	nesteds          map[string]*Structs
	nestedname2alias map[string]string
	//
	anons      map[string]*Structs
	anon2alias map[string]string
}

func (a *Structs) Index() []int {
	return a.index
}
func (a *Structs) HasUnmarshal() bool {
	return a.hasUnmarshal
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

///
func (a *Structs) namealias(name string) string {
	if v, ok := a.name2alias[name]; ok {
		return v
	}
	for _, s := range a.anons {
		if v, ok := s.name2alias[name]; ok {
			return v
		}
	}
	return name
}
func (a *Structs) Field(name string) *Field {
	if a == nil {
		return nil
	}
	f := a.FieldByName(name)
	///
	if f == nil {
		name = a.namealias(name)
		f = a.FieldByName(name)
		return f
	}

	return f
}

////
func (a *Structs) FieldNames() []string {
	///
	size := len(a.fields) + len(a.anons)
	names := make([]string, size)
	pos := 0
	for k, _ := range a.fields {
		names[pos] = k
		pos += 1
	}
	for k, _ := range a.anons {
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
	for _, s := range a.anons {
		if v, ok := s.fields[fieldname]; ok {
			return v
		}
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

///
func (a *Structs) AnonByName(anonname string) *Structs {
	if a == nil {
		return nil
	}
	return a.anons[anonname]
}

func (a *Structs) AnonNames() []string {
	///
	size := len(a.anons)
	names := make([]string, size)
	pos := 0
	for k, _ := range a.anons {
		names[pos] = k
		pos += 1
	}

	return names
}

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
		//
		anons:      map[string]*Structs{},
		anon2alias: map[string]string{},
	}
	for i := 0; i < typ.NumField(); i++ {
		structfield := typ.Field(i)
		s.addStructField(structfield)
	}
	return s
}

///
///

func (a *Structs) addStructField(structfield reflect.StructField) {

	switch structfield.Type.Kind() {
	case reflect.Struct:
		s := a.parseStruct(structfield)
		if structfield.Anonymous {
			a.anons[structfield.Name] = s
			a.anon2alias[s.name] = s.alias
			a.anon2alias[s.alias] = s.name
		} else {
			a.nesteds[structfield.Name] = s
			a.nestedname2alias[s.name] = s.alias
			a.nestedname2alias[s.alias] = s.name
		}
	case reflect.Pointer:
		// a.add(structfield)
	case reflect.Array, reflect.Slice:
		panic("not supply slice/array")
	case reflect.Map:
		panic("not supply Map")
	case reflect.Chan:
		panic("not supply Chan")
	case reflect.Func:
		panic("not supply Func")
	default:
		a.addField(structfield)
	}
}

func (a *Structs) parseStruct(structfield reflect.StructField) *Structs {
	tags := parseTags(string(structfield.Tag))
	_, has := structfield.Type.MethodByName("UnmarshalJSON")
	///
	s := &Structs{
		name:         structfield.Name,
		alias:        tags.Get(AliasTag).Val(),
		tags:         tags,
		hasUnmarshal: has,
		isAnon:       structfield.Anonymous,
		index:        append(a.index[:], structfield.Index...),
		///
		//contain a struct that has UnmarshalJson method
		fields:     map[string]*Field{},
		name2alias: map[string]string{},

		///
		nesteds:          map[string]*Structs{},
		nestedname2alias: map[string]string{},
		//
		anons:      map[string]*Structs{},
		anon2alias: map[string]string{},
	}
	typ := structfield.Type
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		s.addStructField(field)
	}

	return s
}

func (a *Structs) addField(structfield reflect.StructField) {
	f := &Field{
		fieldName:  structfield.Name,
		fieldType:  structfield.Type,
		isAnon:     structfield.Anonymous,
		fieldIndex: append(a.index[:], structfield.Index...),
	}
	///
	f.tags = parseTags(string(structfield.Tag))
	f.alias = f.tags.Get(AliasTag).Val()
	_, has := structfield.Type.MethodByName("UnmarshalJSON")
	f.hasUnmarshal = has
	a.fields[f.Name()] = f
	///
	a.name2alias[f.Alias()] = f.Name()
	a.name2alias[f.Name()] = f.Alias()
}
