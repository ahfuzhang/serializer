package serializer

import (
	"fmt"
	"log"
	"reflect"
	"testing"

	stringutil "github.com/ahfuzhang/serializer/util/strings"
)

func getTestData() []interface{} {
	return []interface{}{
		[]interface{}{
			uint8(1), int8(2), uint16(3), int16(4), uint32(5), int32(6), uint64(7), int64(8), int(9),
			float32(10.1), float64(11.2), true, "aabb", []byte("AABB"),
		},
		[]interface{}{
			uint8(11), int8(12), uint16(13), int16(14), uint32(15), int32(16), uint64(17), int64(18), int(19),
			float32(110.1), float64(111.2), false, "ccdd", []byte("AABB"),
		},
		[]interface{}{
			uint8(21), int8(-22), uint16(23), int16(-24 * 128), uint32(0x7f09), int32(-99), uint64(17), int64(-18), int(-19),
			float32(-110.1), float64(-111.2), false, "eeff", []byte("EEFF"),
		},
		[]interface{}{
			map[string]interface{}{
				"a": 123,
				"b": "abc",
			},
		},
	}
}

func TestEncodeDecode(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	arr := getTestData()
	buf := make([]byte, 0, 1024*4)
	var err error
	buf, err = Encode(buf, 1, arr)
	if err != nil {
		t.Errorf("encode error, err=%+v", err)
		return
	}
	t.Logf("len=%d\n", len(buf))
	fmt.Println(stringutil.HexFormat(buf))
	//
	leftData, values, err := Decode(buf)
	if err != nil {
		t.Errorf("decode error, err=%+v", err)
		return
	}
	t.Logf("leftdata len=%d\n", len(leftData))
	t.Logf("\t type=%T, values = %+v\n", values, values)
	// if !reflect.DeepEqual(arr, values) {
	// 	t.Errorf("not equal")
	// 	return
	// }
}

func TestSreamlyRead(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	arr := getTestData()
	buf := make([]byte, 0, 1024*4)
	var err error
	buf, err = Encode(buf, 1, arr)
	if err != nil {
		t.Errorf("encode error, err=%+v", err)
		return
	}
	t.Logf("len=%d\n", len(buf))
	fmt.Println(stringutil.HexFormat(buf))
	//
	arrayData, headLen, leftData, tag, err := ReadArray(buf)
	if err != nil {
		t.Errorf("decode error, err=%+v", err)
		return
	}
	log.Println(tag, len(arrayData), len(leftData))
	fmt.Println(stringutil.HexFormat(arrayData))
	fmt.Println(stringutil.HexFormat(leftData))
	//
	{
		// first line
		arrayData, headLen, leftData, tag, err = ReadArray(arrayData[headLen:])
		if err != nil {
			t.Errorf("decode error, err=%+v", err)
			return
		}
		log.Println(tag, len(arrayData), len(leftData))
		// fmt.Println(stringutil.HexFormat(arrayData))
		// fmt.Println(stringutil.HexFormat(leftData))
		_, value1, err1 := Decode(arrayData)
		if err1 != nil {
			t.Errorf("decode error, %+v", err)
			return
		}
		t.Logf("%+v", value1)
	}
	//
	{
		arrayData, headLen, leftData, tag, err = ReadArray(leftData)
		if err != nil {
			t.Errorf("decode error, err=%+v", err)
			return
		}
		log.Println(tag, len(arrayData), len(leftData))
		// fmt.Println(stringutil.HexFormat(arrayData))
		// fmt.Println(stringutil.HexFormat(leftData))
		_, value1, err1 := Decode(arrayData)
		if err1 != nil {
			t.Errorf("decode error, %+v", err)
			return
		}
		t.Logf("%+v", value1)
	}
	{
		_, value1, err1 := Decode(leftData)
		if err1 != nil {
			t.Errorf("decode error, %+v", err)
			return
		}
		t.Logf("%+v", value1)
	}
}

func TestAppendArrayStart(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	arr := getTestData()
	buf := make([]byte, 0, 1024*4)
	var err error
	buf, err = Encode(buf, 1, arr)
	if err != nil {
		t.Errorf("encode error, err=%+v", err)
		return
	}
	t.Logf("len=%d\n", len(buf))
	// fmt.Println(stringutil.HexFormat(buf))
	//
	buf1 := make([]byte, 0, 1024*4)
	buf1 = AppendArrayStart(buf1, 1)
	buf1, err = Encode(buf1, 1, arr[0])
	if err != nil {
		t.Errorf("encode error, err=%+v", err)
		return
	}
	fmt.Println(stringutil.HexFormat(buf1))
	//
	// buf1 = AppendArrayStart(buf1, 1)
	buf1, err = Encode(buf1, 2, arr[1])
	if err != nil {
		t.Errorf("encode error, err=%+v", err)
		return
	}
	// buf1 = AppendArrayStart(buf1, 1)
	buf1, err = Encode(buf1, 3, arr[2])
	if err != nil {
		t.Errorf("encode error, err=%+v", err)
		return
	}
	buf1 = AppendArrayEnd(buf1, 1)
	if !reflect.DeepEqual(buf, buf1) {
		fmt.Println("not equal:")
		//t.Errorf("not equal, len1=%d, len2=%d", len(buf), len(buf1))
		fmt.Println(stringutil.HexFormat(buf))
		fmt.Println(stringutil.HexFormat(buf1))
		return
	}
}
