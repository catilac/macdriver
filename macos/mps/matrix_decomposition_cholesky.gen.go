// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixDecompositionCholesky] class.
var MatrixDecompositionCholeskyClass = _MatrixDecompositionCholeskyClass{objc.GetClass("MPSMatrixDecompositionCholesky")}

type _MatrixDecompositionCholeskyClass struct {
	objc.Class
}

// An interface definition for the [MatrixDecompositionCholesky] class.
type IMatrixDecompositionCholesky interface {
	IMatrixUnaryKernel
	EncodeToCommandBufferSourceMatrixResultMatrixStatus(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, resultMatrix IMatrix, status metal.PBuffer)
	EncodeToCommandBufferObjectSourceMatrixResultMatrixStatusObject(commandBufferObject objc.IObject, sourceMatrix IMatrix, resultMatrix IMatrix, statusObject objc.IObject)
}

// A kernel for computing the Cholesky factorization of a matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky?language=objc
type MatrixDecompositionCholesky struct {
	MatrixUnaryKernel
}

func MatrixDecompositionCholeskyFrom(ptr unsafe.Pointer) MatrixDecompositionCholesky {
	return MatrixDecompositionCholesky{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}

func (m_ MatrixDecompositionCholesky) InitWithDeviceLowerOrder(device metal.PDevice, lower bool, order uint) MatrixDecompositionCholesky {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixDecompositionCholesky](m_, objc.Sel("initWithDevice:lower:order:"), po0, lower, order)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867119-initwithdevice?language=objc
func NewMatrixDecompositionCholeskyWithDeviceLowerOrder(device metal.PDevice, lower bool, order uint) MatrixDecompositionCholesky {
	instance := MatrixDecompositionCholeskyClass.Alloc().InitWithDeviceLowerOrder(device, lower, order)
	instance.Autorelease()
	return instance
}

func (mc _MatrixDecompositionCholeskyClass) Alloc() MatrixDecompositionCholesky {
	rv := objc.Call[MatrixDecompositionCholesky](mc, objc.Sel("alloc"))
	return rv
}

func MatrixDecompositionCholesky_Alloc() MatrixDecompositionCholesky {
	return MatrixDecompositionCholeskyClass.Alloc()
}

func (mc _MatrixDecompositionCholeskyClass) New() MatrixDecompositionCholesky {
	rv := objc.Call[MatrixDecompositionCholesky](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixDecompositionCholesky() MatrixDecompositionCholesky {
	return MatrixDecompositionCholeskyClass.New()
}

func (m_ MatrixDecompositionCholesky) Init() MatrixDecompositionCholesky {
	rv := objc.Call[MatrixDecompositionCholesky](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixDecompositionCholesky) InitWithDevice(device metal.PDevice) MatrixDecompositionCholesky {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixDecompositionCholesky](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixDecompositionCholeskyWithDevice(device metal.PDevice) MatrixDecompositionCholesky {
	instance := MatrixDecompositionCholeskyClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixDecompositionCholesky) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixDecompositionCholesky {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixDecompositionCholesky](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixDecompositionCholesky_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixDecompositionCholesky {
	instance := MatrixDecompositionCholeskyClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867004-encodetocommandbuffer?language=objc
func (m_ MatrixDecompositionCholesky) EncodeToCommandBufferSourceMatrixResultMatrixStatus(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, resultMatrix IMatrix, status metal.PBuffer) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po3 := objc.WrapAsProtocol("MTLBuffer", status)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:status:"), po0, objc.Ptr(sourceMatrix), objc.Ptr(resultMatrix), po3)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositioncholesky/2867004-encodetocommandbuffer?language=objc
func (m_ MatrixDecompositionCholesky) EncodeToCommandBufferObjectSourceMatrixResultMatrixStatusObject(commandBufferObject objc.IObject, sourceMatrix IMatrix, resultMatrix IMatrix, statusObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:status:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceMatrix), objc.Ptr(resultMatrix), objc.Ptr(statusObject))
}
