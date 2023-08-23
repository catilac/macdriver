// AUTO-GENERATED CODE, DO NOT MODIFY

package mediaplayer

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NowPlayingInfoLanguageOption] class.
var NowPlayingInfoLanguageOptionClass = _NowPlayingInfoLanguageOptionClass{objc.GetClass("MPNowPlayingInfoLanguageOption")}

type _NowPlayingInfoLanguageOptionClass struct {
	objc.Class
}

// An interface definition for the [NowPlayingInfoLanguageOption] class.
type INowPlayingInfoLanguageOption interface {
	objc.IObject
	IsAutomaticLegibleLanguageOption() bool
	IsAutomaticAudibleLanguageOption() bool
	LanguageOptionType() NowPlayingInfoLanguageOptionType
	LanguageTag() string
	LanguageOptionCharacteristics() []string
	DisplayName() string
	Identifier() string
}

// A set of interfaces for setting the language option for the Now Playing item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption?language=objc
type NowPlayingInfoLanguageOption struct {
	objc.Object
}

func NowPlayingInfoLanguageOptionFrom(ptr unsafe.Pointer) NowPlayingInfoLanguageOption {
	return NowPlayingInfoLanguageOption{
		Object: objc.ObjectFrom(ptr),
	}
}

func (n_ NowPlayingInfoLanguageOption) InitWithTypeLanguageTagCharacteristicsDisplayNameIdentifier(languageOptionType NowPlayingInfoLanguageOptionType, languageTag string, languageOptionCharacteristics []string, displayName string, identifier string) NowPlayingInfoLanguageOption {
	rv := objc.Call[NowPlayingInfoLanguageOption](n_, objc.Sel("initWithType:languageTag:characteristics:displayName:identifier:"), languageOptionType, languageTag, languageOptionCharacteristics, displayName, identifier)
	return rv
}

// Creates a single language option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption/1623159-initwithtype?language=objc
func NewNowPlayingInfoLanguageOptionWithTypeLanguageTagCharacteristicsDisplayNameIdentifier(languageOptionType NowPlayingInfoLanguageOptionType, languageTag string, languageOptionCharacteristics []string, displayName string, identifier string) NowPlayingInfoLanguageOption {
	instance := NowPlayingInfoLanguageOptionClass.Alloc().InitWithTypeLanguageTagCharacteristicsDisplayNameIdentifier(languageOptionType, languageTag, languageOptionCharacteristics, displayName, identifier)
	instance.Autorelease()
	return instance
}

func (nc _NowPlayingInfoLanguageOptionClass) Alloc() NowPlayingInfoLanguageOption {
	rv := objc.Call[NowPlayingInfoLanguageOption](nc, objc.Sel("alloc"))
	return rv
}

func NowPlayingInfoLanguageOption_Alloc() NowPlayingInfoLanguageOption {
	return NowPlayingInfoLanguageOptionClass.Alloc()
}

func (nc _NowPlayingInfoLanguageOptionClass) New() NowPlayingInfoLanguageOption {
	rv := objc.Call[NowPlayingInfoLanguageOption](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNowPlayingInfoLanguageOption() NowPlayingInfoLanguageOption {
	return NowPlayingInfoLanguageOptionClass.New()
}

func (n_ NowPlayingInfoLanguageOption) Init() NowPlayingInfoLanguageOption {
	rv := objc.Call[NowPlayingInfoLanguageOption](n_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that determines whether to use the best legible language option based on the system preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption/1623144-isautomaticlegiblelanguageoption?language=objc
func (n_ NowPlayingInfoLanguageOption) IsAutomaticLegibleLanguageOption() bool {
	rv := objc.Call[bool](n_, objc.Sel("isAutomaticLegibleLanguageOption"))
	return rv
}

// Returns a Boolean value that determines whether to use the best audible language option based on the system preferences. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption/1623151-isautomaticaudiblelanguageoption?language=objc
func (n_ NowPlayingInfoLanguageOption) IsAutomaticAudibleLanguageOption() bool {
	rv := objc.Call[bool](n_, objc.Sel("isAutomaticAudibleLanguageOption"))
	return rv
}

// The type of language option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption/1623153-languageoptiontype?language=objc
func (n_ NowPlayingInfoLanguageOption) LanguageOptionType() NowPlayingInfoLanguageOptionType {
	rv := objc.Call[NowPlayingInfoLanguageOptionType](n_, objc.Sel("languageOptionType"))
	return rv
}

// The abbreviated language code for the language option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption/1623160-languagetag?language=objc
func (n_ NowPlayingInfoLanguageOption) LanguageTag() string {
	rv := objc.Call[string](n_, objc.Sel("languageTag"))
	return rv
}

// The characteristics that describe the content of the language option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption/1623152-languageoptioncharacteristics?language=objc
func (n_ NowPlayingInfoLanguageOption) LanguageOptionCharacteristics() []string {
	rv := objc.Call[[]string](n_, objc.Sel("languageOptionCharacteristics"))
	return rv
}

// The display name for a language option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption/1623145-displayname?language=objc
func (n_ NowPlayingInfoLanguageOption) DisplayName() string {
	rv := objc.Call[string](n_, objc.Sel("displayName"))
	return rv
}

// The unique identifier for the language option. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpnowplayinginfolanguageoption/1623135-identifier?language=objc
func (n_ NowPlayingInfoLanguageOption) Identifier() string {
	rv := objc.Call[string](n_, objc.Sel("identifier"))
	return rv
}
