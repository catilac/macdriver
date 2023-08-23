// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CapturePhoto] class.
var CapturePhotoClass = _CapturePhotoClass{objc.GetClass("AVCapturePhoto")}

type _CapturePhotoClass struct {
	objc.Class
}

// An interface definition for the [CapturePhoto] class.
type ICapturePhoto interface {
	objc.IObject
	FileDataRepresentation() []byte
	CGImageRepresentation() coregraphics.ImageRef
	PixelBuffer() corevideo.PixelBufferRef
	PhotoCount() int
	ResolvedSettings() CaptureResolvedPhotoSettings
	Timestamp() coremedia.Time
}

// A container for image data from a photo capture output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto?language=objc
type CapturePhoto struct {
	objc.Object
}

func CapturePhotoFrom(ptr unsafe.Pointer) CapturePhoto {
	return CapturePhoto{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CapturePhotoClass) Alloc() CapturePhoto {
	rv := objc.Call[CapturePhoto](cc, objc.Sel("alloc"))
	return rv
}

func CapturePhoto_Alloc() CapturePhoto {
	return CapturePhotoClass.Alloc()
}

func (cc _CapturePhotoClass) New() CapturePhoto {
	rv := objc.Call[CapturePhoto](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCapturePhoto() CapturePhoto {
	return CapturePhotoClass.New()
}

func (c_ CapturePhoto) Init() CapturePhoto {
	rv := objc.Call[CapturePhoto](c_, objc.Sel("init"))
	return rv
}

// Generates and returns a flat data representation of the photo and its attachments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/2873919-filedatarepresentation?language=objc
func (c_ CapturePhoto) FileDataRepresentation() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("fileDataRepresentation"))
	return rv
}

// Extracts and returns the captured photo's primary image as a Core Graphics image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/2873963-cgimagerepresentation?language=objc
func (c_ CapturePhoto) CGImageRepresentation() coregraphics.ImageRef {
	rv := objc.Call[coregraphics.ImageRef](c_, objc.Sel("CGImageRepresentation"))
	return rv
}

// The uncompressed or RAW image sample buffer for the photo, if requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/2873914-pixelbuffer?language=objc
func (c_ CapturePhoto) PixelBuffer() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](c_, objc.Sel("pixelBuffer"))
	return rv
}

// The 1-based index of this photo capture relative to other results from the same capture request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/2873906-photocount?language=objc
func (c_ CapturePhoto) PhotoCount() int {
	rv := objc.Call[int](c_, objc.Sel("photoCount"))
	return rv
}

// The settings object that was used to request this photo capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/2873898-resolvedsettings?language=objc
func (c_ CapturePhoto) ResolvedSettings() CaptureResolvedPhotoSettings {
	rv := objc.Call[CaptureResolvedPhotoSettings](c_, objc.Sel("resolvedSettings"))
	return rv
}

// The time at which the image was captured. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturephoto/2873981-timestamp?language=objc
func (c_ CapturePhoto) Timestamp() coremedia.Time {
	rv := objc.Call[coremedia.Time](c_, objc.Sel("timestamp"))
	return rv
}
