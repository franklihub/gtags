package gtags

import (
	"testing"

	"gotest.tools/assert"
)

type NestedParams struct {
	URL     string `json:"host" `
	TimeOut int    `json:"time_out" d:"10" max:"1000" min:"10"`
}
type CfgParams struct {
	NestedParams
	Conn     NestedParams `json:"conn"`
	Host     string       `json:"host" v:"required"`
	Port     int          `json:"port,omitempty"`
	Username string       `yaml:"username,omitempty" v:"required" d:"user"`
	Password string       `toml:"password" v:"required"`
}

func Test_Parse(t *testing.T) {
	stags := ParseStructTags(CfgParams{})
	conn := stags.NestedTags("Conn")
	tags := conn.Tags("URL")
	tag := tags.Get("json")
	assert.Equal(t, "json", tag.Key())
	assert.Equal(t, "host", tag.Val())

	tags = conn.Tags("TimeOut")
	assert.Equal(t, "json", tags.Get("json").Key())
	assert.Equal(t, "time_out", tags.Get("json").Val())
	assert.Equal(t, "1000", tags.Get("max").Val())
	assert.Equal(t, "10", tags.Get("d").Val())
	assert.Equal(t, "10", tags.Get("min").Val())

	tags = conn.Tags("Absent")
	assert.Equal(t, "", tags.Get("json").Key())
	// assert.Equal(t, `json:"username,omitempty" toml:"username" yaml:"username,omitempty"`, tag.String())
	// tag = NewSfTag()
	// tag.Parse(tags["Password"])
	// assert.Equal(t, "password", tag.Get("yaml"))
	// assert.Equal(t, `json:"password" toml:"password" yaml:"password"`, tag.String())
}
