package avro_test

type TestInterface interface {
	SomeFunc() int
}

type TestRecord struct {
	A int64  `avro:"a"`
	B string `avro:"b"`
}

func (*TestRecord) SomeFunc() int {
	return 0
}

type TestPartialRecord struct {
	B string `avro:"b"`
}

type TestNestedRecord struct {
	A TestRecord `avro:"a"`
	B TestRecord `avro:"b"`
}

type TestUnion struct {
	A interface{} `avro:"a"`
}
