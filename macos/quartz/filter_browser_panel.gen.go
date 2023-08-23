// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [FilterBrowserPanel] class.
var FilterBrowserPanelClass = _FilterBrowserPanelClass{objc.GetClass("IKFilterBrowserPanel")}

type _FilterBrowserPanelClass struct {
	objc.Class
}

// An interface definition for the [FilterBrowserPanel] class.
type IFilterBrowserPanel interface {
	appkit.IPanel
	BeginWithOptionsModelessDelegateDidEndSelectorContextInfo(inOptions foundation.Dictionary, modelessDelegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	FilterBrowserViewWithOptions(inOptions foundation.Dictionary) FilterBrowserView
	FilterName() string
	BeginSheetWithOptionsModalForWindowModalDelegateDidEndSelectorContextInfo(inOptions foundation.Dictionary, docWindow appkit.IWindow, modalDelegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer)
	Finish(sender objc.IObject)
	RunModalWithOptions(inOptions foundation.Dictionary) int
}

// The IKFilterBrowserPanel class provides a user interface that allows users to browse Core Image filters (CIFilter), to preview a filter, and to get additional information about the filter, such as its description. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel?language=objc
type FilterBrowserPanel struct {
	appkit.Panel
}

func FilterBrowserPanelFrom(ptr unsafe.Pointer) FilterBrowserPanel {
	return FilterBrowserPanel{
		Panel: appkit.PanelFrom(ptr),
	}
}

func (fc _FilterBrowserPanelClass) Alloc() FilterBrowserPanel {
	rv := objc.Call[FilterBrowserPanel](fc, objc.Sel("alloc"))
	return rv
}

func FilterBrowserPanel_Alloc() FilterBrowserPanel {
	return FilterBrowserPanelClass.Alloc()
}

func (fc _FilterBrowserPanelClass) New() FilterBrowserPanel {
	rv := objc.Call[FilterBrowserPanel](fc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewFilterBrowserPanel() FilterBrowserPanel {
	return FilterBrowserPanelClass.New()
}

func (f_ FilterBrowserPanel) Init() FilterBrowserPanel {
	rv := objc.Call[FilterBrowserPanel](f_, objc.Sel("init"))
	return rv
}

func (fc _FilterBrowserPanelClass) WindowWithContentViewController(contentViewController appkit.IViewController) FilterBrowserPanel {
	rv := objc.Call[FilterBrowserPanel](fc, objc.Sel("windowWithContentViewController:"), objc.Ptr(contentViewController))
	return rv
}

// Creates a titled window that contains the specified content view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419551-windowwithcontentviewcontroller?language=objc
func FilterBrowserPanel_WindowWithContentViewController(contentViewController appkit.IViewController) FilterBrowserPanel {
	return FilterBrowserPanelClass.WindowWithContentViewController(contentViewController)
}

func (f_ FilterBrowserPanel) InitWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style appkit.WindowStyleMask, backingStoreType appkit.BackingStoreType, flag bool, screen appkit.IScreen) FilterBrowserPanel {
	rv := objc.Call[FilterBrowserPanel](f_, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, objc.Ptr(screen))
	return rv
}

// Initializes an allocated window with the specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/1419755-initwithcontentrect?language=objc
func NewFilterBrowserPanelWithContentRectStyleMaskBackingDeferScreen(contentRect foundation.Rect, style appkit.WindowStyleMask, backingStoreType appkit.BackingStoreType, flag bool, screen appkit.IScreen) FilterBrowserPanel {
	instance := FilterBrowserPanelClass.Alloc().InitWithContentRectStyleMaskBackingDeferScreen(contentRect, style, backingStoreType, flag, screen)
	instance.Autorelease()
	return instance
}

// Displays the filter browser in a new utility window, unless the filter browser is already open. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel/1504894-beginwithoptions?language=objc
func (f_ FilterBrowserPanel) BeginWithOptionsModelessDelegateDidEndSelectorContextInfo(inOptions foundation.Dictionary, modelessDelegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](f_, objc.Sel("beginWithOptions:modelessDelegate:didEndSelector:contextInfo:"), inOptions, modelessDelegate, didEndSelector, contextInfo)
}

// Returns a view that contains a filter browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel/1503992-filterbrowserviewwithoptions?language=objc
func (f_ FilterBrowserPanel) FilterBrowserViewWithOptions(inOptions foundation.Dictionary) FilterBrowserView {
	rv := objc.Call[FilterBrowserView](f_, objc.Sel("filterBrowserViewWithOptions:"), inOptions)
	return rv
}

// Creates a shared instance of the IKFilterBrowserPanel class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel/1503504-filterbrowserpanelwithstylemask?language=objc
func (fc _FilterBrowserPanelClass) FilterBrowserPanelWithStyleMask(styleMask int) objc.Object {
	rv := objc.Call[objc.Object](fc, objc.Sel("filterBrowserPanelWithStyleMask:"), styleMask)
	return rv
}

// Creates a shared instance of the IKFilterBrowserPanel class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel/1503504-filterbrowserpanelwithstylemask?language=objc
func FilterBrowserPanel_FilterBrowserPanelWithStyleMask(styleMask int) objc.Object {
	return FilterBrowserPanelClass.FilterBrowserPanelWithStyleMask(styleMask)
}

// Returns the name of the  filter that is currently selected in the filter browser. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel/1504426-filtername?language=objc
func (f_ FilterBrowserPanel) FilterName() string {
	rv := objc.Call[string](f_, objc.Sel("filterName"))
	return rv
}

// Displays the filter browser in a sheet—that is, a dialog that is attached to its parent window and must be dismissed by the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel/1504636-beginsheetwithoptions?language=objc
func (f_ FilterBrowserPanel) BeginSheetWithOptionsModalForWindowModalDelegateDidEndSelectorContextInfo(inOptions foundation.Dictionary, docWindow appkit.IWindow, modalDelegate objc.IObject, didEndSelector objc.Selector, contextInfo unsafe.Pointer) {
	objc.Call[objc.Void](f_, objc.Sel("beginSheetWithOptions:modalForWindow:modalDelegate:didEndSelector:contextInfo:"), inOptions, objc.Ptr(docWindow), modalDelegate, didEndSelector, contextInfo)
}

// Closes a filter browser view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel/1505223-finish?language=objc
func (f_ FilterBrowserPanel) Finish(sender objc.IObject) {
	objc.Call[objc.Void](f_, objc.Sel("finish:"), sender)
}

// Displays the filter browser in a modal dialog that must be dismissed by the user but that is not  attached to a window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/ikfilterbrowserpanel/1504028-runmodalwithoptions?language=objc
func (f_ FilterBrowserPanel) RunModalWithOptions(inOptions foundation.Dictionary) int {
	rv := objc.Call[int](f_, objc.Sel("runModalWithOptions:"), inOptions)
	return rv
}
