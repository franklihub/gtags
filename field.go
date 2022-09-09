package gtags

import "reflect"

///
var AliasTag = "json"

type FieldKind int

// const (
// 	StringField FieldKind = iota
// 	IntField
// 	FloatField
// 	StructField
// 	SliceField
// 	MapField
// 	InterfaceField
// )

type Field struct {
	fieldName    string
	fieldType    reflect.Type
	fieldIndex   int
	isAnon       bool
	hasUnmarshal bool
	//
	alias string
	tags  *Tags
	//
	index []int
}

func (a *Field) HasUnmarshal() bool {
	return a.hasUnmarshal
}

func (a *Field) Index() []int {
	return a.index
}
func (a *Field) FieldType() reflect.Type {
	return a.fieldType
}

func (a *Field) Name() string {
	return a.fieldName
}
func (a *Field) Alias() string {
	return a.alias
}
func (a *Field) Type() reflect.Type {
	return a.fieldType
}

func (a *Field) Tags() *Tags {
	if a == nil {
		return nil
	}
	return a.tags
}
