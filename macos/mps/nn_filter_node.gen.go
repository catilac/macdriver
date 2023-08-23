// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNFilterNode] class.
var NNFilterNodeClass = _NNFilterNodeClass{objc.GetClass("MPSNNFilterNode")}

type _NNFilterNodeClass struct {
	objc.Class
}

// An interface definition for the [NNFilterNode] class.
type INNFilterNode interface {
	objc.IObject
	GradientFiltersWithSources(gradientImages []INNImageNode) []NNGradientFilterNode
	GradientFiltersWithSource(gradientImage INNImageNode) []NNGradientFilterNode
	GradientFilterWithSource(gradientImage INNImageNode) NNGradientFilterNode
	TrainingGraphWithSourceGradientNodeHandler(gradientImage INNImageNode, nodeHandler GradientNodeBlock) []NNFilterNode
	GradientFilterWithSources(gradientImages []INNImageNode) NNGradientFilterNode
	ResultState() NNStateNode
	ResultImage() NNImageNode
	ResultStates() []NNStateNode
	PaddingPolicy() NNPaddingWrapper
	SetPaddingPolicy(value PNNPadding)
	SetPaddingPolicyObject(valueObject objc.IObject)
	Label() string
	SetLabel(value string)
}

// A placeholder node denoting a neural network filter stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode?language=objc
type NNFilterNode struct {
	objc.Object
}

func NNFilterNodeFrom(ptr unsafe.Pointer) NNFilterNode {
	return NNFilterNode{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NNFilterNodeClass) Alloc() NNFilterNode {
	rv := objc.Call[NNFilterNode](nc, objc.Sel("alloc"))
	return rv
}

func NNFilterNode_Alloc() NNFilterNode {
	return NNFilterNodeClass.Alloc()
}

func (nc _NNFilterNodeClass) New() NNFilterNode {
	rv := objc.Call[NNFilterNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNFilterNode() NNFilterNode {
	return NNFilterNodeClass.New()
}

func (n_ NNFilterNode) Init() NNFilterNode {
	rv := objc.Call[NNFilterNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2951955-gradientfilterswithsources?language=objc
func (n_ NNFilterNode) GradientFiltersWithSources(gradientImages []INNImageNode) []NNGradientFilterNode {
	rv := objc.Call[[]NNGradientFilterNode](n_, objc.Sel("gradientFiltersWithSources:"), gradientImages)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2953944-gradientfilterswithsource?language=objc
func (n_ NNFilterNode) GradientFiltersWithSource(gradientImage INNImageNode) []NNGradientFilterNode {
	rv := objc.Call[[]NNGradientFilterNode](n_, objc.Sel("gradientFiltersWithSource:"), objc.Ptr(gradientImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2953941-gradientfilterwithsource?language=objc
func (n_ NNFilterNode) GradientFilterWithSource(gradientImage INNImageNode) NNGradientFilterNode {
	rv := objc.Call[NNGradientFilterNode](n_, objc.Sel("gradientFilterWithSource:"), objc.Ptr(gradientImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/3020688-traininggraphwithsourcegradient?language=objc
func (n_ NNFilterNode) TrainingGraphWithSourceGradientNodeHandler(gradientImage INNImageNode, nodeHandler GradientNodeBlock) []NNFilterNode {
	rv := objc.Call[[]NNFilterNode](n_, objc.Sel("trainingGraphWithSourceGradient:nodeHandler:"), objc.Ptr(gradientImage), nodeHandler)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2948052-gradientfilterwithsources?language=objc
func (n_ NNFilterNode) GradientFilterWithSources(gradientImages []INNImageNode) NNGradientFilterNode {
	rv := objc.Call[NNGradientFilterNode](n_, objc.Sel("gradientFilterWithSources:"), gradientImages)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866503-resultstate?language=objc
func (n_ NNFilterNode) ResultState() NNStateNode {
	rv := objc.Call[NNStateNode](n_, objc.Sel("resultState"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866492-resultimage?language=objc
func (n_ NNFilterNode) ResultImage() NNImageNode {
	rv := objc.Call[NNImageNode](n_, objc.Sel("resultImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866486-resultstates?language=objc
func (n_ NNFilterNode) ResultStates() []NNStateNode {
	rv := objc.Call[[]NNStateNode](n_, objc.Sel("resultStates"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866496-paddingpolicy?language=objc
func (n_ NNFilterNode) PaddingPolicy() NNPaddingWrapper {
	rv := objc.Call[NNPaddingWrapper](n_, objc.Sel("paddingPolicy"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866496-paddingpolicy?language=objc
func (n_ NNFilterNode) SetPaddingPolicy(value PNNPadding) {
	po0 := objc.WrapAsProtocol("MPSNNPadding", value)
	objc.Call[objc.Void](n_, objc.Sel("setPaddingPolicy:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866496-paddingpolicy?language=objc
func (n_ NNFilterNode) SetPaddingPolicyObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setPaddingPolicy:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866465-label?language=objc
func (n_ NNFilterNode) Label() string {
	rv := objc.Call[string](n_, objc.Sel("label"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnfilternode/2866465-label?language=objc
func (n_ NNFilterNode) SetLabel(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setLabel:"), value)
}
