// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TimedMetadataGroup] class.
var TimedMetadataGroupClass = _TimedMetadataGroupClass{objc.GetClass("AVTimedMetadataGroup")}

type _TimedMetadataGroupClass struct {
	objc.Class
}

// An interface definition for the [TimedMetadataGroup] class.
type ITimedMetadataGroup interface {
	IMetadataGroup
	CopyFormatDescription() coremedia.MetadataFormatDescriptionRef
	TimeRange() coremedia.TimeRange
}

// A collection of metadata items that are valid for use during a specific time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup?language=objc
type TimedMetadataGroup struct {
	MetadataGroup
}

func TimedMetadataGroupFrom(ptr unsafe.Pointer) TimedMetadataGroup {
	return TimedMetadataGroup{
		MetadataGroup: MetadataGroupFrom(ptr),
	}
}

func (t_ TimedMetadataGroup) InitWithSampleBuffer(sampleBuffer coremedia.SampleBufferRef) TimedMetadataGroup {
	rv := objc.Call[TimedMetadataGroup](t_, objc.Sel("initWithSampleBuffer:"), sampleBuffer)
	return rv
}

// Creates a timed metadata group with a sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup/1387128-initwithsamplebuffer?language=objc
func NewTimedMetadataGroupWithSampleBuffer(sampleBuffer coremedia.SampleBufferRef) TimedMetadataGroup {
	instance := TimedMetadataGroupClass.Alloc().InitWithSampleBuffer(sampleBuffer)
	instance.Autorelease()
	return instance
}

func (t_ TimedMetadataGroup) InitWithItemsTimeRange(items []IMetadataItem, timeRange coremedia.TimeRange) TimedMetadataGroup {
	rv := objc.Call[TimedMetadataGroup](t_, objc.Sel("initWithItems:timeRange:"), items, timeRange)
	return rv
}

// Creates a timed metadata group initialized with the given metadata items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup/1389632-initwithitems?language=objc
func NewTimedMetadataGroupWithItemsTimeRange(items []IMetadataItem, timeRange coremedia.TimeRange) TimedMetadataGroup {
	instance := TimedMetadataGroupClass.Alloc().InitWithItemsTimeRange(items, timeRange)
	instance.Autorelease()
	return instance
}

func (tc _TimedMetadataGroupClass) Alloc() TimedMetadataGroup {
	rv := objc.Call[TimedMetadataGroup](tc, objc.Sel("alloc"))
	return rv
}

func TimedMetadataGroup_Alloc() TimedMetadataGroup {
	return TimedMetadataGroupClass.Alloc()
}

func (tc _TimedMetadataGroupClass) New() TimedMetadataGroup {
	rv := objc.Call[TimedMetadataGroup](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTimedMetadataGroup() TimedMetadataGroup {
	return TimedMetadataGroupClass.New()
}

func (t_ TimedMetadataGroup) Init() TimedMetadataGroup {
	rv := objc.Call[TimedMetadataGroup](t_, objc.Sel("init"))
	return rv
}

// Creates a format description based on the receiver's items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup/1389461-copyformatdescription?language=objc
func (t_ TimedMetadataGroup) CopyFormatDescription() coremedia.MetadataFormatDescriptionRef {
	rv := objc.Call[coremedia.MetadataFormatDescriptionRef](t_, objc.Sel("copyFormatDescription"))
	return rv
}

// The time range for the timed metadata. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avtimedmetadatagroup/1387992-timerange?language=objc
func (t_ TimedMetadataGroup) TimeRange() coremedia.TimeRange {
	rv := objc.Call[coremedia.TimeRange](t_, objc.Sel("timeRange"))
	return rv
}
