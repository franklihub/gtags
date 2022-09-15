package gtags

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"gotest.tools/assert"
)

type typeint int
type address [40]byte
type combin struct {
	TI    typeint `json:"ti"`
	addr  address `json:"address"`
	IsPtr bool    `json:"is_ptr"`
}
type nested struct {
	combin
	nested *combin `json:"nested"`
}

func Test_Field_combin(t *testing.T) {
	typ := reflect.TypeOf(combin{})
	fs := []*Field{}
	f := newField(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		subf := f.ParseStructField(field)
		fs = append(fs, subf)
	}

	assert.Equal(t, fs[0].fieldType.Kind(), reflect.Int)
	assert.Equal(t, fs[1].fieldType.Kind(), reflect.Array)
	assert.Equal(t, fs[2].fieldType.Kind(), reflect.Bool)
}

func Test_Field_nested(t *testing.T) {
	typ := reflect.TypeOf(nested{})
	fs := []*Field{}
	f := newField(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		subf := f.ParseStructField(field)
		fs = append(fs, subf)
	}
	///
	assert.Equal(t, len(fs), 2)
	assert.Equal(t, len(fs[0].subFields), 3)
	anonsub := fs[0].subFields
	assert.Equal(t, len(fs[1].subFields), 3)
	nested := fs[1].subFields
	///
	///anon
	assert.Equal(t, fs[0].isAnon, true)
	assert.Equal(t, anonsub[0].fieldType.Kind(), reflect.Int)
	assert.Equal(t, anonsub[1].fieldType.Kind(), reflect.Array)
	assert.Equal(t, anonsub[2].fieldType.Kind(), reflect.Bool)
	assert.Equal(t, anonsub[0].tags.Get("json").Val(), "ti")
	assert.Equal(t, anonsub[1].tags.Get("json").Val(), "address")
	assert.Equal(t, anonsub[2].tags.Get("json").Val(), "is_ptr")
	///nested
	assert.Equal(t, fs[1].isAnon, false)
	assert.Equal(t, nested[0].fieldType.Kind(), reflect.Int)
	assert.Equal(t, nested[1].fieldType.Kind(), reflect.Array)
	assert.Equal(t, nested[2].fieldType.Kind(), reflect.Bool)
	//
	assert.Equal(t, nested[0].tags.Get("json").Val(), "ti")
	assert.Equal(t, nested[1].tags.Get("json").Val(), "address")
	assert.Equal(t, nested[2].tags.Get("json").Val(), "is_ptr")
}

///
type sliceStruct struct {
	AddressList []address `json:"address_list"`
	NestedList  []nested  `json:"nested_list`
}

func Test_Field_slice(t *testing.T) {
	// AccessList           string
	typ := reflect.TypeOf(sliceStruct{})
	fs := []*Field{}
	f := newField(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		subf := f.ParseStructField(field)
		fs = append(fs, subf)
	}
	///
	assert.Equal(t, len(fs), 2)
	assert.Equal(t, fs[0].fieldType.Kind(), reflect.Slice)
	assert.Equal(t, fs[0].tags.Get("json").Val(), "address_list")
	assert.Equal(t, len(fs[0].subFields), 0)
	// slicesub := fs[0].subFields
	assert.Equal(t, fs[1].fieldType.Kind(), reflect.Slice)
	assert.Equal(t, fs[1].tags.Get("json").Val(), "nested_list")
	assert.Equal(t, len(fs[1].subFields), 2)
	///nested
	nestedlist := fs[1]
	anoncombin := nestedlist.subFields[0]
	assert.Equal(t, anoncombin.isAnon, true)
	assert.Equal(t, len(anoncombin.subFields), 3)

	nestedcombin := nestedlist.subFields[1]
	assert.Equal(t, nestedcombin.isAnon, false)
	assert.Equal(t, len(nestedcombin.subFields), 3)
}
func Test_ReflectKind(t *testing.T) {
	//
	slice := []rpc.BlockNumberOrHash{}
	typ := reflect.TypeOf(slice)
	fmt.Println("relect.TypeOf:", typ)
	fmt.Println("relect.TypeOf.Kind:", typ.Kind())
	fmt.Println("relect.TypeOf.Elem.Type:", typ.Elem().Kind())
	blocks := []rpc.BlockNumber{}
	typ = reflect.TypeOf(blocks)
	fmt.Println("relect.TypeOf:", typ)
	fmt.Println("relect.TypeOf.Kind:", typ.Kind())
	fmt.Println("relect.TypeOf.Elem.Type:", typ.Elem().Kind())
	addrs := []common.Address{}
	typ = reflect.TypeOf(addrs)
	fmt.Println("relect.TypeOf:", typ)
	fmt.Println("relect.TypeOf.Kind:", typ.Kind())
	fmt.Println("relect.TypeOf.Elem.Type:", typ.Elem().Kind())
	fmt.Println("relect.TypeOf.Elem.Elem.Type:", typ.Elem().Elem().Kind())
}

type URL struct {
	Addr string `json:"addr"`
	Port int    `json:"port" d:"1234"`
}
type Db struct {
	URL
	Pass string `json:"pass"`
}
type Cfg struct {
	Db
	Name string `json:"name"`
}

func Test_AnonAnon(t *testing.T) {
	typ := reflect.TypeOf(Cfg{})
	fs := []*Field{}
	f := newField(typ)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		subf := f.ParseStructField(field)
		fs = append(fs, subf)
	}
	//
	assert.Equal(t, f.subFields[0].Name(), "Db")
	assert.Equal(t, f.subFields[0].subFields[0].Name(), "URL")
	assert.Equal(t, f.subFields[0].subFields[0].subFields[0].Alias(), "addr")
	assert.Equal(t, f.subFields[0].subFields[0].subFields[1].Alias(), "port")

}
