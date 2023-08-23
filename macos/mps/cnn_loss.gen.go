// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLoss] class.
var CNNLossClass = _CNNLossClass{objc.GetClass("MPSCNNLoss")}

type _CNNLossClass struct {
	objc.Class
}

// An interface definition for the [CNNLoss] class.
type ICNNLoss interface {
	ICNNKernel
	EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer metal.PCommandBuffer, sourceImage *foundation.Array, labels *foundation.Array) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceImagesLabels(commandBufferObject objc.IObject, sourceImage *foundation.Array, labels *foundation.Array) *foundation.Array
	EncodeToCommandBufferSourceImageLabels(commandBuffer metal.PCommandBuffer, sourceImage IImage, labels ICNNLossLabels) Image
	EncodeToCommandBufferObjectSourceImageLabels(commandBufferObject objc.IObject, sourceImage IImage, labels ICNNLossLabels) Image
	NumberOfClasses() uint
	Weight() float64
	Epsilon() float64
	Delta() float64
	ReductionType() CNNReductionType
	ReduceAcrossBatch() bool
	LossType() CNNLossType
	LabelSmoothing() float64
}

// A kernel that computes the loss and loss gradient between specified predictions and labels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss?language=objc
type CNNLoss struct {
	CNNKernel
}

func CNNLossFrom(ptr unsafe.Pointer) CNNLoss {
	return CNNLoss{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNLoss) InitWithDeviceLossDescriptor(device metal.PDevice, lossDescriptor ICNNLossDescriptor) CNNLoss {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLoss](c_, objc.Sel("initWithDevice:lossDescriptor:"), po0, objc.Ptr(lossDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942377-initwithdevice?language=objc
func NewCNNLossWithDeviceLossDescriptor(device metal.PDevice, lossDescriptor ICNNLossDescriptor) CNNLoss {
	instance := CNNLossClass.Alloc().InitWithDeviceLossDescriptor(device, lossDescriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNLossClass) Alloc() CNNLoss {
	rv := objc.Call[CNNLoss](cc, objc.Sel("alloc"))
	return rv
}

func CNNLoss_Alloc() CNNLoss {
	return CNNLossClass.Alloc()
}

func (cc _CNNLossClass) New() CNNLoss {
	rv := objc.Call[CNNLoss](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLoss() CNNLoss {
	return CNNLossClass.New()
}

func (c_ CNNLoss) Init() CNNLoss {
	rv := objc.Call[CNNLoss](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNLoss) InitWithDevice(device metal.PDevice) CNNLoss {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLoss](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNLossWithDevice(device metal.PDevice) CNNLoss {
	instance := CNNLossClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNLoss) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLoss {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLoss](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNLoss_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLoss {
	instance := CNNLossClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951839-encodebatchtocommandbuffer?language=objc
func (c_ CNNLoss) EncodeBatchToCommandBufferSourceImagesLabels(commandBuffer metal.PCommandBuffer, sourceImage *foundation.Array, labels *foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), po0, sourceImage, labels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951839-encodebatchtocommandbuffer?language=objc
func (c_ CNNLoss) EncodeBatchToCommandBufferObjectSourceImagesLabels(commandBufferObject objc.IObject, sourceImage *foundation.Array, labels *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:"), objc.Ptr(commandBufferObject), sourceImage, labels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951838-encodetocommandbuffer?language=objc
func (c_ CNNLoss) EncodeToCommandBufferSourceImageLabels(commandBuffer metal.PCommandBuffer, sourceImage IImage, labels ICNNLossLabels) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), po0, objc.Ptr(sourceImage), objc.Ptr(labels))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2951838-encodetocommandbuffer?language=objc
func (c_ CNNLoss) EncodeToCommandBufferObjectSourceImageLabels(commandBufferObject objc.IObject, sourceImage IImage, labels ICNNLossLabels) Image {
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImage:labels:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage), objc.Ptr(labels))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942389-numberofclasses?language=objc
func (c_ CNNLoss) NumberOfClasses() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfClasses"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942387-weight?language=objc
func (c_ CNNLoss) Weight() float64 {
	rv := objc.Call[float64](c_, objc.Sel("weight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942371-epsilon?language=objc
func (c_ CNNLoss) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942360-delta?language=objc
func (c_ CNNLoss) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942365-reductiontype?language=objc
func (c_ CNNLoss) ReductionType() CNNReductionType {
	rv := objc.Call[CNNReductionType](c_, objc.Sel("reductionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/3547981-reduceacrossbatch?language=objc
func (c_ CNNLoss) ReduceAcrossBatch() bool {
	rv := objc.Call[bool](c_, objc.Sel("reduceAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942359-losstype?language=objc
func (c_ CNNLoss) LossType() CNNLossType {
	rv := objc.Call[CNNLossType](c_, objc.Sel("lossType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnloss/2942358-labelsmoothing?language=objc
func (c_ CNNLoss) LabelSmoothing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("labelSmoothing"))
	return rv
}
