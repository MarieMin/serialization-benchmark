# Benchmarks of Go serialization methods

### Serializers

This project contain tests of next serializers: JSON, ProtoBuf, MessagePack, FlatBuffers:
- [encoding/json](http://golang.org/pkg/encoding/json/)
- [golang/protobuf](http://github.com/golang/protobuf)
- [vmihailenco/msgpack/v4](https://github.com/vmihailenco/msgpack)
- [google/flatbuffers](http://github.com/google/flatbuffers)

## Running the benchmarks

```bash
go get -u -t
go test -bench=.
go test -bench=. -cpuprofile=cpu.out .
```

### Test data

```
type SmallStruct struct {
	TestInt32  int32   `json:"testInt32"`
	TestInt64  int64   `json:"testInt64"`
	TestFloat  float32 `json:"testFloat"`
	TestDouble float64 `json:"testDouble"`
	TestBool   bool    `json:"testBool"`
	TestBytes  []byte  `json:"testBytes"`
	TestString string  `json:"testString"`
}

type MediumStruct struct {
	TestInt32    int32       `json:"testInt32"`
	TestInt64    int64       `json:"testInt64"`
	TestFloat    float32     `json:"testFloat"`
	TestDouble   float64     `json:"testDouble"`
	TestBool     bool        `json:"testBool"`
	TestBytes    []byte      `json:"testBytes"`
	TestString   string      `json:"testString"`
	TestIntArray []int64     `json:"testIntArray"`
	Nested       SmallStruct `json:"nested"`
}

type LargeStruct struct {
	NestedMediumMock []MediumStruct `json:"nestedMediumMock"`
	NestedSmallMock  []SmallStruct  `json:"nestedSmallMock"`
	TestString       string         `json:"testString"`
	TestBytes        []byte         `json:"testBytes"`
}
```

### Results

```
Memory allocations

Small structure ---------
Type: JSON                      Data size: 214          Total Allocated: 224             Used/Allocated: 95.54%
Type: ProtoBuf                  Data size: 90           Total Allocated: 90              Used/Allocated: 100.00%
Type: MessagePack               Data size: 167          Total Allocated: 306             Used/Allocated: 54.58%
Type: FlatBuffers               Data size: 144          Total Allocated: 144             Used/Allocated: 100.00%

Medium structure ---------
Type: JSON                      Data size: 1158         Total Allocated: 1280            Used/Allocated: 90.47%
Type: ProtoBuf                  Data size: 763          Total Allocated: 763             Used/Allocated: 100.00%
Type: MessagePack               Data size: 964          Total Allocated: 1449            Used/Allocated: 66.53%
Type: FlatBuffers               Data size: 896          Total Allocated: 896             Used/Allocated: 100.00%

Large structure ---------
Type: JSON                      Data size: 5131         Total Allocated: 5376            Used/Allocated: 95.44%
Type: ProtoBuf                  Data size: 3452         Total Allocated: 3452            Used/Allocated: 100.00%
Type: MessagePack               Data size: 4253         Total Allocated: 6363            Used/Allocated: 66.84%
Type: FlatBuffers               Data size: 3928         Total Allocated: 4696            Used/Allocated: 83.65%


Performance
                                              Iterations              
BenchmarkJSONMarshal/SmallData-4                  998442              1169 ns/op             304 B/op          2 allocs/op
BenchmarkJSONMarshal/MediumData-4                 271840              4241 ns/op            1696 B/op          2 allocs/op
BenchmarkJSONMarshal/LargeData-4                   66457             17422 ns/op            7201 B/op          5 allocs/op

BenchmarkProtobufMarshal/SmallData-4             3412033               340 ns/op              96 B/op          1 allocs/op
BenchmarkProtobufMarshal/MediumData-4            1000000              2014 ns/op             768 B/op          1 allocs/op
BenchmarkProtobufMarshal/LargeData-4              322490              5410 ns/op            3456 B/op          1 allocs/op

BenchmarkMessagePackMarshal/SmallData-4           938059              1328 ns/op             656 B/op          5 allocs/op
BenchmarkMessagePackMarshal/MediumData-4          344727              4389 ns/op            2368 B/op          5 allocs/op
BenchmarkMessagePackMarshal/LargeData-4           105535             18197 ns/op           12097 B/op          7 allocs/op

BenchmarkFlatBuffersMarshal/SmallData-4           904742              1655 ns/op             856 B/op         14 allocs/op
BenchmarkFlatBuffersMarshal/MediumData-4          405182              3100 ns/op            3208 B/op         19 allocs/op
BenchmarkFlatBuffersMarshal/LargeData-4           139308             11596 ns/op           13904 B/op         27 allocs/op

BenchmarkJSONUnmarshal/SmallData-4                274182              4005 ns/op             328 B/op          9 allocs/op
BenchmarkJSONUnmarshal/MediumData-4                72571             16127 ns/op            1480 B/op         24 allocs/op
BenchmarkJSONUnmarshal/LargeData-4                 17272             70339 ns/op            5952 B/op         79 allocs/op

BenchmarkProtobufUnmarshal/SmallData-4           3566096               380 ns/op              64 B/op          2 allocs/op
BenchmarkProtobufUnmarshal/MediumData-4           762135              1468 ns/op             952 B/op          9 allocs/op
BenchmarkProtobufUnmarshal/LargeData-4            181597              6739 ns/op            4792 B/op         43 allocs/op

BenchmarkMessagePackUnmarshal/SmallData-4         914812              1174 ns/op              80 B/op          2 allocs/op
BenchmarkMessagePackUnmarshal/MediumData-4        333187              3496 ns/op             432 B/op          4 allocs/op
BenchmarkMessagePackUnmarshal/LargeData-4          89833             13555 ns/op            1904 B/op         15 allocs/op

BenchmarkFlatBuffersUnmarshal/SmallData-4        1295595               922 ns/op              32 B/op          1 allocs/op
BenchmarkFlatBuffersUnmarshal/MediumData-4        132144              9024 ns/op             672 B/op          3 allocs/op
BenchmarkFlatBuffersUnmarshal/LargeData-4         102468             11721 ns/op            1776 B/op          9 allocs/op
```

### CPU Profiling

Commands for getting profiling information
```bash
go test -bench=. -cpuprofile=cpu.out .
go tool pprof cpu.out
(pprof) top50 -cum
``` 

Results
```bash
       cum   cum%
      0.94s  2.20%  github.com/MarieMin/serialization-benchmark.BenchmarkProtobufMarshal.func2
      0.97s  2.27%  github.com/MarieMin/serialization-benchmark.BenchmarkProtobufUnmarshal.func3
      1s  2.34%  github.com/MarieMin/serialization-benchmark.BenchmarkMessagePackMarshal.func2
      1.02s  2.39%  github.com/MarieMin/serialization-benchmark.BenchmarkFlatBuffersMarshal.func2
      1.07s  2.51%  github.com/MarieMin/serialization-benchmark.BenchmarkFlatBuffersMarshal.func3
      1.09s  2.55%  github.com/MarieMin/serialization-benchmark.BenchmarkProtobufMarshal.func3
      1.11s  2.60%  github.com/MarieMin/serialization-benchmark.BenchmarkJSONUnmarshal.func1
      1.16s  2.72%  github.com/MarieMin/serialization-benchmark.BenchmarkJSONMarshal.func2
      1.18s  2.76%  github.com/MarieMin/serialization-benchmark.BenchmarkFlatBuffersUnmarshal.func2
      1.23s  2.88%  github.com/MarieMin/serialization-benchmark.BenchmarkMessagePackMarshal.func3
      1.24s  2.91%  github.com/MarieMin/serialization-benchmark.BenchmarkFlatBuffersMarshal.func1
      1.36s  3.19%  github.com/MarieMin/serialization-benchmark.BenchmarkJSONUnmarshal.func2
      1.36s  3.19%  github.com/MarieMin/serialization-benchmark.BenchmarkJSONMarshal.func3
      1.41s  3.30%  github.com/MarieMin/serialization-benchmark.BenchmarkFlatBuffersUnmarshal.func3
      1.51s  3.54%  github.com/MarieMin/serialization-benchmark.BenchmarkProtobufMarshal.func1
      1.69s  3.96%  github.com/MarieMin/serialization-benchmark.BenchmarkFlatBuffersUnmarshal.func1
      1.75s  4.10%  github.com/MarieMin/serialization-benchmark.BenchmarkJSONUnmarshal.func3
      1.78s  4.17%  github.com/MarieMin/serialization-benchmark.BenchmarkProtobufUnmarshal.func2
      1.83s  4.29%  github.com/MarieMin/serialization-benchmark.BenchmarkProtobufUnmarshal.func1
      1.88s  4.40%  github.com/MarieMin/serialization-benchmark.BenchmarkJSONMarshal.func1
      2.02s  4.73%  github.com/MarieMin/serialization-benchmark.BenchmarkMessagePackMarshal.func1
      2.14s  5.01%  github.com/MarieMin/serialization-benchmark.BenchmarkMessagePackUnmarshal.func3
      3.69s  8.65%  github.com/MarieMin/serialization-benchmark.BenchmarkMessagePackUnmarshal.func2
      3.82s  8.95%  github.com/MarieMin/serialization-benchmark.BenchmarkMessagePackUnmarshal.func1
```
