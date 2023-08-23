// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageLaplacianPyramidSubtract] class.
var ImageLaplacianPyramidSubtractClass = _ImageLaplacianPyramidSubtractClass{objc.GetClass("MPSImageLaplacianPyramidSubtract")}

type _ImageLaplacianPyramidSubtractClass struct {
	objc.Class
}

// An interface definition for the [ImageLaplacianPyramidSubtract] class.
type IImageLaplacianPyramidSubtract interface {
	IImageLaplacianPyramid
}

// A filter that convolves an image with a subtractive Laplacian pyramid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacianpyramidsubtract?language=objc
type ImageLaplacianPyramidSubtract struct {
	ImageLaplacianPyramid
}

func ImageLaplacianPyramidSubtractFrom(ptr unsafe.Pointer) ImageLaplacianPyramidSubtract {
	return ImageLaplacianPyramidSubtract{
		ImageLaplacianPyramid: ImageLaplacianPyramidFrom(ptr),
	}
}

func (ic _ImageLaplacianPyramidSubtractClass) Alloc() ImageLaplacianPyramidSubtract {
	rv := objc.Call[ImageLaplacianPyramidSubtract](ic, objc.Sel("alloc"))
	return rv
}

func ImageLaplacianPyramidSubtract_Alloc() ImageLaplacianPyramidSubtract {
	return ImageLaplacianPyramidSubtractClass.Alloc()
}

func (ic _ImageLaplacianPyramidSubtractClass) New() ImageLaplacianPyramidSubtract {
	rv := objc.Call[ImageLaplacianPyramidSubtract](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageLaplacianPyramidSubtract() ImageLaplacianPyramidSubtract {
	return ImageLaplacianPyramidSubtractClass.New()
}

func (i_ ImageLaplacianPyramidSubtract) Init() ImageLaplacianPyramidSubtract {
	rv := objc.Call[ImageLaplacianPyramidSubtract](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageLaplacianPyramidSubtract) InitWithDevice(device metal.PDevice) ImageLaplacianPyramidSubtract {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramidSubtract](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a downwards 5-tap image pyramid with the default filter kernel and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648935-initwithdevice?language=objc
func NewImageLaplacianPyramidSubtractWithDevice(device metal.PDevice) ImageLaplacianPyramidSubtract {
	instance := ImageLaplacianPyramidSubtractClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageLaplacianPyramidSubtract) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLaplacianPyramidSubtract {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacianPyramidSubtract](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageLaplacianPyramidSubtract_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLaplacianPyramidSubtract {
	instance := ImageLaplacianPyramidSubtractClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
