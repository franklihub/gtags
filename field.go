package gtags

import "reflect"

///
var AliasTag = "json"

type FieldKind int

const (
	StringField FieldKind = iota
	IntField
	FloatField
	StructField
	SliceField
	MapField
	InterfaceField
)

type Field struct {
	fieldName string
	fieldType reflect.Type
	alias     string
	//
	tags *Tags
	//
	isAnon bool
	kind   FieldKind
}

func parseField(structfield reflect.StructField) *Field {

	field := &Field{
		fieldName: structfield.Name,
		fieldType: structfield.Type,
		//todo: anonymous
	}
	///
	field.tags = parseTags(string(structfield.Tag))
	field.alias = field.tags.Get(AliasTag).Val()
	///
	return field
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
