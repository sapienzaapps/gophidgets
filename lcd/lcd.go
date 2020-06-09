package lcd

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -lphidget22
// #include <stdlib.h>
// #include <phidget22.h>
import "C"
import (
	"errors"
	"unsafe"
)

//PhidgetLCD is the struct that is a phidget lcd sensor
type PhidgetLCD struct {
	handle C.PhidgetLCDHandle
}

//Create creates a phidget lcd sensor
func (t *PhidgetLCD) Create() {
	C.PhidgetLCD_create(&t.handle)
}

//SetText sets the lcd text
func (t *PhidgetLCD) SetText(text string) {
	C.PhidgetLCD_writeText(t.handle, C.FONT_6x12, 40, 25, stringToCCharArray(text))
	C.PhidgetLCD_flush(t.handle)
}

func (t *PhidgetLCD) SetBacklight(brightness float32) {
	C.PhidgetLCD_setBacklight(t.handle, float32ToCdouble(brightness))
}

//Common to all derived phidgets

func (p *PhidgetLCD) getErrorDescription(cerr C.PhidgetReturnCode) string {
	var errorString *C.char
	C.Phidget_getErrorDescription(cerr, &errorString)
	return C.GoString(errorString)
}

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetLCD) SetIsRemote(b bool) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setIsRemote(h, boolToCInt(b))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil

}

//SetDeviceSerialNumber sets a phidget lcd sensor's serial number
func (p *PhidgetLCD) SetDeviceSerialNumber(serial int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//SetHubPort sets a phidget lcd sensor's hub port
func (p *PhidgetLCD) SetHubPort(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setHubPort(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetIsRemote gets a phidget lcd sensor's remote status
func (p *PhidgetLCD) GetIsRemote() (bool, error) {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getIsRemote(h, &r)
	if cerr != C.EPHIDGET_OK {
		return false, errors.New(p.getErrorDescription(cerr))
	}
	return cIntTobool(r), nil
}

//GetDeviceSerialNumber gets a phidget lcd sensor's serial number
func (p *PhidgetLCD) GetDeviceSerialNumber() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getDeviceSerialNumber(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//GetHubPort gets a phidget lcd sensor's hub port
func (p *PhidgetLCD) GetHubPort() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getHubPort(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//OpenWaitForAttachment opens a phidget lcd sensor for attachment
func (p *PhidgetLCD) OpenWaitForAttachment(timeout uint) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//Can't put these in a common module because their type is associated with the module

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