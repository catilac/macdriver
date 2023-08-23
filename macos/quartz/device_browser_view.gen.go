// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DeviceBrowserView] class.
var DeviceBrowserViewClass = _DeviceBrowserViewClass{objc.GetClass("IKDeviceBrowserView")}

type _DeviceBrowserViewClass struct {
	objc.Class
}

// An interface definition for the [DeviceBrowserView] class.
type IDeviceBrowserView interface {
	appkit.IView
	DisplaysNetworkCameras() bool
	SetDisplaysNetworkCameras(value bool)
	Delegate() DeviceBrowserViewDelegateWrapper
	SetDelegate(value PDeviceBrowserViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	DisplaysNetworkScanners() bool
	SetDisplaysNetworkScanners(value bool)
	DisplaysLocalCameras() bool
	SetDisplaysLocalCameras(value bool)
	Mode() DeviceBrowserViewDisplayMode
	SetMode(value DeviceBrowserViewDisplayMode)
	DisplaysLocalScanners() bool
	SetDisplaysLocalScanners(value bool)
}

// The IKDeviceBrowserView allows you to select a camera or scanner from a list of the available devices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview?language=objc
type DeviceBrowserView struct {
	appkit.View
}

func DeviceBrowserViewFrom(ptr unsafe.Pointer) DeviceBrowserView {
	return DeviceBrowserView{
		View: appkit.ViewFrom(ptr),
	}
}

func (dc _DeviceBrowserViewClass) Alloc() DeviceBrowserView {
	rv := objc.Call[DeviceBrowserView](dc, objc.Sel("alloc"))
	return rv
}

func DeviceBrowserView_Alloc() DeviceBrowserView {
	return DeviceBrowserViewClass.Alloc()
}

func (dc _DeviceBrowserViewClass) New() DeviceBrowserView {
	rv := objc.Call[DeviceBrowserView](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDeviceBrowserView() DeviceBrowserView {
	return DeviceBrowserViewClass.New()
}

func (d_ DeviceBrowserView) Init() DeviceBrowserView {
	rv := objc.Call[DeviceBrowserView](d_, objc.Sel("init"))
	return rv
}

func (d_ DeviceBrowserView) InitWithFrame(frameRect foundation.Rect) DeviceBrowserView {
	rv := objc.Call[DeviceBrowserView](d_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewDeviceBrowserViewWithFrame(frameRect foundation.Rect) DeviceBrowserView {
	instance := DeviceBrowserViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Specifies whether network cameras are displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443056-displaysnetworkcameras?language=objc
func (d_ DeviceBrowserView) DisplaysNetworkCameras() bool {
	rv := objc.Call[bool](d_, objc.Sel("displaysNetworkCameras"))
	return rv
}

// Specifies whether network cameras are displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443056-displaysnetworkcameras?language=objc
func (d_ DeviceBrowserView) SetDisplaysNetworkCameras(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDisplaysNetworkCameras:"), value)
}

// Specifies the delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443054-delegate?language=objc
func (d_ DeviceBrowserView) Delegate() DeviceBrowserViewDelegateWrapper {
	rv := objc.Call[DeviceBrowserViewDelegateWrapper](d_, objc.Sel("delegate"))
	return rv
}

// Specifies the delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443054-delegate?language=objc
func (d_ DeviceBrowserView) SetDelegate(value PDeviceBrowserViewDelegate) {
	po0 := objc.WrapAsProtocol("IKDeviceBrowserViewDelegate", value)
	objc.Call[objc.Void](d_, objc.Sel("setDelegate:"), po0)
}

// Specifies the delegate object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443054-delegate?language=objc
func (d_ DeviceBrowserView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](d_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// Specifies whether network scanners are displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443078-displaysnetworkscanners?language=objc
func (d_ DeviceBrowserView) DisplaysNetworkScanners() bool {
	rv := objc.Call[bool](d_, objc.Sel("displaysNetworkScanners"))
	return rv
}

// Specifies whether network scanners are displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443078-displaysnetworkscanners?language=objc
func (d_ DeviceBrowserView) SetDisplaysNetworkScanners(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDisplaysNetworkScanners:"), value)
}

// Specifies whether local cameras are displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443061-displayslocalcameras?language=objc
func (d_ DeviceBrowserView) DisplaysLocalCameras() bool {
	rv := objc.Call[bool](d_, objc.Sel("displaysLocalCameras"))
	return rv
}

// Specifies whether local cameras are displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443061-displayslocalcameras?language=objc
func (d_ DeviceBrowserView) SetDisplaysLocalCameras(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDisplaysLocalCameras:"), value)
}

// Specifies the browser display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443080-mode?language=objc
func (d_ DeviceBrowserView) Mode() DeviceBrowserViewDisplayMode {
	rv := objc.Call[DeviceBrowserViewDisplayMode](d_, objc.Sel("mode"))
	return rv
}

// Specifies the browser display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443080-mode?language=objc
func (d_ DeviceBrowserView) SetMode(value DeviceBrowserViewDisplayMode) {
	objc.Call[objc.Void](d_, objc.Sel("setMode:"), value)
}

// Specifies whether local scanners are displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443052-displayslocalscanners?language=objc
func (d_ DeviceBrowserView) DisplaysLocalScanners() bool {
	rv := objc.Call[bool](d_, objc.Sel("displaysLocalScanners"))
	return rv
}

// Specifies whether local scanners are displayed by the browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikdevicebrowserview/1443052-displayslocalscanners?language=objc
func (d_ DeviceBrowserView) SetDisplaysLocalScanners(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDisplaysLocalScanners:"), value)
}
