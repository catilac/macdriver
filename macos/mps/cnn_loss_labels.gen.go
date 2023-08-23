// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLossLabels] class.
var CNNLossLabelsClass = _CNNLossLabelsClass{objc.GetClass("MPSCNNLossLabels")}

type _CNNLossLabelsClass struct {
	objc.Class
}

// An interface definition for the [CNNLossLabels] class.
type ICNNLossLabels interface {
	IState
	WeightsImage() Image
	LossImage() Image
	LabelsImage() Image
}

// A class that stores the per-element weight buffer used by loss and gradient loss kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels?language=objc
type CNNLossLabels struct {
	State
}

func CNNLossLabelsFrom(ptr unsafe.Pointer) CNNLossLabels {
	return CNNLossLabels{
		State: StateFrom(ptr),
	}
}

func (c_ CNNLossLabels) InitWithDeviceLabelsDescriptor(device metal.PDevice, labelsDescriptor ICNNLossDataDescriptor) CNNLossLabels {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLossLabels](c_, objc.Sel("initWithDevice:labelsDescriptor:"), po0, objc.Ptr(labelsDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2951850-initwithdevice?language=objc
func NewCNNLossLabelsWithDeviceLabelsDescriptor(device metal.PDevice, labelsDescriptor ICNNLossDataDescriptor) CNNLossLabels {
	instance := CNNLossLabelsClass.Alloc().InitWithDeviceLabelsDescriptor(device, labelsDescriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNLossLabelsClass) Alloc() CNNLossLabels {
	rv := objc.Call[CNNLossLabels](cc, objc.Sel("alloc"))
	return rv
}

func CNNLossLabels_Alloc() CNNLossLabels {
	return CNNLossLabelsClass.Alloc()
}

func (cc _CNNLossLabelsClass) New() CNNLossLabels {
	rv := objc.Call[CNNLossLabels](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLossLabels() CNNLossLabels {
	return CNNLossLabelsClass.New()
}

func (c_ CNNLossLabels) Init() CNNLossLabels {
	rv := objc.Call[CNNLossLabels](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNLossLabels) InitWithResources(resources []metal.PResource) CNNLossLabels {
	rv := objc.Call[CNNLossLabels](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNLossLabelsWithResources(resources []metal.PResource) CNNLossLabels {
	instance := CNNLossLabelsClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNLossLabels) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNLossLabels {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLossLabels](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNLossLabelsWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNLossLabels {
	instance := CNNLossLabelsClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNLossLabels) InitWithResource(resource metal.PResource) CNNLossLabels {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNLossLabels](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNLossLabelsWithResource(resource metal.PResource) CNNLossLabels {
	instance := CNNLossLabelsClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNLossLabelsClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNLossLabels {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNLossLabels](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNLossLabels_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNLossLabels {
	return CNNLossLabelsClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2976473-weightsimage?language=objc
func (c_ CNNLossLabels) WeightsImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("weightsImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2951845-lossimage?language=objc
func (c_ CNNLossLabels) LossImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("lossImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlosslabels/2976472-labelsimage?language=objc
func (c_ CNNLossLabels) LabelsImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("labelsImage"))
	return rv
}
