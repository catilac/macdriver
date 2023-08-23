// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScannerDeviceView] class.
var ScannerDeviceViewClass = _ScannerDeviceViewClass{objc.GetClass("IKScannerDeviceView")}

type _ScannerDeviceViewClass struct {
	objc.Class
}

// An interface definition for the [ScannerDeviceView] class.
type IScannerDeviceView interface {
	appkit.IView
	DownloadsDirectory() foundation.URL
	SetDownloadsDirectory(value foundation.IURL)
	ScanControlLabel() string
	SetScanControlLabel(value string)
	HasDisplayModeSimple() bool
	SetHasDisplayModeSimple(value bool)
	DisplaysDownloadsDirectoryControl() bool
	SetDisplaysDownloadsDirectoryControl(value bool)
	DisplaysPostProcessApplicationControl() bool
	SetDisplaysPostProcessApplicationControl(value bool)
	TransferMode() ScannerDeviceViewTransferMode
	SetTransferMode(value ScannerDeviceViewTransferMode)
	Delegate() ScannerDeviceViewDelegateWrapper
	SetDelegate(value PScannerDeviceViewDelegate)
	SetDelegateObject(valueObject objc.IObject)
	PostProcessApplication() foundation.URL
	SetPostProcessApplication(value foundation.IURL)
	DocumentName() string
	SetDocumentName(value string)
	Mode() ScannerDeviceViewDisplayMode
	SetMode(value ScannerDeviceViewDisplayMode)
	OverviewControlLabel() string
	SetOverviewControlLabel(value string)
	HasDisplayModeAdvanced() bool
	SetHasDisplayModeAdvanced(value bool)
}

// The IKScannerDeviceView class displays a view that allows scanning. It can be customized by specifying the display mode. The delegate receives the scanned data and must implement the IKScannerDeviceViewDelegate protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview?language=objc
type ScannerDeviceView struct {
	appkit.View
}

func ScannerDeviceViewFrom(ptr unsafe.Pointer) ScannerDeviceView {
	return ScannerDeviceView{
		View: appkit.ViewFrom(ptr),
	}
}

func (sc _ScannerDeviceViewClass) Alloc() ScannerDeviceView {
	rv := objc.Call[ScannerDeviceView](sc, objc.Sel("alloc"))
	return rv
}

func ScannerDeviceView_Alloc() ScannerDeviceView {
	return ScannerDeviceViewClass.Alloc()
}

func (sc _ScannerDeviceViewClass) New() ScannerDeviceView {
	rv := objc.Call[ScannerDeviceView](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScannerDeviceView() ScannerDeviceView {
	return ScannerDeviceViewClass.New()
}

func (s_ ScannerDeviceView) Init() ScannerDeviceView {
	rv := objc.Call[ScannerDeviceView](s_, objc.Sel("init"))
	return rv
}

func (s_ ScannerDeviceView) InitWithFrame(frameRect foundation.Rect) ScannerDeviceView {
	rv := objc.Call[ScannerDeviceView](s_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes and returns a newly allocated NSView object with a specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/1483458-initwithframe?language=objc
func NewScannerDeviceViewWithFrame(frameRect foundation.Rect) ScannerDeviceView {
	instance := ScannerDeviceViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// The directory where scans are saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1503585-downloadsdirectory?language=objc
func (s_ ScannerDeviceView) DownloadsDirectory() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("downloadsDirectory"))
	return rv
}

// The directory where scans are saved. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1503585-downloadsdirectory?language=objc
func (s_ ScannerDeviceView) SetDownloadsDirectory(value foundation.IURL) {
	objc.Call[objc.Void](s_, objc.Sel("setDownloadsDirectory:"), objc.Ptr(value))
}

// Allows customization of the “Scan” label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504143-scancontrollabel?language=objc
func (s_ ScannerDeviceView) ScanControlLabel() string {
	rv := objc.Call[string](s_, objc.Sel("scanControlLabel"))
	return rv
}

// Allows customization of the “Scan” label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504143-scancontrollabel?language=objc
func (s_ ScannerDeviceView) SetScanControlLabel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setScanControlLabel:"), value)
}

// The property that determines whether the scanner view uses the simple display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504506-hasdisplaymodesimple?language=objc
func (s_ ScannerDeviceView) HasDisplayModeSimple() bool {
	rv := objc.Call[bool](s_, objc.Sel("hasDisplayModeSimple"))
	return rv
}

// The property that determines whether the scanner view uses the simple display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504506-hasdisplaymodesimple?language=objc
func (s_ ScannerDeviceView) SetHasDisplayModeSimple(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setHasDisplayModeSimple:"), value)
}

// Determines whether the downloads directory control is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1503830-displaysdownloadsdirectorycontro?language=objc
func (s_ ScannerDeviceView) DisplaysDownloadsDirectoryControl() bool {
	rv := objc.Call[bool](s_, objc.Sel("displaysDownloadsDirectoryControl"))
	return rv
}

// Determines whether the downloads directory control is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1503830-displaysdownloadsdirectorycontro?language=objc
func (s_ ScannerDeviceView) SetDisplaysDownloadsDirectoryControl(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setDisplaysDownloadsDirectoryControl:"), value)
}

// Specifies whether the post processing application control is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1505053-displayspostprocessapplicationco?language=objc
func (s_ ScannerDeviceView) DisplaysPostProcessApplicationControl() bool {
	rv := objc.Call[bool](s_, objc.Sel("displaysPostProcessApplicationControl"))
	return rv
}

// Specifies whether the post processing application control is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1505053-displayspostprocessapplicationco?language=objc
func (s_ ScannerDeviceView) SetDisplaysPostProcessApplicationControl(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setDisplaysPostProcessApplicationControl:"), value)
}

// Determines how the scanned content is provided to the delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1505017-transfermode?language=objc
func (s_ ScannerDeviceView) TransferMode() ScannerDeviceViewTransferMode {
	rv := objc.Call[ScannerDeviceViewTransferMode](s_, objc.Sel("transferMode"))
	return rv
}

// Determines how the scanned content is provided to the delegate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1505017-transfermode?language=objc
func (s_ ScannerDeviceView) SetTransferMode(value ScannerDeviceViewTransferMode) {
	objc.Call[objc.Void](s_, objc.Sel("setTransferMode:"), value)
}

// The scanner device delegate [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504170-delegate?language=objc
func (s_ ScannerDeviceView) Delegate() ScannerDeviceViewDelegateWrapper {
	rv := objc.Call[ScannerDeviceViewDelegateWrapper](s_, objc.Sel("delegate"))
	return rv
}

// The scanner device delegate [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504170-delegate?language=objc
func (s_ ScannerDeviceView) SetDelegate(value PScannerDeviceViewDelegate) {
	po0 := objc.WrapAsProtocol("IKScannerDeviceViewDelegate", value)
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), po0)
}

// The scanner device delegate [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504170-delegate?language=objc
func (s_ ScannerDeviceView) SetDelegateObject(valueObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("setDelegate:"), objc.Ptr(valueObject))
}

// The URL of the application to use for post processing of the scan. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1505215-postprocessapplication?language=objc
func (s_ ScannerDeviceView) PostProcessApplication() foundation.URL {
	rv := objc.Call[foundation.URL](s_, objc.Sel("postProcessApplication"))
	return rv
}

// The URL of the application to use for post processing of the scan. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1505215-postprocessapplication?language=objc
func (s_ ScannerDeviceView) SetPostProcessApplication(value foundation.IURL) {
	objc.Call[objc.Void](s_, objc.Sel("setPostProcessApplication:"), objc.Ptr(value))
}

// Returns the document name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1503744-documentname?language=objc
func (s_ ScannerDeviceView) DocumentName() string {
	rv := objc.Call[string](s_, objc.Sel("documentName"))
	return rv
}

// Returns the document name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1503744-documentname?language=objc
func (s_ ScannerDeviceView) SetDocumentName(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setDocumentName:"), value)
}

// The display mode used by the device view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504321-mode?language=objc
func (s_ ScannerDeviceView) Mode() ScannerDeviceViewDisplayMode {
	rv := objc.Call[ScannerDeviceViewDisplayMode](s_, objc.Sel("mode"))
	return rv
}

// The display mode used by the device view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504321-mode?language=objc
func (s_ ScannerDeviceView) SetMode(value ScannerDeviceViewDisplayMode) {
	objc.Call[objc.Void](s_, objc.Sel("setMode:"), value)
}

// Allows customization of the “Overview” label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504055-overviewcontrollabel?language=objc
func (s_ ScannerDeviceView) OverviewControlLabel() string {
	rv := objc.Call[string](s_, objc.Sel("overviewControlLabel"))
	return rv
}

// Allows customization of the “Overview” label. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1504055-overviewcontrollabel?language=objc
func (s_ ScannerDeviceView) SetOverviewControlLabel(value string) {
	objc.Call[objc.Void](s_, objc.Sel("setOverviewControlLabel:"), value)
}

// The property that determines whether the scanner view uses the advanced display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1503437-hasdisplaymodeadvanced?language=objc
func (s_ ScannerDeviceView) HasDisplayModeAdvanced() bool {
	rv := objc.Call[bool](s_, objc.Sel("hasDisplayModeAdvanced"))
	return rv
}

// The property that determines whether the scanner view uses the advanced display mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikscannerdeviceview/1503437-hasdisplaymodeadvanced?language=objc
func (s_ ScannerDeviceView) SetHasDisplayModeAdvanced(value bool) {
	objc.Call[objc.Void](s_, objc.Sel("setHasDisplayModeAdvanced:"), value)
}
