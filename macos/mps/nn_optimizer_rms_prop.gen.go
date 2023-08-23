// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNOptimizerRMSProp] class.
var NNOptimizerRMSPropClass = _NNOptimizerRMSPropClass{objc.GetClass("MPSNNOptimizerRMSProp")}

type _NNOptimizerRMSPropClass struct {
	objc.Class
}

// An interface definition for the [NNOptimizerRMSProp] class.
type INNOptimizerRMSProp interface {
	INNOptimizer
	EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix)
	EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix)
	Decay() float64
	Epsilon() float64
}

// An optimization layer that performs a root mean square propagation update. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop?language=objc
type NNOptimizerRMSProp struct {
	NNOptimizer
}

func NNOptimizerRMSPropFrom(ptr unsafe.Pointer) NNOptimizerRMSProp {
	return NNOptimizerRMSProp{
		NNOptimizer: NNOptimizerFrom(ptr),
	}
}

func (n_ NNOptimizerRMSProp) InitWithDeviceLearningRate(device metal.PDevice, learningRate float64) NNOptimizerRMSProp {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerRMSProp](n_, objc.Sel("initWithDevice:learningRate:"), po0, learningRate)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966738-initwithdevice?language=objc
func NewNNOptimizerRMSPropWithDeviceLearningRate(device metal.PDevice, learningRate float64) NNOptimizerRMSProp {
	instance := NNOptimizerRMSPropClass.Alloc().InitWithDeviceLearningRate(device, learningRate)
	instance.Autorelease()
	return instance
}

func (nc _NNOptimizerRMSPropClass) Alloc() NNOptimizerRMSProp {
	rv := objc.Call[NNOptimizerRMSProp](nc, objc.Sel("alloc"))
	return rv
}

func NNOptimizerRMSProp_Alloc() NNOptimizerRMSProp {
	return NNOptimizerRMSPropClass.Alloc()
}

func (nc _NNOptimizerRMSPropClass) New() NNOptimizerRMSProp {
	rv := objc.Call[NNOptimizerRMSProp](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNOptimizerRMSProp() NNOptimizerRMSProp {
	return NNOptimizerRMSPropClass.New()
}

func (n_ NNOptimizerRMSProp) Init() NNOptimizerRMSProp {
	rv := objc.Call[NNOptimizerRMSProp](n_, objc.Sel("init"))
	return rv
}

func (n_ NNOptimizerRMSProp) InitWithDevice(device metal.PDevice) NNOptimizerRMSProp {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerRMSProp](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNNOptimizerRMSPropWithDevice(device metal.PDevice) NNOptimizerRMSProp {
	instance := NNOptimizerRMSPropClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizerRMSProp) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerRMSProp {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizerRMSProp](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNOptimizerRMSProp_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizerRMSProp {
	instance := NNOptimizerRMSPropClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3197827-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBuffer metal.PCommandBuffer, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:"), po0, objc.Ptr(inputGradientMatrix), objc.Ptr(inputValuesMatrix), objc.Ptr(inputSumOfSquaresMatrix), objc.Ptr(resultValuesMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/3197827-encodetocommandbuffer?language=objc
func (n_ NNOptimizerRMSProp) EncodeToCommandBufferObjectInputGradientMatrixInputValuesMatrixInputSumOfSquaresMatrixResultValuesMatrix(commandBufferObject objc.IObject, inputGradientMatrix IMatrix, inputValuesMatrix IMatrix, inputSumOfSquaresMatrix IMatrix, resultValuesMatrix IMatrix) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:inputGradientMatrix:inputValuesMatrix:inputSumOfSquaresMatrix:resultValuesMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputGradientMatrix), objc.Ptr(inputValuesMatrix), objc.Ptr(inputSumOfSquaresMatrix), objc.Ptr(resultValuesMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966734-decay?language=objc
func (n_ NNOptimizerRMSProp) Decay() float64 {
	rv := objc.Call[float64](n_, objc.Sel("decay"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerrmsprop/2966736-epsilon?language=objc
func (n_ NNOptimizerRMSProp) Epsilon() float64 {
	rv := objc.Call[float64](n_, objc.Sel("epsilon"))
	return rv
}
