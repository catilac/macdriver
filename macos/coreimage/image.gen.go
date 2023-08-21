// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/imageio"
	"github.com/progrium/macdriver/macos/iosurface"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Image] class.
var ImageClass = _ImageClass{objc.GetClass("CIImage")}

type _ImageClass struct {
	objc.Class
}

// An interface definition for the [Image] class.
type IImage interface {
	objc.IObject
	AutoAdjustmentFiltersWithOptions(options map[ImageAutoAdjustmentOption]objc.IObject) []Filter
	ImageBySettingAlphaOneInExtent(extent coregraphics.Rect) Image
	DrawAtPointFromRectOperationFraction(point foundation.Point, fromRect foundation.Rect, op objc.IObject, delta float64)
	ImageBySamplingLinear() Image
	ImageBySamplingNearest() Image
	DrawInRectFromRectOperationFraction(rect foundation.Rect, fromRect foundation.Rect, op objc.IObject, delta float64)
	ImageByApplyingGaussianBlurWithSigma(sigma float64) Image
	ImageByClampingToRect(rect coregraphics.Rect) Image
	ImageByUnpremultiplyingAlpha() Image
	ImageTransformForOrientation(orientation int) coregraphics.AffineTransform
	ImageByCompositingOverImage(dest IImage) Image
	ImageByColorMatchingColorSpaceToWorkingSpace(colorSpace coregraphics.ColorSpaceRef) Image
	ImageByClampingToExtent() Image
	ImageByColorMatchingWorkingSpaceToColorSpace(colorSpace coregraphics.ColorSpaceRef) Image
	ImageTransformForCGOrientation(orientation imageio.ImagePropertyOrientation) coregraphics.AffineTransform
	ImageByApplyingCGOrientation(orientation imageio.ImagePropertyOrientation) Image
	ImageByApplyingFilter(filterName string) Image
	ImageByInsertingIntermediate() Image
	ImageBySettingProperties(properties foundation.Dictionary) Image
	ImageByPremultiplyingAlpha() Image
	AutoAdjustmentFilters() []Filter
	ImageByCroppingToRect(rect coregraphics.Rect) Image
	ImageByApplyingTransformHighQualityDownsample(matrix coregraphics.AffineTransform, highQualityDownsample bool) Image
	ImageByApplyingOrientation(orientation int) Image
	RegionOfInterestForImageInRect(image IImage, rect coregraphics.Rect) coregraphics.Rect
	CGImage() coregraphics.ImageRef
	Properties() map[string]objc.Object
	PixelBuffer() corevideo.PixelBufferRef
	ColorSpace() coregraphics.ColorSpaceRef
	Url() foundation.URL
	Extent() coregraphics.Rect
	Definition() FilterShape
}

// A representation of an image to be processed or produced by Core Image filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage?language=objc
type Image struct {
	objc.Object
}

func ImageFrom(ptr unsafe.Pointer) Image {
	return Image{
		Object: objc.ObjectFrom(ptr),
	}
}

func (i_ Image) InitWithCVImageBuffer(imageBuffer corevideo.ImageBufferRef) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithCVImageBuffer:"), imageBuffer)
	return rv
}

// Initializes an image object from the contents of a Core Video image buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1438012-initwithcvimagebuffer?language=objc
func NewImageWithCVImageBuffer(imageBuffer corevideo.ImageBufferRef) Image {
	instance := ImageClass.Alloc().InitWithCVImageBuffer(imageBuffer)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithIOSurface(surface iosurface.Ref) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithIOSurface:"), surface)
	return rv
}

// Initializes an image with the contents of an IOSurface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1438030-initwithiosurface?language=objc
func NewImageWithIOSurface(surface iosurface.Ref) Image {
	instance := ImageClass.Alloc().InitWithIOSurface(surface)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithBitmapDataBytesPerRowSizeFormatColorSpace(data []byte, bytesPerRow uint, size coregraphics.Size, format Format, colorSpace coregraphics.ColorSpaceRef) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	return rv
}

// Initializes an image object with bitmap data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437857-initwithbitmapdata?language=objc
func NewImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data []byte, bytesPerRow uint, size coregraphics.Size, format Format, colorSpace coregraphics.ColorSpaceRef) Image {
	instance := ImageClass.Alloc().InitWithBitmapDataBytesPerRowSizeFormatColorSpace(data, bytesPerRow, size, format, colorSpace)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithColor(color IColor) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithColor:"), objc.Ptr(color))
	return rv
}

// Initializes an image of infinite extent whose entire content is the specified color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437947-initwithcolor?language=objc
func NewImageWithColor(color IColor) Image {
	instance := ImageClass.Alloc().InitWithColor(color)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithData(data []byte) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes an image object with the supplied image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437925-initwithdata?language=objc
func NewImageWithData(data []byte) Image {
	instance := ImageClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithBitmapImageRep(bitmapImageRep objc.IObject) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithBitmapImageRep:"), objc.Ptr(bitmapImageRep))
	return rv
}

// Initializes an image object with the specified bitmap image representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1535335-initwithbitmapimagerep?language=objc
func NewImageWithBitmapImageRep(bitmapImageRep objc.IObject) Image {
	instance := ImageClass.Alloc().InitWithBitmapImageRep(bitmapImageRep)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithMTLTextureOptions(texture metal.PTexture, options map[ImageOption]objc.IObject) Image {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	rv := objc.Call[Image](i_, objc.Sel("initWithMTLTexture:options:"), po0, options)
	return rv
}

// Initializes an image object with data supplied by a Metal texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437890-initwithmtltexture?language=objc
func NewImageWithMTLTextureOptions(texture metal.PTexture, options map[ImageOption]objc.IObject) Image {
	instance := ImageClass.Alloc().InitWithMTLTextureOptions(texture, options)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithContentsOfURL(url foundation.IURL) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Initializes an image object by reading an image from a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437908-initwithcontentsofurl?language=objc
func NewImageWithContentsOfURL(url foundation.IURL) Image {
	instance := ImageClass.Alloc().InitWithContentsOfURL(url)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithImageProviderSize(p objc.IObject, width uint) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithImageProvider:size:"), p, width)
	return rv
}

// Initializes an image object with  data provided by an image provider, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437868-initwithimageprovider?language=objc
func NewImageWithImageProviderSize(p objc.IObject, width uint) Image {
	instance := ImageClass.Alloc().InitWithImageProviderSize(p, width)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithCGImageSourceIndexOptions(source imageio.ImageSourceRef, index uint, dict map[ImageOption]objc.IObject) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithCGImageSource:index:options:"), source, index, dict)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3152399-initwithcgimagesource?language=objc
func NewImageWithCGImageSourceIndexOptions(source imageio.ImageSourceRef, index uint, dict map[ImageOption]objc.IObject) Image {
	instance := ImageClass.Alloc().InitWithCGImageSourceIndexOptions(source, index, dict)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithCGImage(image coregraphics.ImageRef) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithCGImage:"), image)
	return rv
}

// Initializes an image object with a Quartz 2D image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437986-initwithcgimage?language=objc
func NewImageWithCGImage(image coregraphics.ImageRef) Image {
	instance := ImageClass.Alloc().InitWithCGImage(image)
	instance.Autorelease()
	return instance
}

func (i_ Image) InitWithCVPixelBuffer(pixelBuffer corevideo.PixelBufferRef) Image {
	rv := objc.Call[Image](i_, objc.Sel("initWithCVPixelBuffer:"), pixelBuffer)
	return rv
}

// Initializes an image object from the contents of a Core Video pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1438072-initwithcvpixelbuffer?language=objc
func NewImageWithCVPixelBuffer(pixelBuffer corevideo.PixelBufferRef) Image {
	instance := ImageClass.Alloc().InitWithCVPixelBuffer(pixelBuffer)
	instance.Autorelease()
	return instance
}

func (ic _ImageClass) Alloc() Image {
	rv := objc.Call[Image](ic, objc.Sel("alloc"))
	return rv
}

func Image_Alloc() Image {
	return ImageClass.Alloc()
}

func (ic _ImageClass) New() Image {
	rv := objc.Call[Image](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImage() Image {
	return ImageClass.New()
}

func (i_ Image) Init() Image {
	rv := objc.Call[Image](i_, objc.Sel("init"))
	return rv
}

// Returns a subset of automatically selected and configured filters for adjusting the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437792-autoadjustmentfilterswithoptions?language=objc
func (i_ Image) AutoAdjustmentFiltersWithOptions(options map[ImageAutoAdjustmentOption]objc.IObject) []Filter {
	rv := objc.Call[[]Filter](i_, objc.Sel("autoAdjustmentFiltersWithOptions:"), options)
	return rv
}

// Creates and returns an image object from the contents of  CVImageBuffer object, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547028-imagewithcvimagebuffer?language=objc
func (ic _ImageClass) ImageWithCVImageBufferOptions(imageBuffer corevideo.ImageBufferRef, options map[ImageOption]objc.IObject) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithCVImageBuffer:options:"), imageBuffer, options)
	return rv
}

// Creates and returns an image object from the contents of  CVImageBuffer object, using the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547028-imagewithcvimagebuffer?language=objc
func Image_ImageWithCVImageBufferOptions(imageBuffer corevideo.ImageBufferRef, options map[ImageOption]objc.IObject) Image {
	return ImageClass.ImageWithCVImageBufferOptions(imageBuffer, options)
}

// Returns a new image created by setting all alpha values to 1.0 within the specified rectangle and to 0.0 outside of that area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645891-imagebysettingalphaoneinextent?language=objc
func (i_ Image) ImageBySettingAlphaOneInExtent(extent coregraphics.Rect) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageBySettingAlphaOneInExtent:"), extent)
	return rv
}

// Creates and returns an image of infinite extent whose entire content is the specified color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547012-imagewithcolor?language=objc
func (ic _ImageClass) ImageWithColor(color IColor) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithColor:"), objc.Ptr(color))
	return rv
}

// Creates and returns an image of infinite extent whose entire content is the specified color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547012-imagewithcolor?language=objc
func Image_ImageWithColor(color IColor) Image {
	return ImageClass.ImageWithColor(color)
}

// Draws all or part of the image at the specified point in the current coordinate system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1534432-drawatpoint?language=objc
func (i_ Image) DrawAtPointFromRectOperationFraction(point foundation.Point, fromRect foundation.Rect, op objc.IObject, delta float64) {
	objc.Call[objc.Void](i_, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, objc.Ptr(op), delta)
}

// Samples the image using bilinear interpolation and returns the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/2867346-imagebysamplinglinear?language=objc
func (i_ Image) ImageBySamplingLinear() Image {
	rv := objc.Call[Image](i_, objc.Sel("imageBySamplingLinear"))
	return rv
}

// Samples the image using nearest-neighbor and returns the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/2867429-imagebysamplingnearest?language=objc
func (i_ Image) ImageBySamplingNearest() Image {
	rv := objc.Call[Image](i_, objc.Sel("imageBySamplingNearest"))
	return rv
}

// Draws all or part of the image in the specified rectangle in the current coordinate system [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1534407-drawinrect?language=objc
func (i_ Image) DrawInRectFromRectOperationFraction(rect foundation.Rect, fromRect foundation.Rect, op objc.IObject, delta float64) {
	objc.Call[objc.Void](i_, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, objc.Ptr(op), delta)
}

// Returns a new image created by applying a Gaussian Blur filter to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645897-imagebyapplyinggaussianblurwiths?language=objc
func (i_ Image) ImageByApplyingGaussianBlurWithSigma(sigma float64) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByApplyingGaussianBlurWithSigma:"), sigma)
	return rv
}

// Returns a new image created by cropping to a specified area, then making the pixel colors along the edges of the cropped image extend infinitely in all directions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645893-imagebyclampingtorect?language=objc
func (i_ Image) ImageByClampingToRect(rect coregraphics.Rect) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByClampingToRect:"), rect)
	return rv
}

// Creates and returns an image from the contents of an IOSurface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547024-imagewithiosurface?language=objc
func (ic _ImageClass) ImageWithIOSurface(surface iosurface.Ref) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithIOSurface:"), surface)
	return rv
}

// Creates and returns an image from the contents of an IOSurface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547024-imagewithiosurface?language=objc
func Image_ImageWithIOSurface(surface iosurface.Ref) Image {
	return ImageClass.ImageWithIOSurface(surface)
}

// Returns a new image created by dividing the image’s RGB values by its alpha values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645892-imagebyunpremultiplyingalpha?language=objc
func (i_ Image) ImageByUnpremultiplyingAlpha() Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByUnpremultiplyingAlpha"))
	return rv
}

// Returns the transformation needed to reorient the image to the specified orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437930-imagetransformfororientation?language=objc
func (i_ Image) ImageTransformForOrientation(orientation int) coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](i_, objc.Sel("imageTransformForOrientation:"), orientation)
	return rv
}

// Returns a new image created by compositing the original image over the specified destination image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437837-imagebycompositingoverimage?language=objc
func (i_ Image) ImageByCompositingOverImage(dest IImage) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByCompositingOverImage:"), objc.Ptr(dest))
	return rv
}

// Returns a new image created by color matching from the specified color space to the context’s working color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645896-imagebycolormatchingcolorspaceto?language=objc
func (i_ Image) ImageByColorMatchingColorSpaceToWorkingSpace(colorSpace coregraphics.ColorSpaceRef) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByColorMatchingColorSpaceToWorkingSpace:"), colorSpace)
	return rv
}

// Creates and returns an image object with data supplied by a Metal texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1546999-imagewithmtltexture?language=objc
func (ic _ImageClass) ImageWithMTLTextureOptions(texture metal.PTexture, options map[ImageOption]objc.IObject) Image {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	rv := objc.Call[Image](ic, objc.Sel("imageWithMTLTexture:options:"), po0, options)
	return rv
}

// Creates and returns an image object with data supplied by a Metal texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1546999-imagewithmtltexture?language=objc
func Image_ImageWithMTLTextureOptions(texture metal.PTexture, options map[ImageOption]objc.IObject) Image {
	return ImageClass.ImageWithMTLTextureOptions(texture, options)
}

// Returns a new image created by making the pixel colors along its edges extend infinitely in all directions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437628-imagebyclampingtoextent?language=objc
func (i_ Image) ImageByClampingToExtent() Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByClampingToExtent"))
	return rv
}

// Creates and returns an image object initialized with the supplied image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547029-imagewithdata?language=objc
func (ic _ImageClass) ImageWithData(data []byte) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithData:"), data)
	return rv
}

// Creates and returns an image object initialized with the supplied image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547029-imagewithdata?language=objc
func Image_ImageWithData(data []byte) Image {
	return ImageClass.ImageWithData(data)
}

// Creates and returns an image object from the contents of a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547027-imagewithcontentsofurl?language=objc
func (ic _ImageClass) ImageWithContentsOfURL(url foundation.IURL) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithContentsOfURL:"), objc.Ptr(url))
	return rv
}

// Creates and returns an image object from the contents of a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547027-imagewithcontentsofurl?language=objc
func Image_ImageWithContentsOfURL(url foundation.IURL) Image {
	return ImageClass.ImageWithContentsOfURL(url)
}

// Creates and returns an image object initialized with data provided by an image provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1579115-imagewithimageprovider?language=objc
func (ic _ImageClass) ImageWithImageProviderSize(p objc.IObject, width uint) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithImageProvider:size:"), p, width)
	return rv
}

// Creates and returns an image object initialized with data provided by an image provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1579115-imagewithimageprovider?language=objc
func Image_ImageWithImageProviderSize(p objc.IObject, width uint) Image {
	return ImageClass.ImageWithImageProviderSize(p, width)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3152398-imagewithcgimagesource?language=objc
func (ic _ImageClass) ImageWithCGImageSourceIndexOptions(source imageio.ImageSourceRef, index uint, dict map[ImageOption]objc.IObject) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithCGImageSource:index:options:"), source, index, dict)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3152398-imagewithcgimagesource?language=objc
func Image_ImageWithCGImageSourceIndexOptions(source imageio.ImageSourceRef, index uint, dict map[ImageOption]objc.IObject) Image {
	return ImageClass.ImageWithCGImageSourceIndexOptions(source, index, dict)
}

// Returns a new image created by color matching from the context’s working color space to the specified color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645898-imagebycolormatchingworkingspace?language=objc
func (i_ Image) ImageByColorMatchingWorkingSpaceToColorSpace(colorSpace coregraphics.ColorSpaceRef) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByColorMatchingWorkingSpaceToColorSpace:"), colorSpace)
	return rv
}

// The affine transform for changing the image to the given orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/2919726-imagetransformforcgorientation?language=objc
func (i_ Image) ImageTransformForCGOrientation(orientation imageio.ImagePropertyOrientation) coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](i_, objc.Sel("imageTransformForCGOrientation:"), orientation)
	return rv
}

// Transforms the original image by a given CGImagePropertyOrientation and returns the result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/2919727-imagebyapplyingcgorientation?language=objc
func (i_ Image) ImageByApplyingCGOrientation(orientation imageio.ImagePropertyOrientation) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByApplyingCGOrientation:"), orientation)
	return rv
}

// Creates and returns an image object from a Quartz 2D image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547025-imagewithcgimage?language=objc
func (ic _ImageClass) ImageWithCGImage(image coregraphics.ImageRef) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithCGImage:"), image)
	return rv
}

// Creates and returns an image object from a Quartz 2D image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547025-imagewithcgimage?language=objc
func Image_ImageWithCGImage(image coregraphics.ImageRef) Image {
	return ImageClass.ImageWithCGImage(image)
}

// Applies the filter to an image and returns the output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/2915368-imagebyapplyingfilter?language=objc
func (i_ Image) ImageByApplyingFilter(filterName string) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByApplyingFilter:"), filterName)
	return rv
}

// Returns a new image created by inserting an intermediate. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/2966521-imagebyinsertingintermediate?language=objc
func (i_ Image) ImageByInsertingIntermediate() Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByInsertingIntermediate"))
	return rv
}

// Creates and returns an empty image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1438023-emptyimage?language=objc
func (ic _ImageClass) EmptyImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("emptyImage"))
	return rv
}

// Creates and returns an empty image object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1438023-emptyimage?language=objc
func Image_EmptyImage() Image {
	return ImageClass.EmptyImage()
}

// Returns a new image created by adding the specified metadata properties to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645895-imagebysettingproperties?language=objc
func (i_ Image) ImageBySettingProperties(properties foundation.Dictionary) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageBySettingProperties:"), properties)
	return rv
}

// Returns a new image created by multiplying the image’s RGB values by its alpha values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645894-imagebypremultiplyingalpha?language=objc
func (i_ Image) ImageByPremultiplyingAlpha() Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByPremultiplyingAlpha"))
	return rv
}

// Returns all possible automatically selected and configured filters for adjusting the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1645889-autoadjustmentfilters?language=objc
func (i_ Image) AutoAdjustmentFilters() []Filter {
	rv := objc.Call[[]Filter](i_, objc.Sel("autoAdjustmentFilters"))
	return rv
}

// Returns a new image with a cropped portion of the original image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437833-imagebycroppingtorect?language=objc
func (i_ Image) ImageByCroppingToRect(rect coregraphics.Rect) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByCroppingToRect:"), rect)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3334939-imagebyapplyingtransform?language=objc
func (i_ Image) ImageByApplyingTransformHighQualityDownsample(matrix coregraphics.AffineTransform, highQualityDownsample bool) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByApplyingTransform:highQualityDownsample:"), matrix, highQualityDownsample)
	return rv
}

// Creates and returns an image object from bitmap data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547023-imagewithbitmapdata?language=objc
func (ic _ImageClass) ImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data []byte, bytesPerRow uint, size coregraphics.Size, format Format, colorSpace coregraphics.ColorSpaceRef) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithBitmapData:bytesPerRow:size:format:colorSpace:"), data, bytesPerRow, size, format, colorSpace)
	return rv
}

// Creates and returns an image object from bitmap data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547023-imagewithbitmapdata?language=objc
func Image_ImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data []byte, bytesPerRow uint, size coregraphics.Size, format Format, colorSpace coregraphics.ColorSpaceRef) Image {
	return ImageClass.ImageWithBitmapDataBytesPerRowSizeFormatColorSpace(data, bytesPerRow, size, format, colorSpace)
}

// Returns a new image created by transforming the original image to the specified EXIF orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1438223-imagebyapplyingorientation?language=objc
func (i_ Image) ImageByApplyingOrientation(orientation int) Image {
	rv := objc.Call[Image](i_, objc.Sel("imageByApplyingOrientation:"), orientation)
	return rv
}

// Returns the region of interest for the filter chain that generates the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437994-regionofinterestforimage?language=objc
func (i_ Image) RegionOfInterestForImageInRect(image IImage, rect coregraphics.Rect) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](i_, objc.Sel("regionOfInterestForImage:inRect:"), objc.Ptr(image), rect)
	return rv
}

// Creates and returns an image object from the contents of  CVPixelBuffer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547005-imagewithcvpixelbuffer?language=objc
func (ic _ImageClass) ImageWithCVPixelBuffer(pixelBuffer corevideo.PixelBufferRef) Image {
	rv := objc.Call[Image](ic, objc.Sel("imageWithCVPixelBuffer:"), pixelBuffer)
	return rv
}

// Creates and returns an image object from the contents of  CVPixelBuffer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1547005-imagewithcvpixelbuffer?language=objc
func Image_ImageWithCVPixelBuffer(pixelBuffer corevideo.PixelBufferRef) Image {
	return ImageClass.ImageWithCVPixelBuffer(pixelBuffer)
}

// The CoreGraphics image object this image was created from, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1687603-cgimage?language=objc
func (i_ Image) CGImage() coregraphics.ImageRef {
	rv := objc.Call[coregraphics.ImageRef](i_, objc.Sel("CGImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074424-cyanimage?language=objc
func (ic _ImageClass) CyanImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("cyanImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074424-cyanimage?language=objc
func Image_CyanImage() Image {
	return ImageClass.CyanImage()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074422-blueimage?language=objc
func (ic _ImageClass) BlueImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("blueImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074422-blueimage?language=objc
func Image_BlueImage() Image {
	return ImageClass.BlueImage()
}

// A dictionary containing metadata about the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437733-properties?language=objc
func (i_ Image) Properties() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](i_, objc.Sel("properties"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074428-redimage?language=objc
func (ic _ImageClass) RedImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("redImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074428-redimage?language=objc
func Image_RedImage() Image {
	return ImageClass.RedImage()
}

// The CoreVideo pixel buffer this image was created from, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1687604-pixelbuffer?language=objc
func (i_ Image) PixelBuffer() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](i_, objc.Sel("pixelBuffer"))
	return rv
}

// The color space of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437750-colorspace?language=objc
func (i_ Image) ColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](i_, objc.Sel("colorSpace"))
	return rv
}

// The URL from which the image was loaded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1438195-url?language=objc
func (i_ Image) Url() foundation.URL {
	rv := objc.Call[foundation.URL](i_, objc.Sel("url"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074423-clearimage?language=objc
func (ic _ImageClass) ClearImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("clearImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074423-clearimage?language=objc
func Image_ClearImage() Image {
	return ImageClass.ClearImage()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074426-greenimage?language=objc
func (ic _ImageClass) GreenImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("greenImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074426-greenimage?language=objc
func Image_GreenImage() Image {
	return ImageClass.GreenImage()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074430-yellowimage?language=objc
func (ic _ImageClass) YellowImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("yellowImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074430-yellowimage?language=objc
func Image_YellowImage() Image {
	return ImageClass.YellowImage()
}

// A rectangle that specifies the extent of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437996-extent?language=objc
func (i_ Image) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](i_, objc.Sel("extent"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074421-blackimage?language=objc
func (ic _ImageClass) BlackImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("blackImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074421-blackimage?language=objc
func Image_BlackImage() Image {
	return ImageClass.BlackImage()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074425-grayimage?language=objc
func (ic _ImageClass) GrayImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("grayImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074425-grayimage?language=objc
func Image_GrayImage() Image {
	return ImageClass.GrayImage()
}

// Returns a filter shape object that represents the domain of definition of the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/1437804-definition?language=objc
func (i_ Image) Definition() FilterShape {
	rv := objc.Call[FilterShape](i_, objc.Sel("definition"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074427-magentaimage?language=objc
func (ic _ImageClass) MagentaImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("magentaImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074427-magentaimage?language=objc
func Image_MagentaImage() Image {
	return ImageClass.MagentaImage()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074429-whiteimage?language=objc
func (ic _ImageClass) WhiteImage() Image {
	rv := objc.Call[Image](ic, objc.Sel("whiteImage"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimage/3074429-whiteimage?language=objc
func Image_WhiteImage() Image {
	return ImageClass.WhiteImage()
}
