package serialization_benchmark

import (
	flatbuffers "github.com/google/flatbuffers/go"

	serialization_benchmark "github.com/MarieMin/serialization-benchmark/flatBuffers"
)

func MarshalSmallStruct(small SmallStruct) ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)
	builder.Reset()

	s := prepareSmallStructOffset(small, builder)

	builder.Finish(s)

	return builder.Bytes[builder.Head():], nil
}

func MarshalMediumStruct(med MediumStruct) ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)
	builder.Reset()

	m := prepareMediumStructOffset(med, builder)

	builder.Finish(m)
	return builder.Bytes[builder.Head():], nil
}

func MarshalLargeStruct(large LargeStruct) ([]byte, error) {
	builder := flatbuffers.NewBuilder(0)
	builder.Reset()

	var NestedSmallMockLen = len(large.NestedSmallMock)
	smallVectorTemp := make([]flatbuffers.UOffsetT, NestedSmallMockLen)
	for i := 0; i < NestedSmallMockLen; i++ {
		smallVectorTemp = append(smallVectorTemp, prepareSmallStructOffset(large.NestedSmallMock[i], builder))
	}

	var NestedMediumMockLen = len(large.NestedMediumMock)
	mediumVectorTemp := make([]flatbuffers.UOffsetT, NestedMediumMockLen)
	for i := 0; i < NestedMediumMockLen; i++ {
		mediumVectorTemp = append(mediumVectorTemp, prepareMediumStructOffset(large.NestedMediumMock[i], builder))
	}

	serialization_benchmark.LargeStructFBStartNestedSmallStructVector(builder, NestedSmallMockLen)
	for i := NestedSmallMockLen - 1; i >= 0; i-- {
		builder.PrependUOffsetT(smallVectorTemp[i])
	}
	testSmallArray := builder.EndVector(NestedSmallMockLen)

	serialization_benchmark.LargeStructFBStartNestedMediumStructVector(builder, NestedMediumMockLen)
	for i := NestedMediumMockLen - 1; i >= 0; i-- {
		builder.PrependUOffsetT(mediumVectorTemp[i])
	}
	testMediumArray := builder.EndVector(NestedMediumMockLen)

	testString := builder.CreateString(large.TestString)
	testBytesString := builder.CreateByteString(large.TestBytes)

	serialization_benchmark.LargeStructFBStart(builder)
	serialization_benchmark.LargeStructFBAddNestedSmallStruct(builder, testSmallArray)
	serialization_benchmark.LargeStructFBAddNestedMediumStruct(builder, testMediumArray)
	serialization_benchmark.LargeStructFBAddTestBytes(builder, testBytesString)
	serialization_benchmark.LargeStructFBAddTestString(builder, testString)

	builder.Finish(serialization_benchmark.LargeStructFBEnd(builder))
	return builder.Bytes[builder.Head():], nil
}

func UnmarshalSmallStruct(d []byte, small SmallStruct) error {
	fb := serialization_benchmark.SmallStructFB{}
	fb.Init(d, flatbuffers.GetUOffsetT(d))

	smallFBToSmall(&fb, small)

	return nil
}

func UnmarshalMediumStruct(d []byte, med MediumStruct) error {
	fb := serialization_benchmark.MediumStructFB{}
	fb.Init(d, flatbuffers.GetUOffsetT(d))

	mediumFBToMedium(&fb, med)

	return nil
}

func UnmarshalLargeStruct(d []byte, large LargeStruct) error {
	fb := serialization_benchmark.LargeStructFB{}
	fb.Init(d, flatbuffers.GetUOffsetT(d))

	var SmallArrayLen = fb.NestedSmallStructLength()
	smallVector := make([]*SmallStruct, SmallArrayLen)
	for i := 0; i < SmallArrayLen; i++ {
		var smallFB = &serialization_benchmark.SmallStructFB{}
		var small = &SmallStruct{}
		fb.NestedSmallStruct(smallFB, i)
		smallVector[i] = small
	}

	var MediumArrayLen = fb.NestedMediumStructLength()
	medVector := make([]*MediumStruct, MediumArrayLen)
	for i := 0; i < MediumArrayLen; i++ {
		var mediumFB = &serialization_benchmark.MediumStructFB{}
		var med = &MediumStruct{}
		fb.NestedMediumStruct(mediumFB, i)
		medVector[i] = med
	}

	large.TestBytes = make([]byte, fb.TestBytesLength())
	for i := 0; i < fb.TestBytesLength(); i++ {
		large.TestBytes[i] = byte(fb.TestBytes(i))
	}
	large.TestString = string(fb.TestString())

	return nil
}

func prepareSmallStructOffset(small SmallStruct, builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	testString := builder.CreateString(small.TestString)
	testBytesString := builder.CreateByteString(small.TestBytes)

	serialization_benchmark.SmallStructFBStart(builder)
	serialization_benchmark.SmallStructFBAddTestInt32(builder, small.TestInt32)
	serialization_benchmark.SmallStructFBAddTestInt64(builder, small.TestInt64)
	serialization_benchmark.SmallStructFBAddTestFloat(builder, small.TestFloat)
	serialization_benchmark.SmallStructFBAddTestDouble(builder, small.TestDouble)
	serialization_benchmark.SmallStructFBAddTestBool(builder, small.TestBool)
	serialization_benchmark.SmallStructFBAddTestBytes(builder, testBytesString)
	serialization_benchmark.SmallStructFBAddTestString(builder, testString)

	return serialization_benchmark.SmallStructFBEnd(builder)
}

func prepareMediumStructOffset(med MediumStruct, builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	testString := builder.CreateString(med.TestString)
	testBytesString := builder.CreateByteString(med.TestBytes)

	serialization_benchmark.MediumStructFBStartTestIntArrayVector(builder, len(med.TestIntArray))
	for i := 0; i < len(med.TestIntArray); i++ {
		builder.PrependInt64(med.TestIntArray[i])
	}
	testIntArray := builder.EndVector(len(med.TestIntArray))

	testNested := prepareSmallStructOffset(med.Nested, builder)

	serialization_benchmark.MediumStructFBStart(builder)

	serialization_benchmark.MediumStructFBAddTestInt32(builder, med.TestInt32)
	serialization_benchmark.MediumStructFBAddTestInt64(builder, med.TestInt64)
	serialization_benchmark.MediumStructFBAddTestFloat(builder, med.TestFloat)
	serialization_benchmark.MediumStructFBAddTestDouble(builder, med.TestDouble)
	serialization_benchmark.MediumStructFBAddTestBool(builder, med.TestBool)
	serialization_benchmark.MediumStructFBAddTestBytes(builder, testBytesString)
	serialization_benchmark.MediumStructFBAddTestString(builder, testString)
	serialization_benchmark.MediumStructFBAddTestIntArray(builder, testIntArray)
	serialization_benchmark.MediumStructFBAddTestNested(builder, testNested)

	return serialization_benchmark.MediumStructFBEnd(builder)
}

func smallFBToSmall(fb *serialization_benchmark.SmallStructFB, small SmallStruct) {
	small.TestInt32 = fb.TestInt32()
	small.TestInt64 = fb.TestInt64()
	small.TestFloat = fb.TestFloat()
	small.TestDouble = fb.TestDouble()
	small.TestBool = fb.TestBool()
	small.TestBytes = make([]byte, fb.TestBytesLength())
	for i := 0; i < fb.TestBytesLength(); i++ {
		small.TestBytes[i] = byte(fb.TestBytes(i))
	}
	small.TestString = string(fb.TestString())
}

func mediumFBToMedium(fb *serialization_benchmark.MediumStructFB, med MediumStruct) {
	med.TestInt32 = int32(int(fb.TestInt32()))
	med.TestInt64 = int64(int(fb.TestInt64()))
	med.TestFloat = fb.TestFloat()
	med.TestDouble = fb.TestDouble()
	med.TestBool = fb.TestBool()
	med.TestBytes = make([]byte, fb.TestBytesLength())
	for i := 0; i < fb.TestBytesLength(); i++ {
		med.TestBytes[i] = byte(fb.TestBytes(i))
	}
	med.TestString = string(fb.TestString())

	var nestFB = fb.TestNested(&serialization_benchmark.SmallStructFB{})
	smallFBToSmall(nestFB, med.Nested)
}
