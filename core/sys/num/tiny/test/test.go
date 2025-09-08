package main

import (
	"math"
	"math/big"
)

func BuildTestCases() (pass []any, fail []any, providers []any) {
	// Add the underlying types we support

	pass = []any{
		int(1), int8(2), int16(3), int32(4), int64(5),
		uint(6), uint8(7), uint16(8), uint32(9), uint64(10),
		uintptr(11),
		float32(12.34), float64(56.78),
		Realized{}, Natural{}, Measurement{}, Named(1),
		"42", "-44", "11.11", "22.‾123", "22.33‾4", "~1.2345678", "~-8.7654321",
		big.NewInt(13), big.NewInt(-14),
		big.NewFloat(15.67), big.NewFloat(-16.89),
		// "ℯ", "-ℯ", "π", "-π", // We still need to implement the identity system
	}

	// Add pointers to all the above types

	toAdd := make([]any, len(pass))

	for i, c := range pass {
		toAdd[i] = &c
	}

	// Add double pointers to all the above types

	pass = append(pass, toAdd...)

	toAdd2 := make([]any, len(toAdd))
	for i, c := range toAdd {
		toAdd2[i] = &c
	}

	pass = append(pass, toAdd2...)

	// Add provider functions

	// Provider functions are laid out as follows:
	//
	// A -        func() any
	// B - func[T any]() T
	// C -        func() func() any
	// D - func[T any]() func() T
	// E -        func() func() func() any
	// F - func[T any]() func() func() T
	// G -        func() func() func() func() any
	// H - func[T any]() func() func() func() T
	//
	// X - no parameters
	// XX - uint
	// XXX - *uint
	// XXXX - ...uint
	// XXXXX - ...*uint

	providers = []any{
		// These return value types
		ProviderA, ProviderB[int], ProviderC, ProviderD[int], ProviderE, ProviderF[int], ProviderG, ProviderH[int],
		ProviderAA, ProviderBB[int], ProviderCC, ProviderDD[int], ProviderEE, ProviderFF[int], ProviderGG, ProviderHH[int],
		ProviderAAA, ProviderBBB[int], ProviderCCC, ProviderDDD[int], ProviderEEE, ProviderFFF[int], ProviderGGG, ProviderHHH[int],
		ProviderAAAA, ProviderBBBB[int], ProviderCCCC, ProviderDDDD[int], ProviderEEEE, ProviderFFFF[int], ProviderGGGG, ProviderHHHH[int],
		ProviderAAAAA, ProviderBBBBB[int], ProviderCCCCC, ProviderDDDDD[int], ProviderEEEEE, ProviderFFFFF[int], ProviderGGGGG, ProviderHHHHH[int],

		// These return pointer types
		PtrProviderA, PtrProviderB[int], PtrProviderC, PtrProviderD[int], PtrProviderE, PtrProviderF[int], PtrProviderG, PtrProviderH[int],
		PtrProviderAA, PtrProviderBB[int], PtrProviderCC, PtrProviderDD[int], PtrProviderEE, PtrProviderFF[int], PtrProviderGG, PtrProviderHH[int],
		PtrProviderAAA, PtrProviderBBB[int], PtrProviderCCC, PtrProviderDDD[int], PtrProviderEEE, PtrProviderFFF[int], PtrProviderGGG, PtrProviderHHH[int],
		PtrProviderAAAA, PtrProviderBBBB[int], PtrProviderCCCC, PtrProviderDDDD[int], PtrProviderEEEE, PtrProviderFFFF[int], PtrProviderGGGG, PtrProviderHHHH[int],
		PtrProviderAAAAA, PtrProviderBBBBB[int], PtrProviderCCCCC, PtrProviderDDDDD[int], PtrProviderEEEEE, PtrProviderFFFFF[int], PtrProviderGGGGG, PtrProviderHHHHH[int],
	}

	// Build a list of known fail modes

	fail = []any{
		math.NaN(), math.Inf(1), math.Inf(-1),
		"‾123.5", "1‾23.567", "123.-45", "123.45-67", "123.4567-",
		"1~23.456", "123~.456", "123.45~6", "123.456~",
		"-~123.456", "-1~23.456", "-123.~45", "-123.45~67", "-123.4567~",
		"ASDF98u2p0945adf", "", nil, AStruct{},
		big.Int{}, big.Float{}, big.Rat{},
		big.NewRat(1, 2),
	}

	return pass, fail, providers
}

type Named byte

type Realized struct{}
type Natural struct{}
type Measurement struct{}

type AStruct struct {
	A int
	B int
}

func ProviderA() any {
	return 13
}

func ProviderB[T any]() T {
	return any(42).(T)
}

func ProviderC() func() any {
	return ProviderA
}

func ProviderD[T any]() func() T {
	return ProviderB[T]
}

func ProviderE() func() func() any {
	return ProviderC
}

func ProviderF[T any]() func() func() T {
	return ProviderD[T]
}

func ProviderG() func() func() func() any {
	return ProviderE
}

func ProviderH[T any]() func() func() func() T {
	return ProviderF[T]
}

func ProviderAA(uint) any {
	return ProviderA()
}

func ProviderBB[T any](uint) T {
	return ProviderB[T]()
}

func ProviderCC(uint) func() any {
	return ProviderC()
}

func ProviderDD[T any](uint) func() T {
	return ProviderD[T]()
}

func ProviderEE(uint) func() func() any {
	return ProviderE()
}

func ProviderFF[T any](uint) func() func() T {
	return ProviderF[T]()
}

func ProviderGG(uint) func() func() func() any {
	return ProviderG()
}

func ProviderHH[T any](uint) func() func() func() T {
	return ProviderH[T]()
}

func ProviderAAA(*uint) any {
	return ProviderA()
}

func ProviderBBB[T any](*uint) T {
	return ProviderB[T]()
}

func ProviderCCC(*uint) func() any {
	return ProviderC()
}

func ProviderDDD[T any](*uint) func() T {
	return ProviderD[T]()
}

func ProviderEEE(*uint) func() func() any {
	return ProviderE()
}

func ProviderFFF[T any](*uint) func() func() T {
	return ProviderF[T]()
}

func ProviderGGG(*uint) func() func() func() any {
	return ProviderG()
}

func ProviderHHH[T any](*uint) func() func() func() T {
	return ProviderH[T]()
}

func ProviderAAAA(...uint) any {
	return ProviderA()
}

func ProviderBBBB[T any](...uint) T {
	return ProviderB[T]()
}

func ProviderCCCC(...uint) func() any {
	return ProviderC()
}

func ProviderDDDD[T any](...uint) func() T {
	return ProviderD[T]()
}

func ProviderEEEE(...uint) func() func() any {
	return ProviderE()
}

func ProviderFFFF[T any](...uint) func() func() T {
	return ProviderF[T]()
}

func ProviderGGGG(...uint) func() func() func() any {
	return ProviderG()
}

func ProviderHHHH[T any](...uint) func() func() func() T {
	return ProviderH[T]()
}

func ProviderAAAAA(...*uint) any {
	return ProviderA()
}

func ProviderBBBBB[T any](...*uint) T {
	return ProviderB[T]()
}

func ProviderCCCCC(...*uint) func() any {
	return ProviderC()
}

func ProviderDDDDD[T any](...*uint) func() T {
	return ProviderD[T]()
}

func ProviderEEEEE(...*uint) func() func() any {
	return ProviderE()
}

func ProviderFFFFF[T any](...*uint) func() func() T {
	return ProviderF[T]()
}

func ProviderGGGGG(...*uint) func() func() func() any {
	return ProviderG()
}

func ProviderHHHHH[T any](...*uint) func() func() func() T {
	return ProviderH[T]()
}

func PtrProviderA() *any {
	a := ProviderA()
	return &a
}

func PtrProviderB[T any]() *T {
	a := ProviderB[T]()
	return &a
}

func PtrProviderC() func() *any {
	return PtrProviderA
}

func PtrProviderD[T any]() func() *T {
	return PtrProviderB[T]
}

func PtrProviderE() func() func() *any {
	return PtrProviderC
}

func PtrProviderF[T any]() func() func() *T {
	return PtrProviderD[T]
}

func PtrProviderG() func() func() func() *any {
	return PtrProviderE
}

func PtrProviderH[T any]() func() func() func() *T {
	return PtrProviderF[T]
}

func PtrProviderAA(uint) *any {
	a := ProviderA()
	return &a
}

func PtrProviderBB[T any](uint) *T {
	a := ProviderB[T]()
	return &a
}

func PtrProviderCC(uint) func() *any {
	return PtrProviderA
}

func PtrProviderDD[T any](uint) func() *T {
	return PtrProviderB[T]
}

func PtrProviderEE(uint) func() func() *any {
	return PtrProviderC
}

func PtrProviderFF[T any](uint) func() func() *T {
	return PtrProviderD[T]
}

func PtrProviderGG(uint) func() func() func() *any {
	return PtrProviderE
}

func PtrProviderHH[T any](uint) func() func() func() *T {
	return PtrProviderF[T]
}

func PtrProviderAAA(*uint) *any {
	a := ProviderA()
	return &a
}

func PtrProviderBBB[T any](*uint) *T {
	a := ProviderB[T]()
	return &a
}

func PtrProviderCCC(*uint) func() *any {
	return PtrProviderA
}

func PtrProviderDDD[T any](*uint) func() *T {
	return PtrProviderB[T]
}

func PtrProviderEEE(*uint) func() func() *any {
	return PtrProviderC
}

func PtrProviderFFF[T any](*uint) func() func() *T {
	return PtrProviderD[T]
}

func PtrProviderGGG(*uint) func() func() func() *any {
	return PtrProviderE
}

func PtrProviderHHH[T any](*uint) func() func() func() *T {
	return PtrProviderF[T]
}

func PtrProviderAAAA(...uint) *any {
	a := ProviderA()
	return &a
}

func PtrProviderBBBB[T any](...uint) *T {
	a := ProviderB[T]()
	return &a
}

func PtrProviderCCCC(...uint) func() *any {
	return PtrProviderA
}

func PtrProviderDDDD[T any](...uint) func() *T {
	return PtrProviderB[T]
}

func PtrProviderEEEE(...uint) func() func() *any {
	return PtrProviderC
}

func PtrProviderFFFF[T any](...uint) func() func() *T {
	return PtrProviderD[T]
}

func PtrProviderGGGG(...uint) func() func() func() *any {
	return PtrProviderE
}

func PtrProviderHHHH[T any](...uint) func() func() func() *T {
	return PtrProviderF[T]
}

func PtrProviderAAAAA(...*uint) *any {
	a := ProviderA()
	return &a
}

func PtrProviderBBBBB[T any](...*uint) *T {
	a := ProviderB[T]()
	return &a
}

func PtrProviderCCCCC(...*uint) func() *any {
	return PtrProviderA
}

func PtrProviderDDDDD[T any](...*uint) func() *T {
	return PtrProviderB[T]
}

func PtrProviderEEEEE(...*uint) func() func() *any {
	return PtrProviderC
}

func PtrProviderFFFFF[T any](...*uint) func() func() *T {
	return PtrProviderD[T]
}

func PtrProviderGGGGG(...*uint) func() func() func() *any {
	return PtrProviderE
}

func PtrProviderHHHHH[T any](...*uint) func() func() func() *T {
	return PtrProviderF[T]
}
