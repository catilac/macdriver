// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TitlebarAccessoryViewController] class.
var TitlebarAccessoryViewControllerClass = _TitlebarAccessoryViewControllerClass{objc.GetClass("NSTitlebarAccessoryViewController")}

type _TitlebarAccessoryViewControllerClass struct {
	objc.Class
}

// An interface definition for the [TitlebarAccessoryViewController] class.
type ITitlebarAccessoryViewController interface {
	IViewController
	AutomaticallyAdjustsSize() bool
	SetAutomaticallyAdjustsSize(value bool)
	IsHidden() bool
	SetHidden(value bool)
	LayoutAttribute() LayoutAttribute
	SetLayoutAttribute(value LayoutAttribute)
	FullScreenMinHeight() float64
	SetFullScreenMinHeight(value float64)
}

// An object that manages a custom view—known as an accessory view—in the title bar–toolbar area of a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller?language=objc
type TitlebarAccessoryViewController struct {
	ViewController
}

func TitlebarAccessoryViewControllerFrom(ptr unsafe.Pointer) TitlebarAccessoryViewController {
	return TitlebarAccessoryViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}

func (tc _TitlebarAccessoryViewControllerClass) Alloc() TitlebarAccessoryViewController {
	rv := objc.Call[TitlebarAccessoryViewController](tc, objc.Sel("alloc"))
	return rv
}

func TitlebarAccessoryViewController_Alloc() TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.Alloc()
}

func (tc _TitlebarAccessoryViewControllerClass) New() TitlebarAccessoryViewController {
	rv := objc.Call[TitlebarAccessoryViewController](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTitlebarAccessoryViewController() TitlebarAccessoryViewController {
	return TitlebarAccessoryViewControllerClass.New()
}

func (t_ TitlebarAccessoryViewController) Init() TitlebarAccessoryViewController {
	rv := objc.Call[TitlebarAccessoryViewController](t_, objc.Sel("init"))
	return rv
}

func (t_ TitlebarAccessoryViewController) InitWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) TitlebarAccessoryViewController {
	rv := objc.Call[TitlebarAccessoryViewController](t_, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, objc.Ptr(nibBundleOrNil))
	return rv
}

// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434481-initwithnibname?language=objc
func NewTitlebarAccessoryViewControllerWithNibNameBundle(nibNameOrNil NibName, nibBundleOrNil foundation.IBundle) TitlebarAccessoryViewController {
	instance := TitlebarAccessoryViewControllerClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/3656518-automaticallyadjustssize?language=objc
func (t_ TitlebarAccessoryViewController) AutomaticallyAdjustsSize() bool {
	rv := objc.Call[bool](t_, objc.Sel("automaticallyAdjustsSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/3656518-automaticallyadjustssize?language=objc
func (t_ TitlebarAccessoryViewController) SetAutomaticallyAdjustsSize(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setAutomaticallyAdjustsSize:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/2097084-hidden?language=objc
func (t_ TitlebarAccessoryViewController) IsHidden() bool {
	rv := objc.Call[bool](t_, objc.Sel("isHidden"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/2097084-hidden?language=objc
func (t_ TitlebarAccessoryViewController) SetHidden(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setHidden:"), value)
}

// The location of the accessory view, in relation to the window’s title bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/1397778-layoutattribute?language=objc
func (t_ TitlebarAccessoryViewController) LayoutAttribute() LayoutAttribute {
	rv := objc.Call[LayoutAttribute](t_, objc.Sel("layoutAttribute"))
	return rv
}

// The location of the accessory view, in relation to the window’s title bar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/1397778-layoutattribute?language=objc
func (t_ TitlebarAccessoryViewController) SetLayoutAttribute(value LayoutAttribute) {
	objc.Call[objc.Void](t_, objc.Sel("setLayoutAttribute:"), value)
}

// The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/1397782-fullscreenminheight?language=objc
func (t_ TitlebarAccessoryViewController) FullScreenMinHeight() float64 {
	rv := objc.Call[float64](t_, objc.Sel("fullScreenMinHeight"))
	return rv
}

// The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstitlebaraccessoryviewcontroller/1397782-fullscreenminheight?language=objc
func (t_ TitlebarAccessoryViewController) SetFullScreenMinHeight(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setFullScreenMinHeight:"), value)
}
