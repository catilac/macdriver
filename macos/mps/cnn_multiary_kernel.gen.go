// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNMultiaryKernel] class.
var CNNMultiaryKernelClass = _CNNMultiaryKernelClass{objc.GetClass("MPSCNNMultiaryKernel")}

type _CNNMultiaryKernelClass struct {
	objc.Class
}

// An interface definition for the [CNNMultiaryKernel] class.
type ICNNMultiaryKernel interface {
	IKernel
	EdgeModeAtIndex(index uint) ImageEdgeMode
	StrideInPixelsXatIndex(index uint) uint
	TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage []IImage, sourceStates []IState, destinationImage IImage) State
	TemporaryResultStateForCommandBufferObjectSourceImagesSourceStatesDestinationImage(commandBufferObject objc.IObject, sourceImage []IImage, sourceStates []IState, destinationImage IImage) State
	ResultStateForSourceImagesSourceStatesDestinationImage(sourceImages []IImage, sourceStates []IState, destinationImage IImage) State
	AppendBatchBarrier() bool
	TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage []*foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	TemporaryResultStateBatchForCommandBufferObjectSourceImagesSourceStatesDestinationImage(commandBufferObject objc.IObject, sourceImage []*foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []IImage, sourceStates []IState) ImageDescriptor
	EncodeBatchToCommandBufferSourceImages(commandBuffer metal.PCommandBuffer, sourceImageBatches []*foundation.Array) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceImages(commandBufferObject objc.IObject, sourceImageBatches []*foundation.Array) *foundation.Array
	SetStrideInPixelsYAtIndex(stride uint, index uint)
	SetKernelWidthAtIndex(width uint, index uint)
	ResultStateBatchForSourceImagesSourceStatesDestinationImage(sourceImages []*foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	SetOffsetAtIndex(offset Offset, index uint)
	SetSourceFeatureChannelMaxCountAtIndex(count uint, index uint)
	SetStrideInPixelsXAtIndex(stride uint, index uint)
	DilationRateXatIndex(index uint) uint
	SetDilationRateXAtIndex(dilationRate uint, index uint)
	OffsetAtIndex(index uint) Offset
	StrideInPixelsYatIndex(index uint) uint
	SetDilationRateYAtIndex(dilationRate uint, index uint)
	SourceFeatureChannelMaxCountAtIndex(index uint) uint
	DilationRateYatIndex(index uint) uint
	KernelWidthAtIndex(index uint) uint
	EncodeToCommandBufferSourceImages(commandBuffer metal.PCommandBuffer, sourceImages []IImage) Image
	EncodeToCommandBufferObjectSourceImages(commandBufferObject objc.IObject, sourceImages []IImage) Image
	SetKernelHeightAtIndex(height uint, index uint)
	IsResultStateReusedAcrossBatch() bool
	SetEdgeModeAtIndex(edgeMode ImageEdgeMode, index uint)
	KernelHeightAtIndex(index uint) uint
	SetSourceFeatureChannelOffsetAtIndex(offset uint, index uint)
	SourceFeatureChannelOffsetAtIndex(index uint) uint
	IsStateModified() bool
	Padding() NNPaddingWrapper
	SetPadding(value PNNPadding)
	SetPaddingObject(valueObject objc.IObject)
	SourceCount() uint
	IsBackwards() bool
	ClipRect() metal.Region
	SetClipRect(value metal.Region)
	DestinationImageAllocator() ImageAllocatorWrapper
	SetDestinationImageAllocator(value PImageAllocator)
	SetDestinationImageAllocatorObject(valueObject objc.IObject)
	DestinationFeatureChannelOffset() uint
	SetDestinationFeatureChannelOffset(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel?language=objc
type CNNMultiaryKernel struct {
	Kernel
}

func CNNMultiaryKernelFrom(ptr unsafe.Pointer) CNNMultiaryKernel {
	return CNNMultiaryKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (c_ CNNMultiaryKernel) InitWithDeviceSourceCount(device metal.PDevice, sourceCount uint) CNNMultiaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNMultiaryKernel](c_, objc.Sel("initWithDevice:sourceCount:"), po0, sourceCount)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043426-initwithdevice?language=objc
func NewCNNMultiaryKernelWithDeviceSourceCount(device metal.PDevice, sourceCount uint) CNNMultiaryKernel {
	instance := CNNMultiaryKernelClass.Alloc().InitWithDeviceSourceCount(device, sourceCount)
	instance.Autorelease()
	return instance
}

func (cc _CNNMultiaryKernelClass) Alloc() CNNMultiaryKernel {
	rv := objc.Call[CNNMultiaryKernel](cc, objc.Sel("alloc"))
	return rv
}

func CNNMultiaryKernel_Alloc() CNNMultiaryKernel {
	return CNNMultiaryKernelClass.Alloc()
}

func (cc _CNNMultiaryKernelClass) New() CNNMultiaryKernel {
	rv := objc.Call[CNNMultiaryKernel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNMultiaryKernel() CNNMultiaryKernel {
	return CNNMultiaryKernelClass.New()
}

func (c_ CNNMultiaryKernel) Init() CNNMultiaryKernel {
	rv := objc.Call[CNNMultiaryKernel](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNMultiaryKernel) InitWithDevice(device metal.PDevice) CNNMultiaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNMultiaryKernel](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewCNNMultiaryKernelWithDevice(device metal.PDevice) CNNMultiaryKernel {
	instance := CNNMultiaryKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNMultiaryKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNMultiaryKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNMultiaryKernel](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNMultiaryKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNMultiaryKernel {
	instance := CNNMultiaryKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043418-edgemodeatindex?language=objc
func (c_ CNNMultiaryKernel) EdgeModeAtIndex(index uint) ImageEdgeMode {
	rv := objc.Call[ImageEdgeMode](c_, objc.Sel("edgeModeAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043449-strideinpixelsxatindex?language=objc
func (c_ CNNMultiaryKernel) StrideInPixelsXatIndex(index uint) uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsXatIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043452-temporaryresultstateforcommandbu?language=objc
func (c_ CNNMultiaryKernel) TemporaryResultStateForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage []IImage, sourceStates []IState, destinationImage IImage) State {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[State](c_, objc.Sel("temporaryResultStateForCommandBuffer:sourceImages:sourceStates:destinationImage:"), po0, sourceImage, sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043452-temporaryresultstateforcommandbu?language=objc
func (c_ CNNMultiaryKernel) TemporaryResultStateForCommandBufferObjectSourceImagesSourceStatesDestinationImage(commandBufferObject objc.IObject, sourceImage []IImage, sourceStates []IState, destinationImage IImage) State {
	rv := objc.Call[State](c_, objc.Sel("temporaryResultStateForCommandBuffer:sourceImages:sourceStates:destinationImage:"), objc.Ptr(commandBufferObject), sourceImage, sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043435-resultstateforsourceimages?language=objc
func (c_ CNNMultiaryKernel) ResultStateForSourceImagesSourceStatesDestinationImage(sourceImages []IImage, sourceStates []IState, destinationImage IImage) State {
	rv := objc.Call[State](c_, objc.Sel("resultStateForSourceImages:sourceStates:destinationImage:"), sourceImages, sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043411-appendbatchbarrier?language=objc
func (c_ CNNMultiaryKernel) AppendBatchBarrier() bool {
	rv := objc.Call[bool](c_, objc.Sel("appendBatchBarrier"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043451-temporaryresultstatebatchforcomm?language=objc
func (c_ CNNMultiaryKernel) TemporaryResultStateBatchForCommandBufferSourceImagesSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage []*foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImages:sourceStates:destinationImage:"), po0, sourceImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043451-temporaryresultstatebatchforcomm?language=objc
func (c_ CNNMultiaryKernel) TemporaryResultStateBatchForCommandBufferObjectSourceImagesSourceStatesDestinationImage(commandBufferObject objc.IObject, sourceImage []*foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImages:sourceStates:destinationImage:"), objc.Ptr(commandBufferObject), sourceImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043415-destinationimagedescriptorforsou?language=objc
func (c_ CNNMultiaryKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []IImage, sourceStates []IState) ImageDescriptor {
	rv := objc.Call[ImageDescriptor](c_, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), sourceImages, sourceStates)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043419-encodebatchtocommandbuffer?language=objc
func (c_ CNNMultiaryKernel) EncodeBatchToCommandBufferSourceImages(commandBuffer metal.PCommandBuffer, sourceImageBatches []*foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:"), po0, sourceImageBatches)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043419-encodebatchtocommandbuffer?language=objc
func (c_ CNNMultiaryKernel) EncodeBatchToCommandBufferObjectSourceImages(commandBufferObject objc.IObject, sourceImageBatches []*foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:"), objc.Ptr(commandBufferObject), sourceImageBatches)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043445-setstrideinpixelsy?language=objc
func (c_ CNNMultiaryKernel) SetStrideInPixelsYAtIndex(stride uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setStrideInPixelsY:atIndex:"), stride, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043440-setkernelwidth?language=objc
func (c_ CNNMultiaryKernel) SetKernelWidthAtIndex(width uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelWidth:atIndex:"), width, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043434-resultstatebatchforsourceimages?language=objc
func (c_ CNNMultiaryKernel) ResultStateBatchForSourceImagesSourceStatesDestinationImage(sourceImages []*foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("resultStateBatchForSourceImages:sourceStates:destinationImage:"), sourceImages, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043441-setoffset?language=objc
func (c_ CNNMultiaryKernel) SetOffsetAtIndex(offset Offset, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setOffset:atIndex:"), offset, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043442-setsourcefeaturechannelmaxcount?language=objc
func (c_ CNNMultiaryKernel) SetSourceFeatureChannelMaxCountAtIndex(count uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSourceFeatureChannelMaxCount:atIndex:"), count, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043444-setstrideinpixelsx?language=objc
func (c_ CNNMultiaryKernel) SetStrideInPixelsXAtIndex(stride uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setStrideInPixelsX:atIndex:"), stride, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043416-dilationratexatindex?language=objc
func (c_ CNNMultiaryKernel) DilationRateXatIndex(index uint) uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateXatIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043436-setdilationratex?language=objc
func (c_ CNNMultiaryKernel) SetDilationRateXAtIndex(dilationRate uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDilationRateX:atIndex:"), dilationRate, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043432-offsetatindex?language=objc
func (c_ CNNMultiaryKernel) OffsetAtIndex(index uint) Offset {
	rv := objc.Call[Offset](c_, objc.Sel("offsetAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043450-strideinpixelsyatindex?language=objc
func (c_ CNNMultiaryKernel) StrideInPixelsYatIndex(index uint) uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsYatIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043437-setdilationratey?language=objc
func (c_ CNNMultiaryKernel) SetDilationRateYAtIndex(dilationRate uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDilationRateY:atIndex:"), dilationRate, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043447-sourcefeaturechannelmaxcountatin?language=objc
func (c_ CNNMultiaryKernel) SourceFeatureChannelMaxCountAtIndex(index uint) uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceFeatureChannelMaxCountAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043417-dilationrateyatindex?language=objc
func (c_ CNNMultiaryKernel) DilationRateYatIndex(index uint) uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateYatIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043431-kernelwidthatindex?language=objc
func (c_ CNNMultiaryKernel) KernelWidthAtIndex(index uint) uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidthAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043422-encodetocommandbuffer?language=objc
func (c_ CNNMultiaryKernel) EncodeToCommandBufferSourceImages(commandBuffer metal.PCommandBuffer, sourceImages []IImage) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImages:"), po0, sourceImages)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043422-encodetocommandbuffer?language=objc
func (c_ CNNMultiaryKernel) EncodeToCommandBufferObjectSourceImages(commandBufferObject objc.IObject, sourceImages []IImage) Image {
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImages:"), objc.Ptr(commandBufferObject), sourceImages)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043439-setkernelheight?language=objc
func (c_ CNNMultiaryKernel) SetKernelHeightAtIndex(height uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelHeight:atIndex:"), height, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043428-isresultstatereusedacrossbatch?language=objc
func (c_ CNNMultiaryKernel) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Call[bool](c_, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043438-setedgemode?language=objc
func (c_ CNNMultiaryKernel) SetEdgeModeAtIndex(edgeMode ImageEdgeMode, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setEdgeMode:atIndex:"), edgeMode, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043430-kernelheightatindex?language=objc
func (c_ CNNMultiaryKernel) KernelHeightAtIndex(index uint) uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeightAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043443-setsourcefeaturechanneloffset?language=objc
func (c_ CNNMultiaryKernel) SetSourceFeatureChannelOffsetAtIndex(offset uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSourceFeatureChannelOffset:atIndex:"), offset, index)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043448-sourcefeaturechanneloffsetatinde?language=objc
func (c_ CNNMultiaryKernel) SourceFeatureChannelOffsetAtIndex(index uint) uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceFeatureChannelOffsetAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043429-isstatemodified?language=objc
func (c_ CNNMultiaryKernel) IsStateModified() bool {
	rv := objc.Call[bool](c_, objc.Sel("isStateModified"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043433-padding?language=objc
func (c_ CNNMultiaryKernel) Padding() NNPaddingWrapper {
	rv := objc.Call[NNPaddingWrapper](c_, objc.Sel("padding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043433-padding?language=objc
func (c_ CNNMultiaryKernel) SetPadding(value PNNPadding) {
	po0 := objc.WrapAsProtocol("MPSNNPadding", value)
	objc.Call[objc.Void](c_, objc.Sel("setPadding:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043433-padding?language=objc
func (c_ CNNMultiaryKernel) SetPaddingObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setPadding:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043446-sourcecount?language=objc
func (c_ CNNMultiaryKernel) SourceCount() uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043427-isbackwards?language=objc
func (c_ CNNMultiaryKernel) IsBackwards() bool {
	rv := objc.Call[bool](c_, objc.Sel("isBackwards"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043412-cliprect?language=objc
func (c_ CNNMultiaryKernel) ClipRect() metal.Region {
	rv := objc.Call[metal.Region](c_, objc.Sel("clipRect"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043412-cliprect?language=objc
func (c_ CNNMultiaryKernel) SetClipRect(value metal.Region) {
	objc.Call[objc.Void](c_, objc.Sel("setClipRect:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043414-destinationimageallocator?language=objc
func (c_ CNNMultiaryKernel) DestinationImageAllocator() ImageAllocatorWrapper {
	rv := objc.Call[ImageAllocatorWrapper](c_, objc.Sel("destinationImageAllocator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043414-destinationimageallocator?language=objc
func (c_ CNNMultiaryKernel) SetDestinationImageAllocator(value PImageAllocator) {
	po0 := objc.WrapAsProtocol("MPSImageAllocator", value)
	objc.Call[objc.Void](c_, objc.Sel("setDestinationImageAllocator:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043414-destinationimageallocator?language=objc
func (c_ CNNMultiaryKernel) SetDestinationImageAllocatorObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDestinationImageAllocator:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043413-destinationfeaturechanneloffset?language=objc
func (c_ CNNMultiaryKernel) DestinationFeatureChannelOffset() uint {
	rv := objc.Call[uint](c_, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiarykernel/3043413-destinationfeaturechanneloffset?language=objc
func (c_ CNNMultiaryKernel) SetDestinationFeatureChannelOffset(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}
