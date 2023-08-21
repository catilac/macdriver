// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptCommandDescription] class.
var ScriptCommandDescriptionClass = _ScriptCommandDescriptionClass{objc.GetClass("NSScriptCommandDescription")}

type _ScriptCommandDescriptionClass struct {
	objc.Class
}

// An interface definition for the [ScriptCommandDescription] class.
type IScriptCommandDescription interface {
	objc.IObject
	IsOptionalArgumentWithName(argumentName string) bool
	CreateCommandInstance() ScriptCommand
	CreateCommandInstanceWithZone(zone unsafe.Pointer) ScriptCommand
	TypeForArgumentWithName(argumentName string) string
	AppleEventCodeForArgumentWithName(argumentName string) uint
	ArgumentNames() []string
	AppleEventCode() uint
	AppleEventClassCode() uint
	AppleEventCodeForReturnType() uint
	SuiteName() string
	CommandClassName() string
	CommandName() string
	ReturnType() string
}

// A script command that a macOS app supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription?language=objc
type ScriptCommandDescription struct {
	objc.Object
}

func ScriptCommandDescriptionFrom(ptr unsafe.Pointer) ScriptCommandDescription {
	return ScriptCommandDescription{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ ScriptCommandDescription) InitWithSuiteNameCommandNameDictionary(suiteName string, commandName string, commandDeclaration Dictionary) ScriptCommandDescription {
	rv := objc.Call[ScriptCommandDescription](s_, objc.Sel("initWithSuiteName:commandName:dictionary:"), suiteName, commandName, commandDeclaration)
	return rv
}

// Initializes and returns a newly allocated instance of NSScriptCommandDescription. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1410038-initwithsuitename?language=objc
func NewScriptCommandDescriptionWithSuiteNameCommandNameDictionary(suiteName string, commandName string, commandDeclaration Dictionary) ScriptCommandDescription {
	instance := ScriptCommandDescriptionClass.Alloc().InitWithSuiteNameCommandNameDictionary(suiteName, commandName, commandDeclaration)
	instance.Autorelease()
	return instance
}

func (sc _ScriptCommandDescriptionClass) Alloc() ScriptCommandDescription {
	rv := objc.Call[ScriptCommandDescription](sc, objc.Sel("alloc"))
	return rv
}

func ScriptCommandDescription_Alloc() ScriptCommandDescription {
	return ScriptCommandDescriptionClass.Alloc()
}

func (sc _ScriptCommandDescriptionClass) New() ScriptCommandDescription {
	rv := objc.Call[ScriptCommandDescription](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptCommandDescription() ScriptCommandDescription {
	return ScriptCommandDescriptionClass.New()
}

func (s_ ScriptCommandDescription) Init() ScriptCommandDescription {
	rv := objc.Call[ScriptCommandDescription](s_, objc.Sel("init"))
	return rv
}

// Returns a Boolean value that indicates whether the command argument identified by the specified argument key is an optional argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1415798-isoptionalargumentwithname?language=objc
func (s_ ScriptCommandDescription) IsOptionalArgumentWithName(argumentName string) bool {
	rv := objc.Call[bool](s_, objc.Sel("isOptionalArgumentWithName:"), argumentName)
	return rv
}

// Creates and returns an instance of the command object described by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1418415-createcommandinstance?language=objc
func (s_ ScriptCommandDescription) CreateCommandInstance() ScriptCommand {
	rv := objc.Call[ScriptCommand](s_, objc.Sel("createCommandInstance"))
	return rv
}

// Creates and returns an instance of the command object described by the receiver in the specified memory zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1413755-createcommandinstancewithzone?language=objc
func (s_ ScriptCommandDescription) CreateCommandInstanceWithZone(zone unsafe.Pointer) ScriptCommand {
	rv := objc.Call[ScriptCommand](s_, objc.Sel("createCommandInstanceWithZone:"), zone)
	return rv
}

// Returns the type of the command argument identified by the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1416163-typeforargumentwithname?language=objc
func (s_ ScriptCommandDescription) TypeForArgumentWithName(argumentName string) string {
	rv := objc.Call[string](s_, objc.Sel("typeForArgumentWithName:"), argumentName)
	return rv
}

// Returns the Apple event code for the specified command argument of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1414752-appleeventcodeforargumentwithnam?language=objc
func (s_ ScriptCommandDescription) AppleEventCodeForArgumentWithName(argumentName string) uint {
	rv := objc.Call[uint](s_, objc.Sel("appleEventCodeForArgumentWithName:"), argumentName)
	return rv
}

// Returns the names (or keys) for all arguments of the receiver’s command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1409125-argumentnames?language=objc
func (s_ ScriptCommandDescription) ArgumentNames() []string {
	rv := objc.Call[[]string](s_, objc.Sel("argumentNames"))
	return rv
}

// Returns the four-character code for the Apple event ID of the receiver’s command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1408972-appleeventcode?language=objc
func (s_ ScriptCommandDescription) AppleEventCode() uint {
	rv := objc.Call[uint](s_, objc.Sel("appleEventCode"))
	return rv
}

// Returns the four-character code for the Apple event class of the receiver’s command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1416191-appleeventclasscode?language=objc
func (s_ ScriptCommandDescription) AppleEventClassCode() uint {
	rv := objc.Call[uint](s_, objc.Sel("appleEventClassCode"))
	return rv
}

// Returns the Apple event code that identifies the command’s return type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1408166-appleeventcodeforreturntype?language=objc
func (s_ ScriptCommandDescription) AppleEventCodeForReturnType() uint {
	rv := objc.Call[uint](s_, objc.Sel("appleEventCodeForReturnType"))
	return rv
}

// Returns the name of the suite that contains the command described by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1413657-suitename?language=objc
func (s_ ScriptCommandDescription) SuiteName() string {
	rv := objc.Call[string](s_, objc.Sel("suiteName"))
	return rv
}

// Returns the name of the class that will be instantiated to handle the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1417478-commandclassname?language=objc
func (s_ ScriptCommandDescription) CommandClassName() string {
	rv := objc.Call[string](s_, objc.Sel("commandClassName"))
	return rv
}

// Returns the name of the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1407512-commandname?language=objc
func (s_ ScriptCommandDescription) CommandName() string {
	rv := objc.Call[string](s_, objc.Sel("commandName"))
	return rv
}

// Returns the return type of the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsscriptcommanddescription/1410754-returntype?language=objc
func (s_ ScriptCommandDescription) ReturnType() string {
	rv := objc.Call[string](s_, objc.Sel("returnType"))
	return rv
}
