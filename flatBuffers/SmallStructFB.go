// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatBuffers

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type SmallStructFB struct {
	_tab flatbuffers.Table
}

func GetRootAsSmallStructFB(buf []byte, offset flatbuffers.UOffsetT) *SmallStructFB {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &SmallStructFB{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *SmallStructFB) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *SmallStructFB) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *SmallStructFB) TestInt32() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SmallStructFB) MutateTestInt32(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func (rcv *SmallStructFB) TestInt64() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *SmallStructFB) MutateTestInt64(n int64) bool {
	return rcv._tab.MutateInt64Slot(6, n)
}

func (rcv *SmallStructFB) TestFloat() float32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetFloat32(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *SmallStructFB) MutateTestFloat(n float32) bool {
	return rcv._tab.MutateFloat32Slot(8, n)
}

func (rcv *SmallStructFB) TestDouble() float64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetFloat64(o + rcv._tab.Pos)
	}
	return 0.0
}

func (rcv *SmallStructFB) MutateTestDouble(n float64) bool {
	return rcv._tab.MutateFloat64Slot(10, n)
}

func (rcv *SmallStructFB) TestBool() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

func (rcv *SmallStructFB) MutateTestBool(n bool) bool {
	return rcv._tab.MutateBoolSlot(12, n)
}

func (rcv *SmallStructFB) TestBytes(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *SmallStructFB) TestBytesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *SmallStructFB) MutateTestBytes(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *SmallStructFB) TestString() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func SmallStructFBStart(builder *flatbuffers.Builder) {
	builder.StartObject(7)
}
func SmallStructFBAddTestInt32(builder *flatbuffers.Builder, testInt32 int32) {
	builder.PrependInt32Slot(0, testInt32, 0)
}
func SmallStructFBAddTestInt64(builder *flatbuffers.Builder, testInt64 int64) {
	builder.PrependInt64Slot(1, testInt64, 0)
}
func SmallStructFBAddTestFloat(builder *flatbuffers.Builder, testFloat float32) {
	builder.PrependFloat32Slot(2, testFloat, 0.0)
}
func SmallStructFBAddTestDouble(builder *flatbuffers.Builder, testDouble float64) {
	builder.PrependFloat64Slot(3, testDouble, 0.0)
}
func SmallStructFBAddTestBool(builder *flatbuffers.Builder, testBool bool) {
	builder.PrependBoolSlot(4, testBool, false)
}
func SmallStructFBAddTestBytes(builder *flatbuffers.Builder, testBytes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(testBytes), 0)
}
func SmallStructFBStartTestBytesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func SmallStructFBAddTestString(builder *flatbuffers.Builder, testString flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(testString), 0)
}
func SmallStructFBEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
