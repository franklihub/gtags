package gtags

type StructTags struct {
	fieldtags     map[string]*Tags
	anonestedtags map[string]*Tags
	nestedtags    map[string]*StructTags
}

func (a *StructTags) Tags(field string) *Tags {
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
func (a *StructTags) NestedTags(nestedfield string) *StructTags {
	if a == nil {
		return nil
	}
	return a.nestedtags[nestedfield]
}
