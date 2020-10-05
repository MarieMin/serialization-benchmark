package serialization_benchmark

import (
	"github.com/MarieMin/serialization-benchmark/proto_buf"
)

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

var (

	// Mocks for test JSON, MessagePack and FlatBuffers

	SmallMock = SmallStruct{
		TestInt32:  1024,
		TestInt64:  20482048,
		TestFloat:  0.14285715,
		TestDouble: 0.14285714285714285,
		TestBool:   true,
		TestBytes:  []byte("string in bytes representation"),
		TestString: "protobuf-serialization-benchmark",
	}

	MediumMock = MediumStruct{
		TestInt32:  134217728,
		TestInt64:  107374182400,
		TestFloat:  0.14285715,
		TestDouble: 0.14285714285714285,
		TestBool:   false,
		TestBytes: []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras eu consequat enim. Aenean in 
						lacus quis sapien consectetur pulvinar. Cras libero massa, facilisis at tincidunt quis, maximus 
						id mi. Aenean ultricies maximus sem vel volutpat. Nulla malesuada scelerisque risus et luctus.`),
		TestString: `Nunc convallis mi sem, sit amet finibus felis fermentum id. Nulla auctor rutrum fermentum. Aenean
						in lacus quis sapien consectetur pulvinar. Cras libero massa, facilisis at tincidunt quis, 
						maximus id mi. Aenean ultricies maximus sem vel volutpat. Nulla malesuada scelerisque risus et 
						luctus.`,
		TestIntArray: []int64{2, 67, 10000, 4546546, 345678, -1, -345678},
		Nested:       SmallMock,
	}

	LargeMock = LargeStruct{
		NestedMediumMock: []MediumStruct{MediumMock, MediumMock, MediumMock},
		NestedSmallMock:  []SmallStruct{SmallMock, SmallMock},
		TestString: `Nulla auctor rutrum fermentum. Donec rhoncus malesuada odio, ut porttitor leo pulvinar 
								vel. Curabitur vitae ex vestibulum risus sagittis consequat in ac urna.  Mauris blandit 
								convallis hendrerit.Suspendisse potenti. Ut bibendum nunc a nisl aliquet, ac malesuada 
								urna suscipit. Morbi et magna sem. Donec non diam eu nisi eleifend scelerisque. 	
								Suspendisse eu leo vehicula, fermentum ipsum non, tincidunt lorem. Pellentesque justo 
								urna, malesuada eu dui porta, semper dictum ex. Curabitur ut convallis sem.`,
		TestBytes: []byte(`Nunc finibus congue augue, id porttitor tortor condimentum et. Curabitur ut 
								convallis sem. Quisque quis nulla dignissim, viverra lacus ut, luctus turpis. 
								Suspendisse sed tincidunt nulla, nec elementum mi. Curabitur vel elementum sapien, vitae 
								mattis lectus. Integer id pharetra mi. Donec lacinia est ut vestibulum imperdiet. Sed 
								dolor erat, lacinia ut pretium vel, porta non est.`),
	}

	// Mocks for test ProtoBuf
	SmallMockPB = proto_buf.SmallMessage{
		TestInt32:  1024,
		TestInt64:  20482048,
		TestFloat:  0.14285715,
		TestDouble: 0.14285714285714285,
		TestBool:   true,
		TestBytes:  []byte("string in bytes representation"),
		TestString: "protobuf-serialization-benchmark",
	}

	MediumMockPB = proto_buf.MediumMessage{
		TestInt32:  134217728,
		TestInt64:  107374182400,
		TestFloat:  0.14285715,
		TestDouble: 0.14285714285714285,
		TestBool:   false,
		TestBytes: []byte(`Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras eu consequat enim. Aenean in 
						lacus quis sapien consectetur pulvinar. Cras libero massa, facilisis at tincidunt quis, maximus 
						id mi. Aenean ultricies maximus sem vel volutpat. Nulla malesuada scelerisque risus et luctus.`),
		TestString: `Nunc convallis mi sem, sit amet finibus felis fermentum id. Nulla auctor rutrum fermentum. Aenean
						in lacus quis sapien consectetur pulvinar. Cras libero massa, facilisis at tincidunt quis, 
						maximus id mi. Aenean ultricies maximus sem vel volutpat. Nulla malesuada scelerisque risus et 
						luctus.`,
		TestIntArray: []int64{2, 67, 10000, 4546546, 345678, -1, -345678},
		Nested:       &SmallMockPB,
	}

	LargeMockPB = proto_buf.LargeMessage{
		MediumNestedStructure: []*proto_buf.MediumMessage{&MediumMockPB, &MediumMockPB, &MediumMockPB},
		SmallNestedStructure:  []*proto_buf.SmallMessage{&SmallMockPB, &SmallMockPB},
		TestString: `Nulla auctor rutrum fermentum. Donec rhoncus malesuada odio, ut porttitor leo pulvinar 
								vel. Curabitur vitae ex vestibulum risus sagittis consequat in ac urna.  Mauris blandit 
								convallis hendrerit.Suspendisse potenti. Ut bibendum nunc a nisl aliquet, ac malesuada 
								urna suscipit. Morbi et magna sem. Donec non diam eu nisi eleifend scelerisque. 	
								Suspendisse eu leo vehicula, fermentum ipsum non, tincidunt lorem. Pellentesque justo 
								urna, malesuada eu dui porta, semper dictum ex. Curabitur ut convallis sem.`,
		TestBytes: []byte(`Nunc finibus congue augue, id porttitor tortor condimentum et. Curabitur ut 
								convallis sem. Quisque quis nulla dignissim, viverra lacus ut, luctus turpis. 
								Suspendisse sed tincidunt nulla, nec elementum mi. Curabitur vel elementum sapien, vitae 
								mattis lectus. Integer id pharetra mi. Donec lacinia est ut vestibulum imperdiet. Sed 
								dolor erat, lacinia ut pretium vel, porta non est.`),
	}
)
