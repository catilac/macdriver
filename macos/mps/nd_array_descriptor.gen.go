// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/kernel"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayDescriptor] class.
var NDArrayDescriptorClass = _NDArrayDescriptorClass{objc.GetClass("MPSNDArrayDescriptor")}

type _NDArrayDescriptorClass struct {
	objc.Class
}

// An interface definition for the [NDArrayDescriptor] class.
type INDArrayDescriptor interface {
	objc.IObject
	SliceDimensionWithSubrange(dimensionIndex uint, subRange DimensionSlice)
	TransposeDimensionWithDimension(dimensionIndex uint, dimensionIndex2 uint)
	SliceRangeForDimension(dimensionIndex uint) DimensionSlice
	DimensionOrder() kernel.Vector_uchar16
	LengthOfDimension(dimensionIndex uint) uint
	ReshapeWithDimensionCountDimensionSizes(numberOfDimensions uint, dimensionSizes *uint)
	ReshapeWithShape(shape []foundation.INumber)
	DataType() DataType
	SetDataType(value DataType)
	NumberOfDimensions() uint
	SetNumberOfDimensions(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor?language=objc
type NDArrayDescriptor struct {
	objc.Object
}

func NDArrayDescriptorFrom(ptr unsafe.Pointer) NDArrayDescriptor {
	return NDArrayDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NDArrayDescriptorClass) DescriptorWithDataTypeDimensionSizes(dataType DataType, dimension0 uint, args ...any) NDArrayDescriptor {
	rv := objc.Call[NDArrayDescriptor](nc, objc.Sel("descriptorWithDataType:dimensionSizes:"), append([]any{dataType, dimension0}, args...)...)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114064-descriptorwithdatatype?language=objc
func NDArrayDescriptor_DescriptorWithDataTypeDimensionSizes(dataType DataType, dimension0 uint, args ...any) NDArrayDescriptor {
	return NDArrayDescriptorClass.DescriptorWithDataTypeDimensionSizes(dataType, dimension0, args...)
}

func (nc _NDArrayDescriptorClass) Alloc() NDArrayDescriptor {
	rv := objc.Call[NDArrayDescriptor](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayDescriptor_Alloc() NDArrayDescriptor {
	return NDArrayDescriptorClass.Alloc()
}

func (nc _NDArrayDescriptorClass) New() NDArrayDescriptor {
	rv := objc.Call[NDArrayDescriptor](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayDescriptor() NDArrayDescriptor {
	return NDArrayDescriptorClass.New()
}

func (n_ NDArrayDescriptor) Init() NDArrayDescriptor {
	rv := objc.Call[NDArrayDescriptor](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114069-slicedimension?language=objc
func (n_ NDArrayDescriptor) SliceDimensionWithSubrange(dimensionIndex uint, subRange DimensionSlice) {
	objc.Call[objc.Void](n_, objc.Sel("sliceDimension:withSubrange:"), dimensionIndex, subRange)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114071-transposedimension?language=objc
func (n_ NDArrayDescriptor) TransposeDimensionWithDimension(dimensionIndex uint, dimensionIndex2 uint) {
	objc.Call[objc.Void](n_, objc.Sel("transposeDimension:withDimension:"), dimensionIndex, dimensionIndex2)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114070-slicerangefordimension?language=objc
func (n_ NDArrayDescriptor) SliceRangeForDimension(dimensionIndex uint) DimensionSlice {
	rv := objc.Call[DimensionSlice](n_, objc.Sel("sliceRangeForDimension:"), dimensionIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114065-dimensionorder?language=objc
func (n_ NDArrayDescriptor) DimensionOrder() kernel.Vector_uchar16 {
	rv := objc.Call[kernel.Vector_uchar16](n_, objc.Sel("dimensionOrder"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114066-lengthofdimension?language=objc
func (n_ NDArrayDescriptor) LengthOfDimension(dimensionIndex uint) uint {
	rv := objc.Call[uint](n_, objc.Sel("lengthOfDimension:"), dimensionIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143492-reshapewithdimensioncount?language=objc
func (n_ NDArrayDescriptor) ReshapeWithDimensionCountDimensionSizes(numberOfDimensions uint, dimensionSizes *uint) {
	objc.Call[objc.Void](n_, objc.Sel("reshapeWithDimensionCount:dimensionSizes:"), numberOfDimensions, dimensionSizes)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3143493-reshapewithshape?language=objc
func (n_ NDArrayDescriptor) ReshapeWithShape(shape []foundation.INumber) {
	objc.Call[objc.Void](n_, objc.Sel("reshapeWithShape:"), shape)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114062-datatype?language=objc
func (n_ NDArrayDescriptor) DataType() DataType {
	rv := objc.Call[DataType](n_, objc.Sel("dataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114062-datatype?language=objc
func (n_ NDArrayDescriptor) SetDataType(value DataType) {
	objc.Call[objc.Void](n_, objc.Sel("setDataType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114067-numberofdimensions?language=objc
func (n_ NDArrayDescriptor) NumberOfDimensions() uint {
	rv := objc.Call[uint](n_, objc.Sel("numberOfDimensions"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraydescriptor/3114067-numberofdimensions?language=objc
func (n_ NDArrayDescriptor) SetNumberOfDimensions(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setNumberOfDimensions:"), value)
}
