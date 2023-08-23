// AUTO-GENERATED CODE, DO NOT MODIFY

package coremediaio

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionProviderProperties] class.
var ExtensionProviderPropertiesClass = _ExtensionProviderPropertiesClass{objc.GetClass("CMIOExtensionProviderProperties")}

type _ExtensionProviderPropertiesClass struct {
	objc.Class
}

// An interface definition for the [ExtensionProviderProperties] class.
type IExtensionProviderProperties interface {
	objc.IObject
	SetPropertyStateForProperty(propertyState IExtensionPropertyState, property ExtensionProperty)
	Name() string
	SetName(value string)
	Manufacturer() string
	SetManufacturer(value string)
	PropertiesDictionary() map[ExtensionProperty]ExtensionPropertyState
	SetPropertiesDictionary(value map[ExtensionProperty]IExtensionPropertyState)
}

// An object that manages the properties of an extension provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties?language=objc
type ExtensionProviderProperties struct {
	objc.Object
}

func ExtensionProviderPropertiesFrom(ptr unsafe.Pointer) ExtensionProviderProperties {
	return ExtensionProviderProperties{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExtensionProviderPropertiesClass) ProviderPropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionProviderProperties {
	rv := objc.Call[ExtensionProviderProperties](ec, objc.Sel("providerPropertiesWithDictionary:"), propertiesDictionary)
	return rv
}

// Returns a new provider properties object with the specified properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915921-providerpropertieswithdictionary?language=objc
func ExtensionProviderProperties_ProviderPropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionProviderProperties {
	return ExtensionProviderPropertiesClass.ProviderPropertiesWithDictionary(propertiesDictionary)
}

func (e_ ExtensionProviderProperties) InitWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionProviderProperties {
	rv := objc.Call[ExtensionProviderProperties](e_, objc.Sel("initWithDictionary:"), propertiesDictionary)
	return rv
}

// Creates a provider properties object with the specified properties. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915917-initwithdictionary?language=objc
func NewExtensionProviderPropertiesWithDictionary(propertiesDictionary map[ExtensionProperty]IExtensionPropertyState) ExtensionProviderProperties {
	instance := ExtensionProviderPropertiesClass.Alloc().InitWithDictionary(propertiesDictionary)
	instance.Autorelease()
	return instance
}

func (ec _ExtensionProviderPropertiesClass) Alloc() ExtensionProviderProperties {
	rv := objc.Call[ExtensionProviderProperties](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionProviderProperties_Alloc() ExtensionProviderProperties {
	return ExtensionProviderPropertiesClass.Alloc()
}

func (ec _ExtensionProviderPropertiesClass) New() ExtensionProviderProperties {
	rv := objc.Call[ExtensionProviderProperties](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionProviderProperties() ExtensionProviderProperties {
	return ExtensionProviderPropertiesClass.New()
}

func (e_ ExtensionProviderProperties) Init() ExtensionProviderProperties {
	rv := objc.Call[ExtensionProviderProperties](e_, objc.Sel("init"))
	return rv
}

// Sets a state value for the specified property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915922-setpropertystate?language=objc
func (e_ ExtensionProviderProperties) SetPropertyStateForProperty(propertyState IExtensionPropertyState, property ExtensionProperty) {
	objc.Call[objc.Void](e_, objc.Sel("setPropertyState:forProperty:"), objc.Ptr(propertyState), property)
}

// The provider name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915919-name?language=objc
func (e_ ExtensionProviderProperties) Name() string {
	rv := objc.Call[string](e_, objc.Sel("name"))
	return rv
}

// The provider name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915919-name?language=objc
func (e_ ExtensionProviderProperties) SetName(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setName:"), value)
}

// The provider manufacturer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915918-manufacturer?language=objc
func (e_ ExtensionProviderProperties) Manufacturer() string {
	rv := objc.Call[string](e_, objc.Sel("manufacturer"))
	return rv
}

// The provider manufacturer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915918-manufacturer?language=objc
func (e_ ExtensionProviderProperties) SetManufacturer(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setManufacturer:"), value)
}

// A dictionary of properties for a provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915920-propertiesdictionary?language=objc
func (e_ ExtensionProviderProperties) PropertiesDictionary() map[ExtensionProperty]ExtensionPropertyState {
	rv := objc.Call[map[ExtensionProperty]ExtensionPropertyState](e_, objc.Sel("propertiesDictionary"))
	return rv
}

// A dictionary of properties for a provider. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremediaio/cmioextensionproviderproperties/3915920-propertiesdictionary?language=objc
func (e_ ExtensionProviderProperties) SetPropertiesDictionary(value map[ExtensionProperty]IExtensionPropertyState) {
	objc.Call[objc.Void](e_, objc.Sel("setPropertiesDictionary:"), value)
}
