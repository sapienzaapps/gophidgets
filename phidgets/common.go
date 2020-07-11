package phidgets

/*
#cgo CFLAGS: -I . -g -Wall
#cgo LDFLAGS: -L . -lphidget22
#include <stdlib.h>
#include <phidget22.h>
*/
import "C"
import (
	"fmt"
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

//export callback
func callback(handle unsafe.Pointer, ctx unsafe.Pointer, value C.double) {
	p2 := gopointer.Restore(ctx).(func())
	p2()
	fmt.Println("Callback is getting the value", cDoubleTofloat32(value))
}

//Common functions that convert different types for this package

func float32ToCdouble(f float32) C.double {
	return C.double(f)
}

func boolToCInt(b bool) C.int {
	var r C.int
	if b {
		r = 1
	} else {
		r = 0
	}
	return r
}

func intToBool(i int) bool {
	var b bool
	if i > 0 {
		b = true
	} else {
		b = false
	}
	return b
}

func intToCInt(i int) C.int {
	var c C.int
	c = (C.int)(i)
	return c
}

func cIntToint(c C.int) int {
	var i int
	i = (int)(c)
	return i
}

func cIntTobool(c C.int) bool {
	i := cIntToint(c)
	return intToBool(i)
}

func uintToCUInt(i uint) C.uint {
	var c C.uint
	c = (C.uint)(i)
	return c
}

func cDoubleTofloat32(d C.double) float32 {
	var f float32
	f = (float32)(d)
	return f
}

func stringToCCharArray(s string) *C.char {
	return C.CString(s)
}
