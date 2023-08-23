// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionTransposeGradientState] class.
var CNNConvolutionTransposeGradientStateClass = _CNNConvolutionTransposeGradientStateClass{objc.GetClass("MPSCNNConvolutionTransposeGradientState")}

type _CNNConvolutionTransposeGradientStateClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionTransposeGradientState] class.
type ICNNConvolutionTransposeGradientState interface {
	ICNNConvolutionGradientState
	ConvolutionTranspose() CNNConvolutionTranspose
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientstate?language=objc
type CNNConvolutionTransposeGradientState struct {
	CNNConvolutionGradientState
}

func CNNConvolutionTransposeGradientStateFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientState{
		CNNConvolutionGradientState: CNNConvolutionGradientStateFrom(ptr),
	}
}

func (cc _CNNConvolutionTransposeGradientStateClass) Alloc() CNNConvolutionTransposeGradientState {
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionTransposeGradientState_Alloc() CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientStateClass.Alloc()
}

func (cc _CNNConvolutionTransposeGradientStateClass) New() CNNConvolutionTransposeGradientState {
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionTransposeGradientState() CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientStateClass.New()
}

func (c_ CNNConvolutionTransposeGradientState) Init() CNNConvolutionTransposeGradientState {
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNConvolutionTransposeGradientState) InitWithResources(resources []metal.PResource) CNNConvolutionTransposeGradientState {
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNConvolutionTransposeGradientStateWithResources(resources []metal.PResource) CNNConvolutionTransposeGradientState {
	instance := CNNConvolutionTransposeGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionTransposeGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNConvolutionTransposeGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNConvolutionTransposeGradientState {
	instance := CNNConvolutionTransposeGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionTransposeGradientState) InitWithResource(resource metal.PResource) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNConvolutionTransposeGradientState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNConvolutionTransposeGradientStateWithResource(resource metal.PResource) CNNConvolutionTransposeGradientState {
	instance := CNNConvolutionTransposeGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionTransposeGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNConvolutionTransposeGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNConvolutionTransposeGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionTransposeGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNConvolutionTransposeGradientState {
	return CNNConvolutionTransposeGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientstate/3131790-convolutiontranspose?language=objc
func (c_ CNNConvolutionTransposeGradientState) ConvolutionTranspose() CNNConvolutionTranspose {
	rv := objc.Call[CNNConvolutionTranspose](c_, objc.Sel("convolutionTranspose"))
	return rv
}
