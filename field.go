package gtags

import "reflect"

///
var AliasTag = "json"

type Field struct {
	fieldName string
	fieldType reflect.Type
	//
	isAnon   bool
	anonName string
	//
	alias string
	tags  *Tags
}

func parseField(structfield reflect.StructField) *Field {

	field := &Field{
		fieldName: structfield.Name,
		fieldType: structfield.Type,
		//todo: anonymous
		isAnon:   structfield.Anonymous,
		anonName: structfield.Type.Name(),
	}
	///
	field.tags = parseTags(string(structfield.Tag))
	field.anonName = field.tags.Get(AliasTag).Val()
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
func (a *Field) AnonName() string {
	if a.isAnon {
		return a.anonName
	}
	return ""
}
func (a *Field) Tags() *Tags {
	if a == nil {
		return nil
	}
	return a.tags
}
