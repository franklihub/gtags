package gtags

import "reflect"

///
var AliasTag = "json"

type Field struct {
	fieldName string
	fieldType reflect.Type
	alias     string
	///
	tags *Tags
	//
	isAnon   bool
	anonName string
}

func (a *Field) Name() string {
	return a.fieldName
}
func (a *Field) Type() reflect.Type {
	return a.fieldType
}
func (a *Field) Aliase() string {
	return a.alias
}
func (a *Field) AnonName() string {
	if a.isAnon {
		return a.anonName
	}
	return ""
}
func (a *Field) Tags() *Tags {
	return a.tags
}
