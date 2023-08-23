// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetResourceLoadingDataRequest] class.
var AssetResourceLoadingDataRequestClass = _AssetResourceLoadingDataRequestClass{objc.GetClass("AVAssetResourceLoadingDataRequest")}

type _AssetResourceLoadingDataRequestClass struct {
	objc.Class
}

// An interface definition for the [AssetResourceLoadingDataRequest] class.
type IAssetResourceLoadingDataRequest interface {
	objc.IObject
	RespondWithData(data []byte)
	RequestedLength() int
	RequestsAllDataToEndOfResource() bool
	RequestedOffset() int64
	CurrentOffset() int64
}

// An object for requesting data from a resource that an asset resource-loading request references. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingdatarequest?language=objc
type AssetResourceLoadingDataRequest struct {
	objc.Object
}

func AssetResourceLoadingDataRequestFrom(ptr unsafe.Pointer) AssetResourceLoadingDataRequest {
	return AssetResourceLoadingDataRequest{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AssetResourceLoadingDataRequestClass) Alloc() AssetResourceLoadingDataRequest {
	rv := objc.Call[AssetResourceLoadingDataRequest](ac, objc.Sel("alloc"))
	return rv
}

func AssetResourceLoadingDataRequest_Alloc() AssetResourceLoadingDataRequest {
	return AssetResourceLoadingDataRequestClass.Alloc()
}

func (ac _AssetResourceLoadingDataRequestClass) New() AssetResourceLoadingDataRequest {
	rv := objc.Call[AssetResourceLoadingDataRequest](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetResourceLoadingDataRequest() AssetResourceLoadingDataRequest {
	return AssetResourceLoadingDataRequestClass.New()
}

func (a_ AssetResourceLoadingDataRequest) Init() AssetResourceLoadingDataRequest {
	rv := objc.Call[AssetResourceLoadingDataRequest](a_, objc.Sel("init"))
	return rv
}

// Provides data to the loading request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingdatarequest/1390581-respondwithdata?language=objc
func (a_ AssetResourceLoadingDataRequest) RespondWithData(data []byte) {
	objc.Call[objc.Void](a_, objc.Sel("respondWithData:"), data)
}

// The length, in bytes, of the data requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingdatarequest/1387720-requestedlength?language=objc
func (a_ AssetResourceLoadingDataRequest) RequestedLength() int {
	rv := objc.Call[int](a_, objc.Sel("requestedLength"))
	return rv
}

// A Boolean value that indicates the entire remaining length of the resource from the offest to the end of the resource is being requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingdatarequest/1386864-requestsalldatatoendofresource?language=objc
func (a_ AssetResourceLoadingDataRequest) RequestsAllDataToEndOfResource() bool {
	rv := objc.Call[bool](a_, objc.Sel("requestsAllDataToEndOfResource"))
	return rv
}

// The position within the resource of the first byte requested. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingdatarequest/1388428-requestedoffset?language=objc
func (a_ AssetResourceLoadingDataRequest) RequestedOffset() int64 {
	rv := objc.Call[int64](a_, objc.Sel("requestedOffset"))
	return rv
}

// The position within the resource of the next byte. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloadingdatarequest/1385945-currentoffset?language=objc
func (a_ AssetResourceLoadingDataRequest) CurrentOffset() int64 {
	rv := objc.Call[int64](a_, objc.Sel("currentOffset"))
	return rv
}
