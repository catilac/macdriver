// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingBilinear] class.
var CNNUpsamplingBilinearClass = _CNNUpsamplingBilinearClass{objc.GetClass("MPSCNNUpsamplingBilinear")}

type _CNNUpsamplingBilinearClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingBilinear] class.
type ICNNUpsamplingBilinear interface {
	ICNNUpsampling
}

// A bilinear spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinear?language=objc
type CNNUpsamplingBilinear struct {
	CNNUpsampling
}

func CNNUpsamplingBilinearFrom(ptr unsafe.Pointer) CNNUpsamplingBilinear {
	return CNNUpsamplingBilinear{
		CNNUpsampling: CNNUpsamplingFrom(ptr),
	}
}

func (c_ CNNUpsamplingBilinear) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.PDevice, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinear {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingBilinear](c_, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), po0, integerScaleFactorX, integerScaleFactorY)
	return rv
}

// Initializes a bilinear spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinear/2875160-initwithdevice?language=objc
func NewCNNUpsamplingBilinearWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.PDevice, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinear {
	instance := CNNUpsamplingBilinearClass.Alloc().InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device, integerScaleFactorX, integerScaleFactorY)
	instance.Autorelease()
	return instance
}

func (cc _CNNUpsamplingBilinearClass) Alloc() CNNUpsamplingBilinear {
	rv := objc.Call[CNNUpsamplingBilinear](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsamplingBilinear_Alloc() CNNUpsamplingBilinear {
	return CNNUpsamplingBilinearClass.Alloc()
}

func (cc _CNNUpsamplingBilinearClass) New() CNNUpsamplingBilinear {
	rv := objc.Call[CNNUpsamplingBilinear](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingBilinear() CNNUpsamplingBilinear {
	return CNNUpsamplingBilinearClass.New()
}

func (c_ CNNUpsamplingBilinear) Init() CNNUpsamplingBilinear {
	rv := objc.Call[CNNUpsamplingBilinear](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNUpsamplingBilinear) InitWithDevice(device metal.PDevice) CNNUpsamplingBilinear {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingBilinear](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNUpsamplingBilinearWithDevice(device metal.PDevice) CNNUpsamplingBilinear {
	instance := CNNUpsamplingBilinearClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNUpsamplingBilinear) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingBilinear {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingBilinear](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNUpsamplingBilinear_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingBilinear {
	instance := CNNUpsamplingBilinearClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
