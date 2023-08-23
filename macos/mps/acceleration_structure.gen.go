// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccelerationStructure] class.
var AccelerationStructureClass = _AccelerationStructureClass{objc.GetClass("MPSAccelerationStructure")}

type _AccelerationStructureClass struct {
	objc.Class
}

// An interface definition for the [AccelerationStructure] class.
type IAccelerationStructure interface {
	IKernel
}

// The base class for data structures that are built over geometry and used to accelerate ray tracing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsaccelerationstructure?language=objc
type AccelerationStructure struct {
	Kernel
}

func AccelerationStructureFrom(ptr unsafe.Pointer) AccelerationStructure {
	return AccelerationStructure{
		Kernel: KernelFrom(ptr),
	}
}

func (ac _AccelerationStructureClass) Alloc() AccelerationStructure {
	rv := objc.Call[AccelerationStructure](ac, objc.Sel("alloc"))
	return rv
}

func AccelerationStructure_Alloc() AccelerationStructure {
	return AccelerationStructureClass.Alloc()
}

func (ac _AccelerationStructureClass) New() AccelerationStructure {
	rv := objc.Call[AccelerationStructure](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccelerationStructure() AccelerationStructure {
	return AccelerationStructureClass.New()
}

func (a_ AccelerationStructure) Init() AccelerationStructure {
	rv := objc.Call[AccelerationStructure](a_, objc.Sel("init"))
	return rv
}

func (a_ AccelerationStructure) InitWithDevice(device metal.PDevice) AccelerationStructure {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[AccelerationStructure](a_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewAccelerationStructureWithDevice(device metal.PDevice) AccelerationStructure {
	instance := AccelerationStructureClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (a_ AccelerationStructure) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) AccelerationStructure {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[AccelerationStructure](a_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func AccelerationStructure_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) AccelerationStructure {
	instance := AccelerationStructureClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
