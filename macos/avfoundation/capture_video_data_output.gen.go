// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureVideoDataOutput] class.
var CaptureVideoDataOutputClass = _CaptureVideoDataOutputClass{objc.GetClass("AVCaptureVideoDataOutput")}

type _CaptureVideoDataOutputClass struct {
	objc.Class
}

// An interface definition for the [CaptureVideoDataOutput] class.
type ICaptureVideoDataOutput interface {
	ICaptureOutput
	AvailableVideoCodecTypesForAssetWriterWithOutputFileType(outputFileType FileType) []VideoCodecType
	RecommendedVideoSettingsForAssetWriterWithOutputFileType(outputFileType FileType) map[string]objc.Object
	RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType, outputFileType FileType) map[string]objc.Object
	SetSampleBufferDelegateQueue(sampleBufferDelegate PCaptureVideoDataOutputSampleBufferDelegate, sampleBufferCallbackQueue dispatch.Queue)
	SetSampleBufferDelegateObjectQueue(sampleBufferDelegateObject objc.IObject, sampleBufferCallbackQueue dispatch.Queue)
	AvailableVideoCVPixelFormatTypes() []foundation.Number
	VideoSettings() map[string]objc.Object
	SetVideoSettings(value map[string]objc.IObject)
	SampleBufferCallbackQueue() dispatch.Queue
	AvailableVideoCodecTypes() []VideoCodecType
	AlwaysDiscardsLateVideoFrames() bool
	SetAlwaysDiscardsLateVideoFrames(value bool)
	SampleBufferDelegate() CaptureVideoDataOutputSampleBufferDelegateWrapper
}

// A capture output that records video and provides access to video frames for processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput?language=objc
type CaptureVideoDataOutput struct {
	CaptureOutput
}

func CaptureVideoDataOutputFrom(ptr unsafe.Pointer) CaptureVideoDataOutput {
	return CaptureVideoDataOutput{
		CaptureOutput: CaptureOutputFrom(ptr),
	}
}

func (cc _CaptureVideoDataOutputClass) New() CaptureVideoDataOutput {
	rv := objc.Call[CaptureVideoDataOutput](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureVideoDataOutput() CaptureVideoDataOutput {
	return CaptureVideoDataOutputClass.New()
}

func (c_ CaptureVideoDataOutput) Init() CaptureVideoDataOutput {
	rv := objc.Call[CaptureVideoDataOutput](c_, objc.Sel("init"))
	return rv
}

func (cc _CaptureVideoDataOutputClass) Alloc() CaptureVideoDataOutput {
	rv := objc.Call[CaptureVideoDataOutput](cc, objc.Sel("alloc"))
	return rv
}

func CaptureVideoDataOutput_Alloc() CaptureVideoDataOutput {
	return CaptureVideoDataOutputClass.Alloc()
}

// The video codecs that the output supports for writing video to the output file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/2867901-availablevideocodectypesforasset?language=objc
func (c_ CaptureVideoDataOutput) AvailableVideoCodecTypesForAssetWriterWithOutputFileType(outputFileType FileType) []VideoCodecType {
	rv := objc.Call[[]VideoCodecType](c_, objc.Sel("availableVideoCodecTypesForAssetWriterWithOutputFileType:"), outputFileType)
	return rv
}

// Specifies the recommended settings for use with an AVAssetWriterInput. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1616290-recommendedvideosettingsforasset?language=objc
func (c_ CaptureVideoDataOutput) RecommendedVideoSettingsForAssetWriterWithOutputFileType(outputFileType FileType) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("recommendedVideoSettingsForAssetWriterWithOutputFileType:"), outputFileType)
	return rv
}

// Returns a video settings dictionary appropriate for capturing video to be recorded to a file with the specified codec and type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/2867900-recommendedvideosettingsforvideo?language=objc
func (c_ CaptureVideoDataOutput) RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType(videoCodecType VideoCodecType, outputFileType FileType) map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("recommendedVideoSettingsForVideoCodecType:assetWriterOutputFileType:"), videoCodecType, outputFileType)
	return rv
}

// Sets the sample buffer delegate and the queue for invoking callbacks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1389008-setsamplebufferdelegate?language=objc
func (c_ CaptureVideoDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate PCaptureVideoDataOutputSampleBufferDelegate, sampleBufferCallbackQueue dispatch.Queue) {
	po0 := objc.WrapAsProtocol("AVCaptureVideoDataOutputSampleBufferDelegate", sampleBufferDelegate)
	objc.Call[objc.Void](c_, objc.Sel("setSampleBufferDelegate:queue:"), po0, sampleBufferCallbackQueue)
}

// Sets the sample buffer delegate and the queue for invoking callbacks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1389008-setsamplebufferdelegate?language=objc
func (c_ CaptureVideoDataOutput) SetSampleBufferDelegateObjectQueue(sampleBufferDelegateObject objc.IObject, sampleBufferCallbackQueue dispatch.Queue) {
	objc.Call[objc.Void](c_, objc.Sel("setSampleBufferDelegate:queue:"), objc.Ptr(sampleBufferDelegateObject), sampleBufferCallbackQueue)
}

// The video pixel formats that the output supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1387050-availablevideocvpixelformattypes?language=objc
func (c_ CaptureVideoDataOutput) AvailableVideoCVPixelFormatTypes() []foundation.Number {
	rv := objc.Call[[]foundation.Number](c_, objc.Sel("availableVideoCVPixelFormatTypes"))
	return rv
}

// A dictionary that contains the compression settings for the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1389945-videosettings?language=objc
func (c_ CaptureVideoDataOutput) VideoSettings() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("videoSettings"))
	return rv
}

// A dictionary that contains the compression settings for the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1389945-videosettings?language=objc
func (c_ CaptureVideoDataOutput) SetVideoSettings(value map[string]objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setVideoSettings:"), value)
}

// The queue on which the system invokes delegate callbacks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1385831-samplebuffercallbackqueue?language=objc
func (c_ CaptureVideoDataOutput) SampleBufferCallbackQueue() dispatch.Queue {
	rv := objc.Call[dispatch.Queue](c_, objc.Sel("sampleBufferCallbackQueue"))
	return rv
}

// The video codecs that the output supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1389227-availablevideocodectypes?language=objc
func (c_ CaptureVideoDataOutput) AvailableVideoCodecTypes() []VideoCodecType {
	rv := objc.Call[[]VideoCodecType](c_, objc.Sel("availableVideoCodecTypes"))
	return rv
}

// Indicates whether to drop video frames if they arrive late. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1385780-alwaysdiscardslatevideoframes?language=objc
func (c_ CaptureVideoDataOutput) AlwaysDiscardsLateVideoFrames() bool {
	rv := objc.Call[bool](c_, objc.Sel("alwaysDiscardsLateVideoFrames"))
	return rv
}

// Indicates whether to drop video frames if they arrive late. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1385780-alwaysdiscardslatevideoframes?language=objc
func (c_ CaptureVideoDataOutput) SetAlwaysDiscardsLateVideoFrames(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setAlwaysDiscardsLateVideoFrames:"), value)
}

// The capture object’s delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/1385886-samplebufferdelegate?language=objc
func (c_ CaptureVideoDataOutput) SampleBufferDelegate() CaptureVideoDataOutputSampleBufferDelegateWrapper {
	rv := objc.Call[CaptureVideoDataOutputSampleBufferDelegateWrapper](c_, objc.Sel("sampleBufferDelegate"))
	return rv
}
