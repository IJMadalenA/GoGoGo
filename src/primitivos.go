package main

func main() {
	println(boolFun())
	println(stringFun())
	println(intFun())
	println(uintFun())
	println(floatFun())
	println(complexFun())
}

// primitive data types.
// bool.
func boolFun() bool {
	return true
}

// string.
func stringFun() string {
	return "Hello World."
}

// int  int8  int16  int32  int64
func intFun() int {
	return 12
}

// uint uint8 uint16 uint32 uint64 uintptr
func uintFun() uint {
	return 12
}

// float32 float64
func floatFun() float64 {
	return 12.12
}

// complex64 complex128
func complexFun() complex128 {
	return 12i
}
