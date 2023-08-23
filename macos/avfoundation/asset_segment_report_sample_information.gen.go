// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetSegmentReportSampleInformation] class.
var AssetSegmentReportSampleInformationClass = _AssetSegmentReportSampleInformationClass{objc.GetClass("AVAssetSegmentReportSampleInformation")}

type _AssetSegmentReportSampleInformationClass struct {
	objc.Class
}

// An interface definition for the [AssetSegmentReportSampleInformation] class.
type IAssetSegmentReportSampleInformation interface {
	objc.IObject
	Length() int
	IsSyncSample() bool
	Offset() int
	PresentationTimeStamp() coremedia.Time
}

// An object that provides information about sample data in a track. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreportsampleinformation?language=objc
type AssetSegmentReportSampleInformation struct {
	objc.Object
}

func AssetSegmentReportSampleInformationFrom(ptr unsafe.Pointer) AssetSegmentReportSampleInformation {
	return AssetSegmentReportSampleInformation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetSegmentReportSampleInformationClass) Alloc() AssetSegmentReportSampleInformation {
	rv := objc.Call[AssetSegmentReportSampleInformation](ac, objc.Sel("alloc"))
	return rv
}

func AssetSegmentReportSampleInformation_Alloc() AssetSegmentReportSampleInformation {
	return AssetSegmentReportSampleInformationClass.Alloc()
}

func (ac _AssetSegmentReportSampleInformationClass) New() AssetSegmentReportSampleInformation {
	rv := objc.Call[AssetSegmentReportSampleInformation](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetSegmentReportSampleInformation() AssetSegmentReportSampleInformation {
	return AssetSegmentReportSampleInformationClass.New()
}

func (a_ AssetSegmentReportSampleInformation) Init() AssetSegmentReportSampleInformation {
	rv := objc.Call[AssetSegmentReportSampleInformation](a_, objc.Sel("init"))
	return rv
}

// The length of the sample data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreportsampleinformation/3546573-length?language=objc
func (a_ AssetSegmentReportSampleInformation) Length() int {
	rv := objc.Call[int](a_, objc.Sel("length"))
	return rv
}

// A Boolean value that indicates whether the sample is a key frame. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreportsampleinformation/3563932-issyncsample?language=objc
func (a_ AssetSegmentReportSampleInformation) IsSyncSample() bool {
	rv := objc.Call[bool](a_, objc.Sel("isSyncSample"))
	return rv
}

// The offset of a sample in the segment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreportsampleinformation/3546574-offset?language=objc
func (a_ AssetSegmentReportSampleInformation) Offset() int {
	rv := objc.Call[int](a_, objc.Sel("offset"))
	return rv
}

// The presentation timestamp (PTS) of a sample. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreportsampleinformation/3546575-presentationtimestamp?language=objc
func (a_ AssetSegmentReportSampleInformation) PresentationTimeStamp() coremedia.Time {
	rv := objc.Call[coremedia.Time](a_, objc.Sel("presentationTimeStamp"))
	return rv
}
