// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RNNRecurrentMatrixState] class.
var RNNRecurrentMatrixStateClass = _RNNRecurrentMatrixStateClass{objc.GetClass("MPSRNNRecurrentMatrixState")}

type _RNNRecurrentMatrixStateClass struct {
	objc.Class
}

// An interface definition for the [RNNRecurrentMatrixState] class.
type IRNNRecurrentMatrixState interface {
	IState
	GetMemoryCellMatrixForLayerIndex(layerIndex uint) Matrix
	GetRecurrentOutputMatrixForLayerIndex(layerIndex uint) Matrix
}

// A class holds all the data that's passed from one sequence iteration of the matrix-based recurrent neural network layer to the next. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate?language=objc
type RNNRecurrentMatrixState struct {
	State
}

func RNNRecurrentMatrixStateFrom(ptr unsafe.Pointer) RNNRecurrentMatrixState {
	return RNNRecurrentMatrixState{
		State: StateFrom(ptr),
	}
}

func (rc _RNNRecurrentMatrixStateClass) Alloc() RNNRecurrentMatrixState {
	rv := objc.Call[RNNRecurrentMatrixState](rc, objc.Sel("alloc"))
	return rv
}

func RNNRecurrentMatrixState_Alloc() RNNRecurrentMatrixState {
	return RNNRecurrentMatrixStateClass.Alloc()
}

func (rc _RNNRecurrentMatrixStateClass) New() RNNRecurrentMatrixState {
	rv := objc.Call[RNNRecurrentMatrixState](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRNNRecurrentMatrixState() RNNRecurrentMatrixState {
	return RNNRecurrentMatrixStateClass.New()
}

func (r_ RNNRecurrentMatrixState) Init() RNNRecurrentMatrixState {
	rv := objc.Call[RNNRecurrentMatrixState](r_, objc.Sel("init"))
	return rv
}

func (r_ RNNRecurrentMatrixState) InitWithResources(resources []metal.PResource) RNNRecurrentMatrixState {
	rv := objc.Call[RNNRecurrentMatrixState](r_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewRNNRecurrentMatrixStateWithResources(resources []metal.PResource) RNNRecurrentMatrixState {
	instance := RNNRecurrentMatrixStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (r_ RNNRecurrentMatrixState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) RNNRecurrentMatrixState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNRecurrentMatrixState](r_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewRNNRecurrentMatrixStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) RNNRecurrentMatrixState {
	instance := RNNRecurrentMatrixStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (r_ RNNRecurrentMatrixState) InitWithResource(resource metal.PResource) RNNRecurrentMatrixState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[RNNRecurrentMatrixState](r_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewRNNRecurrentMatrixStateWithResource(resource metal.PResource) RNNRecurrentMatrixState {
	instance := RNNRecurrentMatrixStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (rc _RNNRecurrentMatrixStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) RNNRecurrentMatrixState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[RNNRecurrentMatrixState](rc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func RNNRecurrentMatrixState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) RNNRecurrentMatrixState {
	return RNNRecurrentMatrixStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873390-getmemorycellmatrixforlayerindex?language=objc
func (r_ RNNRecurrentMatrixState) GetMemoryCellMatrixForLayerIndex(layerIndex uint) Matrix {
	rv := objc.Call[Matrix](r_, objc.Sel("getMemoryCellMatrixForLayerIndex:"), layerIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentmatrixstate/2873339-getrecurrentoutputmatrixforlayer?language=objc
func (r_ RNNRecurrentMatrixState) GetRecurrentOutputMatrixForLayerIndex(layerIndex uint) Matrix {
	rv := objc.Call[Matrix](r_, objc.Sel("getRecurrentOutputMatrixForLayerIndex:"), layerIndex)
	return rv
}
