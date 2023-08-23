// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetReaderSampleReferenceOutput] class.
var AssetReaderSampleReferenceOutputClass = _AssetReaderSampleReferenceOutputClass{objc.GetClass("AVAssetReaderSampleReferenceOutput")}

type _AssetReaderSampleReferenceOutputClass struct {
	objc.Class
}

// An interface definition for the [AssetReaderSampleReferenceOutput] class.
type IAssetReaderSampleReferenceOutput interface {
	IAssetReaderOutput
	Track() AssetTrack
}

// An object that reads sample references from an asset track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadersamplereferenceoutput?language=objc
type AssetReaderSampleReferenceOutput struct {
	AssetReaderOutput
}

func AssetReaderSampleReferenceOutputFrom(ptr unsafe.Pointer) AssetReaderSampleReferenceOutput {
	return AssetReaderSampleReferenceOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}

func (ac _AssetReaderSampleReferenceOutputClass) AssetReaderSampleReferenceOutputWithTrack(track IAssetTrack) AssetReaderSampleReferenceOutput {
	rv := objc.Call[AssetReaderSampleReferenceOutput](ac, objc.Sel("assetReaderSampleReferenceOutputWithTrack:"), objc.Ptr(track))
	return rv
}

// Returns a new object that supplies sample references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadersamplereferenceoutput/1490320-assetreadersamplereferenceoutput?language=objc
func AssetReaderSampleReferenceOutput_AssetReaderSampleReferenceOutputWithTrack(track IAssetTrack) AssetReaderSampleReferenceOutput {
	return AssetReaderSampleReferenceOutputClass.AssetReaderSampleReferenceOutputWithTrack(track)
}

func (a_ AssetReaderSampleReferenceOutput) InitWithTrack(track IAssetTrack) AssetReaderSampleReferenceOutput {
	rv := objc.Call[AssetReaderSampleReferenceOutput](a_, objc.Sel("initWithTrack:"), objc.Ptr(track))
	return rv
}

// Creates an object that supplies sample references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadersamplereferenceoutput/1387339-initwithtrack?language=objc
func NewAssetReaderSampleReferenceOutputWithTrack(track IAssetTrack) AssetReaderSampleReferenceOutput {
	instance := AssetReaderSampleReferenceOutputClass.Alloc().InitWithTrack(track)
	instance.Autorelease()
	return instance
}

func (ac _AssetReaderSampleReferenceOutputClass) Alloc() AssetReaderSampleReferenceOutput {
	rv := objc.Call[AssetReaderSampleReferenceOutput](ac, objc.Sel("alloc"))
	return rv
}

func AssetReaderSampleReferenceOutput_Alloc() AssetReaderSampleReferenceOutput {
	return AssetReaderSampleReferenceOutputClass.Alloc()
}

func (ac _AssetReaderSampleReferenceOutputClass) New() AssetReaderSampleReferenceOutput {
	rv := objc.Call[AssetReaderSampleReferenceOutput](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetReaderSampleReferenceOutput() AssetReaderSampleReferenceOutput {
	return AssetReaderSampleReferenceOutputClass.New()
}

func (a_ AssetReaderSampleReferenceOutput) Init() AssetReaderSampleReferenceOutput {
	rv := objc.Call[AssetReaderSampleReferenceOutput](a_, objc.Sel("init"))
	return rv
}

// The track from which the output reads sample references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadersamplereferenceoutput/1390057-track?language=objc
func (a_ AssetReaderSampleReferenceOutput) Track() AssetTrack {
	rv := objc.Call[AssetTrack](a_, objc.Sel("track"))
	return rv
}
