// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixBinaryKernel] class.
var MatrixBinaryKernelClass = _MatrixBinaryKernelClass{objc.GetClass("MPSMatrixBinaryKernel")}

type _MatrixBinaryKernelClass struct {
	objc.Class
}

// An interface definition for the [MatrixBinaryKernel] class.
type IMatrixBinaryKernel interface {
	IKernel
	ResultMatrixOrigin() metal.Origin
	SetResultMatrixOrigin(value metal.Origin)
	BatchStart() uint
	SetBatchStart(value uint)
	BatchSize() uint
	SetBatchSize(value uint)
	SecondarySourceMatrixOrigin() metal.Origin
	SetSecondarySourceMatrixOrigin(value metal.Origin)
	PrimarySourceMatrixOrigin() metal.Origin
	SetPrimarySourceMatrixOrigin(value metal.Origin)
}

// A kernel that consumes two matrices and produces one matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel?language=objc
type MatrixBinaryKernel struct {
	Kernel
}

func MatrixBinaryKernelFrom(ptr unsafe.Pointer) MatrixBinaryKernel {
	return MatrixBinaryKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (mc _MatrixBinaryKernelClass) Alloc() MatrixBinaryKernel {
	rv := objc.Call[MatrixBinaryKernel](mc, objc.Sel("alloc"))
	return rv
}

func MatrixBinaryKernel_Alloc() MatrixBinaryKernel {
	return MatrixBinaryKernelClass.Alloc()
}

func (mc _MatrixBinaryKernelClass) New() MatrixBinaryKernel {
	rv := objc.Call[MatrixBinaryKernel](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixBinaryKernel() MatrixBinaryKernel {
	return MatrixBinaryKernelClass.New()
}

func (m_ MatrixBinaryKernel) Init() MatrixBinaryKernel {
	rv := objc.Call[MatrixBinaryKernel](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixBinaryKernel) InitWithDevice(device metal.PDevice) MatrixBinaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixBinaryKernel](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixBinaryKernelWithDevice(device metal.PDevice) MatrixBinaryKernel {
	instance := MatrixBinaryKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixBinaryKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixBinaryKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixBinaryKernel](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixBinaryKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixBinaryKernel {
	instance := MatrixBinaryKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867193-resultmatrixorigin?language=objc
func (m_ MatrixBinaryKernel) ResultMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("resultMatrixOrigin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867193-resultmatrixorigin?language=objc
func (m_ MatrixBinaryKernel) SetResultMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setResultMatrixOrigin:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867152-batchstart?language=objc
func (m_ MatrixBinaryKernel) BatchStart() uint {
	rv := objc.Call[uint](m_, objc.Sel("batchStart"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867152-batchstart?language=objc
func (m_ MatrixBinaryKernel) SetBatchStart(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setBatchStart:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867089-batchsize?language=objc
func (m_ MatrixBinaryKernel) BatchSize() uint {
	rv := objc.Call[uint](m_, objc.Sel("batchSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867089-batchsize?language=objc
func (m_ MatrixBinaryKernel) SetBatchSize(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setBatchSize:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867096-secondarysourcematrixorigin?language=objc
func (m_ MatrixBinaryKernel) SecondarySourceMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("secondarySourceMatrixOrigin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867096-secondarysourcematrixorigin?language=objc
func (m_ MatrixBinaryKernel) SetSecondarySourceMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setSecondarySourceMatrixOrigin:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867182-primarysourcematrixorigin?language=objc
func (m_ MatrixBinaryKernel) PrimarySourceMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("primarySourceMatrixOrigin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbinarykernel/2867182-primarysourcematrixorigin?language=objc
func (m_ MatrixBinaryKernel) SetPrimarySourceMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setPrimarySourceMatrixOrigin:"), value)
}
