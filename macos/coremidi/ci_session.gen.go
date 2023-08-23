// AUTO-GENERATED CODE, DO NOT MODIFY

package coremidi

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CISession] class.
var CISessionClass = _CISessionClass{objc.GetClass("MIDICISession")}

type _CISessionClass struct {
	objc.Class
}

// An interface definition for the [CISession] class.
type ICISession interface {
	objc.IObject
	EnableProfileOnChannelError(profile ICIProfile, channel ChannelNumber, outError foundation.IError) bool
	ProfileStateForChannel(channel ChannelNumber) CIProfileState
	DisableProfileOnChannelError(profile ICIProfile, channel ChannelNumber, outError foundation.IError) bool
	SendProfileOnChannelProfileData(profile ICIProfile, channel ChannelNumber, profileSpecificData []byte) bool
	MaxPropertyRequests() foundation.Number
	MaxSysExSize() foundation.Number
	DeviceInfo() CIDeviceInfo
	ProfileChangedCallback() CIProfileChangedBlock
	SetProfileChangedCallback(value CIProfileChangedBlock)
	SupportsPropertyCapability() bool
	MidiDestination() EntityRef
	ProfileSpecificDataHandler() CIProfileSpecificDataBlock
	SetProfileSpecificDataHandler(value CIProfileSpecificDataBlock)
	SupportsProfileCapability() bool
}

// An object that represents a MIDI-CI session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession?language=objc
type CISession struct {
	objc.Object
}

func CISessionFrom(ptr unsafe.Pointer) CISession {
	return CISession{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ CISession) InitWithDiscoveredNodeDataReadyHandlerDisconnectHandler(discoveredNode ICIDiscoveredNode, handler func(), disconnectHandler CISessionDisconnectBlock) CISession {
	rv := objc.Call[CISession](c_, objc.Sel("initWithDiscoveredNode:dataReadyHandler:disconnectHandler:"), objc.Ptr(discoveredNode), handler, disconnectHandler)
	return rv
}

// Creates a MIDI-CI session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/3580335-initwithdiscoverednode?language=objc
func NewCISessionWithDiscoveredNodeDataReadyHandlerDisconnectHandler(discoveredNode ICIDiscoveredNode, handler func(), disconnectHandler CISessionDisconnectBlock) CISession {
	instance := CISessionClass.Alloc().InitWithDiscoveredNodeDataReadyHandlerDisconnectHandler(discoveredNode, handler, disconnectHandler)
	instance.Autorelease()
	return instance
}

func (cc _CISessionClass) Alloc() CISession {
	rv := objc.Call[CISession](cc, objc.Sel("alloc"))
	return rv
}

func CISession_Alloc() CISession {
	return CISessionClass.Alloc()
}

func (cc _CISessionClass) New() CISession {
	rv := objc.Call[CISession](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCISession() CISession {
	return CISessionClass.New()
}

func (c_ CISession) Init() CISession {
	rv := objc.Call[CISession](c_, objc.Sel("init"))
	return rv
}

// Performs an asynchronous request to enable a profile for a specific MIDI channel number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/2977117-enableprofile?language=objc
func (c_ CISession) EnableProfileOnChannelError(profile ICIProfile, channel ChannelNumber, outError foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("enableProfile:onChannel:error:"), objc.Ptr(profile), channel, objc.Ptr(outError))
	return rv
}

// Returns the profile state for the specified MIDI channel number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/2977123-profilestateforchannel?language=objc
func (c_ CISession) ProfileStateForChannel(channel ChannelNumber) CIProfileState {
	rv := objc.Call[CIProfileState](c_, objc.Sel("profileStateForChannel:"), channel)
	return rv
}

// Performs an asynchronous request to disable a profile for a specific MIDI channel number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/2977116-disableprofile?language=objc
func (c_ CISession) DisableProfileOnChannelError(profile ICIProfile, channel ChannelNumber, outError foundation.IError) bool {
	rv := objc.Call[bool](c_, objc.Sel("disableProfile:onChannel:error:"), objc.Ptr(profile), channel, objc.Ptr(outError))
	return rv
}

// Sends profile-specific data to the MIDI-CI session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/3553276-sendprofile?language=objc
func (c_ CISession) SendProfileOnChannelProfileData(profile ICIProfile, channel ChannelNumber, profileSpecificData []byte) bool {
	rv := objc.Call[bool](c_, objc.Sel("sendProfile:onChannel:profileData:"), objc.Ptr(profile), channel, profileSpecificData)
	return rv
}

// The maximum number of simultaneous property exchange requests, if supported. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/3580336-maxpropertyrequests?language=objc
func (c_ CISession) MaxPropertyRequests() foundation.Number {
	rv := objc.Call[foundation.Number](c_, objc.Sel("maxPropertyRequests"))
	return rv
}

// The maximum size of System Exclusive (SysEx) messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/3580337-maxsysexsize?language=objc
func (c_ CISession) MaxSysExSize() foundation.Number {
	rv := objc.Call[foundation.Number](c_, objc.Sel("maxSysExSize"))
	return rv
}

// Information about a MIDI-CI device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/3553274-deviceinfo?language=objc
func (c_ CISession) DeviceInfo() CIDeviceInfo {
	rv := objc.Call[CIDeviceInfo](c_, objc.Sel("deviceInfo"))
	return rv
}

// An optional block the system calls after it enables or disables a profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/2977122-profilechangedcallback?language=objc
func (c_ CISession) ProfileChangedCallback() CIProfileChangedBlock {
	rv := objc.Call[CIProfileChangedBlock](c_, objc.Sel("profileChangedCallback"))
	return rv
}

// An optional block the system calls after it enables or disables a profile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/2977122-profilechangedcallback?language=objc
func (c_ CISession) SetProfileChangedCallback(value CIProfileChangedBlock) {
	objc.Call[objc.Void](c_, objc.Sel("setProfileChangedCallback:"), value)
}

// A Boolean value that indicates whether the entity supports the MIDI-CI property exchange capability. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/2977127-supportspropertycapability?language=objc
func (c_ CISession) SupportsPropertyCapability() bool {
	rv := objc.Call[bool](c_, objc.Sel("supportsPropertyCapability"))
	return rv
}

// The MIDI destination with which the session is communicating. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/3580338-mididestination?language=objc
func (c_ CISession) MidiDestination() EntityRef {
	rv := objc.Call[EntityRef](c_, objc.Sel("midiDestination"))
	return rv
}

// An optional block the system calls when a device sends profile-specific data to the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/3580339-profilespecificdatahandler?language=objc
func (c_ CISession) ProfileSpecificDataHandler() CIProfileSpecificDataBlock {
	rv := objc.Call[CIProfileSpecificDataBlock](c_, objc.Sel("profileSpecificDataHandler"))
	return rv
}

// An optional block the system calls when a device sends profile-specific data to the session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/3580339-profilespecificdatahandler?language=objc
func (c_ CISession) SetProfileSpecificDataHandler(value CIProfileSpecificDataBlock) {
	objc.Call[objc.Void](c_, objc.Sel("setProfileSpecificDataHandler:"), value)
}

// A Boolean value that indicates whether the entity supports the MIDI-CI profile’s capability. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midicisession/2977126-supportsprofilecapability?language=objc
func (c_ CISession) SupportsProfileCapability() bool {
	rv := objc.Call[bool](c_, objc.Sel("supportsProfileCapability"))
	return rv
}
