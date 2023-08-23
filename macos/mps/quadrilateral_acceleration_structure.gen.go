// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [QuadrilateralAccelerationStructure] class.
var QuadrilateralAccelerationStructureClass = _QuadrilateralAccelerationStructureClass{objc.GetClass("MPSQuadrilateralAccelerationStructure")}

type _QuadrilateralAccelerationStructureClass struct {
	objc.Class
}

// An interface definition for the [QuadrilateralAccelerationStructure] class.
type IQuadrilateralAccelerationStructure interface {
	IPolygonAccelerationStructure
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsquadrilateralaccelerationstructure?language=objc
type QuadrilateralAccelerationStructure struct {
	PolygonAccelerationStructure
}

func QuadrilateralAccelerationStructureFrom(ptr unsafe.Pointer) QuadrilateralAccelerationStructure {
	return QuadrilateralAccelerationStructure{
		PolygonAccelerationStructure: PolygonAccelerationStructureFrom(ptr),
	}
}

func (qc _QuadrilateralAccelerationStructureClass) Alloc() QuadrilateralAccelerationStructure {
	rv := objc.Call[QuadrilateralAccelerationStructure](qc, objc.Sel("alloc"))
	return rv
}

func QuadrilateralAccelerationStructure_Alloc() QuadrilateralAccelerationStructure {
	return QuadrilateralAccelerationStructureClass.Alloc()
}

func (qc _QuadrilateralAccelerationStructureClass) New() QuadrilateralAccelerationStructure {
	rv := objc.Call[QuadrilateralAccelerationStructure](qc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewQuadrilateralAccelerationStructure() QuadrilateralAccelerationStructure {
	return QuadrilateralAccelerationStructureClass.New()
}

func (q_ QuadrilateralAccelerationStructure) Init() QuadrilateralAccelerationStructure {
	rv := objc.Call[QuadrilateralAccelerationStructure](q_, objc.Sel("init"))
	return rv
}

func (q_ QuadrilateralAccelerationStructure) InitWithDevice(device metal.PDevice) QuadrilateralAccelerationStructure {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[QuadrilateralAccelerationStructure](q_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewQuadrilateralAccelerationStructureWithDevice(device metal.PDevice) QuadrilateralAccelerationStructure {
	instance := QuadrilateralAccelerationStructureClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (q_ QuadrilateralAccelerationStructure) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) QuadrilateralAccelerationStructure {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[QuadrilateralAccelerationStructure](q_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func QuadrilateralAccelerationStructure_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) QuadrilateralAccelerationStructure {
	instance := QuadrilateralAccelerationStructureClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
