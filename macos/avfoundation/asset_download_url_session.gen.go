// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AssetDownloadURLSession] class.
var AssetDownloadURLSessionClass = _AssetDownloadURLSessionClass{objc.GetClass("AVAssetDownloadURLSession")}

type _AssetDownloadURLSessionClass struct {
	objc.Class
}

// An interface definition for the [AssetDownloadURLSession] class.
type IAssetDownloadURLSession interface {
	foundation.IURLSession
	AggregateAssetDownloadTaskWithURLAssetMediaSelectionsAssetTitleAssetArtworkDataOptions(URLAsset IURLAsset, mediaSelections []IMediaSelection, title string, artworkData []byte, options map[string]objc.IObject) AggregateAssetDownloadTask
	AssetDownloadTaskWithConfiguration(downloadConfiguration IAssetDownloadConfiguration) AssetDownloadTask
}

// A URL session that creates and executes asset download tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadurlsession?language=objc
type AssetDownloadURLSession struct {
	foundation.URLSession
}

func AssetDownloadURLSessionFrom(ptr unsafe.Pointer) AssetDownloadURLSession {
	return AssetDownloadURLSession{
		URLSession: foundation.URLSessionFrom(ptr),
	}
}

func (ac _AssetDownloadURLSessionClass) Alloc() AssetDownloadURLSession {
	rv := objc.Call[AssetDownloadURLSession](ac, objc.Sel("alloc"))
	return rv
}

func AssetDownloadURLSession_Alloc() AssetDownloadURLSession {
	return AssetDownloadURLSessionClass.Alloc()
}

func (ac _AssetDownloadURLSessionClass) New() AssetDownloadURLSession {
	rv := objc.Call[AssetDownloadURLSession](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAssetDownloadURLSession() AssetDownloadURLSession {
	return AssetDownloadURLSessionClass.New()
}

func (a_ AssetDownloadURLSession) Init() AssetDownloadURLSession {
	rv := objc.Call[AssetDownloadURLSession](a_, objc.Sel("init"))
	return rv
}

// Creates a URL session to download assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadurlsession/1621015-sessionwithconfiguration?language=objc
func (ac _AssetDownloadURLSessionClass) SessionWithConfigurationAssetDownloadDelegateDelegateQueue(configuration foundation.IURLSessionConfiguration, delegate PAssetDownloadDelegate, delegateQueue foundation.IOperationQueue) AssetDownloadURLSession {
	po1 := objc.WrapAsProtocol("AVAssetDownloadDelegate", delegate)
	rv := objc.Call[AssetDownloadURLSession](ac, objc.Sel("sessionWithConfiguration:assetDownloadDelegate:delegateQueue:"), objc.Ptr(configuration), po1, objc.Ptr(delegateQueue))
	return rv
}

// Creates a URL session to download assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadurlsession/1621015-sessionwithconfiguration?language=objc
func AssetDownloadURLSession_SessionWithConfigurationAssetDownloadDelegateDelegateQueue(configuration foundation.IURLSessionConfiguration, delegate PAssetDownloadDelegate, delegateQueue foundation.IOperationQueue) AssetDownloadURLSession {
	return AssetDownloadURLSessionClass.SessionWithConfigurationAssetDownloadDelegateDelegateQueue(configuration, delegate, delegateQueue)
}

// Creates a URL session to download assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadurlsession/1621015-sessionwithconfiguration?language=objc
func (ac _AssetDownloadURLSessionClass) SessionWithConfigurationAssetDownloadDelegateObjectDelegateQueue(configuration foundation.IURLSessionConfiguration, delegateObject objc.IObject, delegateQueue foundation.IOperationQueue) AssetDownloadURLSession {
	rv := objc.Call[AssetDownloadURLSession](ac, objc.Sel("sessionWithConfiguration:assetDownloadDelegate:delegateQueue:"), objc.Ptr(configuration), objc.Ptr(delegateObject), objc.Ptr(delegateQueue))
	return rv
}

// Creates a URL session to download assets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadurlsession/1621015-sessionwithconfiguration?language=objc
func AssetDownloadURLSession_SessionWithConfigurationAssetDownloadDelegateObjectDelegateQueue(configuration foundation.IURLSessionConfiguration, delegateObject objc.IObject, delegateQueue foundation.IOperationQueue) AssetDownloadURLSession {
	return AssetDownloadURLSessionClass.SessionWithConfigurationAssetDownloadDelegateObjectDelegateQueue(configuration, delegateObject, delegateQueue)
}

// Creates a download task to download the asset and media selections. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadurlsession/2897242-aggregateassetdownloadtaskwithur?language=objc
func (a_ AssetDownloadURLSession) AggregateAssetDownloadTaskWithURLAssetMediaSelectionsAssetTitleAssetArtworkDataOptions(URLAsset IURLAsset, mediaSelections []IMediaSelection, title string, artworkData []byte, options map[string]objc.IObject) AggregateAssetDownloadTask {
	rv := objc.Call[AggregateAssetDownloadTask](a_, objc.Sel("aggregateAssetDownloadTaskWithURLAsset:mediaSelections:assetTitle:assetArtworkData:options:"), objc.Ptr(URLAsset), mediaSelections, title, artworkData, options)
	return rv
}

// Creates a download task that uses the specified configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetdownloadurlsession/3751761-assetdownloadtaskwithconfigurati?language=objc
func (a_ AssetDownloadURLSession) AssetDownloadTaskWithConfiguration(downloadConfiguration IAssetDownloadConfiguration) AssetDownloadTask {
	rv := objc.Call[AssetDownloadTask](a_, objc.Sel("assetDownloadTaskWithConfiguration:"), objc.Ptr(downloadConfiguration))
	return rv
}
