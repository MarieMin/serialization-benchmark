package serialization_benchmark

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/vmihailenco/msgpack/v5"

	"github.com/MarieMin/serialization-benchmark/proto_buf"
)

func TestDataAllocationsSmall(_ *testing.T) {
	fmt.Printf("Small structure ---------\n")
	j, _ := json.Marshal(&SmallMock)
	p, _ := proto.Marshal(&SmallMockPB)
	m, _ := msgpack.Marshal(&SmallMock)
	f, _ := MarshalSmallStruct(SmallMock)

	printInfo(j, "JSON")
	printInfo(p, "ProtoBuf")
	printInfo(m, "MessagePack")
	printInfo(f, "FlatBuffers")
	fmt.Printf("\n")
}

func TestDataAllocationsMedium(_ *testing.T) {
	fmt.Printf("Medium structure ---------\n")
	j, _ := json.Marshal(&MediumMock)
	p, _ := proto.Marshal(&MediumMockPB)
	m, _ := msgpack.Marshal(&MediumMock)
	f, _ := MarshalMediumStruct(MediumMock)

	printInfo(j, "JSON")
	printInfo(p, "ProtoBuf")
	printInfo(m, "MessagePack")
	printInfo(f, "FlatBuffers")
	fmt.Printf("\n")
}

func TestDataAllocationsLarge(_ *testing.T) {
	fmt.Printf("Large structure ---------\n")
	j, _ := json.Marshal(&LargeMock)
	p, _ := proto.Marshal(&LargeMockPB)
	m, _ := msgpack.Marshal(&LargeMock)
	f, _ := MarshalLargeStruct(LargeMock)

	printInfo(j, "JSON")
	printInfo(p, "ProtoBuf")
	printInfo(m, "MessagePack")
	printInfo(f, "FlatBuffers")
	fmt.Printf("\n")
}

func BenchmarkJSONMarshal(b *testing.B) {
	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(SmallMock)
			_ = d
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&MediumMock)
			_ = d
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := json.Marshal(&LargeMock)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkProtobufMarshal(b *testing.B) {
	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&SmallMockPB)
			_ = d
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&MediumMockPB)
			_ = d
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := proto.Marshal(&LargeMockPB)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkMessagePackMarshal(b *testing.B) {
	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := msgpack.Marshal(SmallMock)
			_ = d
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := msgpack.Marshal(&MediumMock)
			_ = d
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := msgpack.Marshal(&LargeMock)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkFlatBuffersMarshal(b *testing.B) {
	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := MarshalSmallStruct(SmallMock)
			_ = d
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := MarshalMediumStruct(MediumMock)
			_ = d
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			d, _ := MarshalLargeStruct(LargeMock)
			_ = d
		}
	})
	fmt.Printf("\n")
}

func BenchmarkJSONUnmarshal(b *testing.B) {
	sSerialized, _ := json.Marshal(&SmallMock)
	mSerialized, _ := json.Marshal(&MediumMock)
	lSerialized, _ := json.Marshal(&LargeMock)

	var sDeserialized SmallStruct
	var mDeserialized MediumStruct
	var lDeserialized LargeStruct

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(sSerialized, &sDeserialized)
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(mSerialized, &mDeserialized)
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = json.Unmarshal(lSerialized, &lDeserialized)
		}
	})
	fmt.Printf("\n")
}

func BenchmarkProtobufUnmarshal(b *testing.B) {
	sSerialized, _ := proto.Marshal(&SmallMockPB)
	mSerialized, _ := proto.Marshal(&MediumMockPB)
	lSerialized, _ := proto.Marshal(&LargeMockPB)

	var sDeserialized proto_buf.SmallMessage
	var mDeserialized proto_buf.MediumMessage
	var lDeserialized proto_buf.LargeMessage

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(sSerialized, &sDeserialized)
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(mSerialized, &mDeserialized)
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = proto.Unmarshal(lSerialized, &lDeserialized)
		}
	})
	fmt.Printf("\n")
}

func BenchmarkMessagePackUnmarshal(b *testing.B) {
	sSerialized, _ := msgpack.Marshal(&SmallMock)
	mSerialized, _ := msgpack.Marshal(&MediumMock)
	lSerialized, _ := msgpack.Marshal(&LargeMock)

	var sDeserialized SmallStruct
	var mDeserialized MediumStruct
	var lDeserialized LargeStruct

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = msgpack.Unmarshal(sSerialized, &sDeserialized)
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = msgpack.Unmarshal(mSerialized, &mDeserialized)
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = msgpack.Unmarshal(lSerialized, &lDeserialized)
		}
	})
	fmt.Printf("\n")
}

func BenchmarkFlatBuffersUnmarshal(b *testing.B) {
	sSerialized, _ := MarshalSmallStruct(SmallMock)
	mSerialized, _ := MarshalMediumStruct(MediumMock)
	lSerialized, _ := MarshalLargeStruct(LargeMock)

	var sDeserialized SmallStruct
	var mDeserialized MediumStruct
	var lDeserialized LargeStruct

	b.ResetTimer()

	b.Run("SmallData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = UnmarshalSmallStruct(sSerialized, sDeserialized)
		}
	})
	b.Run("MediumData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = UnmarshalMediumStruct(mSerialized, mDeserialized)
		}
	})
	b.Run("LargeData", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			_ = UnmarshalLargeStruct(lSerialized, lDeserialized)
		}
	})
	fmt.Printf("\n")
}

func printInfo(d []byte, ser string) {
	used := len(d)
	allocated := cap(d)
	fmt.Printf("Type: %s \t\tData size: %d \t\tTotal Allocated: %d \t\t Used/Allocated: %.2f%%\n", ser, used, allocated, percentUsed(used, allocated)*100)
}

func percentUsed(used, allocated int) float32 {
	return float32(used) / float32(allocated)
}
