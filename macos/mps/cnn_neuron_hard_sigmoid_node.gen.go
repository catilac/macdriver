// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronHardSigmoidNode] class.
var CNNNeuronHardSigmoidNodeClass = _CNNNeuronHardSigmoidNodeClass{objc.GetClass("MPSCNNNeuronHardSigmoidNode")}

type _CNNNeuronHardSigmoidNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronHardSigmoidNode] class.
type ICNNNeuronHardSigmoidNode interface {
	ICNNNeuronNode
}

// A representation of a hard sigmoid neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoidnode?language=objc
type CNNNeuronHardSigmoidNode struct {
	CNNNeuronNode
}

func CNNNeuronHardSigmoidNodeFrom(ptr unsafe.Pointer) CNNNeuronHardSigmoidNode {
	return CNNNeuronHardSigmoidNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronHardSigmoidNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronHardSigmoidNode {
	rv := objc.Call[CNNNeuronHardSigmoidNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoidnode/2921453-nodewithsource?language=objc
func CNNNeuronHardSigmoidNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronHardSigmoidNode {
	return CNNNeuronHardSigmoidNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronHardSigmoidNode) InitWithSource(sourceNode INNImageNode) CNNNeuronHardSigmoidNode {
	rv := objc.Call[CNNNeuronHardSigmoidNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoidnode/2921455-initwithsource?language=objc
func NewCNNNeuronHardSigmoidNodeWithSource(sourceNode INNImageNode) CNNNeuronHardSigmoidNode {
	instance := CNNNeuronHardSigmoidNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronHardSigmoidNodeClass) Alloc() CNNNeuronHardSigmoidNode {
	rv := objc.Call[CNNNeuronHardSigmoidNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronHardSigmoidNode_Alloc() CNNNeuronHardSigmoidNode {
	return CNNNeuronHardSigmoidNodeClass.Alloc()
}

func (cc _CNNNeuronHardSigmoidNodeClass) New() CNNNeuronHardSigmoidNode {
	rv := objc.Call[CNNNeuronHardSigmoidNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronHardSigmoidNode() CNNNeuronHardSigmoidNode {
	return CNNNeuronHardSigmoidNodeClass.New()
}

func (c_ CNNNeuronHardSigmoidNode) Init() CNNNeuronHardSigmoidNode {
	rv := objc.Call[CNNNeuronHardSigmoidNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronHardSigmoidNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronHardSigmoidNode {
	rv := objc.Call[CNNNeuronHardSigmoidNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronHardSigmoidNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronHardSigmoidNode {
	return CNNNeuronHardSigmoidNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
