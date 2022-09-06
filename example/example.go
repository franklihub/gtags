package main

import (
	"fmt"
	"gtags"
	// "github.com/cryptoless/chain-raw-api-server/util/gtags"
)

type NestedParams struct {
	URL     string `json:"host"`
	TimeOut int    `json:"time_out" d:"10" max:"1000" min:"10"`
}
type CfgParams struct {
	NestedParams
	Conn NestedParams `json:"conn"`
	Host string       `json:"host" v:"required"`
	Port int          `json:"port,omitempty"`
}

func main() {
	cfg := CfgParams{}
	stags := gtags.ParseStructTags(cfg)
	fmt.Println(stags.FieldByName("URL").Tags().Get("json").Val())
	fmt.Println(stags.FieldByName("TimeOut").Tags().Get("max").Val())
	fmt.Println(stags.FieldByName("AbsentField").Tags().Get("json").Val())
}
