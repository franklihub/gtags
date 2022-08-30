package main

import (
	"fmt"
	"gtags"
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
	fmt.Println(stags.Tags("URL").Get("json").Val())
	fmt.Println(stags.Tags("TimeOut").Get("max").Val())
	fmt.Println(stags.Tags("AbsentField").Get("json").Val())
}
