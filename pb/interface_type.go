// Package pb for gogo proto, support a custom type for golang interface{} data
package pb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"

	"github.com/ahfuzhang/serializer"
	"github.com/ahfuzhang/serializer/util/debugs"
)

type InterfaceType struct {
	Value interface{}
	temp  []byte // todo: ugly way to solve MarshalTo problem
}

const (
	defaultBufferSize = 1024
	defaultArrayTag   = 1
)

func (t InterfaceType) Marshal() ([]byte, error) {
	arr := []interface{}{t.Value}
	out := make([]byte, 0, defaultBufferSize)
	var err error
	out, err = serializer.Encode(out, defaultArrayTag, arr)
	if err != nil {
		return out, debugs.WarpError(err, "serializer.Encode error")
	}
	return out, nil
}

func (t *InterfaceType) MarshalTo(data []byte) (n int, err error) {
	if len(t.temp) == len(data) {
		copy(data, t.temp)
		t.temp = nil
		return len(t.temp), nil
	}
	arr := []interface{}{t.Value}
	x := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	h := reflect.SliceHeader{Data: x.Data, Len: 0, Cap: x.Len}
	buf := *(*[]byte)(unsafe.Pointer(&h))
	buf, err = serializer.Encode(buf, 1, arr)
	if err != nil {
		return 0, debugs.WarpError(err, "serializer.Encode error")
	}
	if x.Data != h.Data {
		return 0, fmt.Errorf("[%s]buffer size not enought", debugs.SourceCodeLoc(1))
	}
	return len(buf), nil
}

func (t *InterfaceType) Unmarshal(data []byte) error {
	_, arr, err := serializer.Decode(data)
	if err != nil {
		return debugs.WarpError(err, "serializer.Decode error")
	}
	v, ok := arr.([]interface{})
	if !ok {
		return fmt.Errorf("[%s]decode data not a []interface{}", debugs.SourceCodeLoc(1))
	}
	if len(v) != 1 {
		return fmt.Errorf("[%s]decode []interface{} count not 1", debugs.SourceCodeLoc(1))
	}
	t.Value = v[0]
	return nil
}

func (t *InterfaceType) Size() int {
	buf, err := t.Marshal()
	if err != nil {
		return 0
	}
	t.temp = buf // todo: ugly
	return len(buf)
}

func (t InterfaceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Value)
}

func (t *InterfaceType) UnmarshalJSON(data []byte) error {
	decode := json.NewDecoder(bytes.NewReader(data))
	decode.UseNumber()
	return decode.Decode(&t.Value)
}

func (t InterfaceType) String() string {
	buf, err := json.Marshal(t.Value)
	if err != nil {
		return fmt.Sprintf("[%s]encode json error, err=%+v", debugs.SourceCodeLoc(1), err)
	}
	return string(buf)
}
