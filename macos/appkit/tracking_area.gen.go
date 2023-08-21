// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TrackingArea] class.
var TrackingAreaClass = _TrackingAreaClass{objc.GetClass("NSTrackingArea")}

type _TrackingAreaClass struct {
	objc.Class
}

// An interface definition for the [TrackingArea] class.
type ITrackingArea interface {
	objc.IObject
	Options() TrackingAreaOptions
	Owner() objc.Object
	Rect() foundation.Rect
	UserInfo() foundation.Dictionary
}

// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea?language=objc
type TrackingArea struct {
	objc.Object
}

func TrackingAreaFrom(ptr unsafe.Pointer) TrackingArea {
	return TrackingArea{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TrackingArea) InitWithRectOptionsOwnerUserInfo(rect foundation.Rect, options TrackingAreaOptions, owner objc.IObject, userInfo foundation.Dictionary) TrackingArea {
	rv := objc.Call[TrackingArea](t_, objc.Sel("initWithRect:options:owner:userInfo:"), rect, options, owner, userInfo)
	return rv
}

// Initializes and returns an object defining a region of a view to receive mouse-tracking events, mouse-moved events, cursor-update events, or possibly all these events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/1524488-initwithrect?language=objc
func NewTrackingAreaWithRectOptionsOwnerUserInfo(rect foundation.Rect, options TrackingAreaOptions, owner objc.IObject, userInfo foundation.Dictionary) TrackingArea {
	instance := TrackingAreaClass.Alloc().InitWithRectOptionsOwnerUserInfo(rect, options, owner, userInfo)
	instance.Autorelease()
	return instance
}

func (tc _TrackingAreaClass) Alloc() TrackingArea {
	rv := objc.Call[TrackingArea](tc, objc.Sel("alloc"))
	return rv
}

func TrackingArea_Alloc() TrackingArea {
	return TrackingAreaClass.Alloc()
}

func (tc _TrackingAreaClass) New() TrackingArea {
	rv := objc.Call[TrackingArea](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTrackingArea() TrackingArea {
	return TrackingAreaClass.New()
}

func (t_ TrackingArea) Init() TrackingArea {
	rv := objc.Call[TrackingArea](t_, objc.Sel("init"))
	return rv
}

// The options specified for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/1533013-options?language=objc
func (t_ TrackingArea) Options() TrackingAreaOptions {
	rv := objc.Call[TrackingAreaOptions](t_, objc.Sel("options"))
	return rv
}

// The object owning the receiver, which is the recipient of mouse-tracking, mouse-movement, and cursor-update messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/1525965-owner?language=objc
func (t_ TrackingArea) Owner() objc.Object {
	rv := objc.Call[objc.Object](t_, objc.Sel("owner"))
	return rv
}

// The rectangle defining the area encompassed by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/1525874-rect?language=objc
func (t_ TrackingArea) Rect() foundation.Rect {
	rv := objc.Call[foundation.Rect](t_, objc.Sel("rect"))
	return rv
}

// The dictionary containing the data associated with the receiver when it was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstrackingarea/1527949-userinfo?language=objc
func (t_ TrackingArea) UserInfo() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](t_, objc.Sel("userInfo"))
	return rv
}
