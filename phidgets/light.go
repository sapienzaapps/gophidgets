package phidgets

/*
#cgo CFLAGS: -g -Wall
#cgo LDFLAGS: -lphidget22
#include <stdlib.h>
#include <phidget22.h>
typedef void (*callback_fcn)(void* handle, void* ctx, double b);
void ccallback(void* handle, void* ctx, double b);  // Forward declaration.
*/
import "C"
import (
	"errors"
	"reflect"
	"unsafe"

	gopointer "github.com/mattn/go-pointer"
)

//PhidgetLightSensor is the struct that is a phidget lumenance sensor
type PhidgetLightSensor struct {
	handle C.PhidgetLightSensorHandle
}

//Create creates a phidget lumenance sensor
func (p *PhidgetLightSensor) Create() {
	C.PhidgetLightSensor_create(&p.handle)
}

//GetValue gets the lumenance from a phidget lumenance sensor
func (p *PhidgetLightSensor) GetValue() (float32, error) {
	var r C.double
	cerr := C.PhidgetLightSensor_getIlluminance(p.handle, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cDoubleTofloat32(r), nil
}

//SetOnIlluminanceChangeHandler - interrupt for illumiance changes calls a function
func (p *PhidgetLightSensor) SetOnIlluminanceChangeHandler(f func(Phidget, interface{}, float32), ctx interface{}) error {
	//make a c function pointer to a go function pointer and pass it through the phidget context
	var passthrough Passthrough
	passthrough.f = f
	passthrough.ctx = ctx
	passthrough.handle = p
	pt := gopointer.Save(passthrough)
	cerr := C.PhidgetLightSensor_setOnIlluminanceChangeHandler(p.handle, (C.callback_fcn)(unsafe.Pointer(C.ccallback)), pt)
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//Common to all derived phidgets

func (p *PhidgetLightSensor) getErrorDescription(cerr C.PhidgetReturnCode) string {
	var errorString *C.char
	C.Phidget_getErrorDescription(cerr, &errorString)
	//Get the name of our class
	t := reflect.TypeOf(p)
	return t.Elem().Name() + ": " + C.GoString(errorString)
}

//SetIsRemote sets a phidget sensor as a remote device
func (p *PhidgetLightSensor) SetIsRemote(b bool) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setIsRemote(h, boolToCInt(b))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil

}

//SetDeviceSerialNumber sets a phidget light sensor's serial number
func (p *PhidgetLightSensor) SetDeviceSerialNumber(serial int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setDeviceSerialNumber(h, intToCInt(serial))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//SetHubPort sets a phidget light sensor's hub port
func (p *PhidgetLightSensor) SetHubPort(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setHubPort(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetHubPort gets a phidget light sensor's hub port
func (p *PhidgetLightSensor) GetHubPort() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getHubPort(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//SetChannel sets a phidget light sensor's channel port
func (p *PhidgetLightSensor) SetChannel(port int) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_setChannel(h, intToCInt(port))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//GetChannel gets a phidget light sensor's channel port
func (p *PhidgetLightSensor) GetChannel() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getChannel(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//GetIsRemote gets a phidget light sensor's remote status
func (p *PhidgetLightSensor) GetIsRemote() (bool, error) {
	//Cast TemperatureHandle to PhidgetHandle
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getIsRemote(h, &r)
	if cerr != C.EPHIDGET_OK {
		return false, errors.New(p.getErrorDescription(cerr))
	}
	return cIntTobool(r), nil
}

//GetDeviceSerialNumber gets a phidget light sensor's serial number
func (p *PhidgetLightSensor) GetDeviceSerialNumber() (int, error) {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	var r C.int
	cerr := C.Phidget_getDeviceSerialNumber(h, &r)
	if cerr != C.EPHIDGET_OK {
		return 0, errors.New(p.getErrorDescription(cerr))
	}
	return cIntToint(r), nil
}

//OpenWaitForAttachment opens a phidget light sensor for attachment
func (p *PhidgetLightSensor) OpenWaitForAttachment(timeout uint) error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_openWaitForAttachment(h, uintToCUInt(timeout))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}

//Close - close the handle and delete it
func (p *PhidgetLightSensor) Close() error {
	h := (*C.struct__Phidget)(unsafe.Pointer(p.handle))
	cerr := C.Phidget_close(h)
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	cerr = C.PhidgetLightSensor_delete((*C.PhidgetLightSensorHandle)(&p.handle))
	if cerr != C.EPHIDGET_OK {
		return errors.New(p.getErrorDescription(cerr))
	}
	return nil
}
