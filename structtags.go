package gtags

import (
	"fmt"
	"reflect"
)

type StructTags struct {
	fieldtags     map[string]*Tags
	fields        map[string]*Field
	anonfields    map[string]*Field
	anonstruct    map[string]*StructTags
	anonestedtags map[string]*Tags
	nestedtags    map[string]*StructTags
	isAnon        bool
}

func (a *StructTags) FieldNames() []string {
	///
	size := len(a.fieldtags) + len(a.anonestedtags) + len(a.nestedtags)
	names := make([]string, size)
	pos := 0
	for k, _ := range a.fieldtags {
		names[pos] = k
		pos += 1
	}
	for k, _ := range a.anonestedtags {
		names[pos] = k
		pos += 1
	}
	for k, _ := range a.nestedtags {
		names[pos] = k
		pos += 1
	}
	return names
}
func (a *StructTags) FieldByName(fieldname string) *Field {
	if a == nil {
		return nil
	}
	if v, ok := a.fields[fieldname]; ok {
		return v
	}
	if v, ok := a.anonfields[fieldname]; ok {
		return v
	}
	return nil
}
func (a *StructTags) Field(field string) *Tags {
	if a == nil {
		return nil
	}
	if v, ok := a.fieldtags[field]; ok {
		return v
	}
	if v, ok := a.anonestedtags[field]; ok {
		return v
	}
	return nil
}
func (a *StructTags) Nested(nestedfield string) *StructTags {
	if a == nil {
		return nil
	}
	return a.nestedtags[nestedfield]
}

///
////

/////
///notice: must be elem
func parseStructType(val reflect.Value, isAnon bool) *StructTags {
	// func parseStructType(typ reflect.Type, isAnon bool) *StructTags {
	// typ = indirectType(typ)
	typ := reflect.TypeOf(val)

	stags := &StructTags{
		fieldtags:     map[string]*Tags{},
		fields:        map[string]*Field{},
		anonfields:    map[string]*Field{},
		anonstruct:    map[string]*StructTags{},
		anonestedtags: map[string]*Tags{},
		nestedtags:    map[string]*StructTags{},
		isAnon:        isAnon,
	}
	///
	for i := 0; i < val.NumField(); i++ {
		valfield := val.Field(i)
		if !valfield.CanAddr() {
			fmt.Println(valfield.Type())
			fmt.Println(valfield.Type().Kind())
			fmt.Println(valfield)
		}
		// typfield := typ.Field(i)
		///
		// if typfield.Type.Kind() == reflect.Struct {
		if valfield.Kind() == reflect.Struct {
			typfield := typ.Field(i)
			if typfield.Anonymous == true {
				// if field.Name == field.Type.Name() {
				//anonymous
				anotags := parseStructType(valfield, true)
				for k, v := range anotags.fieldtags {
					stags.anonestedtags[k] = v
				}
				for k, v := range anotags.fields {
					stags.anonfields[k] = v
				}
			} else {
				//nested
				nesttags := parseStructType(valfield, false)
				stags.nestedtags[typfield.Name] = nesttags
			}
		} else {
			typfield := typ.Field(i)
			tags := parseTags(string(typfield.Tag))

			stags.fields[typfield.Name] = &Field{
				fieldName: typfield.Name,
				fieldType: typfield.Type,
				aliase:    tags.Get(AliaseTag).Val(),
				tags:      tags,
				isAnon:    isAnon,
				anonName:  typ.Name(),
			}
			stags.fieldtags[typfield.Name] = tags
		}
	}

	// typ.Name()
	// typ.PkgPath()
	// typ.Kind().String()
	return stags
}
