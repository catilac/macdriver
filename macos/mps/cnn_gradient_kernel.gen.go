// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNGradientKernel] class.
var CNNGradientKernelClass = _CNNGradientKernelClass{objc.GetClass("MPSCNNGradientKernel")}

type _CNNGradientKernelClass struct {
	objc.Class
}

// An interface definition for the [CNNGradientKernel] class.
type ICNNGradientKernel interface {
	ICNNBinaryKernel
	EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates(commandBuffer metal.PCommandBuffer, sourceGradients *foundation.Array, sourceImages *foundation.Array, gradientStates *foundation.Array) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceGradientsSourceImagesGradientStates(commandBufferObject objc.IObject, sourceGradients *foundation.Array, sourceImages *foundation.Array, gradientStates *foundation.Array) *foundation.Array
	EncodeToCommandBufferSourceGradientSourceImageGradientState(commandBuffer metal.PCommandBuffer, sourceGradient IImage, sourceImage IImage, gradientState IState) Image
	EncodeToCommandBufferObjectSourceGradientSourceImageGradientState(commandBufferObject objc.IObject, sourceGradient IImage, sourceImage IImage, gradientState IState) Image
	KernelOffsetX() int
	SetKernelOffsetX(value int)
	KernelOffsetY() int
	SetKernelOffsetY(value int)
}

// The base class for gradient layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel?language=objc
type CNNGradientKernel struct {
	CNNBinaryKernel
}

func CNNGradientKernelFrom(ptr unsafe.Pointer) CNNGradientKernel {
	return CNNGradientKernel{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}

func (c_ CNNGradientKernel) InitWithDevice(device metal.PDevice) CNNGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNGradientKernel](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNGradientKernelWithDevice(device metal.PDevice) CNNGradientKernel {
	instance := CNNGradientKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNGradientKernelClass) Alloc() CNNGradientKernel {
	rv := objc.Call[CNNGradientKernel](cc, objc.Sel("alloc"))
	return rv
}

func CNNGradientKernel_Alloc() CNNGradientKernel {
	return CNNGradientKernelClass.Alloc()
}

func (cc _CNNGradientKernelClass) New() CNNGradientKernel {
	rv := objc.Call[CNNGradientKernel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNGradientKernel() CNNGradientKernel {
	return CNNGradientKernelClass.New()
}

func (c_ CNNGradientKernel) Init() CNNGradientKernel {
	rv := objc.Call[CNNGradientKernel](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNGradientKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNGradientKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNGradientKernel](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNGradientKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNGradientKernel {
	instance := CNNGradientKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942668-encodebatchtocommandbuffer?language=objc
func (c_ CNNGradientKernel) EncodeBatchToCommandBufferSourceGradientsSourceImagesGradientStates(commandBuffer metal.PCommandBuffer, sourceGradients *foundation.Array, sourceImages *foundation.Array, gradientStates *foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:gradientStates:"), po0, sourceGradients, sourceImages, gradientStates)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942668-encodebatchtocommandbuffer?language=objc
func (c_ CNNGradientKernel) EncodeBatchToCommandBufferObjectSourceGradientsSourceImagesGradientStates(commandBufferObject objc.IObject, sourceGradients *foundation.Array, sourceImages *foundation.Array, gradientStates *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:gradientStates:"), objc.Ptr(commandBufferObject), sourceGradients, sourceImages, gradientStates)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942663-encodetocommandbuffer?language=objc
func (c_ CNNGradientKernel) EncodeToCommandBufferSourceGradientSourceImageGradientState(commandBuffer metal.PCommandBuffer, sourceGradient IImage, sourceImage IImage, gradientState IState) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:gradientState:"), po0, objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942663-encodetocommandbuffer?language=objc
func (c_ CNNGradientKernel) EncodeToCommandBufferObjectSourceGradientSourceImageGradientState(commandBufferObject objc.IObject, sourceGradient IImage, sourceImage IImage, gradientState IState) Image {
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:gradientState:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942676-kerneloffsetx?language=objc
func (c_ CNNGradientKernel) KernelOffsetX() int {
	rv := objc.Call[int](c_, objc.Sel("kernelOffsetX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942676-kerneloffsetx?language=objc
func (c_ CNNGradientKernel) SetKernelOffsetX(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelOffsetX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942644-kerneloffsety?language=objc
func (c_ CNNGradientKernel) KernelOffsetY() int {
	rv := objc.Call[int](c_, objc.Sel("kernelOffsetY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942644-kerneloffsety?language=objc
func (c_ CNNGradientKernel) SetKernelOffsetY(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelOffsetY:"), value)
}
