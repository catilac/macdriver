// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SVGFDefaultTextureAllocator] class.
var SVGFDefaultTextureAllocatorClass = _SVGFDefaultTextureAllocatorClass{objc.GetClass("MPSSVGFDefaultTextureAllocator")}

type _SVGFDefaultTextureAllocatorClass struct {
	objc.Class
}

// An interface definition for the [SVGFDefaultTextureAllocator] class.
type ISVGFDefaultTextureAllocator interface {
	objc.IObject
	ReturnTexture(texture metal.PTexture)
	ReturnTextureObject(textureObject objc.IObject)
	TextureWithPixelFormatWidthHeight(pixelFormat metal.PixelFormat, width uint, height uint) metal.TextureWrapper
	Reset()
	Device() metal.DeviceWrapper
	AllocatedTextureCount() uint
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator?language=objc
type SVGFDefaultTextureAllocator struct {
	objc.Object
}

func SVGFDefaultTextureAllocatorFrom(ptr unsafe.Pointer) SVGFDefaultTextureAllocator {
	return SVGFDefaultTextureAllocator{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SVGFDefaultTextureAllocator) InitWithDevice(device metal.PDevice) SVGFDefaultTextureAllocator {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[SVGFDefaultTextureAllocator](s_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242897-initwithdevice?language=objc
func NewSVGFDefaultTextureAllocatorWithDevice(device metal.PDevice) SVGFDefaultTextureAllocator {
	instance := SVGFDefaultTextureAllocatorClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (sc _SVGFDefaultTextureAllocatorClass) Alloc() SVGFDefaultTextureAllocator {
	rv := objc.Call[SVGFDefaultTextureAllocator](sc, objc.Sel("alloc"))
	return rv
}

func SVGFDefaultTextureAllocator_Alloc() SVGFDefaultTextureAllocator {
	return SVGFDefaultTextureAllocatorClass.Alloc()
}

func (sc _SVGFDefaultTextureAllocatorClass) New() SVGFDefaultTextureAllocator {
	rv := objc.Call[SVGFDefaultTextureAllocator](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSVGFDefaultTextureAllocator() SVGFDefaultTextureAllocator {
	return SVGFDefaultTextureAllocatorClass.New()
}

func (s_ SVGFDefaultTextureAllocator) Init() SVGFDefaultTextureAllocator {
	rv := objc.Call[SVGFDefaultTextureAllocator](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242899-returntexture?language=objc
func (s_ SVGFDefaultTextureAllocator) ReturnTexture(texture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](s_, objc.Sel("returnTexture:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242899-returntexture?language=objc
func (s_ SVGFDefaultTextureAllocator) ReturnTextureObject(textureObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("returnTexture:"), objc.Ptr(textureObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242900-texturewithpixelformat?language=objc
func (s_ SVGFDefaultTextureAllocator) TextureWithPixelFormatWidthHeight(pixelFormat metal.PixelFormat, width uint, height uint) metal.TextureWrapper {
	rv := objc.Call[metal.TextureWrapper](s_, objc.Sel("textureWithPixelFormat:width:height:"), pixelFormat, width, height)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242898-reset?language=objc
func (s_ SVGFDefaultTextureAllocator) Reset() {
	objc.Call[objc.Void](s_, objc.Sel("reset"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242896-device?language=objc
func (s_ SVGFDefaultTextureAllocator) Device() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](s_, objc.Sel("device"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdefaulttextureallocator/3242895-allocatedtexturecount?language=objc
func (s_ SVGFDefaultTextureAllocator) AllocatedTextureCount() uint {
	rv := objc.Call[uint](s_, objc.Sel("allocatedTextureCount"))
	return rv
}
