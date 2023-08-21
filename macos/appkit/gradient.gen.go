// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Gradient] class.
var GradientClass = _GradientClass{objc.GetClass("NSGradient")}

type _GradientClass struct {
	objc.Class
}

// An interface definition for the [Gradient] class.
type IGradient interface {
	objc.IObject
	GetColorLocationAtIndex(color IColor, location *float64, index int)
	DrawFromCenterRadiusToCenterRadiusOptions(startCenter foundation.Point, startRadius float64, endCenter foundation.Point, endRadius float64, options GradientDrawingOptions)
	DrawInRectRelativeCenterPosition(rect foundation.Rect, relativeCenterPosition foundation.Point)
	InterpolatedColorAtLocation(location float64) Color
	DrawFromPointToPointOptions(startingPoint foundation.Point, endingPoint foundation.Point, options GradientDrawingOptions)
	DrawInBezierPathAngle(path IBezierPath, angle float64)
	ColorSpace() ColorSpace
	NumberOfColorStops() int
}

// An object that can draw gradient fill colors [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient?language=objc
type Gradient struct {
	objc.Object
}

func GradientFrom(ptr unsafe.Pointer) Gradient {
	return Gradient{
		Object: objc.ObjectFrom(ptr),
	}
}

func (g_ Gradient) InitWithColorsAndLocations(firstColor IColor, args ...any) Gradient {
	rv := objc.Call[Gradient](g_, objc.Sel("initWithColorsAndLocations:"), append([]any{objc.Ptr(firstColor)}, args...)...)
	return rv
}

// Initializes a newly allocated gradient object with a comma-separated list of arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1555387-initwithcolorsandlocations?language=objc
func NewGradientWithColorsAndLocations(firstColor IColor, args ...any) Gradient {
	instance := GradientClass.Alloc().InitWithColorsAndLocations(firstColor, args...)
	instance.Autorelease()
	return instance
}

func (g_ Gradient) InitWithStartingColorEndingColor(startingColor IColor, endingColor IColor) Gradient {
	rv := objc.Call[Gradient](g_, objc.Sel("initWithStartingColor:endingColor:"), objc.Ptr(startingColor), objc.Ptr(endingColor))
	return rv
}

// Initializes a newly allocated gradient object with two colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1525448-initwithstartingcolor?language=objc
func NewGradientWithStartingColorEndingColor(startingColor IColor, endingColor IColor) Gradient {
	instance := GradientClass.Alloc().InitWithStartingColorEndingColor(startingColor, endingColor)
	instance.Autorelease()
	return instance
}

func (g_ Gradient) InitWithColors(colorArray []IColor) Gradient {
	rv := objc.Call[Gradient](g_, objc.Sel("initWithColors:"), colorArray)
	return rv
}

// Initializes a newly allocated gradient object with an array of colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1535315-initwithcolors?language=objc
func NewGradientWithColors(colorArray []IColor) Gradient {
	instance := GradientClass.Alloc().InitWithColors(colorArray)
	instance.Autorelease()
	return instance
}

func (gc _GradientClass) Alloc() Gradient {
	rv := objc.Call[Gradient](gc, objc.Sel("alloc"))
	return rv
}

func Gradient_Alloc() Gradient {
	return GradientClass.Alloc()
}

func (gc _GradientClass) New() Gradient {
	rv := objc.Call[Gradient](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGradient() Gradient {
	return GradientClass.New()
}

func (g_ Gradient) Init() Gradient {
	rv := objc.Call[Gradient](g_, objc.Sel("init"))
	return rv
}

// Returns information about the color stop at the specified index in the receiver’s color array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1533505-getcolor?language=objc
func (g_ Gradient) GetColorLocationAtIndex(color IColor, location *float64, index int) {
	objc.Call[objc.Void](g_, objc.Sel("getColor:location:atIndex:"), objc.Ptr(color), location, index)
}

// Draws a radial gradient between the specified circles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1530677-drawfromcenter?language=objc
func (g_ Gradient) DrawFromCenterRadiusToCenterRadiusOptions(startCenter foundation.Point, startRadius float64, endCenter foundation.Point, endRadius float64, options GradientDrawingOptions) {
	objc.Call[objc.Void](g_, objc.Sel("drawFromCenter:radius:toCenter:radius:options:"), startCenter, startRadius, endCenter, endRadius, options)
}

// Draws a radial gradient starting at the center of the specified rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1533703-drawinrect?language=objc
func (g_ Gradient) DrawInRectRelativeCenterPosition(rect foundation.Rect, relativeCenterPosition foundation.Point) {
	objc.Call[objc.Void](g_, objc.Sel("drawInRect:relativeCenterPosition:"), rect, relativeCenterPosition)
}

// Returns the color of the rendered gradient at the specified relative location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1526409-interpolatedcoloratlocation?language=objc
func (g_ Gradient) InterpolatedColorAtLocation(location float64) Color {
	rv := objc.Call[Color](g_, objc.Sel("interpolatedColorAtLocation:"), location)
	return rv
}

// Draws a linear gradient between the specified start and end points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1532511-drawfrompoint?language=objc
func (g_ Gradient) DrawFromPointToPointOptions(startingPoint foundation.Point, endingPoint foundation.Point, options GradientDrawingOptions) {
	objc.Call[objc.Void](g_, objc.Sel("drawFromPoint:toPoint:options:"), startingPoint, endingPoint, options)
}

// Fills the specified path with a linear gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1534785-drawinbezierpath?language=objc
func (g_ Gradient) DrawInBezierPathAngle(path IBezierPath, angle float64) {
	objc.Call[objc.Void](g_, objc.Sel("drawInBezierPath:angle:"), objc.Ptr(path), angle)
}

// The color space of the colors associated with the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1531310-colorspace?language=objc
func (g_ Gradient) ColorSpace() ColorSpace {
	rv := objc.Call[ColorSpace](g_, objc.Sel("colorSpace"))
	return rv
}

// The number of color stops associated with the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgradient/1535846-numberofcolorstops?language=objc
func (g_ Gradient) NumberOfColorStops() int {
	rv := objc.Call[int](g_, objc.Sel("numberOfColorStops"))
	return rv
}
