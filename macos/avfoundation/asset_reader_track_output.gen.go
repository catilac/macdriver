// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetReaderTrackOutput] class.
var AssetReaderTrackOutputClass = _AssetReaderTrackOutputClass{objc.GetClass("AVAssetReaderTrackOutput")}

type _AssetReaderTrackOutputClass struct {
	objc.Class
}

// An interface definition for the [AssetReaderTrackOutput] class.
type IAssetReaderTrackOutput interface {
	IAssetReaderOutput
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm)
	OutputSettings() map[string]objc.Object
	Track() AssetTrack
}

// An object that reads media data from a single track of an asset. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput?language=objc
type AssetReaderTrackOutput struct {
	AssetReaderOutput
}

func AssetReaderTrackOutputFrom(ptr unsafe.Pointer) AssetReaderTrackOutput {
	return AssetReaderTrackOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}

func (ac _AssetReaderTrackOutputClass) AssetReaderTrackOutputWithTrackOutputSettings(track IAssetTrack, outputSettings map[string]objc.IObject) AssetReaderTrackOutput {
	rv := objc.Call[AssetReaderTrackOutput](ac, objc.Sel("assetReaderTrackOutputWithTrack:outputSettings:"), objc.Ptr(track), outputSettings)
	return rv
}

// Returns a new object that reads media data from an asset track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/1490322-assetreadertrackoutputwithtrack?language=objc
func AssetReaderTrackOutput_AssetReaderTrackOutputWithTrackOutputSettings(track IAssetTrack, outputSettings map[string]objc.IObject) AssetReaderTrackOutput {
	return AssetReaderTrackOutputClass.AssetReaderTrackOutputWithTrackOutputSettings(track, outputSettings)
}

func (a_ AssetReaderTrackOutput) InitWithTrackOutputSettings(track IAssetTrack, outputSettings map[string]objc.IObject) AssetReaderTrackOutput {
	rv := objc.Call[AssetReaderTrackOutput](a_, objc.Sel("initWithTrack:outputSettings:"), objc.Ptr(track), outputSettings)
	return rv
}

// Creates an object that reads media data from an asset track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/1385807-initwithtrack?language=objc
func NewAssetReaderTrackOutputWithTrackOutputSettings(track IAssetTrack, outputSettings map[string]objc.IObject) AssetReaderTrackOutput {
	instance := AssetReaderTrackOutputClass.Alloc().InitWithTrackOutputSettings(track, outputSettings)
	instance.Autorelease()
	return instance
}

func (ac _AssetReaderTrackOutputClass) Alloc() AssetReaderTrackOutput {
	rv := objc.Call[AssetReaderTrackOutput](ac, objc.Sel("alloc"))
	return rv
}

func AssetReaderTrackOutput_Alloc() AssetReaderTrackOutput {
	return AssetReaderTrackOutputClass.Alloc()
}

func (ac _AssetReaderTrackOutputClass) New() AssetReaderTrackOutput {
	rv := objc.Call[AssetReaderTrackOutput](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetReaderTrackOutput() AssetReaderTrackOutput {
	return AssetReaderTrackOutputClass.New()
}

func (a_ AssetReaderTrackOutput) Init() AssetReaderTrackOutput {
	rv := objc.Call[AssetReaderTrackOutput](a_, objc.Sel("init"))
	return rv
}

// The processing algorithm to use for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/1387851-audiotimepitchalgorithm?language=objc
func (a_ AssetReaderTrackOutput) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm {
	rv := objc.Call[AudioTimePitchAlgorithm](a_, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}

// The processing algorithm to use for scaled audio edits. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/1387851-audiotimepitchalgorithm?language=objc
func (a_ AssetReaderTrackOutput) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm) {
	objc.Call[objc.Void](a_, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}

// The output settings for this track output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/1387163-outputsettings?language=objc
func (a_ AssetReaderTrackOutput) OutputSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](a_, objc.Sel("outputSettings"))
	return rv
}

// The track from which the output reads sample buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreadertrackoutput/1386921-track?language=objc
func (a_ AssetReaderTrackOutput) Track() AssetTrack {
	rv := objc.Call[AssetTrack](a_, objc.Sel("track"))
	return rv
}
