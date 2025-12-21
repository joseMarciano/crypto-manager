package testdata

type TestStruct struct {
	Field1 string  `validate:"required"`
	Field2 int     `validate:"gte=2"`
	Field3 int     `validate:"gt=1"`
	Field4 float64 `validate:"two-decimals"`
}

func DefaultTestStruct() TestStruct {
	return TestStruct{
		Field1: "default",
		Field2: 2,
		Field3: 2,
		Field4: 1.13,
	}
}
