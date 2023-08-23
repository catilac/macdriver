// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TemporaryImage] class.
var TemporaryImageClass = _TemporaryImageClass{objc.GetClass("MPSTemporaryImage")}

type _TemporaryImageClass struct {
	objc.Class
}

// An interface definition for the [TemporaryImage] class.
type ITemporaryImage interface {
	IImage
	ReadCount() uint
	SetReadCount(value uint)
}

// A texture for use in convolutional neural networks that stores transient data to be used and discarded promptly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage?language=objc
type TemporaryImage struct {
	Image
}

func TemporaryImageFrom(ptr unsafe.Pointer) TemporaryImage {
	return TemporaryImage{
		Image: ImageFrom(ptr),
	}
}

func (tc _TemporaryImageClass) TemporaryImageWithCommandBufferTextureDescriptorFeatureChannels(commandBuffer metal.PCommandBuffer, textureDescriptor metal.ITextureDescriptor, featureChannels uint) TemporaryImage {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[TemporaryImage](tc, objc.Sel("temporaryImageWithCommandBuffer:textureDescriptor:featureChannels:"), po0, objc.Ptr(textureDescriptor), featureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2942489-temporaryimagewithcommandbuffer?language=objc
func TemporaryImage_TemporaryImageWithCommandBufferTextureDescriptorFeatureChannels(commandBuffer metal.PCommandBuffer, textureDescriptor metal.ITextureDescriptor, featureChannels uint) TemporaryImage {
	return TemporaryImageClass.TemporaryImageWithCommandBufferTextureDescriptorFeatureChannels(commandBuffer, textureDescriptor, featureChannels)
}

func (tc _TemporaryImageClass) Alloc() TemporaryImage {
	rv := objc.Call[TemporaryImage](tc, objc.Sel("alloc"))
	return rv
}

func TemporaryImage_Alloc() TemporaryImage {
	return TemporaryImageClass.Alloc()
}

func (tc _TemporaryImageClass) New() TemporaryImage {
	rv := objc.Call[TemporaryImage](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTemporaryImage() TemporaryImage {
	return TemporaryImageClass.New()
}

func (t_ TemporaryImage) Init() TemporaryImage {
	rv := objc.Call[TemporaryImage](t_, objc.Sel("init"))
	return rv
}

func (t_ TemporaryImage) InitWithParentImageSliceRangeFeatureChannels(parent IImage, sliceRange foundation.Range, featureChannels uint) TemporaryImage {
	rv := objc.Call[TemporaryImage](t_, objc.Sel("initWithParentImage:sliceRange:featureChannels:"), objc.Ptr(parent), sliceRange, featureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2942493-initwithparentimage?language=objc
func NewTemporaryImageWithParentImageSliceRangeFeatureChannels(parent IImage, sliceRange foundation.Range, featureChannels uint) TemporaryImage {
	instance := TemporaryImageClass.Alloc().InitWithParentImageSliceRangeFeatureChannels(parent, sliceRange, featureChannels)
	instance.Autorelease()
	return instance
}

func (t_ TemporaryImage) InitWithTextureFeatureChannels(texture metal.PTexture, featureChannels uint) TemporaryImage {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	rv := objc.Call[TemporaryImage](t_, objc.Sel("initWithTexture:featureChannels:"), po0, featureChannels)
	return rv
}

// Initializes an image from a texture. The user-allocated texture has been created for a specific number of feature channels and number of images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/2097547-initwithtexture?language=objc
func NewTemporaryImageWithTextureFeatureChannels(texture metal.PTexture, featureChannels uint) TemporaryImage {
	instance := TemporaryImageClass.Alloc().InitWithTextureFeatureChannels(texture, featureChannels)
	instance.Autorelease()
	return instance
}

func (t_ TemporaryImage) InitWithDeviceImageDescriptor(device metal.PDevice, imageDescriptor IImageDescriptor) TemporaryImage {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TemporaryImage](t_, objc.Sel("initWithDevice:imageDescriptor:"), po0, objc.Ptr(imageDescriptor))
	return rv
}

// Initializes an empty image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimage/1648920-initwithdevice?language=objc
func NewTemporaryImageWithDeviceImageDescriptor(device metal.PDevice, imageDescriptor IImageDescriptor) TemporaryImage {
	instance := TemporaryImageClass.Alloc().InitWithDeviceImageDescriptor(device, imageDescriptor)
	instance.Autorelease()
	return instance
}

// A method that helps the framework decide which allocations to make ahead of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097544-prefetchstoragewithcommandbuffer?language=objc
func (tc _TemporaryImageClass) PrefetchStorageWithCommandBufferImageDescriptorList(commandBuffer metal.PCommandBuffer, descriptorList []IImageDescriptor) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](tc, objc.Sel("prefetchStorageWithCommandBuffer:imageDescriptorList:"), po0, descriptorList)
}

// A method that helps the framework decide which allocations to make ahead of time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097544-prefetchstoragewithcommandbuffer?language=objc
func TemporaryImage_PrefetchStorageWithCommandBufferImageDescriptorList(commandBuffer metal.PCommandBuffer, descriptorList []IImageDescriptor) {
	TemporaryImageClass.PrefetchStorageWithCommandBufferImageDescriptorList(commandBuffer, descriptorList)
}

// The number of times a temporary image may be read by a CNN kernel before its contents become undefined. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097546-readcount?language=objc
func (t_ TemporaryImage) ReadCount() uint {
	rv := objc.Call[uint](t_, objc.Sel("readCount"))
	return rv
}

// The number of times a temporary image may be read by a CNN kernel before its contents become undefined. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporaryimage/2097546-readcount?language=objc
func (t_ TemporaryImage) SetReadCount(value uint) {
	objc.Call[objc.Void](t_, objc.Sel("setReadCount:"), value)
}
