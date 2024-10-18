// This package provides helper functions for creating pointers to basic types.
//
// The functions in this package are useful when you need to pass a pointer to a basic type to a function that expects a pointer.
// For example, when you need to pass a pointer to a basic type to a function that expects a pointer.
//
// Example:
//
//	 package main
//
//	 import "github.com/scheiblingco/gofn/pointers"
//
//	 func main() {
//	 	// Create a slice of pointers to int
//	 	slice := []*int{
//	 		pointers.Int(1),
//	 		pointers.Int(2),
//	 		pointers.Int(3),
//	 	}
//	}
package pointers

// Returns a pointer to a string
func String(s string) *string {
	return &s
}

// Returns a pointer to an int
func Int(i int) *int {
	return &i
}

// Returns a pointer to an int8
func Int8(i int8) *int8 {
	return &i
}

// Returns a pointer to an int16
func Int16(i int16) *int16 {
	return &i
}

// Returns a pointer to an int32
func Int32(i int32) *int32 {
	return &i
}

// Returns a pointer to an int64
func Int64(i int64) *int64 {
	return &i
}

// Returns a pointer to a uint
func Uint(i uint) *uint {
	return &i
}

// Returns a pointer to a uint8
func Uint8(i uint8) *uint8 {
	return &i
}

// Returns a pointer to a uint16
func Uint16(i uint16) *uint16 {
	return &i
}

// Returns a pointer to a uint32
func Uint32(i uint32) *uint32 {
	return &i
}

// Returns a pointer to a uint64
func Uint64(i uint64) *uint64 {
	return &i
}

// Returns a pointer to a float32
func Float32(f float32) *float32 {
	return &f
}

// Returns a pointer to a float64
func Float64(f float64) *float64 {
	return &f
}

// Returns a pointer to a bool
func Bool(b bool) *bool {
	return &b
}

// Returns a pointer to a byte
func Byte(b byte) *byte {
	return &b
}

// Returns a pointer to a rune
func Rune(r rune) *rune {
	return &r
}

// Returns a pointer to a complex64
func Complex64(c complex64) *complex64 {
	return &c
}

// Returns a pointer to a complex128
func Complex128(c complex128) *complex128 {
	return &c
}

// Returns a pointer to an interface{}
func Interface(i interface{}) *interface{} {
	return &i
}

// Returns a pointer to a slice
func Slice(s interface{}) *interface{} {
	return &s
}

// Returns a pointer to any
func Any[T any](t T) *T {
	return &t
}
