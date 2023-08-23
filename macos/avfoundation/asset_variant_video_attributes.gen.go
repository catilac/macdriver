// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetVariantVideoAttributes] class.
var AssetVariantVideoAttributesClass = _AssetVariantVideoAttributesClass{objc.GetClass("AVAssetVariantVideoAttributes")}

type _AssetVariantVideoAttributesClass struct {
	objc.Class
}

// An interface definition for the [AssetVariantVideoAttributes] class.
type IAssetVariantVideoAttributes interface {
	objc.IObject
	VideoRange() VideoRange
	CodecTypes() []foundation.Number
	PresentationSize() coregraphics.Size
	NominalFrameRate() float64
}

// An object that defines the video attributes for an asset variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantvideoattributes?language=objc
type AssetVariantVideoAttributes struct {
	objc.Object
}

func AssetVariantVideoAttributesFrom(ptr unsafe.Pointer) AssetVariantVideoAttributes {
	return AssetVariantVideoAttributes{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetVariantVideoAttributesClass) Alloc() AssetVariantVideoAttributes {
	rv := objc.Call[AssetVariantVideoAttributes](ac, objc.Sel("alloc"))
	return rv
}

func AssetVariantVideoAttributes_Alloc() AssetVariantVideoAttributes {
	return AssetVariantVideoAttributesClass.Alloc()
}

func (ac _AssetVariantVideoAttributesClass) New() AssetVariantVideoAttributes {
	rv := objc.Call[AssetVariantVideoAttributes](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetVariantVideoAttributes() AssetVariantVideoAttributes {
	return AssetVariantVideoAttributesClass.New()
}

func (a_ AssetVariantVideoAttributes) Init() AssetVariantVideoAttributes {
	rv := objc.Call[AssetVariantVideoAttributes](a_, objc.Sel("init"))
	return rv
}

// The video range of the variant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantvideoattributes/3746556-videorange?language=objc
func (a_ AssetVariantVideoAttributes) VideoRange() VideoRange {
	rv := objc.Call[VideoRange](a_, objc.Sel("videoRange"))
	return rv
}

// The video sample codec types present in the variant’s renditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantvideoattributes/3746553-codectypes?language=objc
func (a_ AssetVariantVideoAttributes) CodecTypes() []foundation.Number {
	rv := objc.Call[[]foundation.Number](a_, objc.Sel("codecTypes"))
	return rv
}

// The presentation size of the variant’s renditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantvideoattributes/3746555-presentationsize?language=objc
func (a_ AssetVariantVideoAttributes) PresentationSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](a_, objc.Sel("presentationSize"))
	return rv
}

// The nominal frame rate of the variant’s renditions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariantvideoattributes/3746554-nominalframerate?language=objc
func (a_ AssetVariantVideoAttributes) NominalFrameRate() float64 {
	rv := objc.Call[float64](a_, objc.Sel("nominalFrameRate"))
	return rv
}
