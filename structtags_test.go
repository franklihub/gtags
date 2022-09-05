package gtags

import (
	"testing"

	"gotest.tools/assert"
)

type HeartCfg int
type NestedParams struct {
	URL     string `p:"host" `
	TimeOut int    `p:"time_out" d:"10" max:"1000" min:"10"`
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
	AliaseTag = "p"
	//
	stags := ParseStructTags(CfgParams{})
	conn := stags.Nested("Conn")
	tags := conn.Field("URL")
	tag := tags.Get("p")
	assert.Equal(t, "p", tag.Key())
	assert.Equal(t, "host", tag.Val())

	tags = conn.Field("TimeOut")
	assert.Equal(t, "p", tags.Get("p").Key())
	assert.Equal(t, "time_out", tags.Get("p").Val())
	assert.Equal(t, "1000", tags.Get("max").Val())
	assert.Equal(t, "10", tags.Get("d").Val())
	assert.Equal(t, "10", tags.Get("min").Val())

	tags = conn.Field("Absent")
	assert.Equal(t, "", tags.Get("p").Key())
	// assert.Equal(t, `p:"username,omitempty" toml:"username" yaml:"username,omitempty"`, tag.String())
	// tag = NewSfTag()
	// tag.Parse(tags["Password"])
	// assert.Equal(t, "password", tag.Get("yaml"))
	// assert.Equal(t, `p:"password" toml:"password" yaml:"password"`, tag.String())
}
