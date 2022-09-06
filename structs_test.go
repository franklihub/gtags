package gtags

import (
	"testing"

	"gotest.tools/assert"
)

type HeartCfg int
type NestedParams struct {
	URL     string `json:"host" `
	TimeOut int    `json:"time_out" d:"10" max:"1000" min:"10"`
}
type CfgParams struct {
	NestedParams
	Conn     NestedParams `json:"conn"`
	Host     string       `json:"host" v:"required"`
	Port     int          `json:"port,omitempty"`
	Username string       `json:"username,omitempty" v:"required" d:"user"`
	Password string       `json:"password" v:"required"`
	Heart    HeartCfg     `json:"heart" d:"allways"`
}

func Test_Parse(t *testing.T) {
	//
	stags := ParseStructTags(CfgParams{})
	conn := stags.NestedByName("Conn")
	field := conn.FieldByName("URL")
	tag := field.Tags().Get("json")
	assert.Equal(t, "json", tag.Key())
	assert.Equal(t, "host", tag.Val())

	field = conn.FieldByName("TimeOut")
	tags := field.Tags()
	assert.Equal(t, "json", tags.Get("json").Key())
	assert.Equal(t, "time_out", tags.Get("json").Val())
	assert.Equal(t, "1000", tags.Get("max").Val())
	assert.Equal(t, "10", tags.Get("d").Val())
	assert.Equal(t, "10", tags.Get("min").Val())

	field = conn.FieldByName("Absent")
	tags = field.Tags()
	assert.Equal(t, "", tags.Get("json").Key())
	// assert.Equal(t, `p:"username,omitempty" toml:"username" yaml:"username,omitempty"`, tag.String())
	// tag = NewSfTag()
	// tag.Parse(tags["Password"])
	// assert.Equal(t, "password", tag.Get("yaml"))
	// assert.Equal(t, `p:"password" toml:"password" yaml:"password"`, tag.String())
}
