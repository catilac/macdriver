// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VideoComposition] class.
var VideoCompositionClass = _VideoCompositionClass{objc.GetClass("AVVideoComposition")}

type _VideoCompositionClass struct {
	objc.Class
}

// An interface definition for the [VideoComposition] class.
type IVideoComposition interface {
	objc.IObject
	ColorPrimaries() string
	ColorYCbCrMatrix() string
	CustomVideoCompositorClass() objc.Class
	FrameDuration() coremedia.Time
	Instructions() []objc.Object
	SourceSampleDataTrackIDs() []foundation.Number
	AnimationTool() VideoCompositionCoreAnimationTool
	ColorTransferFunction() string
	SourceTrackIDForFrameTiming() objc.Object
	RenderSize() coregraphics.Size
	RenderScale() float64
}

// An object that describes how to compose video frames at particular points in time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition?language=objc
type VideoComposition struct {
	objc.Object
}

func VideoCompositionFrom(ptr unsafe.Pointer) VideoComposition {
	return VideoComposition{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VideoCompositionClass) Alloc() VideoComposition {
	rv := objc.Call[VideoComposition](vc, objc.Sel("alloc"))
	return rv
}

func VideoComposition_Alloc() VideoComposition {
	return VideoCompositionClass.Alloc()
}

func (vc _VideoCompositionClass) New() VideoComposition {
	rv := objc.Call[VideoComposition](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVideoComposition() VideoComposition {
	return VideoCompositionClass.New()
}

func (v_ VideoComposition) Init() VideoComposition {
	rv := objc.Call[VideoComposition](v_, objc.Sel("init"))
	return rv
}

// The color primaries used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1643235-colorprimaries?language=objc
func (v_ VideoComposition) ColorPrimaries() string {
	rv := objc.Call[string](v_, objc.Sel("colorPrimaries"))
	return rv
}

// The YCbCr matrix used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1643236-colorycbcrmatrix?language=objc
func (v_ VideoComposition) ColorYCbCrMatrix() string {
	rv := objc.Call[string](v_, objc.Sel("colorYCbCrMatrix"))
	return rv
}

// A custom compositor class to use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1389622-customvideocompositorclass?language=objc
func (v_ VideoComposition) CustomVideoCompositorClass() objc.Class {
	rv := objc.Call[objc.Class](v_, objc.Sel("customVideoCompositorClass"))
	return rv
}

// A time interval for which the video composition should render composed video frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1388013-frameduration?language=objc
func (v_ VideoComposition) FrameDuration() coremedia.Time {
	rv := objc.Call[coremedia.Time](v_, objc.Sel("frameDuration"))
	return rv
}

// The video composition instructions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1389211-instructions?language=objc
func (v_ VideoComposition) Instructions() []objc.Object {
	rv := objc.Call[[]objc.Object](v_, objc.Sel("instructions"))
	return rv
}

// The identifiers of source sample data tracks in the composition that the compositor requires to compose frames. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/3750318-sourcesampledatatrackids?language=objc
func (v_ VideoComposition) SourceSampleDataTrackIDs() []foundation.Number {
	rv := objc.Call[[]foundation.Number](v_, objc.Sel("sourceSampleDataTrackIDs"))
	return rv
}

// A video composition tool to use with Core Animation in offline rendering. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1387030-animationtool?language=objc
func (v_ VideoComposition) AnimationTool() VideoCompositionCoreAnimationTool {
	rv := objc.Call[VideoCompositionCoreAnimationTool](v_, objc.Sel("animationTool"))
	return rv
}

// The transfer function used for video composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1643230-colortransferfunction?language=objc
func (v_ VideoComposition) ColorTransferFunction() string {
	rv := objc.Call[string](v_, objc.Sel("colorTransferFunction"))
	return rv
}

// An identifier of the source track from which the video composition derives frame timing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/2873798-sourcetrackidforframetiming?language=objc
func (v_ VideoComposition) SourceTrackIDForFrameTiming() objc.Object {
	rv := objc.Call[objc.Object](v_, objc.Sel("sourceTrackIDForFrameTiming"))
	return rv
}

// The size at which the video composition should render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1388705-rendersize?language=objc
func (v_ VideoComposition) RenderSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](v_, objc.Sel("renderSize"))
	return rv
}

// The scale at which the video composition should render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avvideocomposition/1615786-renderscale?language=objc
func (v_ VideoComposition) RenderScale() float64 {
	rv := objc.Call[float64](v_, objc.Sel("renderScale"))
	return rv
}
