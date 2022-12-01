# serializer
Use protocol buffers binary format to encode/decode golang `[]interface{}`. To avoid alloc too many small object.

binary format is like:
```
ArrayStart: tag 1, type protowire.StartGroupType, 1 byte
  ArrayStart: tag $index+1, type protowire.StartGroupType, 1 byte
    col1:
     data type: tag 15, type protowire.VarintType, 1 byte
                type protowire.VarintType, value 1~14
     data: tag $col_index+1, type is decide by interface{} type, value is encoded data
    col2:
      ....             
  ArrayEnd: tag $index+1, type protowire.EndGroupType, 1 byte
ArrayEnd: tag 1, type protowire.EndGroupType, 1 byte
```

Those golang data:
```go
	arr := []interface{}{
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
	}
```
will encoded as:
```
0b 0b 78 03 08 01 78 02 10 02 78 05 18 03 78 04  |   x   x   x   x 
20 04 78 07 28 05 78 06 30 06 78 09 38 07 78 08  |   x ( x 0 x 8 x 
40 08 78 0a 48 09 78 0b 55 9a 99 21 41 78 0c 59  | @ x H x U  !Ax Y
66 66 66 66 66 66 26 40 78 01 60 01 78 0d 6a 04  | ffffff&@x ` x j 
61 61 62 62 78 0e 72 04 41 41 42 42 0c 13 78 03  | aabbx r AABB  x 
08 0b 78 02 10 0c 78 05 18 0d 78 04 20 0e 78 07  |   x   x   x   x 
28 0f 78 06 30 10 78 09 38 11 78 08 40 12 78 0a  | ( x 0 x 8 x @ x 
48 13 78 0b 55 33 33 dc 42 78 0c 59 cd cc cc cc  | H x U33 Bx Y    
cc cc 5b 40 78 01 60 00 78 0d 6a 04 63 63 64 64  |   [@x ` x j ccdd
78 0e 72 04 41 41 42 42 14 1b 78 03 08 15 78 02  | x r AABB  x   x 
10 ea ff ff ff ff ff ff ff ff 01 78 05 18 17 78  |            x   x
04 20 80 e8 ff ff ff ff ff ff ff 01 78 07 28 89  |             x ( 
fe 01 78 06 30 9d ff ff ff ff ff ff ff ff 01 78  |   x 0          x
09 38 11 78 08 40 ee ff ff ff ff ff ff ff ff 01  |  8 x @          
78 0a 48 ed ff ff ff ff ff ff ff ff 01 78 0b 55  | x H          x U
33 33 dc c2 78 0c 59 cd cc cc cc cc cc 5b c0 78  | 33  x Y      [ x
01 60 00 78 0d 6a 04 65 65 66 66 78 0e 72 04 45  |  ` x j eeffx r E
45 46 46 1c 0c                                   | EFF           
```
