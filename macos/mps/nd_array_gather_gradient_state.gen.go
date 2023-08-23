// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayGatherGradientState] class.
var NDArrayGatherGradientStateClass = _NDArrayGatherGradientStateClass{objc.GetClass("MPSNDArrayGatherGradientState")}

type _NDArrayGatherGradientStateClass struct {
	objc.Class
}

// An interface definition for the [NDArrayGatherGradientState] class.
type INDArrayGatherGradientState interface {
	INDArrayGradientState
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygathergradientstate?language=objc
type NDArrayGatherGradientState struct {
	NDArrayGradientState
}

func NDArrayGatherGradientStateFrom(ptr unsafe.Pointer) NDArrayGatherGradientState {
	return NDArrayGatherGradientState{
		NDArrayGradientState: NDArrayGradientStateFrom(ptr),
	}
}

func (nc _NDArrayGatherGradientStateClass) Alloc() NDArrayGatherGradientState {
	rv := objc.Call[NDArrayGatherGradientState](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayGatherGradientState_Alloc() NDArrayGatherGradientState {
	return NDArrayGatherGradientStateClass.Alloc()
}

func (nc _NDArrayGatherGradientStateClass) New() NDArrayGatherGradientState {
	rv := objc.Call[NDArrayGatherGradientState](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayGatherGradientState() NDArrayGatherGradientState {
	return NDArrayGatherGradientStateClass.New()
}

func (n_ NDArrayGatherGradientState) Init() NDArrayGatherGradientState {
	rv := objc.Call[NDArrayGatherGradientState](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayGatherGradientState) InitWithResources(resources []metal.PResource) NDArrayGatherGradientState {
	rv := objc.Call[NDArrayGatherGradientState](n_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewNDArrayGatherGradientStateWithResources(resources []metal.PResource) NDArrayGatherGradientState {
	instance := NDArrayGatherGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGatherGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NDArrayGatherGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGatherGradientState](n_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewNDArrayGatherGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NDArrayGatherGradientState {
	instance := NDArrayGatherGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGatherGradientState) InitWithResource(resource metal.PResource) NDArrayGatherGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[NDArrayGatherGradientState](n_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewNDArrayGatherGradientStateWithResource(resource metal.PResource) NDArrayGatherGradientState {
	instance := NDArrayGatherGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayGatherGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NDArrayGatherGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[NDArrayGatherGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func NDArrayGatherGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NDArrayGatherGradientState {
	return NDArrayGatherGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}
