// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionNode] class.
var CNNConvolutionNodeClass = _CNNConvolutionNodeClass{objc.GetClass("MPSCNNConvolutionNode")}

type _CNNConvolutionNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionNode] class.
type ICNNConvolutionNode interface {
	INNFilterNode
	ConvolutionGradientState() CNNConvolutionGradientStateNode
	TrainingStyle() NNTrainingStyle
	SetTrainingStyle(value NNTrainingStyle)
	AccumulatorPrecision() NNConvolutionAccumulatorPrecisionOption
	SetAccumulatorPrecision(value NNConvolutionAccumulatorPrecisionOption)
}

// A representation of a convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode?language=objc
type CNNConvolutionNode struct {
	NNFilterNode
}

func CNNConvolutionNodeFrom(ptr unsafe.Pointer) CNNConvolutionNode {
	return CNNConvolutionNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNConvolutionNodeClass) NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNConvolutionNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionNode](cc, objc.Sel("nodeWithSource:weights:"), objc.Ptr(sourceNode), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866436-nodewithsource?language=objc
func CNNConvolutionNode_NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNConvolutionNode {
	return CNNConvolutionNodeClass.NodeWithSourceWeights(sourceNode, weights)
}

func (c_ CNNConvolutionNode) InitWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNConvolutionNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionNode](c_, objc.Sel("initWithSource:weights:"), objc.Ptr(sourceNode), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866470-initwithsource?language=objc
func NewCNNConvolutionNodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNConvolutionNode {
	instance := CNNConvolutionNodeClass.Alloc().InitWithSourceWeights(sourceNode, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionNodeClass) Alloc() CNNConvolutionNode {
	rv := objc.Call[CNNConvolutionNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionNode_Alloc() CNNConvolutionNode {
	return CNNConvolutionNodeClass.Alloc()
}

func (cc _CNNConvolutionNodeClass) New() CNNConvolutionNode {
	rv := objc.Call[CNNConvolutionNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionNode() CNNConvolutionNode {
	return CNNConvolutionNodeClass.New()
}

func (c_ CNNConvolutionNode) Init() CNNConvolutionNode {
	rv := objc.Call[CNNConvolutionNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2942634-convolutiongradientstate?language=objc
func (c_ CNNConvolutionNode) ConvolutionGradientState() CNNConvolutionGradientStateNode {
	rv := objc.Call[CNNConvolutionGradientStateNode](c_, objc.Sel("convolutionGradientState"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/3197822-trainingstyle?language=objc
func (c_ CNNConvolutionNode) TrainingStyle() NNTrainingStyle {
	rv := objc.Call[NNTrainingStyle](c_, objc.Sel("trainingStyle"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/3197822-trainingstyle?language=objc
func (c_ CNNConvolutionNode) SetTrainingStyle(value NNTrainingStyle) {
	objc.Call[objc.Void](c_, objc.Sel("setTrainingStyle:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2980757-accumulatorprecision?language=objc
func (c_ CNNConvolutionNode) AccumulatorPrecision() NNConvolutionAccumulatorPrecisionOption {
	rv := objc.Call[NNConvolutionAccumulatorPrecisionOption](c_, objc.Sel("accumulatorPrecision"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2980757-accumulatorprecision?language=objc
func (c_ CNNConvolutionNode) SetAccumulatorPrecision(value NNConvolutionAccumulatorPrecisionOption) {
	objc.Call[objc.Void](c_, objc.Sel("setAccumulatorPrecision:"), value)
}
