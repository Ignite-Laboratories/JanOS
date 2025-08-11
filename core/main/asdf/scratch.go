package main

func MyFunc() MyStruct[int] {
	return MyStruct[int]{
		MyField: func() int {
			return 0
		},
	}
}
