package gtags

import "strings"

type Tag struct {
	key  string
	val  string
	opts []string
}

func (a *Tag) Key() string {
	if a == nil {
		return ""
	}
	return a.key
}
func (a *Tag) Val() string {
	if a == nil {
		return ""
	}
	return a.val
}
func (a *Tag) Opts() []string {
	if a == nil {
		return nil
	}
	return a.opts
}

func (a *Tag) HasOpt(opt string) bool {
	if a == nil {
		return false
	}
	for _, o := range a.opts {
		if o == opt {
			return true
		}
	}
	return false
}

func parseTag(s string) *Tag {
	keyIndex := strings.Index(s, ":")
	if keyIndex == -1 {
		return nil
	}

	values := strings.Split(strings.Trim(s[keyIndex+1:], "\""), ",")
	if len(values) == 0 {
		return nil
	}

	return &Tag{
		key:  s[:keyIndex],
		val:  values[0],
		opts: values[1:],
	}
}
