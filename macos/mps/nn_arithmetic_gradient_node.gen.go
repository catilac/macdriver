// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNArithmeticGradientNode] class.
var NNArithmeticGradientNodeClass = _NNArithmeticGradientNodeClass{objc.GetClass("MPSNNArithmeticGradientNode")}

type _NNArithmeticGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNArithmeticGradientNode] class.
type INNArithmeticGradientNode interface {
	INNGradientFilterNode
	SecondaryScale() float64
	SetSecondaryScale(value float64)
	IsSecondarySourceFilter() bool
	SecondaryStrideInPixelsY() uint
	SetSecondaryStrideInPixelsY(value uint)
	SecondaryStrideInPixelsX() uint
	SetSecondaryStrideInPixelsX(value uint)
	MaximumValue() float64
	SetMaximumValue(value float64)
	PrimaryScale() float64
	SetPrimaryScale(value float64)
	MinimumValue() float64
	SetMinimumValue(value float64)
	Bias() float64
	SetBias(value float64)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
}

// A representation of the base class for gradient arithmetic operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode?language=objc
type NNArithmeticGradientNode struct {
	NNGradientFilterNode
}

func NNArithmeticGradientNodeFrom(ptr unsafe.Pointer) NNArithmeticGradientNode {
	return NNArithmeticGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (n_ NNArithmeticGradientNode) InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](n_, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), gradientImages, objc.Ptr(filter), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952980-initwithgradientimages?language=objc
func NewNNArithmeticGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	instance := NNArithmeticGradientNodeClass.Alloc().InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages, filter, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (n_ NNArithmeticGradientNode) InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956166-initwithsourcegradient?language=objc
func NewNNArithmeticGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	instance := NNArithmeticGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (nc _NNArithmeticGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956167-nodewithsourcegradient?language=objc
func NNArithmeticGradientNode_NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNArithmeticGradientNode {
	return NNArithmeticGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
}

func (nc _NNArithmeticGradientNodeClass) Alloc() NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNArithmeticGradientNode_Alloc() NNArithmeticGradientNode {
	return NNArithmeticGradientNodeClass.Alloc()
}

func (nc _NNArithmeticGradientNodeClass) New() NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNArithmeticGradientNode() NNArithmeticGradientNode {
	return NNArithmeticGradientNodeClass.New()
}

func (n_ NNArithmeticGradientNode) Init() NNArithmeticGradientNode {
	rv := objc.Call[NNArithmeticGradientNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952981-secondaryscale?language=objc
func (n_ NNArithmeticGradientNode) SecondaryScale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("secondaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952981-secondaryscale?language=objc
func (n_ NNArithmeticGradientNode) SetSecondaryScale(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952987-issecondarysourcefilter?language=objc
func (n_ NNArithmeticGradientNode) IsSecondarySourceFilter() bool {
	rv := objc.Call[bool](n_, objc.Sel("isSecondarySourceFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952994-secondarystrideinpixelsy?language=objc
func (n_ NNArithmeticGradientNode) SecondaryStrideInPixelsY() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952994-secondarystrideinpixelsy?language=objc
func (n_ NNArithmeticGradientNode) SetSecondaryStrideInPixelsY(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952968-secondarystrideinpixelsx?language=objc
func (n_ NNArithmeticGradientNode) SecondaryStrideInPixelsX() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952968-secondarystrideinpixelsx?language=objc
func (n_ NNArithmeticGradientNode) SetSecondaryStrideInPixelsX(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952986-maximumvalue?language=objc
func (n_ NNArithmeticGradientNode) MaximumValue() float64 {
	rv := objc.Call[float64](n_, objc.Sel("maximumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952986-maximumvalue?language=objc
func (n_ NNArithmeticGradientNode) SetMaximumValue(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setMaximumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952993-primaryscale?language=objc
func (n_ NNArithmeticGradientNode) PrimaryScale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("primaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952993-primaryscale?language=objc
func (n_ NNArithmeticGradientNode) SetPrimaryScale(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setPrimaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952989-minimumvalue?language=objc
func (n_ NNArithmeticGradientNode) MinimumValue() float64 {
	rv := objc.Call[float64](n_, objc.Sel("minimumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952989-minimumvalue?language=objc
func (n_ NNArithmeticGradientNode) SetMinimumValue(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setMinimumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952988-bias?language=objc
func (n_ NNArithmeticGradientNode) Bias() float64 {
	rv := objc.Call[float64](n_, objc.Sel("bias"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952988-bias?language=objc
func (n_ NNArithmeticGradientNode) SetBias(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setBias:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952984-secondarystrideinfeaturechannels?language=objc
func (n_ NNArithmeticGradientNode) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952984-secondarystrideinfeaturechannels?language=objc
func (n_ NNArithmeticGradientNode) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}
