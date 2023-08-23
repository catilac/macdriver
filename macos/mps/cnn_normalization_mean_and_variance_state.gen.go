// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNormalizationMeanAndVarianceState] class.
var CNNNormalizationMeanAndVarianceStateClass = _CNNNormalizationMeanAndVarianceStateClass{objc.GetClass("MPSCNNNormalizationMeanAndVarianceState")}

type _CNNNormalizationMeanAndVarianceStateClass struct {
	objc.Class
}

// An interface definition for the [CNNNormalizationMeanAndVarianceState] class.
type ICNNNormalizationMeanAndVarianceState interface {
	IState
	Variance() metal.BufferWrapper
	Mean() metal.BufferWrapper
}

// An object that stores mean and variance terms used to execute batch normalization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate?language=objc
type CNNNormalizationMeanAndVarianceState struct {
	State
}

func CNNNormalizationMeanAndVarianceStateFrom(ptr unsafe.Pointer) CNNNormalizationMeanAndVarianceState {
	return CNNNormalizationMeanAndVarianceState{
		State: StateFrom(ptr),
	}
}

func (c_ CNNNormalizationMeanAndVarianceState) InitWithMeanVariance(mean metal.PBuffer, variance metal.PBuffer) CNNNormalizationMeanAndVarianceState {
	po0 := objc.WrapAsProtocol("MTLBuffer", mean)
	po1 := objc.WrapAsProtocol("MTLBuffer", variance)
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](c_, objc.Sel("initWithMean:variance:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002363-initwithmean?language=objc
func NewCNNNormalizationMeanAndVarianceStateWithMeanVariance(mean metal.PBuffer, variance metal.PBuffer) CNNNormalizationMeanAndVarianceState {
	instance := CNNNormalizationMeanAndVarianceStateClass.Alloc().InitWithMeanVariance(mean, variance)
	instance.Autorelease()
	return instance
}

func (cc _CNNNormalizationMeanAndVarianceStateClass) TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer metal.PCommandBuffer, numberOfFeatureChannels uint) CNNNormalizationMeanAndVarianceState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](cc, objc.Sel("temporaryStateWithCommandBuffer:numberOfFeatureChannels:"), po0, numberOfFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002365-temporarystatewithcommandbuffer?language=objc
func CNNNormalizationMeanAndVarianceState_TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer metal.PCommandBuffer, numberOfFeatureChannels uint) CNNNormalizationMeanAndVarianceState {
	return CNNNormalizationMeanAndVarianceStateClass.TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer, numberOfFeatureChannels)
}

func (cc _CNNNormalizationMeanAndVarianceStateClass) Alloc() CNNNormalizationMeanAndVarianceState {
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](cc, objc.Sel("alloc"))
	return rv
}

func CNNNormalizationMeanAndVarianceState_Alloc() CNNNormalizationMeanAndVarianceState {
	return CNNNormalizationMeanAndVarianceStateClass.Alloc()
}

func (cc _CNNNormalizationMeanAndVarianceStateClass) New() CNNNormalizationMeanAndVarianceState {
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNormalizationMeanAndVarianceState() CNNNormalizationMeanAndVarianceState {
	return CNNNormalizationMeanAndVarianceStateClass.New()
}

func (c_ CNNNormalizationMeanAndVarianceState) Init() CNNNormalizationMeanAndVarianceState {
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNormalizationMeanAndVarianceState) InitWithResources(resources []metal.PResource) CNNNormalizationMeanAndVarianceState {
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNNormalizationMeanAndVarianceStateWithResources(resources []metal.PResource) CNNNormalizationMeanAndVarianceState {
	instance := CNNNormalizationMeanAndVarianceStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNNormalizationMeanAndVarianceState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNNormalizationMeanAndVarianceState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNNormalizationMeanAndVarianceStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNNormalizationMeanAndVarianceState {
	instance := CNNNormalizationMeanAndVarianceStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNNormalizationMeanAndVarianceState) InitWithResource(resource metal.PResource) CNNNormalizationMeanAndVarianceState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNNormalizationMeanAndVarianceStateWithResource(resource metal.PResource) CNNNormalizationMeanAndVarianceState {
	instance := CNNNormalizationMeanAndVarianceStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNNormalizationMeanAndVarianceStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNNormalizationMeanAndVarianceState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationMeanAndVarianceState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNNormalizationMeanAndVarianceState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNNormalizationMeanAndVarianceState {
	return CNNNormalizationMeanAndVarianceStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002366-variance?language=objc
func (c_ CNNNormalizationMeanAndVarianceState) Variance() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("variance"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationmeanandvariancestate/3002364-mean?language=objc
func (c_ CNNNormalizationMeanAndVarianceState) Mean() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("mean"))
	return rv
}
