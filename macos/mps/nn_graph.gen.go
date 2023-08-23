// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGraph] class.
var NNGraphClass = _NNGraphClass{objc.GetClass("MPSNNGraph")}

type _NNGraphClass struct {
	objc.Class
}

// An interface definition for the [NNGraph] class.
type INNGraph interface {
	IKernel
	EncodeBatchToCommandBufferSourceImagesSourceStates(commandBuffer metal.PCommandBuffer, sourceImages []*foundation.Array, sourceStates []*foundation.Array) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceImagesSourceStates(commandBufferObject objc.IObject, sourceImages []*foundation.Array, sourceStates []*foundation.Array) *foundation.Array
	ReadCountForSourceStateAtIndex(index uint) uint
	ReadCountForSourceImageAtIndex(index uint) uint
	ExecuteAsyncWithSourceImagesCompletionHandler(sourceImages []IImage, handler NNGraphCompletionHandler) Image
	EncodeToCommandBufferSourceImages(commandBuffer metal.PCommandBuffer, sourceImages []IImage) Image
	EncodeToCommandBufferObjectSourceImages(commandBufferObject objc.IObject, sourceImages []IImage) Image
	ReloadFromDataSources()
	IntermediateImageHandles() []HandleWrapper
	SourceStateHandles() []HandleWrapper
	ResultStateHandles() []HandleWrapper
	SourceImageHandles() []HandleWrapper
	ResultHandle() HandleWrapper
	DestinationImageAllocator() ImageAllocatorWrapper
	SetDestinationImageAllocator(value PImageAllocator)
	SetDestinationImageAllocatorObject(valueObject objc.IObject)
	ResultImageIsNeeded() bool
	Format() ImageFeatureChannelFormat
	SetFormat(value ImageFeatureChannelFormat)
	OutputStateIsTemporary() bool
	SetOutputStateIsTemporary(value bool)
}

// An optimized representation of a graph of neural network image and filter nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph?language=objc
type NNGraph struct {
	Kernel
}

func NNGraphFrom(ptr unsafe.Pointer) NNGraph {
	return NNGraph{
		Kernel: KernelFrom(ptr),
	}
}

func (nc _NNGraphClass) Alloc() NNGraph {
	rv := objc.Call[NNGraph](nc, objc.Sel("alloc"))
	return rv
}

func NNGraph_Alloc() NNGraph {
	return NNGraphClass.Alloc()
}

func (nc _NNGraphClass) New() NNGraph {
	rv := objc.Call[NNGraph](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGraph() NNGraph {
	return NNGraphClass.New()
}

func (n_ NNGraph) Init() NNGraph {
	rv := objc.Call[NNGraph](n_, objc.Sel("init"))
	return rv
}

func (n_ NNGraph) InitWithDevice(device metal.PDevice) NNGraph {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGraph](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNNGraphWithDevice(device metal.PDevice) NNGraph {
	instance := NNGraphClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNGraph) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNGraph {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGraph](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNGraph_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNGraph {
	instance := NNGraphClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953952-encodebatchtocommandbuffer?language=objc
func (n_ NNGraph) EncodeBatchToCommandBufferSourceImagesSourceStates(commandBuffer metal.PCommandBuffer, sourceImages []*foundation.Array, sourceStates []*foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](n_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:sourceStates:"), po0, sourceImages, sourceStates)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953952-encodebatchtocommandbuffer?language=objc
func (n_ NNGraph) EncodeBatchToCommandBufferObjectSourceImagesSourceStates(commandBufferObject objc.IObject, sourceImages []*foundation.Array, sourceStates []*foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](n_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:sourceStates:"), objc.Ptr(commandBufferObject), sourceImages, sourceStates)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/3037387-readcountforsourcestateatindex?language=objc
func (n_ NNGraph) ReadCountForSourceStateAtIndex(index uint) uint {
	rv := objc.Call[uint](n_, objc.Sel("readCountForSourceStateAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/3037386-readcountforsourceimageatindex?language=objc
func (n_ NNGraph) ReadCountForSourceImageAtIndex(index uint) uint {
	rv := objc.Call[uint](n_, objc.Sel("readCountForSourceImageAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2890826-executeasyncwithsourceimages?language=objc
func (n_ NNGraph) ExecuteAsyncWithSourceImagesCompletionHandler(sourceImages []IImage, handler NNGraphCompletionHandler) Image {
	rv := objc.Call[Image](n_, objc.Sel("executeAsyncWithSourceImages:completionHandler:"), sourceImages, handler)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867036-encodetocommandbuffer?language=objc
func (n_ NNGraph) EncodeToCommandBufferSourceImages(commandBuffer metal.PCommandBuffer, sourceImages []IImage) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](n_, objc.Sel("encodeToCommandBuffer:sourceImages:"), po0, sourceImages)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867036-encodetocommandbuffer?language=objc
func (n_ NNGraph) EncodeToCommandBufferObjectSourceImages(commandBufferObject objc.IObject, sourceImages []IImage) Image {
	rv := objc.Call[Image](n_, objc.Sel("encodeToCommandBuffer:sourceImages:"), objc.Ptr(commandBufferObject), sourceImages)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2976512-reloadfromdatasources?language=objc
func (n_ NNGraph) ReloadFromDataSources() {
	objc.Call[objc.Void](n_, objc.Sel("reloadFromDataSources"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867000-intermediateimagehandles?language=objc
func (n_ NNGraph) IntermediateImageHandles() []HandleWrapper {
	rv := objc.Call[[]HandleWrapper](n_, objc.Sel("intermediateImageHandles"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867056-sourcestatehandles?language=objc
func (n_ NNGraph) SourceStateHandles() []HandleWrapper {
	rv := objc.Call[[]HandleWrapper](n_, objc.Sel("sourceStateHandles"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867149-resultstatehandles?language=objc
func (n_ NNGraph) ResultStateHandles() []HandleWrapper {
	rv := objc.Call[[]HandleWrapper](n_, objc.Sel("resultStateHandles"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867012-sourceimagehandles?language=objc
func (n_ NNGraph) SourceImageHandles() []HandleWrapper {
	rv := objc.Call[[]HandleWrapper](n_, objc.Sel("sourceImageHandles"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867123-resulthandle?language=objc
func (n_ NNGraph) ResultHandle() HandleWrapper {
	rv := objc.Call[HandleWrapper](n_, objc.Sel("resultHandle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2866998-destinationimageallocator?language=objc
func (n_ NNGraph) DestinationImageAllocator() ImageAllocatorWrapper {
	rv := objc.Call[ImageAllocatorWrapper](n_, objc.Sel("destinationImageAllocator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2866998-destinationimageallocator?language=objc
func (n_ NNGraph) SetDestinationImageAllocator(value PImageAllocator) {
	po0 := objc.WrapAsProtocol("MPSImageAllocator", value)
	objc.Call[objc.Void](n_, objc.Sel("setDestinationImageAllocator:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2866998-destinationimageallocator?language=objc
func (n_ NNGraph) SetDestinationImageAllocatorObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setDestinationImageAllocator:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953954-resultimageisneeded?language=objc
func (n_ NNGraph) ResultImageIsNeeded() bool {
	rv := objc.Call[bool](n_, objc.Sel("resultImageIsNeeded"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953133-format?language=objc
func (n_ NNGraph) Format() ImageFeatureChannelFormat {
	rv := objc.Call[ImageFeatureChannelFormat](n_, objc.Sel("format"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2953133-format?language=objc
func (n_ NNGraph) SetFormat(value ImageFeatureChannelFormat) {
	objc.Call[objc.Void](n_, objc.Sel("setFormat:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867094-outputstateistemporary?language=objc
func (n_ NNGraph) OutputStateIsTemporary() bool {
	rv := objc.Call[bool](n_, objc.Sel("outputStateIsTemporary"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngraph/2867094-outputstateistemporary?language=objc
func (n_ NNGraph) SetOutputStateIsTemporary(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setOutputStateIsTemporary:"), value)
}
