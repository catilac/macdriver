// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LevelIndicator] class.
var LevelIndicatorClass = _LevelIndicatorClass{objc.GetClass("NSLevelIndicator")}

type _LevelIndicatorClass struct {
	objc.Class
}

// An interface definition for the [LevelIndicator] class.
type ILevelIndicator interface {
	IControl
	RectOfTickMarkAtIndex(index int) foundation.Rect
	TickMarkValueAtIndex(index int) float64
	DrawsTieredCapacityLevels() bool
	SetDrawsTieredCapacityLevels(value bool)
	RatingPlaceholderImage() Image
	SetRatingPlaceholderImage(value IImage)
	TickMarkPosition() TickMarkPosition
	SetTickMarkPosition(value TickMarkPosition)
	CriticalValue() float64
	SetCriticalValue(value float64)
	NumberOfMajorTickMarks() int
	SetNumberOfMajorTickMarks(value int)
	WarningValue() float64
	SetWarningValue(value float64)
	IsEditable() bool
	SetEditable(value bool)
	LevelIndicatorStyle() LevelIndicatorStyle
	SetLevelIndicatorStyle(value LevelIndicatorStyle)
	CriticalFillColor() Color
	SetCriticalFillColor(value IColor)
	WarningFillColor() Color
	SetWarningFillColor(value IColor)
	MinValue() float64
	SetMinValue(value float64)
	FillColor() Color
	SetFillColor(value IColor)
	MaxValue() float64
	SetMaxValue(value float64)
	RatingImage() Image
	SetRatingImage(value IImage)
	PlaceholderVisibility() LevelIndicatorPlaceholderVisibility
	SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility)
	NumberOfTickMarks() int
	SetNumberOfTickMarks(value int)
}

// A visual representation of a level or quantity, using discrete values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator?language=objc
type LevelIndicator struct {
	Control
}

func LevelIndicatorFrom(ptr unsafe.Pointer) LevelIndicator {
	return LevelIndicator{
		Control: ControlFrom(ptr),
	}
}

func (lc _LevelIndicatorClass) Alloc() LevelIndicator {
	rv := objc.Call[LevelIndicator](lc, objc.Sel("alloc"))
	return rv
}

func LevelIndicator_Alloc() LevelIndicator {
	return LevelIndicatorClass.Alloc()
}

func (lc _LevelIndicatorClass) New() LevelIndicator {
	rv := objc.Call[LevelIndicator](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLevelIndicator() LevelIndicator {
	return LevelIndicatorClass.New()
}

func (l_ LevelIndicator) Init() LevelIndicator {
	rv := objc.Call[LevelIndicator](l_, objc.Sel("init"))
	return rv
}

func (l_ LevelIndicator) InitWithFrame(frameRect foundation.Rect) LevelIndicator {
	rv := objc.Call[LevelIndicator](l_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

// Initializes a control with the specified frame rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/1428900-initwithframe?language=objc
func NewLevelIndicatorWithFrame(frameRect foundation.Rect) LevelIndicator {
	instance := LevelIndicatorClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Returns the bounding rectangle of the tick mark identified by the specified index (the minimum-value tick mark is at index 0). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388825-rectoftickmarkatindex?language=objc
func (l_ LevelIndicator) RectOfTickMarkAtIndex(index int) foundation.Rect {
	rv := objc.Call[foundation.Rect](l_, objc.Sel("rectOfTickMarkAtIndex:"), index)
	return rv
}

// Returns the receiver’s value represented by the tick mark at the specified index (the minimum-value tick mark has an index of 0). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388823-tickmarkvalueatindex?language=objc
func (l_ LevelIndicator) TickMarkValueAtIndex(index int) float64 {
	rv := objc.Call[float64](l_, objc.Sel("tickMarkValueAtIndex:"), index)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902322-drawstieredcapacitylevels?language=objc
func (l_ LevelIndicator) DrawsTieredCapacityLevels() bool {
	rv := objc.Call[bool](l_, objc.Sel("drawsTieredCapacityLevels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902322-drawstieredcapacitylevels?language=objc
func (l_ LevelIndicator) SetDrawsTieredCapacityLevels(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setDrawsTieredCapacityLevels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902328-ratingplaceholderimage?language=objc
func (l_ LevelIndicator) RatingPlaceholderImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("ratingPlaceholderImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902328-ratingplaceholderimage?language=objc
func (l_ LevelIndicator) SetRatingPlaceholderImage(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setRatingPlaceholderImage:"), objc.Ptr(value))
}

// Determines how the receiver’s tick marks are aligned with it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388837-tickmarkposition?language=objc
func (l_ LevelIndicator) TickMarkPosition() TickMarkPosition {
	rv := objc.Call[TickMarkPosition](l_, objc.Sel("tickMarkPosition"))
	return rv
}

// Determines how the receiver’s tick marks are aligned with it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388837-tickmarkposition?language=objc
func (l_ LevelIndicator) SetTickMarkPosition(value TickMarkPosition) {
	objc.Call[objc.Void](l_, objc.Sel("setTickMarkPosition:"), value)
}

// The receiver’s critical value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388821-criticalvalue?language=objc
func (l_ LevelIndicator) CriticalValue() float64 {
	rv := objc.Call[float64](l_, objc.Sel("criticalValue"))
	return rv
}

// The receiver’s critical value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388821-criticalvalue?language=objc
func (l_ LevelIndicator) SetCriticalValue(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setCriticalValue:"), value)
}

// The number of major tick marks associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388819-numberofmajortickmarks?language=objc
func (l_ LevelIndicator) NumberOfMajorTickMarks() int {
	rv := objc.Call[int](l_, objc.Sel("numberOfMajorTickMarks"))
	return rv
}

// The number of major tick marks associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388819-numberofmajortickmarks?language=objc
func (l_ LevelIndicator) SetNumberOfMajorTickMarks(value int) {
	objc.Call[objc.Void](l_, objc.Sel("setNumberOfMajorTickMarks:"), value)
}

// The receiver’s warning value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388835-warningvalue?language=objc
func (l_ LevelIndicator) WarningValue() float64 {
	rv := objc.Call[float64](l_, objc.Sel("warningValue"))
	return rv
}

// The receiver’s warning value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388835-warningvalue?language=objc
func (l_ LevelIndicator) SetWarningValue(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setWarningValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2919732-editable?language=objc
func (l_ LevelIndicator) IsEditable() bool {
	rv := objc.Call[bool](l_, objc.Sel("isEditable"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2919732-editable?language=objc
func (l_ LevelIndicator) SetEditable(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setEditable:"), value)
}

// The appearance of the indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388833-levelindicatorstyle?language=objc
func (l_ LevelIndicator) LevelIndicatorStyle() LevelIndicatorStyle {
	rv := objc.Call[LevelIndicatorStyle](l_, objc.Sel("levelIndicatorStyle"))
	return rv
}

// The appearance of the indicator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388833-levelindicatorstyle?language=objc
func (l_ LevelIndicator) SetLevelIndicatorStyle(value LevelIndicatorStyle) {
	objc.Call[objc.Void](l_, objc.Sel("setLevelIndicatorStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902316-criticalfillcolor?language=objc
func (l_ LevelIndicator) CriticalFillColor() Color {
	rv := objc.Call[Color](l_, objc.Sel("criticalFillColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902316-criticalfillcolor?language=objc
func (l_ LevelIndicator) SetCriticalFillColor(value IColor) {
	objc.Call[objc.Void](l_, objc.Sel("setCriticalFillColor:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902307-warningfillcolor?language=objc
func (l_ LevelIndicator) WarningFillColor() Color {
	rv := objc.Call[Color](l_, objc.Sel("warningFillColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902307-warningfillcolor?language=objc
func (l_ LevelIndicator) SetWarningFillColor(value IColor) {
	objc.Call[objc.Void](l_, objc.Sel("setWarningFillColor:"), objc.Ptr(value))
}

// The receiver’s minimum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388829-minvalue?language=objc
func (l_ LevelIndicator) MinValue() float64 {
	rv := objc.Call[float64](l_, objc.Sel("minValue"))
	return rv
}

// The receiver’s minimum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388829-minvalue?language=objc
func (l_ LevelIndicator) SetMinValue(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setMinValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902325-fillcolor?language=objc
func (l_ LevelIndicator) FillColor() Color {
	rv := objc.Call[Color](l_, objc.Sel("fillColor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902325-fillcolor?language=objc
func (l_ LevelIndicator) SetFillColor(value IColor) {
	objc.Call[objc.Void](l_, objc.Sel("setFillColor:"), objc.Ptr(value))
}

// The receiver’s maximum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388839-maxvalue?language=objc
func (l_ LevelIndicator) MaxValue() float64 {
	rv := objc.Call[float64](l_, objc.Sel("maxValue"))
	return rv
}

// The receiver’s maximum value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388839-maxvalue?language=objc
func (l_ LevelIndicator) SetMaxValue(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setMaxValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902327-ratingimage?language=objc
func (l_ LevelIndicator) RatingImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("ratingImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902327-ratingimage?language=objc
func (l_ LevelIndicator) SetRatingImage(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setRatingImage:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902323-placeholdervisibility?language=objc
func (l_ LevelIndicator) PlaceholderVisibility() LevelIndicatorPlaceholderVisibility {
	rv := objc.Call[LevelIndicatorPlaceholderVisibility](l_, objc.Sel("placeholderVisibility"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/2902323-placeholdervisibility?language=objc
func (l_ LevelIndicator) SetPlaceholderVisibility(value LevelIndicatorPlaceholderVisibility) {
	objc.Call[objc.Void](l_, objc.Sel("setPlaceholderVisibility:"), value)
}

// The number of tick marks associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388827-numberoftickmarks?language=objc
func (l_ LevelIndicator) NumberOfTickMarks() int {
	rv := objc.Call[int](l_, objc.Sel("numberOfTickMarks"))
	return rv
}

// The number of tick marks associated with the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nslevelindicator/1388827-numberoftickmarks?language=objc
func (l_ LevelIndicator) SetNumberOfTickMarks(value int) {
	objc.Call[objc.Void](l_, objc.Sel("setNumberOfTickMarks:"), value)
}
