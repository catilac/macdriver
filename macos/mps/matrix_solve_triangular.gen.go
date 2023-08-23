// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixSolveTriangular] class.
var MatrixSolveTriangularClass = _MatrixSolveTriangularClass{objc.GetClass("MPSMatrixSolveTriangular")}

type _MatrixSolveTriangularClass struct {
	objc.Class
}

// An interface definition for the [MatrixSolveTriangular] class.
type IMatrixSolveTriangular interface {
	IMatrixBinaryKernel
	EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix)
	EncodeToCommandBufferObjectSourceMatrixRightHandSideMatrixSolutionMatrix(commandBufferObject objc.IObject, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix)
}

// A kernel for computing the solution of a linear system of equations using a triangular coefficient matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvetriangular?language=objc
type MatrixSolveTriangular struct {
	MatrixBinaryKernel
}

func MatrixSolveTriangularFrom(ptr unsafe.Pointer) MatrixSolveTriangular {
	return MatrixSolveTriangular{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixSolveTriangular) InitWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha(device metal.PDevice, right bool, upper bool, transpose bool, unit bool, order uint, numberOfRightHandSides uint, alpha float64) MatrixSolveTriangular {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveTriangular](m_, objc.Sel("initWithDevice:right:upper:transpose:unit:order:numberOfRightHandSides:alpha:"), po0, right, upper, transpose, unit, order, numberOfRightHandSides, alpha)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvetriangular/2873007-initwithdevice?language=objc
func NewMatrixSolveTriangularWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha(device metal.PDevice, right bool, upper bool, transpose bool, unit bool, order uint, numberOfRightHandSides uint, alpha float64) MatrixSolveTriangular {
	instance := MatrixSolveTriangularClass.Alloc().InitWithDeviceRightUpperTransposeUnitOrderNumberOfRightHandSidesAlpha(device, right, upper, transpose, unit, order, numberOfRightHandSides, alpha)
	instance.Autorelease()
	return instance
}

func (mc _MatrixSolveTriangularClass) Alloc() MatrixSolveTriangular {
	rv := objc.Call[MatrixSolveTriangular](mc, objc.Sel("alloc"))
	return rv
}

func MatrixSolveTriangular_Alloc() MatrixSolveTriangular {
	return MatrixSolveTriangularClass.Alloc()
}

func (mc _MatrixSolveTriangularClass) New() MatrixSolveTriangular {
	rv := objc.Call[MatrixSolveTriangular](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixSolveTriangular() MatrixSolveTriangular {
	return MatrixSolveTriangularClass.New()
}

func (m_ MatrixSolveTriangular) Init() MatrixSolveTriangular {
	rv := objc.Call[MatrixSolveTriangular](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixSolveTriangular) InitWithDevice(device metal.PDevice) MatrixSolveTriangular {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveTriangular](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixSolveTriangularWithDevice(device metal.PDevice) MatrixSolveTriangular {
	instance := MatrixSolveTriangularClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixSolveTriangular) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSolveTriangular {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveTriangular](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixSolveTriangular_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSolveTriangular {
	instance := MatrixSolveTriangularClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvetriangular/2867027-encodetocommandbuffer?language=objc
func (m_ MatrixSolveTriangular) EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:"), po0, objc.Ptr(sourceMatrix), objc.Ptr(rightHandSideMatrix), objc.Ptr(solutionMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvetriangular/2867027-encodetocommandbuffer?language=objc
func (m_ MatrixSolveTriangular) EncodeToCommandBufferObjectSourceMatrixRightHandSideMatrixSolutionMatrix(commandBufferObject objc.IObject, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceMatrix), objc.Ptr(rightHandSideMatrix), objc.Ptr(solutionMatrix))
}
