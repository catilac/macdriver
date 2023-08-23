// AUTO-GENERATED CODE, DO NOT MODIFY

package quartz

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/quartzcore"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CompositionLayer] class.
var CompositionLayerClass = _CompositionLayerClass{objc.GetClass("QCCompositionLayer")}

type _CompositionLayerClass struct {
	objc.Class
}

// An interface definition for the [CompositionLayer] class.
type ICompositionLayer interface {
	quartzcore.IOpenGLLayer
}

// A layer that loads, plays, and controls Quartz Composer compositions in a Core Animation layer hierarchy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartz/qccompositionlayer?language=objc
type CompositionLayer struct {
	quartzcore.OpenGLLayer
}

func CompositionLayerFrom(ptr unsafe.Pointer) CompositionLayer {
	return CompositionLayer{
		OpenGLLayer: quartzcore.OpenGLLayerFrom(ptr),
	}
}

func (cc _CompositionLayerClass) Alloc() CompositionLayer {
	rv := objc.Call[CompositionLayer](cc, objc.Sel("alloc"))
	return rv
}

func CompositionLayer_Alloc() CompositionLayer {
	return CompositionLayerClass.Alloc()
}

func (cc _CompositionLayerClass) New() CompositionLayer {
	rv := objc.Call[CompositionLayer](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCompositionLayer() CompositionLayer {
	return CompositionLayerClass.New()
}

func (c_ CompositionLayer) Init() CompositionLayer {
	rv := objc.Call[CompositionLayer](c_, objc.Sel("init"))
	return rv
}

func (cc _CompositionLayerClass) Layer() CompositionLayer {
	rv := objc.Call[CompositionLayer](cc, objc.Sel("layer"))
	return rv
}

// Creates and returns an instance of the layer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410793-layer?language=objc
func CompositionLayer_Layer() CompositionLayer {
	return CompositionLayerClass.Layer()
}

func (c_ CompositionLayer) InitWithLayer(layer objc.IObject) CompositionLayer {
	rv := objc.Call[CompositionLayer](c_, objc.Sel("initWithLayer:"), layer)
	return rv
}

// Override to copy or initialize custom fields of the specified layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410842-initwithlayer?language=objc
func NewCompositionLayerWithLayer(layer objc.IObject) CompositionLayer {
	instance := CompositionLayerClass.Alloc().InitWithLayer(layer)
	instance.Autorelease()
	return instance
}

func (c_ CompositionLayer) ModelLayer() CompositionLayer {
	rv := objc.Call[CompositionLayer](c_, objc.Sel("modelLayer"))
	return rv
}

// Returns the model layer object associated with the receiver, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410853-modellayer?language=objc
func CompositionLayer_ModelLayer() CompositionLayer {
	instance := CompositionLayerClass.Alloc().ModelLayer()
	instance.Autorelease()
	return instance
}

func (c_ CompositionLayer) PresentationLayer() CompositionLayer {
	rv := objc.Call[CompositionLayer](c_, objc.Sel("presentationLayer"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/1410744-presentationlayer?language=objc
func CompositionLayer_PresentationLayer() CompositionLayer {
	instance := CompositionLayerClass.Alloc().PresentationLayer()
	instance.Autorelease()
	return instance
}
