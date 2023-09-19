// Package serializer use protocol buffers binary format to encode/decode `[]interface{}`
package serializer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"strconv"

	"google.golang.org/protobuf/encoding/protowire"

	"github.com/ahfuzhang/serializer/util/debugs"
)

const (
	tagOfDataType = 15 // use binary(01111) as special tag id, for data type info
)

const defaultArrayCount = 10

// all golang basic data types
const (
	tBool = iota + 1
	tInt8
	tUint8
	tInt16
	tUint16
	tInt32
	tUint32
	tInt64
	tUint64
	tInt
	tFloat32
	tFloat64
	tString
	tBytes
	tJSON
)

// Encode encode []interface{} to binary
func Encode(buf []byte, tag int, v interface{}) ([]byte, error) {
	switch v1 := v.(type) {
	case bool:
		buf = setType(buf, tBool)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, protowire.EncodeBool(v1))
	case *bool:
		buf = setType(buf, tBool)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, protowire.EncodeBool(*v1))
	case int8:
		buf = setType(buf, tInt8)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(v1))
	case *int8:
		buf = setType(buf, tInt8)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*v1))
	case uint8:
		buf = setType(buf, tUint8)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(v1))
	case *uint8:
		buf = setType(buf, tUint8)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*v1))
	case int16:
		buf = setType(buf, tInt16)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(v1))
	case *int16:
		buf = setType(buf, tInt16)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*v1))
	case uint16:
		buf = setType(buf, tUint16)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(v1))
	case *uint16:
		buf = setType(buf, tUint16)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*v1))
	case int32:
		buf = setType(buf, tInt32)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(v1))
	case *int32:
		buf = setType(buf, tInt32)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*v1))
	case uint32:
		buf = setType(buf, tUint32)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(v1))
	case *uint32:
		buf = setType(buf, tUint32)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*v1))
	case int64:
		buf = setType(buf, tInt64)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(v1))
	case *int64:
		buf = setType(buf, tInt64)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*v1))
	case uint64:
		buf = setType(buf, tUint64)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, v1)
	case *uint64:
		buf = setType(buf, tUint64)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, *v1)
	case int:
		buf = setType(buf, tInt)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(v1))
	case *int:
		buf = setType(buf, tInt)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.VarintType)
		buf = protowire.AppendVarint(buf, uint64(*v1))
	case float32:
		buf = setType(buf, tFloat32)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.Fixed32Type)
		buf = protowire.AppendFixed32(buf, math.Float32bits(v1))
	case *float32:
		buf = setType(buf, tFloat32)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.Fixed32Type)
		buf = protowire.AppendFixed32(buf, math.Float32bits(*v1))
	case float64:
		buf = setType(buf, tFloat64)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.Fixed64Type)
		buf = protowire.AppendFixed64(buf, math.Float64bits(v1))
	case *float64:
		buf = setType(buf, tFloat64)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.Fixed64Type)
		buf = protowire.AppendFixed64(buf, math.Float64bits(*v1))
	case string:
		buf = setType(buf, tString)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.BytesType)
		buf = protowire.AppendString(buf, v1)
	case *string:
		buf = setType(buf, tString)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.BytesType)
		buf = protowire.AppendString(buf, *v1)
	case []byte:
		buf = setType(buf, tBytes)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.BytesType)
		buf = protowire.AppendBytes(buf, v1)
	case *[]byte:
		buf = setType(buf, tBytes)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.BytesType)
		buf = protowire.AppendBytes(buf, *v1)
	case []interface{}:
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.StartGroupType)
		var err error
		for idx, item := range v1 {
			buf, err = Encode(buf, idx+1, item)
			if err != nil {
				return buf, debugs.WarpError(err, fmt.Sprintf("encode item %d error, value=%+v", idx, item))
			}
		}
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.EndGroupType)
	default:
		// try to use json encode
		temp, err := json.Marshal(v)
		if err != nil {
			return buf, fmt.Errorf("[%s]not support datatype(and can not encode to JSON), %T, value=%+v",
				debugs.SourceCodeLoc(1), v, v)
		}
		buf = setType(buf, tJSON)
		buf = protowire.AppendTag(buf, protowire.Number(tag), protowire.BytesType)
		buf = protowire.AppendBytes(buf, temp)
	}
	return buf, nil
}

// Decode decode binary to []interface{}
func Decode(buf []byte) ([]byte, interface{}, error) {
	tag, typeOfField, totalLen := protowire.ConsumeField(buf)
	if totalLen < 0 {
		return buf, nil, fmt.Errorf("[%s]read field error,code=%d", debugs.SourceCodeLoc(1), totalLen)
	}
	golangType := uint64(0)
	if tag == tagOfDataType && typeOfField == protowire.VarintType {
		buf = buf[totalLen-1:]
		var headLen int
		golangType, headLen = protowire.ConsumeVarint(buf)
		if headLen < 0 {
			return buf, nil, fmt.Errorf("[%s]read field error,code=%d", debugs.SourceCodeLoc(1), headLen)
		}
		buf = buf[headLen:]
		tag, typeOfField, totalLen = protowire.ConsumeField(buf)
		if totalLen < 0 {
			return buf, nil, fmt.Errorf("[%s]read field error,code=%d", debugs.SourceCodeLoc(1), totalLen)
		}
	}
	_, headLen := protowire.ConsumeVarint(buf)
	buf = buf[headLen:]
	switch typeOfField {
	case protowire.VarintType:
		value, dataLen := protowire.ConsumeVarint(buf)
		if dataLen < 0 {
			return buf, nil, fmt.Errorf("[%s]read VarintType error,code=%d", debugs.SourceCodeLoc(1), dataLen)
		}
		buf = buf[dataLen:]
		v, err := uint64ToInterfaceType(value, golangType)
		if err != nil {
			return buf, v, debugs.WarpError(err, "uint64ToInterfaceType error")
		}
		return buf, v, nil
	case protowire.Fixed32Type:
		value, dataLen := protowire.ConsumeFixed32(buf)
		if dataLen < 0 {
			return buf, nil, fmt.Errorf("[%s]read Fixed32Type error,code=%d", debugs.SourceCodeLoc(1), dataLen)
		}
		buf = buf[dataLen:]
		v, err := uint64ToInterfaceType(uint64(value), golangType)
		if err != nil {
			return buf, v, debugs.WarpError(err, "uint64ToInterfaceType error")
		}
		return buf, v, nil
	case protowire.Fixed64Type:
		value, dataLen := protowire.ConsumeFixed64(buf)
		if dataLen < 0 {
			return buf, nil, fmt.Errorf("[%s]read Fixed64Type error,code=%d", debugs.SourceCodeLoc(1), dataLen)
		}
		buf = buf[dataLen:]
		v, err := uint64ToInterfaceType(value, golangType)
		if err != nil {
			return buf, v, debugs.WarpError(err, "uint64ToInterfaceType error")
		}
		return buf, v, nil
	case protowire.BytesType:
		switch golangType {
		case tString:
			value, dataLen := protowire.ConsumeString(buf)
			if dataLen < 0 {
				return buf, nil, fmt.Errorf("[%s]read BytesType error,code=%d", debugs.SourceCodeLoc(1), dataLen)
			}
			buf = buf[dataLen:]
			return buf, value, nil
		case tBytes:
			value, dataLen := protowire.ConsumeBytes(buf)
			if dataLen < 0 {
				return buf, nil, fmt.Errorf("[%s]read BytesType error,code=%d", debugs.SourceCodeLoc(1), dataLen)
			}
			buf = buf[dataLen:]
			return buf, value, nil
		case tJSON:
			value, dataLen := protowire.ConsumeBytes(buf)
			if dataLen < 0 {
				return buf, nil, fmt.Errorf("[%s]read BytesType error,code=%d", debugs.SourceCodeLoc(1), dataLen)
			}
			buf = buf[dataLen:]
			decoder := json.NewDecoder(bytes.NewBuffer(value))
			decoder.UseNumber()
			var out interface{}
			if err := decoder.Decode(&out); err != nil {
				return buf, nil, fmt.Errorf("[%s]decode json error,err=%s", debugs.SourceCodeLoc(1), err.Error())
			}
			return buf, out, nil
		default:
			return buf, nil, fmt.Errorf("[%s]read BytesType error,golangType=%d", debugs.SourceCodeLoc(1), golangType)
		}
	case protowire.StartGroupType:
		out := make([]interface{}, 0, defaultArrayCount)
		for len(buf) > 0 {
			leftData, value, err := Decode(buf)
			if err != nil {
				return buf, out, debugs.WarpError(err, "decode array item error")
			}
			buf = leftData
			out = append(out, value)
			_, nextType, nextHeadLen := protowire.ConsumeTag(buf)
			if nextHeadLen < 0 {
				return buf, out, debugs.WarpError(err, "decode array item end flag error")
			}
			if nextType == protowire.EndGroupType {
				buf = buf[headLen:]
				return buf, out, nil
			}
		}
		return buf, out, nil
	default:
		return buf, nil, fmt.Errorf("[%s]unknown field tag=%d", debugs.SourceCodeLoc(1), typeOfField)
	}
}

// AppendArrayStart add the array header to buffer, for serialize data streamly
func AppendArrayStart(buf []byte, tag int) []byte {
	return protowire.AppendTag(buf, protowire.Number(tag), protowire.StartGroupType)
}

// AppendArrayEnd add the array end flag to buffer
func AppendArrayEnd(buf []byte, tag int) []byte {
	return protowire.AppendTag(buf, protowire.Number(tag), protowire.EndGroupType)
}

// ReadArray read whole array item as separate []byte, to read data streamly
func ReadArray(buf []byte) (arrayData []byte, headLen int, leftData []byte, tag int, err error) {
	arrTag, typeOfField, totalLen := protowire.ConsumeField(buf)
	if totalLen < 0 {
		err = fmt.Errorf("[%s]read field error,code=%d", debugs.SourceCodeLoc(1), totalLen)
		return
	}
	if typeOfField != protowire.StartGroupType {
		err = fmt.Errorf("[%s]not a array, type=%d", debugs.SourceCodeLoc(1), typeOfField)
		return
	}
	_, headLen = protowire.ConsumeVarint(buf)
	arrayData = buf[:totalLen]
	leftData = buf[totalLen:]
	tag = int(arrTag)
	return
}

// RowCallback func type to read rows
type RowCallback func(tag int, cols ...interface{}) error

// ReadEachRow read rows, send data to callback func
func ReadEachRow(buf []byte, callback RowCallback) error {
	arrayData, headLen, leftData, tag, err := ReadArray(buf)
	if err != nil {
		return debugs.WarpError(err, "ReadArray error")
	}
	arrayData = arrayData[headLen:]
	for len(arrayData) > 2 {
		arrayData, _, leftData, tag, err = ReadArray(arrayData)
		if err != nil {
			return debugs.WarpError(err, "ReadArray read row error")
		}
		var values interface{}
		_, values, err = Decode(arrayData)
		if err != nil {
			return debugs.WarpError(err, "Decode row error")
		}
		arrayData = leftData
		//
		cols, ok := values.([]interface{})
		if !ok {
			if err = callback(tag, values); err != nil {
				return debugs.WarpError(err, "callback with one value error")
			}
		} else {
			if err = callback(tag, cols...); err != nil {
				return debugs.WarpError(err, "callback with multi value error")
			}
		}
	}
	return nil
}

func setType(buf []byte, t uint64) []byte {
	buf = protowire.AppendTag(buf, protowire.Number(tagOfDataType), protowire.VarintType)
	buf = protowire.AppendVarint(buf, t)
	// buf = protowire.AppendTag(buf, protowire.Number(t), protowire.VarintType) // use this way to reduce 1 byte, but read become complex
	return buf
}

func uint64ToInterfaceType(v uint64, golangType uint64) (interface{}, error) {
	switch golangType {
	case tBool:
		switch v {
		case 0:
			return false, nil
		case 1:
			return true, nil
		default:
			return nil, fmt.Errorf("[%s]not a bool value, %d", debugs.SourceCodeLoc(1), v)
		}
	case tInt8:
		return int8(v), nil
	case tUint8:
		return uint8(v), nil
	case tInt16:
		return int16(v), nil
	case tUint16:
		return uint16(v), nil
	case tInt32:
		return int32(v), nil
	case tUint32:
		return uint32(v), nil
	case tInt64:
		return int64(v), nil
	case tUint64:
		return v, nil
	case tInt:
		return int(v), nil
	case tFloat32:
		return math.Float32frombits(uint32(v)), nil
	case tFloat64:
		return math.Float64frombits(v), nil
	default:
		return nil, fmt.Errorf("[%s]not a number type, %d", debugs.SourceCodeLoc(1), golangType)
	}
}

// BasicTypeToString format a basic type to string
func BasicTypeToString(v interface{}) (string, error) {
	switch r := v.(type) {
	case bool:
		return strconv.FormatBool(r), nil
	case int8:
		return strconv.FormatInt(int64(r), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(r), 10), nil
	case int16:
		return strconv.FormatInt(int64(r), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(r), 10), nil
	case int32:
		return strconv.FormatInt(int64(r), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(r), 10), nil
	case int64:
		return strconv.FormatInt(r, 10), nil
	case uint64:
		return strconv.FormatUint(r, 10), nil
	case int:
		return strconv.FormatInt(int64(r), 10), nil
	case float32:
		return fmt.Sprintf("%f", r), nil
	case float64:
		return fmt.Sprintf("%f", r), nil
	case string:
		return r, nil
	default:
		return "", fmt.Errorf("[%s]not support type %T", debugs.SourceCodeLoc(1), v)
	}
}
