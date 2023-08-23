// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNSubtractionNode] class.
var NNSubtractionNodeClass = _NNSubtractionNodeClass{objc.GetClass("MPSNNSubtractionNode")}

type _NNSubtractionNodeClass struct {
	objc.Class
}

// An interface definition for the [NNSubtractionNode] class.
type INNSubtractionNode interface {
	INNBinaryArithmeticNode
}

// A representation of an subtraction operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnsubtractionnode?language=objc
type NNSubtractionNode struct {
	NNBinaryArithmeticNode
}

func NNSubtractionNodeFrom(ptr unsafe.Pointer) NNSubtractionNode {
	return NNSubtractionNode{
		NNBinaryArithmeticNode: NNBinaryArithmeticNodeFrom(ptr),
	}
}

func (nc _NNSubtractionNodeClass) Alloc() NNSubtractionNode {
	rv := objc.Call[NNSubtractionNode](nc, objc.Sel("alloc"))
	return rv
}

func NNSubtractionNode_Alloc() NNSubtractionNode {
	return NNSubtractionNodeClass.Alloc()
}

func (nc _NNSubtractionNodeClass) New() NNSubtractionNode {
	rv := objc.Call[NNSubtractionNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNSubtractionNode() NNSubtractionNode {
	return NNSubtractionNodeClass.New()
}

func (n_ NNSubtractionNode) Init() NNSubtractionNode {
	rv := objc.Call[NNSubtractionNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNSubtractionNodeClass) NodeWithSources(sourceNodes []INNImageNode) NNSubtractionNode {
	rv := objc.Call[NNSubtractionNode](nc, objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890829-nodewithsources?language=objc
func NNSubtractionNode_NodeWithSources(sourceNodes []INNImageNode) NNSubtractionNode {
	return NNSubtractionNodeClass.NodeWithSources(sourceNodes)
}

func (nc _NNSubtractionNodeClass) NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNSubtractionNode {
	rv := objc.Call[NNSubtractionNode](nc, objc.Sel("nodeWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890830-nodewithleftsource?language=objc
func NNSubtractionNode_NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNSubtractionNode {
	return NNSubtractionNodeClass.NodeWithLeftSourceRightSource(left, right)
}

func (n_ NNSubtractionNode) InitWithSources(sourceNodes []INNImageNode) NNSubtractionNode {
	rv := objc.Call[NNSubtractionNode](n_, objc.Sel("initWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890820-initwithsources?language=objc
func NewNNSubtractionNodeWithSources(sourceNodes []INNImageNode) NNSubtractionNode {
	instance := NNSubtractionNodeClass.Alloc().InitWithSources(sourceNodes)
	instance.Autorelease()
	return instance
}

func (n_ NNSubtractionNode) InitWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNSubtractionNode {
	rv := objc.Call[NNSubtractionNode](n_, objc.Sel("initWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890825-initwithleftsource?language=objc
func NewNNSubtractionNodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNSubtractionNode {
	instance := NNSubtractionNodeClass.Alloc().InitWithLeftSourceRightSource(left, right)
	instance.Autorelease()
	return instance
}
