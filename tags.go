package gtags

import "strings"

type Tags struct {
	tags map[string]*Tag
}

func parseTags(s string) *Tags {
	tags := &Tags{
		tags: map[string]*Tag{},
	}

	for _, t := range strings.Split(s, " ") {
		if tag := parseTag(t); tag != nil {
			tags.tags[tag.key] = tag
		}
	}

	return tags
}
func (a *Tags) Get(field string) *Tag {
	if a == nil {
		return nil
	}
	return a.tags[field]
}
