// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayUnaryKernel] class.
var NDArrayUnaryKernelClass = _NDArrayUnaryKernelClass{objc.GetClass("MPSNDArrayUnaryKernel")}

type _NDArrayUnaryKernelClass struct {
	objc.Class
}

// An interface definition for the [NDArrayUnaryKernel] class.
type INDArrayUnaryKernel interface {
	INDArrayMultiaryKernel
	EncodeToCommandBufferSourceArrayResultStateDestinationArray(cmdBuf metal.PCommandBuffer, sourceArray INDArray, outGradientState IState, destination INDArray)
	EncodeToCommandBufferObjectSourceArrayResultStateDestinationArray(cmdBufObject objc.IObject, sourceArray INDArray, outGradientState IState, destination INDArray)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel?language=objc
type NDArrayUnaryKernel struct {
	NDArrayMultiaryKernel
}

func NDArrayUnaryKernelFrom(ptr unsafe.Pointer) NDArrayUnaryKernel {
	return NDArrayUnaryKernel{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}

func (n_ NDArrayUnaryKernel) InitWithDevice(device metal.PDevice) NDArrayUnaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayUnaryKernel](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143540-initwithdevice?language=objc
func NewNDArrayUnaryKernelWithDevice(device metal.PDevice) NDArrayUnaryKernel {
	instance := NDArrayUnaryKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayUnaryKernelClass) Alloc() NDArrayUnaryKernel {
	rv := objc.Call[NDArrayUnaryKernel](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayUnaryKernel_Alloc() NDArrayUnaryKernel {
	return NDArrayUnaryKernelClass.Alloc()
}

func (nc _NDArrayUnaryKernelClass) New() NDArrayUnaryKernel {
	rv := objc.Call[NDArrayUnaryKernel](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayUnaryKernel() NDArrayUnaryKernel {
	return NDArrayUnaryKernelClass.New()
}

func (n_ NDArrayUnaryKernel) Init() NDArrayUnaryKernel {
	rv := objc.Call[NDArrayUnaryKernel](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayUnaryKernel) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayUnaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayUnaryKernel](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175010-initwithdevice?language=objc
func NewNDArrayUnaryKernelWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayUnaryKernel {
	instance := NDArrayUnaryKernelClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayUnaryKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayUnaryKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayUnaryKernel](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayUnaryKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayUnaryKernel {
	instance := NDArrayUnaryKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143538-encodetocommandbuffer?language=objc
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferSourceArrayResultStateDestinationArray(cmdBuf metal.PCommandBuffer, sourceArray INDArray, outGradientState IState, destination INDArray) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:sourceArray:resultState:destinationArray:"), po0, objc.Ptr(sourceArray), objc.Ptr(outGradientState), objc.Ptr(destination))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143538-encodetocommandbuffer?language=objc
func (n_ NDArrayUnaryKernel) EncodeToCommandBufferObjectSourceArrayResultStateDestinationArray(cmdBufObject objc.IObject, sourceArray INDArray, outGradientState IState, destination INDArray) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:sourceArray:resultState:destinationArray:"), objc.Ptr(cmdBufObject), objc.Ptr(sourceArray), objc.Ptr(outGradientState), objc.Ptr(destination))
}
