#### go不同进制之间的转换

- 10进制转2进制
``` go
fmt.SPrintf("%b",9)
```
- 10进制转8进制
``` go
fmt.SPrintf("%o",9)
```
- 10进制转16进制
``` go
fmt.SPrintf("%x",9)
```
- 2进制转10进制
``` go
strconv.ParseInt("1001", 2, 8)  // int8
strconv.ParseInt("1001", 2, 16) // int16
strconv.ParseInt("1001", 2, 32) // int32
strconv.ParseInt("1001", 2, 64) // int64
```
- 8进制转10进制
``` go
strconv.ParseInt("0d1001", 2, 8)  // int8
strconv.ParseInt("0d1001", 2, 16) // int16
strconv.ParseInt("0d1001", 2, 32) // int32
strconv.ParseInt("0d1001", 2, 64) // int64
```
- 16进制转10进制
``` go
strconv.ParseInt("0x1001", 2, 8)  // int8
strconv.ParseInt("0x1001", 2, 16) // int16
strconv.ParseInt("0x1001", 2, 32) // int32
strconv.ParseInt("0x1001", 2, 64) // int64
```