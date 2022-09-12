package gtags

import "reflect"

///
var AliasTag = "json"

type FieldKind int

type Field struct {
	fieldName     string
	fieldType     reflect.Type
	fieldTypeName string
	fieldIndex    []int
	isAnon        bool
	hasUnmarshal  bool
	alias         string
	tags          *Tags
	//
	subFields []*Field
}

///
func (a *Field) Fields() []*Field {
	fs := []*Field{}
	for _, f := range a.subFields {
		if f.IsSlice() || f.IsStruct() {
			continue
		}
		///
		fs = append(fs, f)
	}
	return fs
}
func (a *Field) AnonFields() []*Field {
	fs := []*Field{}
	for _, f := range a.subFields {
		if !f.IsStruct() {
			continue
		}
		if !f.isAnon {
			continue
		}
		///
		fss := f.Fields()
		fs = append(fs, fss...)
	}
	return fs
}
func (a *Field) Nesteds() []*Field {
	fs := []*Field{}
	for _, f := range a.subFields {
		if !f.IsStruct() {
			continue
		}
		if f.isAnon {
			continue
		}
		///
		fs = append(fs, f)
	}
	return fs
}

///
///
func (a *Field) FieldNames() []string {
	names := []string{}
	for _, f := range a.Fields() {
		names = append(names, f.fieldName)
	}

	return names
}
func (a *Field) FieldAlias() []string {
	names := []string{}
	for _, f := range a.Fields() {
		names = append(names, f.alias)
	}

	return names
}

//
func (a *Field) FieldByName(name string) *Field {
	for _, f := range a.subFields {
		if f.IsSlice() || f.IsStruct() {
			continue
		}
		///
		if f.fieldName == name ||
			f.alias == name {
			return f
		}
	}
	///
	sf := a.AnonByName(name)
	if sf != nil {
		return sf
	}

	return nil
}

//
func (a *Field) AnonByName(name string) *Field {
	for _, f := range a.subFields {
		if !f.IsStruct() {
			continue
		}
		if !f.isAnon {
			continue
		}
		///
		sf := f.FieldByName(name)
		if sf != nil {
			return sf
		}
		///
	}
	return nil
}

///
func (a *Field) NestedByName(name string) *Field {
	for _, f := range a.subFields {
		if !f.IsStruct() {
			continue
		}
		if f.isAnon {
			continue
		}
		///
		if f.fieldName == name ||
			f.alias == name {
			return f
		}
	}

	return nil
}

///
///
func (a *Field) IsStruct() bool {
	return a.fieldType.Kind() == reflect.Struct
}
func (a *Field) IsSlice() bool {
	return a.fieldType.Kind() == reflect.Slice
}

//
func (a *Field) HasUnmarshal() bool {
	return a.hasUnmarshal
}

func (a *Field) IsAnon() bool {
	return a.isAnon
}

///
func (a *Field) Index() []int {
	return a.fieldIndex
}
func (a *Field) Type() reflect.Type {
	return a.fieldType
}

func (a *Field) Name() string {
	return a.fieldName
}
func (a *Field) Alias() string {
	return a.alias
}

func (a *Field) Tags() *Tags {
	if a == nil {
		return nil
	}
	return a.tags
}

///
////
func (a *Field) addStruct(typ reflect.Type) {

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		f := ParseStructField(field)
		//
		f.fieldIndex = append(a.fieldIndex, f.fieldIndex...)
		///
		a.subFields = append(a.subFields, f)
		///
	}
}

func (a *Field) addSlice(typ reflect.Type) {
	// func (a *Field) addSlice(structfield reflect.StructField) {

	a.parseType(typ.Elem())
	// f := newField(structfield.Type.Elem())
	// if f != nil {
	// 	a.subFields = append(a.subFields, f)
	// }
	///
}
