// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingGradient] class.
var CNNPoolingGradientClass = _CNNPoolingGradientClass{objc.GetClass("MPSCNNPoolingGradient")}

type _CNNPoolingGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingGradient] class.
type ICNNPoolingGradient interface {
	ICNNGradientKernel
	SourceSize() metal.Size
	SetSourceSize(value metal.Size)
}

// A gradient pooling kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient?language=objc
type CNNPoolingGradient struct {
	CNNGradientKernel
}

func CNNPoolingGradientFrom(ptr unsafe.Pointer) CNNPoolingGradient {
	return CNNPoolingGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNPoolingGradient) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNPoolingGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942337-initwithdevice?language=objc
func NewCNNPoolingGradientWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNPoolingGradient {
	instance := CNNPoolingGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingGradientClass) Alloc() CNNPoolingGradient {
	rv := objc.Call[CNNPoolingGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingGradient_Alloc() CNNPoolingGradient {
	return CNNPoolingGradientClass.Alloc()
}

func (cc _CNNPoolingGradientClass) New() CNNPoolingGradient {
	rv := objc.Call[CNNPoolingGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingGradient() CNNPoolingGradient {
	return CNNPoolingGradientClass.New()
}

func (c_ CNNPoolingGradient) Init() CNNPoolingGradient {
	rv := objc.Call[CNNPoolingGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingGradient) InitWithDevice(device metal.PDevice) CNNPoolingGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNPoolingGradientWithDevice(device metal.PDevice) CNNPoolingGradient {
	instance := CNNPoolingGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNPoolingGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingGradient {
	instance := CNNPoolingGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942343-sourcesize?language=objc
func (c_ CNNPoolingGradient) SourceSize() metal.Size {
	rv := objc.Call[metal.Size](c_, objc.Sel("sourceSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942343-sourcesize?language=objc
func (c_ CNNPoolingGradient) SetSourceSize(value metal.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setSourceSize:"), value)
}
