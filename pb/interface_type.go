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

func estimateSize(v interface{}, s int) int {
	switch v1 := v.(type) {
	case bool, int8, uint8, *bool, *int8, *uint8:
		return s + 3
	case int16, uint16, *int16, *uint16:
		return s + 5
	case int32, uint32, *int32, *uint32:
		return s + 2 + 4
	case int64, uint64, int, *int64, *uint64, *int:
		return s + 2 + 8
	case float32, *float32:
		return s + 2 + 4
	case float64, *float64:
		return s + 2 + 8
	case string:
		return s + 2 + 2 + len(v1)
	case *string:
		return s + 2 + 2 + len(*v1)
	case []byte:
		return s + 2 + 2 + len(v1)
	case *[]byte:
		return s + 2 + 2 + len(*v1)
	case []interface{}:
		for _, item := range v1 {
			s += estimateSize(item, s)
		}
		return s + 4
	case *[]interface{}:
		for _, item := range *v1 {
			s += estimateSize(item, s)
		}
		return s + 4
	default:
		panic(fmt.Sprintf("unknown data type %T, value=%+v", v1, v1))
	}
}

func (t *InterfaceType) Size() int {
	buf, err := t.Marshal()
	if err != nil {
		return -1
	}
	t.temp = buf // todo: ugly
	return len(buf)
	// return estimateSize(t.Value, 0)
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
