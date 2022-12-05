package pb

import (
	"fmt"
	"testing"

	"github.com/ahfuzhang/serializer/util/strings"
)

func TestInterfaceType_Marshal(t *testing.T) {
	v := &InterfaceType{Value: 123}
	buf, err := v.Marshal()
	if err != nil {
		t.Errorf("encode error, err=%s", err.Error())
		return
	}
	fmt.Println(strings.HexFormat(buf))
	fmt.Println(v.Size())
	//
	v1 := &InterfaceType{}
	err = v1.Unmarshal(buf)
	if err != nil {
		t.Errorf("decode error, err=%s", err.Error())
		return
	}
	fmt.Printf("%T, %+v\n\n", v1.Value, v1.Value)
	//
	buf = make([]byte, 100)
	n, err := v.MarshalTo(buf)
	if err != nil {
		t.Errorf("encode error, err=%s", err.Error())
		return
	}
	fmt.Println(len(buf), n)
	fmt.Println(strings.HexFormat(buf))
}
