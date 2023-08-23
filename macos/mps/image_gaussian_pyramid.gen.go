// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageGaussianPyramid] class.
var ImageGaussianPyramidClass = _ImageGaussianPyramidClass{objc.GetClass("MPSImageGaussianPyramid")}

type _ImageGaussianPyramidClass struct {
	objc.Class
}

// An interface definition for the [ImageGaussianPyramid] class.
type IImageGaussianPyramid interface {
	IImagePyramid
}

// A filter that convolves an image with a Gaussian pyramid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianpyramid?language=objc
type ImageGaussianPyramid struct {
	ImagePyramid
}

func ImageGaussianPyramidFrom(ptr unsafe.Pointer) ImageGaussianPyramid {
	return ImageGaussianPyramid{
		ImagePyramid: ImagePyramidFrom(ptr),
	}
}

func (ic _ImageGaussianPyramidClass) Alloc() ImageGaussianPyramid {
	rv := objc.Call[ImageGaussianPyramid](ic, objc.Sel("alloc"))
	return rv
}

func ImageGaussianPyramid_Alloc() ImageGaussianPyramid {
	return ImageGaussianPyramidClass.Alloc()
}

func (ic _ImageGaussianPyramidClass) New() ImageGaussianPyramid {
	rv := objc.Call[ImageGaussianPyramid](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageGaussianPyramid() ImageGaussianPyramid {
	return ImageGaussianPyramidClass.New()
}

func (i_ ImageGaussianPyramid) Init() ImageGaussianPyramid {
	rv := objc.Call[ImageGaussianPyramid](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageGaussianPyramid) InitWithDevice(device metal.PDevice) ImageGaussianPyramid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageGaussianPyramid](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a downwards 5-tap image pyramid with the default filter kernel and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648935-initwithdevice?language=objc
func NewImageGaussianPyramidWithDevice(device metal.PDevice) ImageGaussianPyramid {
	instance := ImageGaussianPyramidClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageGaussianPyramid) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageGaussianPyramid {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageGaussianPyramid](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageGaussianPyramid_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageGaussianPyramid {
	instance := ImageGaussianPyramidClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
