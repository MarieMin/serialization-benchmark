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
go test -bench='.*' ./
```


command for compile proto buf
```
protoc --go_out=. *.proto
```

command for compile flat buf
```
flatc --go schema.fbs
```