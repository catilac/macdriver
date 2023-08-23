// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronGradient] class.
var CNNNeuronGradientClass = _CNNNeuronGradientClass{objc.GetClass("MPSCNNNeuronGradient")}

type _CNNNeuronGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronGradient] class.
type ICNNNeuronGradient interface {
	ICNNGradientKernel
	A() float64
	Data() []byte
	NeuronType() CNNNeuronType
	C() float64
	B() float64
}

// A gradient neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient?language=objc
type CNNNeuronGradient struct {
	CNNGradientKernel
}

func CNNNeuronGradientFrom(ptr unsafe.Pointer) CNNNeuronGradient {
	return CNNNeuronGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNNeuronGradient) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronGradient](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942293-initwithdevice?language=objc
func NewCNNNeuronGradientWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronGradient {
	instance := CNNNeuronGradientClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronGradientClass) Alloc() CNNNeuronGradient {
	rv := objc.Call[CNNNeuronGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronGradient_Alloc() CNNNeuronGradient {
	return CNNNeuronGradientClass.Alloc()
}

func (cc _CNNNeuronGradientClass) New() CNNNeuronGradient {
	rv := objc.Call[CNNNeuronGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronGradient() CNNNeuronGradient {
	return CNNNeuronGradientClass.New()
}

func (c_ CNNNeuronGradient) Init() CNNNeuronGradient {
	rv := objc.Call[CNNNeuronGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronGradient) InitWithDevice(device metal.PDevice) CNNNeuronGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNNeuronGradientWithDevice(device metal.PDevice) CNNNeuronGradient {
	instance := CNNNeuronGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronGradient {
	instance := CNNNeuronGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942312-a?language=objc
func (c_ CNNNeuronGradient) A() float64 {
	rv := objc.Call[float64](c_, objc.Sel("a"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942314-data?language=objc
func (c_ CNNNeuronGradient) Data() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("data"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942300-neurontype?language=objc
func (c_ CNNNeuronGradient) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](c_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942310-c?language=objc
func (c_ CNNNeuronGradient) C() float64 {
	rv := objc.Call[float64](c_, objc.Sel("c"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongradient/2942313-b?language=objc
func (c_ CNNNeuronGradient) B() float64 {
	rv := objc.Call[float64](c_, objc.Sel("b"))
	return rv
}
