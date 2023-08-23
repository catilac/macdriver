// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PlayerItemTrack] class.
var PlayerItemTrackClass = _PlayerItemTrackClass{objc.GetClass("AVPlayerItemTrack")}

type _PlayerItemTrackClass struct {
	objc.Class
}

// An interface definition for the [PlayerItemTrack] class.
type IPlayerItemTrack interface {
	objc.IObject
	VideoFieldMode() string
	SetVideoFieldMode(value string)
	CurrentVideoFrameRate() float64
	IsEnabled() bool
	SetEnabled(value bool)
	AssetTrack() AssetTrack
}

// An object that represents the presentation state of an asset track during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack?language=objc
type PlayerItemTrack struct {
	objc.Object
}

func PlayerItemTrackFrom(ptr unsafe.Pointer) PlayerItemTrack {
	return PlayerItemTrack{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PlayerItemTrackClass) Alloc() PlayerItemTrack {
	rv := objc.Call[PlayerItemTrack](pc, objc.Sel("alloc"))
	return rv
}

func PlayerItemTrack_Alloc() PlayerItemTrack {
	return PlayerItemTrackClass.Alloc()
}

func (pc _PlayerItemTrackClass) New() PlayerItemTrack {
	rv := objc.Call[PlayerItemTrack](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPlayerItemTrack() PlayerItemTrack {
	return PlayerItemTrackClass.New()
}

func (p_ PlayerItemTrack) Init() PlayerItemTrack {
	rv := objc.Call[PlayerItemTrack](p_, objc.Sel("init"))
	return rv
}

// A mode that specifies the handling of video frames that contain multiple fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack/1388045-videofieldmode?language=objc
func (p_ PlayerItemTrack) VideoFieldMode() string {
	rv := objc.Call[string](p_, objc.Sel("videoFieldMode"))
	return rv
}

// A mode that specifies the handling of video frames that contain multiple fields. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack/1388045-videofieldmode?language=objc
func (p_ PlayerItemTrack) SetVideoFieldMode(value string) {
	objc.Call[objc.Void](p_, objc.Sel("setVideoFieldMode:"), value)
}

// The current frame rate of the video track as it plays. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack/1388956-currentvideoframerate?language=objc
func (p_ PlayerItemTrack) CurrentVideoFrameRate() float64 {
	rv := objc.Call[float64](p_, objc.Sel("currentVideoFrameRate"))
	return rv
}

// A Boolean value that indicates whether the player item presents the track’s media during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack/1387062-enabled?language=objc
func (p_ PlayerItemTrack) IsEnabled() bool {
	rv := objc.Call[bool](p_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether the player item presents the track’s media during playback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack/1387062-enabled?language=objc
func (p_ PlayerItemTrack) SetEnabled(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setEnabled:"), value)
}

// An asset track that provides the media for the player item track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avplayeritemtrack/1390701-assettrack?language=objc
func (p_ PlayerItemTrack) AssetTrack() AssetTrack {
	rv := objc.Call[AssetTrack](p_, objc.Sel("assetTrack"))
	return rv
}
