// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PrintInfo] class.
var PrintInfoClass = _PrintInfoClass{objc.GetClass("NSPrintInfo")}

type _PrintInfoClass struct {
	objc.Class
}

// An interface definition for the [PrintInfo] class.
type IPrintInfo interface {
	objc.IObject
	UpdateFromPMPageFormat()
	PMPageFormat() unsafe.Pointer
	Dictionary() foundation.MutableDictionary
	PMPrintSession() unsafe.Pointer
	TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo)
	SetUpPrintOperationDefaultValues()
	PMPrintSettings() unsafe.Pointer
	UpdateFromPMPrintSettings()
	RightMargin() float64
	SetRightMargin(value float64)
	BottomMargin() float64
	SetBottomMargin(value float64)
	JobDisposition() PrintJobDispositionValue
	SetJobDisposition(value PrintJobDispositionValue)
	Printer() Printer
	SetPrinter(value IPrinter)
	IsHorizontallyCentered() bool
	SetHorizontallyCentered(value bool)
	PaperSize() foundation.Size
	SetPaperSize(value foundation.Size)
	VerticalPagination() PrintingPaginationMode
	SetVerticalPagination(value PrintingPaginationMode)
	LeftMargin() float64
	SetLeftMargin(value float64)
	TopMargin() float64
	SetTopMargin(value float64)
	ImageablePageBounds() foundation.Rect
	ScalingFactor() float64
	SetScalingFactor(value float64)
	LocalizedPaperName() string
	HorizontalPagination() PrintingPaginationMode
	SetHorizontalPagination(value PrintingPaginationMode)
	PrintSettings() foundation.MutableDictionary
	IsSelectionOnly() bool
	SetSelectionOnly(value bool)
	IsVerticallyCentered() bool
	SetVerticallyCentered(value bool)
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperName() PrinterPaperName
	SetPaperName(value PrinterPaperName)
}

// An object that stores information that’s used to generate printed output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo?language=objc
type PrintInfo struct {
	objc.Object
}

func PrintInfoFrom(ptr unsafe.Pointer) PrintInfo {
	return PrintInfo{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PrintInfo) InitWithDictionary(attributes map[PrintInfoAttributeKey]objc.IObject) PrintInfo {
	rv := objc.Call[PrintInfo](p_, objc.Sel("initWithDictionary:"), attributes)
	return rv
}

// Returns a printing information object initialized with the parameters in the specified dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1526768-initwithdictionary?language=objc
func NewPrintInfoWithDictionary(attributes map[PrintInfoAttributeKey]objc.IObject) PrintInfo {
	instance := PrintInfoClass.Alloc().InitWithDictionary(attributes)
	instance.Autorelease()
	return instance
}

func (p_ PrintInfo) Init() PrintInfo {
	rv := objc.Call[PrintInfo](p_, objc.Sel("init"))
	return rv
}

func (pc _PrintInfoClass) Alloc() PrintInfo {
	rv := objc.Call[PrintInfo](pc, objc.Sel("alloc"))
	return rv
}

func PrintInfo_Alloc() PrintInfo {
	return PrintInfoClass.Alloc()
}

func (pc _PrintInfoClass) New() PrintInfo {
	rv := objc.Call[PrintInfo](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPrintInfo() PrintInfo {
	return PrintInfoClass.New()
}

// Synchronizes the print info’s page format information with information from its associated page format object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1534429-updatefrompmpageformat?language=objc
func (p_ PrintInfo) UpdateFromPMPageFormat() {
	objc.Call[objc.Void](p_, objc.Sel("updateFromPMPageFormat"))
}

// Returns a Core Printing object configured with the print info’s page format information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1528578-pmpageformat?language=objc
func (p_ PrintInfo) PMPageFormat() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](p_, objc.Sel("PMPageFormat"))
	return rv
}

// Returns the print info’s dictionary that contains the printing attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1524990-dictionary?language=objc
func (p_ PrintInfo) Dictionary() foundation.MutableDictionary {
	rv := objc.Call[foundation.MutableDictionary](p_, objc.Sel("dictionary"))
	return rv
}

// Returns a Core Printing object configured with the print info’s session information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1534414-pmprintsession?language=objc
func (p_ PrintInfo) PMPrintSession() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](p_, objc.Sel("PMPrintSession"))
	return rv
}

// Updates the print info with all the settings and attributes in the specified PDF info object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1530099-takesettingsfrompdfinfo?language=objc
func (p_ PrintInfo) TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo) {
	objc.Call[objc.Void](p_, objc.Sel("takeSettingsFromPDFInfo:"), objc.Ptr(inPDFInfo))
}

// Validates the attributes encapsulated by the print info. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1529453-setupprintoperationdefaultvalues?language=objc
func (p_ PrintInfo) SetUpPrintOperationDefaultValues() {
	objc.Call[objc.Void](p_, objc.Sel("setUpPrintOperationDefaultValues"))
}

// Returns a Core Printing object configured with the print info’s print settings information [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1533537-pmprintsettings?language=objc
func (p_ PrintInfo) PMPrintSettings() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](p_, objc.Sel("PMPrintSettings"))
	return rv
}

// Synchronizes the print info’s print settings information with information from its associated print settings object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1525003-updatefrompmprintsettings?language=objc
func (p_ PrintInfo) UpdateFromPMPrintSettings() {
	objc.Call[objc.Void](p_, objc.Sel("updateFromPMPrintSettings"))
}

// The width of the right margin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1530882-rightmargin?language=objc
func (p_ PrintInfo) RightMargin() float64 {
	rv := objc.Call[float64](p_, objc.Sel("rightMargin"))
	return rv
}

// The width of the right margin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1530882-rightmargin?language=objc
func (p_ PrintInfo) SetRightMargin(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setRightMargin:"), value)
}

// The height of the bottom margin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1528397-bottommargin?language=objc
func (p_ PrintInfo) BottomMargin() float64 {
	rv := objc.Call[float64](p_, objc.Sel("bottomMargin"))
	return rv
}

// The height of the bottom margin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1528397-bottommargin?language=objc
func (p_ PrintInfo) SetBottomMargin(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setBottomMargin:"), value)
}

// Deprecated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1530521-defaultprinter?language=objc
func (pc _PrintInfoClass) DefaultPrinter() Printer {
	rv := objc.Call[Printer](pc, objc.Sel("defaultPrinter"))
	return rv
}

// Deprecated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1530521-defaultprinter?language=objc
func PrintInfo_DefaultPrinter() Printer {
	return PrintInfoClass.DefaultPrinter()
}

// The action specified for the job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1528717-jobdisposition?language=objc
func (p_ PrintInfo) JobDisposition() PrintJobDispositionValue {
	rv := objc.Call[PrintJobDispositionValue](p_, objc.Sel("jobDisposition"))
	return rv
}

// The action specified for the job. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1528717-jobdisposition?language=objc
func (p_ PrintInfo) SetJobDisposition(value PrintJobDispositionValue) {
	objc.Call[objc.Void](p_, objc.Sel("setJobDisposition:"), value)
}

// The printer object to be used for printing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1524495-printer?language=objc
func (p_ PrintInfo) Printer() Printer {
	rv := objc.Call[Printer](p_, objc.Sel("printer"))
	return rv
}

// The printer object to be used for printing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1524495-printer?language=objc
func (p_ PrintInfo) SetPrinter(value IPrinter) {
	objc.Call[objc.Void](p_, objc.Sel("setPrinter:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the image is centered horizontally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1534703-horizontallycentered?language=objc
func (p_ PrintInfo) IsHorizontallyCentered() bool {
	rv := objc.Call[bool](p_, objc.Sel("isHorizontallyCentered"))
	return rv
}

// A Boolean value that indicates whether the image is centered horizontally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1534703-horizontallycentered?language=objc
func (p_ PrintInfo) SetHorizontallyCentered(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setHorizontallyCentered:"), value)
}

// The size of the paper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1534030-papersize?language=objc
func (p_ PrintInfo) PaperSize() foundation.Size {
	rv := objc.Call[foundation.Size](p_, objc.Sel("paperSize"))
	return rv
}

// The size of the paper. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1534030-papersize?language=objc
func (p_ PrintInfo) SetPaperSize(value foundation.Size) {
	objc.Call[objc.Void](p_, objc.Sel("setPaperSize:"), value)
}

// The vertical pagination to the specified mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1526743-verticalpagination?language=objc
func (p_ PrintInfo) VerticalPagination() PrintingPaginationMode {
	rv := objc.Call[PrintingPaginationMode](p_, objc.Sel("verticalPagination"))
	return rv
}

// The vertical pagination to the specified mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1526743-verticalpagination?language=objc
func (p_ PrintInfo) SetVerticalPagination(value PrintingPaginationMode) {
	objc.Call[objc.Void](p_, objc.Sel("setVerticalPagination:"), value)
}

// The width of the left margin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1533430-leftmargin?language=objc
func (p_ PrintInfo) LeftMargin() float64 {
	rv := objc.Call[float64](p_, objc.Sel("leftMargin"))
	return rv
}

// The width of the left margin. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1533430-leftmargin?language=objc
func (p_ PrintInfo) SetLeftMargin(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setLeftMargin:"), value)
}

// The top margin to the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1529662-topmargin?language=objc
func (p_ PrintInfo) TopMargin() float64 {
	rv := objc.Call[float64](p_, objc.Sel("topMargin"))
	return rv
}

// The top margin to the specified size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1529662-topmargin?language=objc
func (p_ PrintInfo) SetTopMargin(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setTopMargin:"), value)
}

// The imageable area of a sheet of paper specified by the print info. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1526570-imageablepagebounds?language=objc
func (p_ PrintInfo) ImageablePageBounds() foundation.Rect {
	rv := objc.Call[foundation.Rect](p_, objc.Sel("imageablePageBounds"))
	return rv
}

// The shared printing information object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1535610-sharedprintinfo?language=objc
func (pc _PrintInfoClass) SharedPrintInfo() PrintInfo {
	rv := objc.Call[PrintInfo](pc, objc.Sel("sharedPrintInfo"))
	return rv
}

// The shared printing information object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1535610-sharedprintinfo?language=objc
func PrintInfo_SharedPrintInfo() PrintInfo {
	return PrintInfoClass.SharedPrintInfo()
}

// The shared printing information object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1535610-sharedprintinfo?language=objc
func (pc _PrintInfoClass) SetSharedPrintInfo(value IPrintInfo) {
	objc.Call[objc.Void](pc, objc.Sel("setSharedPrintInfo:"), objc.Ptr(value))
}

// The shared printing information object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1535610-sharedprintinfo?language=objc
func PrintInfo_SetSharedPrintInfo(value IPrintInfo) {
	PrintInfoClass.SetSharedPrintInfo(value)
}

// The current scaling factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1529753-scalingfactor?language=objc
func (p_ PrintInfo) ScalingFactor() float64 {
	rv := objc.Call[float64](p_, objc.Sel("scalingFactor"))
	return rv
}

// The current scaling factor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1529753-scalingfactor?language=objc
func (p_ PrintInfo) SetScalingFactor(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setScalingFactor:"), value)
}

// The human-readable name of the currently selected paper size, suitable for presentation in user interfaces. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1524573-localizedpapername?language=objc
func (p_ PrintInfo) LocalizedPaperName() string {
	rv := objc.Call[string](p_, objc.Sel("localizedPaperName"))
	return rv
}

// The horizontal pagination mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1532693-horizontalpagination?language=objc
func (p_ PrintInfo) HorizontalPagination() PrintingPaginationMode {
	rv := objc.Call[PrintingPaginationMode](p_, objc.Sel("horizontalPagination"))
	return rv
}

// The horizontal pagination mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1532693-horizontalpagination?language=objc
func (p_ PrintInfo) SetHorizontalPagination(value PrintingPaginationMode) {
	objc.Call[objc.Void](p_, objc.Sel("setHorizontalPagination:"), value)
}

// A mutable dictionary containing the print settings from Core Printing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1529413-printsettings?language=objc
func (p_ PrintInfo) PrintSettings() foundation.MutableDictionary {
	rv := objc.Call[foundation.MutableDictionary](p_, objc.Sel("printSettings"))
	return rv
}

// A Boolean value that indicates whether only the currently selected contents should be printed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1534094-selectiononly?language=objc
func (p_ PrintInfo) IsSelectionOnly() bool {
	rv := objc.Call[bool](p_, objc.Sel("isSelectionOnly"))
	return rv
}

// A Boolean value that indicates whether only the currently selected contents should be printed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1534094-selectiononly?language=objc
func (p_ PrintInfo) SetSelectionOnly(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setSelectionOnly:"), value)
}

// A Boolean value that indicates whether the image is centered vertically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1530330-verticallycentered?language=objc
func (p_ PrintInfo) IsVerticallyCentered() bool {
	rv := objc.Call[bool](p_, objc.Sel("isVerticallyCentered"))
	return rv
}

// A Boolean value that indicates whether the image is centered vertically. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1530330-verticallycentered?language=objc
func (p_ PrintInfo) SetVerticallyCentered(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setVerticallyCentered:"), value)
}

// The orientation attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1533755-orientation?language=objc
func (p_ PrintInfo) Orientation() PaperOrientation {
	rv := objc.Call[PaperOrientation](p_, objc.Sel("orientation"))
	return rv
}

// The orientation attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1533755-orientation?language=objc
func (p_ PrintInfo) SetOrientation(value PaperOrientation) {
	objc.Call[objc.Void](p_, objc.Sel("setOrientation:"), value)
}

// The name of the currently selected paper size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1532124-papername?language=objc
func (p_ PrintInfo) PaperName() PrinterPaperName {
	rv := objc.Call[PrinterPaperName](p_, objc.Sel("paperName"))
	return rv
}

// The name of the currently selected paper size. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/1532124-papername?language=objc
func (p_ PrintInfo) SetPaperName(value PrinterPaperName) {
	objc.Call[objc.Void](p_, objc.Sel("setPaperName:"), value)
}
