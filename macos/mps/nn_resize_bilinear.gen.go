// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNResizeBilinear] class.
var NNResizeBilinearClass = _NNResizeBilinearClass{objc.GetClass("MPSNNResizeBilinear")}

type _NNResizeBilinearClass struct {
	objc.Class
}

// An interface definition for the [NNResizeBilinear] class.
type INNResizeBilinear interface {
	ICNNKernel
	ResizeHeight() uint
	AlignCorners() bool
	ResizeWidth() uint
}

// A bilinear resizing filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear?language=objc
type NNResizeBilinear struct {
	CNNKernel
}

func NNResizeBilinearFrom(ptr unsafe.Pointer) NNResizeBilinear {
	return NNResizeBilinear{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (n_ NNResizeBilinear) InitWithDeviceResizeWidthResizeHeightAlignCorners(device metal.PDevice, resizeWidth uint, resizeHeight uint, alignCorners bool) NNResizeBilinear {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNResizeBilinear](n_, objc.Sel("initWithDevice:resizeWidth:resizeHeight:alignCorners:"), po0, resizeWidth, resizeHeight, alignCorners)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012966-initwithdevice?language=objc
func NewNNResizeBilinearWithDeviceResizeWidthResizeHeightAlignCorners(device metal.PDevice, resizeWidth uint, resizeHeight uint, alignCorners bool) NNResizeBilinear {
	instance := NNResizeBilinearClass.Alloc().InitWithDeviceResizeWidthResizeHeightAlignCorners(device, resizeWidth, resizeHeight, alignCorners)
	instance.Autorelease()
	return instance
}

func (nc _NNResizeBilinearClass) Alloc() NNResizeBilinear {
	rv := objc.Call[NNResizeBilinear](nc, objc.Sel("alloc"))
	return rv
}

func NNResizeBilinear_Alloc() NNResizeBilinear {
	return NNResizeBilinearClass.Alloc()
}

func (nc _NNResizeBilinearClass) New() NNResizeBilinear {
	rv := objc.Call[NNResizeBilinear](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNResizeBilinear() NNResizeBilinear {
	return NNResizeBilinearClass.New()
}

func (n_ NNResizeBilinear) Init() NNResizeBilinear {
	rv := objc.Call[NNResizeBilinear](n_, objc.Sel("init"))
	return rv
}

func (n_ NNResizeBilinear) InitWithDevice(device metal.PDevice) NNResizeBilinear {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNResizeBilinear](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewNNResizeBilinearWithDevice(device metal.PDevice) NNResizeBilinear {
	instance := NNResizeBilinearClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNResizeBilinear) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNResizeBilinear {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNResizeBilinear](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNResizeBilinear_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNResizeBilinear {
	instance := NNResizeBilinearClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012968-resizeheight?language=objc
func (n_ NNResizeBilinear) ResizeHeight() uint {
	rv := objc.Call[uint](n_, objc.Sel("resizeHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012965-aligncorners?language=objc
func (n_ NNResizeBilinear) AlignCorners() bool {
	rv := objc.Call[bool](n_, objc.Sel("alignCorners"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnresizebilinear/3012969-resizewidth?language=objc
func (n_ NNResizeBilinear) ResizeWidth() uint {
	rv := objc.Call[uint](n_, objc.Sel("resizeWidth"))
	return rv
}
