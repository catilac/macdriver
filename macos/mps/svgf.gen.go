// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SVGF] class.
var SVGFClass = _SVGFClass{objc.GetClass("MPSSVGF")}

type _SVGFClass struct {
	objc.Class
}

// An interface definition for the [SVGF] class.
type ISVGF interface {
	IKernel
	EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureDepthNormalTexture(commandBuffer metal.PCommandBuffer, stepDistance uint, sourceTexture metal.PTexture, destinationTexture metal.PTexture, depthNormalTexture metal.PTexture)
	EncodeBilateralFilterToCommandBufferObjectStepDistanceSourceTextureObjectDestinationTextureObjectDepthNormalTextureObject(commandBufferObject objc.IObject, stepDistance uint, sourceTextureObject objc.IObject, destinationTextureObject objc.IObject, depthNormalTextureObject objc.IObject)
	EncodeWithCoder(coder foundation.ICoder)
	EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureFrameCountTextureDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, luminanceMomentsTexture metal.PTexture, destinationTexture metal.PTexture, frameCountTexture metal.PTexture, depthNormalTexture metal.PTexture)
	EncodeVarianceEstimationToCommandBufferObjectSourceTextureObjectLuminanceMomentsTextureObjectDestinationTextureObjectFrameCountTextureObjectDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, luminanceMomentsTextureObject objc.IObject, destinationTextureObject objc.IObject, frameCountTextureObject objc.IObject, depthNormalTextureObject objc.IObject)
	EncodeReprojectionToCommandBufferSourceTexturePreviousTextureDestinationTexturePreviousLuminanceMomentsTextureDestinationLuminanceMomentsTexturePreviousFrameCountTextureDestinationFrameCountTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, previousTexture metal.PTexture, destinationTexture metal.PTexture, previousLuminanceMomentsTexture metal.PTexture, destinationLuminanceMomentsTexture metal.PTexture, previousFrameCountTexture metal.PTexture, destinationFrameCountTexture metal.PTexture, motionVectorTexture metal.PTexture, depthNormalTexture metal.PTexture, previousDepthNormalTexture metal.PTexture)
	EncodeReprojectionToCommandBufferObjectSourceTextureObjectPreviousTextureObjectDestinationTextureObjectPreviousLuminanceMomentsTextureObjectDestinationLuminanceMomentsTextureObjectPreviousFrameCountTextureObjectDestinationFrameCountTextureObjectMotionVectorTextureObjectDepthNormalTextureObjectPreviousDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, previousTextureObject objc.IObject, destinationTextureObject objc.IObject, previousLuminanceMomentsTextureObject objc.IObject, destinationLuminanceMomentsTextureObject objc.IObject, previousFrameCountTextureObject objc.IObject, destinationFrameCountTextureObject objc.IObject, motionVectorTextureObject objc.IObject, depthNormalTextureObject objc.IObject, previousDepthNormalTextureObject objc.IObject)
	ChannelCount() uint
	SetChannelCount(value uint)
	VarianceEstimationSigma() float64
	SetVarianceEstimationSigma(value float64)
	BilateralFilterSigma() float64
	SetBilateralFilterSigma(value float64)
	MinimumFramesForVarianceEstimation() uint
	SetMinimumFramesForVarianceEstimation(value uint)
	BilateralFilterRadius() uint
	SetBilateralFilterRadius(value uint)
	ChannelCount2() uint
	SetChannelCount2(value uint)
	VariancePrefilterRadius() uint
	SetVariancePrefilterRadius(value uint)
	TemporalReprojectionBlendFactor() float64
	SetTemporalReprojectionBlendFactor(value float64)
	LuminanceWeight() float64
	SetLuminanceWeight(value float64)
	NormalWeight() float64
	SetNormalWeight(value float64)
	TemporalWeighting() TemporalWeighting
	SetTemporalWeighting(value TemporalWeighting)
	VarianceEstimationRadius() uint
	SetVarianceEstimationRadius(value uint)
	VariancePrefilterSigma() float64
	SetVariancePrefilterSigma(value float64)
	DepthWeight() float64
	SetDepthWeight(value float64)
	ReprojectionThreshold() float64
	SetReprojectionThreshold(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf?language=objc
type SVGF struct {
	Kernel
}

func SVGFFrom(ptr unsafe.Pointer) SVGF {
	return SVGF{
		Kernel: KernelFrom(ptr),
	}
}

func (s_ SVGF) InitWithDevice(device metal.PDevice) SVGF {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[SVGF](s_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143566-initwithdevice?language=objc
func NewSVGFWithDevice(device metal.PDevice) SVGF {
	instance := SVGFClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (s_ SVGF) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) SVGF {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[SVGF](s_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143559-copywithzone?language=objc
func SVGF_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) SVGF {
	instance := SVGFClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (sc _SVGFClass) Alloc() SVGF {
	rv := objc.Call[SVGF](sc, objc.Sel("alloc"))
	return rv
}

func SVGF_Alloc() SVGF {
	return SVGFClass.Alloc()
}

func (sc _SVGFClass) New() SVGF {
	rv := objc.Call[SVGF](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSVGF() SVGF {
	return SVGFClass.New()
}

func (s_ SVGF) Init() SVGF {
	rv := objc.Call[SVGF](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242891-encodebilateralfiltertocommandbu?language=objc
func (s_ SVGF) EncodeBilateralFilterToCommandBufferStepDistanceSourceTextureDestinationTextureDepthNormalTexture(commandBuffer metal.PCommandBuffer, stepDistance uint, sourceTexture metal.PTexture, destinationTexture metal.PTexture, depthNormalTexture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po2 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", destinationTexture)
	po4 := objc.WrapAsProtocol("MTLTexture", depthNormalTexture)
	objc.Call[objc.Void](s_, objc.Sel("encodeBilateralFilterToCommandBuffer:stepDistance:sourceTexture:destinationTexture:depthNormalTexture:"), po0, stepDistance, po2, po3, po4)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242891-encodebilateralfiltertocommandbu?language=objc
func (s_ SVGF) EncodeBilateralFilterToCommandBufferObjectStepDistanceSourceTextureObjectDestinationTextureObjectDepthNormalTextureObject(commandBufferObject objc.IObject, stepDistance uint, sourceTextureObject objc.IObject, destinationTextureObject objc.IObject, depthNormalTextureObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("encodeBilateralFilterToCommandBuffer:stepDistance:sourceTexture:destinationTexture:depthNormalTexture:"), objc.Ptr(commandBufferObject), stepDistance, objc.Ptr(sourceTextureObject), objc.Ptr(destinationTextureObject), objc.Ptr(depthNormalTextureObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143564-encodewithcoder?language=objc
func (s_ SVGF) EncodeWithCoder(coder foundation.ICoder) {
	objc.Call[objc.Void](s_, objc.Sel("encodeWithCoder:"), objc.Ptr(coder))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242893-encodevarianceestimationtocomman?language=objc
func (s_ SVGF) EncodeVarianceEstimationToCommandBufferSourceTextureLuminanceMomentsTextureDestinationTextureFrameCountTextureDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, luminanceMomentsTexture metal.PTexture, destinationTexture metal.PTexture, frameCountTexture metal.PTexture, depthNormalTexture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po2 := objc.WrapAsProtocol("MTLTexture", luminanceMomentsTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", destinationTexture)
	po4 := objc.WrapAsProtocol("MTLTexture", frameCountTexture)
	po5 := objc.WrapAsProtocol("MTLTexture", depthNormalTexture)
	objc.Call[objc.Void](s_, objc.Sel("encodeVarianceEstimationToCommandBuffer:sourceTexture:luminanceMomentsTexture:destinationTexture:frameCountTexture:depthNormalTexture:"), po0, po1, po2, po3, po4, po5)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242893-encodevarianceestimationtocomman?language=objc
func (s_ SVGF) EncodeVarianceEstimationToCommandBufferObjectSourceTextureObjectLuminanceMomentsTextureObjectDestinationTextureObjectFrameCountTextureObjectDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, luminanceMomentsTextureObject objc.IObject, destinationTextureObject objc.IObject, frameCountTextureObject objc.IObject, depthNormalTextureObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("encodeVarianceEstimationToCommandBuffer:sourceTexture:luminanceMomentsTexture:destinationTexture:frameCountTexture:depthNormalTexture:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceTextureObject), objc.Ptr(luminanceMomentsTextureObject), objc.Ptr(destinationTextureObject), objc.Ptr(frameCountTextureObject), objc.Ptr(depthNormalTextureObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242892-encodereprojectiontocommandbuffe?language=objc
func (s_ SVGF) EncodeReprojectionToCommandBufferSourceTexturePreviousTextureDestinationTexturePreviousLuminanceMomentsTextureDestinationLuminanceMomentsTexturePreviousFrameCountTextureDestinationFrameCountTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, previousTexture metal.PTexture, destinationTexture metal.PTexture, previousLuminanceMomentsTexture metal.PTexture, destinationLuminanceMomentsTexture metal.PTexture, previousFrameCountTexture metal.PTexture, destinationFrameCountTexture metal.PTexture, motionVectorTexture metal.PTexture, depthNormalTexture metal.PTexture, previousDepthNormalTexture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po2 := objc.WrapAsProtocol("MTLTexture", previousTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", destinationTexture)
	po4 := objc.WrapAsProtocol("MTLTexture", previousLuminanceMomentsTexture)
	po5 := objc.WrapAsProtocol("MTLTexture", destinationLuminanceMomentsTexture)
	po6 := objc.WrapAsProtocol("MTLTexture", previousFrameCountTexture)
	po7 := objc.WrapAsProtocol("MTLTexture", destinationFrameCountTexture)
	po8 := objc.WrapAsProtocol("MTLTexture", motionVectorTexture)
	po9 := objc.WrapAsProtocol("MTLTexture", depthNormalTexture)
	po10 := objc.WrapAsProtocol("MTLTexture", previousDepthNormalTexture)
	objc.Call[objc.Void](s_, objc.Sel("encodeReprojectionToCommandBuffer:sourceTexture:previousTexture:destinationTexture:previousLuminanceMomentsTexture:destinationLuminanceMomentsTexture:previousFrameCountTexture:destinationFrameCountTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), po0, po1, po2, po3, po4, po5, po6, po7, po8, po9, po10)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3242892-encodereprojectiontocommandbuffe?language=objc
func (s_ SVGF) EncodeReprojectionToCommandBufferObjectSourceTextureObjectPreviousTextureObjectDestinationTextureObjectPreviousLuminanceMomentsTextureObjectDestinationLuminanceMomentsTextureObjectPreviousFrameCountTextureObjectDestinationFrameCountTextureObjectMotionVectorTextureObjectDepthNormalTextureObjectPreviousDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, previousTextureObject objc.IObject, destinationTextureObject objc.IObject, previousLuminanceMomentsTextureObject objc.IObject, destinationLuminanceMomentsTextureObject objc.IObject, previousFrameCountTextureObject objc.IObject, destinationFrameCountTextureObject objc.IObject, motionVectorTextureObject objc.IObject, depthNormalTextureObject objc.IObject, previousDepthNormalTextureObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("encodeReprojectionToCommandBuffer:sourceTexture:previousTexture:destinationTexture:previousLuminanceMomentsTexture:destinationLuminanceMomentsTexture:previousFrameCountTexture:destinationFrameCountTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceTextureObject), objc.Ptr(previousTextureObject), objc.Ptr(destinationTextureObject), objc.Ptr(previousLuminanceMomentsTextureObject), objc.Ptr(destinationLuminanceMomentsTextureObject), objc.Ptr(previousFrameCountTextureObject), objc.Ptr(destinationFrameCountTextureObject), objc.Ptr(motionVectorTextureObject), objc.Ptr(depthNormalTextureObject), objc.Ptr(previousDepthNormalTextureObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143557-channelcount?language=objc
func (s_ SVGF) ChannelCount() uint {
	rv := objc.Call[uint](s_, objc.Sel("channelCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143557-channelcount?language=objc
func (s_ SVGF) SetChannelCount(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setChannelCount:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143574-varianceestimationsigma?language=objc
func (s_ SVGF) VarianceEstimationSigma() float64 {
	rv := objc.Call[float64](s_, objc.Sel("varianceEstimationSigma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143574-varianceestimationsigma?language=objc
func (s_ SVGF) SetVarianceEstimationSigma(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setVarianceEstimationSigma:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143556-bilateralfiltersigma?language=objc
func (s_ SVGF) BilateralFilterSigma() float64 {
	rv := objc.Call[float64](s_, objc.Sel("bilateralFilterSigma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143556-bilateralfiltersigma?language=objc
func (s_ SVGF) SetBilateralFilterSigma(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setBilateralFilterSigma:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143568-minimumframesforvarianceestimati?language=objc
func (s_ SVGF) MinimumFramesForVarianceEstimation() uint {
	rv := objc.Call[uint](s_, objc.Sel("minimumFramesForVarianceEstimation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143568-minimumframesforvarianceestimati?language=objc
func (s_ SVGF) SetMinimumFramesForVarianceEstimation(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setMinimumFramesForVarianceEstimation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143555-bilateralfilterradius?language=objc
func (s_ SVGF) BilateralFilterRadius() uint {
	rv := objc.Call[uint](s_, objc.Sel("bilateralFilterRadius"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143555-bilateralfilterradius?language=objc
func (s_ SVGF) SetBilateralFilterRadius(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setBilateralFilterRadius:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143558-channelcount2?language=objc
func (s_ SVGF) ChannelCount2() uint {
	rv := objc.Call[uint](s_, objc.Sel("channelCount2"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143558-channelcount2?language=objc
func (s_ SVGF) SetChannelCount2(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setChannelCount2:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143575-varianceprefilterradius?language=objc
func (s_ SVGF) VariancePrefilterRadius() uint {
	rv := objc.Call[uint](s_, objc.Sel("variancePrefilterRadius"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143575-varianceprefilterradius?language=objc
func (s_ SVGF) SetVariancePrefilterRadius(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setVariancePrefilterRadius:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143571-temporalreprojectionblendfactor?language=objc
func (s_ SVGF) TemporalReprojectionBlendFactor() float64 {
	rv := objc.Call[float64](s_, objc.Sel("temporalReprojectionBlendFactor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143571-temporalreprojectionblendfactor?language=objc
func (s_ SVGF) SetTemporalReprojectionBlendFactor(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setTemporalReprojectionBlendFactor:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143567-luminanceweight?language=objc
func (s_ SVGF) LuminanceWeight() float64 {
	rv := objc.Call[float64](s_, objc.Sel("luminanceWeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143567-luminanceweight?language=objc
func (s_ SVGF) SetLuminanceWeight(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setLuminanceWeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143569-normalweight?language=objc
func (s_ SVGF) NormalWeight() float64 {
	rv := objc.Call[float64](s_, objc.Sel("normalWeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143569-normalweight?language=objc
func (s_ SVGF) SetNormalWeight(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setNormalWeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143572-temporalweighting?language=objc
func (s_ SVGF) TemporalWeighting() TemporalWeighting {
	rv := objc.Call[TemporalWeighting](s_, objc.Sel("temporalWeighting"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143572-temporalweighting?language=objc
func (s_ SVGF) SetTemporalWeighting(value TemporalWeighting) {
	objc.Call[objc.Void](s_, objc.Sel("setTemporalWeighting:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143573-varianceestimationradius?language=objc
func (s_ SVGF) VarianceEstimationRadius() uint {
	rv := objc.Call[uint](s_, objc.Sel("varianceEstimationRadius"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143573-varianceestimationradius?language=objc
func (s_ SVGF) SetVarianceEstimationRadius(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setVarianceEstimationRadius:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143576-varianceprefiltersigma?language=objc
func (s_ SVGF) VariancePrefilterSigma() float64 {
	rv := objc.Call[float64](s_, objc.Sel("variancePrefilterSigma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143576-varianceprefiltersigma?language=objc
func (s_ SVGF) SetVariancePrefilterSigma(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setVariancePrefilterSigma:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143560-depthweight?language=objc
func (s_ SVGF) DepthWeight() float64 {
	rv := objc.Call[float64](s_, objc.Sel("depthWeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143560-depthweight?language=objc
func (s_ SVGF) SetDepthWeight(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setDepthWeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143570-reprojectionthreshold?language=objc
func (s_ SVGF) ReprojectionThreshold() float64 {
	rv := objc.Call[float64](s_, objc.Sel("reprojectionThreshold"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgf/3143570-reprojectionthreshold?language=objc
func (s_ SVGF) SetReprojectionThreshold(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setReprojectionThreshold:"), value)
}
