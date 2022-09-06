package gtags

import "reflect"

///
var AliasTag = "json"

type Field struct {
	fieldName string
	fieldType reflect.Type
	//
	isAnon bool
	//
	alias string
	tags  *Tags
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
