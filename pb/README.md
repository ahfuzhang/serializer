pb.Any is hard to use.
but we can use [gogo proto](https://github.com/gogo/protobuf) to define a custom type.
Look this example:
```protobuf
// QueryRsp Query response format
message QueryRsp {
  int32 code = 1;  // error code
  string msg = 2;  // error msg
  repeated bytes datas = 6[(gogoproto.customtype) = "InterfaceType"];   // return a tow dimension interface{} array
}
```

then copy interface_type.go to xx.pb.go dir, and your pb type will support interface{} now.
see: https://www.cnblogs.com/ahfuzhang/p/16858832.html (Chinese)

