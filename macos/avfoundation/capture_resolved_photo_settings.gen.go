// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureResolvedPhotoSettings] class.
var CaptureResolvedPhotoSettingsClass = _CaptureResolvedPhotoSettingsClass{objc.GetClass("AVCaptureResolvedPhotoSettings")}

type _CaptureResolvedPhotoSettingsClass struct {
	objc.Class
}

// An interface definition for the [CaptureResolvedPhotoSettings] class.
type ICaptureResolvedPhotoSettings interface {
	objc.IObject
	PhotoDimensions() coremedia.VideoDimensions
	ExpectedPhotoCount() uint
	UniqueID() int64
}

// A description of the features and settings in use for an in-progress or complete photo capture request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings?language=objc
type CaptureResolvedPhotoSettings struct {
	objc.Object
}

func CaptureResolvedPhotoSettingsFrom(ptr unsafe.Pointer) CaptureResolvedPhotoSettings {
	return CaptureResolvedPhotoSettings{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureResolvedPhotoSettingsClass) Alloc() CaptureResolvedPhotoSettings {
	rv := objc.Call[CaptureResolvedPhotoSettings](cc, objc.Sel("alloc"))
	return rv
}

func CaptureResolvedPhotoSettings_Alloc() CaptureResolvedPhotoSettings {
	return CaptureResolvedPhotoSettingsClass.Alloc()
}

func (cc _CaptureResolvedPhotoSettingsClass) New() CaptureResolvedPhotoSettings {
	rv := objc.Call[CaptureResolvedPhotoSettings](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureResolvedPhotoSettings() CaptureResolvedPhotoSettings {
	return CaptureResolvedPhotoSettingsClass.New()
}

func (c_ CaptureResolvedPhotoSettings) Init() CaptureResolvedPhotoSettings {
	rv := objc.Call[CaptureResolvedPhotoSettings](c_, objc.Sel("init"))
	return rv
}

// The size, in pixels, of the photo image (in a processed format, such as JPEG) that the capture delivers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/1648782-photodimensions?language=objc
func (c_ CaptureResolvedPhotoSettings) PhotoDimensions() coremedia.VideoDimensions {
	rv := objc.Call[coremedia.VideoDimensions](c_, objc.Sel("photoDimensions"))
	return rv
}

// The number of photo capture results in the capture request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/2873973-expectedphotocount?language=objc
func (c_ CaptureResolvedPhotoSettings) ExpectedPhotoCount() uint {
	rv := objc.Call[uint](c_, objc.Sel("expectedPhotoCount"))
	return rv
}

// The unique identifier for the photo capture this settings object corresponds to. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcaptureresolvedphotosettings/1648656-uniqueid?language=objc
func (c_ CaptureResolvedPhotoSettings) UniqueID() int64 {
	rv := objc.Call[int64](c_, objc.Sel("uniqueID"))
	return rv
}
