// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MenuItem] class.
var MenuItemClass = _MenuItemClass{objc.GetClass("NSMenuItem")}

type _MenuItemClass struct {
	objc.Class
}

// An interface definition for the [MenuItem] class.
type IMenuItem interface {
	objc.IObject
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	OffStateImage() Image
	SetOffStateImage(value IImage)
	IsHighlighted() bool
	UserKeyEquivalent() string
	Target() objc.Object
	SetTarget(value objc.IObject)
	State() ControlStateValue
	SetState(value ControlStateValue)
	IsHidden() bool
	SetHidden(value bool)
	IndentationLevel() int
	SetIndentationLevel(value int)
	Action() objc.Selector
	SetAction(value objc.Selector)
	ToolTip() string
	SetToolTip(value string)
	IsAlternate() bool
	SetAlternate(value bool)
	View() View
	SetView(value IView)
	AllowsAutomaticKeyEquivalentLocalization() bool
	SetAllowsAutomaticKeyEquivalentLocalization(value bool)
	AttributedTitle() foundation.AttributedString
	SetAttributedTitle(value foundation.IAttributedString)
	OnStateImage() Image
	SetOnStateImage(value IImage)
	AllowsKeyEquivalentWhenHidden() bool
	SetAllowsKeyEquivalentWhenHidden(value bool)
	AllowsAutomaticKeyEquivalentMirroring() bool
	SetAllowsAutomaticKeyEquivalentMirroring(value bool)
	Menu() Menu
	SetMenu(value IMenu)
	IsSeparatorItem() bool
	RepresentedObject() objc.Object
	SetRepresentedObject(value objc.IObject)
	Tag() int
	SetTag(value int)
	Title() string
	SetTitle(value string)
	MixedStateImage() Image
	SetMixedStateImage(value IImage)
	Submenu() Menu
	SetSubmenu(value IMenu)
	HasSubmenu() bool
	KeyEquivalentModifierMask() EventModifierFlags
	SetKeyEquivalentModifierMask(value EventModifierFlags)
	IsHiddenOrHasHiddenAncestor() bool
	ParentItem() MenuItem
	IsEnabled() bool
	SetEnabled(value bool)
	Image() Image
	SetImage(value IImage)
}

// A command item in an app menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem?language=objc
type MenuItem struct {
	objc.Object
}

func MenuItemFrom(ptr unsafe.Pointer) MenuItem {
	return MenuItem{
		Object: objc.ObjectFrom(ptr),
	}
}

func (m_ MenuItem) InitWithTitleActionKeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("initWithTitle:action:keyEquivalent:"), string_, selector, charCode)
	return rv
}

// Returns an initialized instance of NSMenuItem. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514858-initwithtitle?language=objc
func NewMenuItemWithTitleActionKeyEquivalent(string_ string, selector objc.Selector, charCode string) MenuItem {
	instance := MenuItemClass.Alloc().InitWithTitleActionKeyEquivalent(string_, selector, charCode)
	instance.Autorelease()
	return instance
}

func (mc _MenuItemClass) Alloc() MenuItem {
	rv := objc.Call[MenuItem](mc, objc.Sel("alloc"))
	return rv
}

func MenuItem_Alloc() MenuItem {
	return MenuItemClass.Alloc()
}

func (mc _MenuItemClass) New() MenuItem {
	rv := objc.Call[MenuItem](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMenuItem() MenuItem {
	return MenuItemClass.New()
}

func (m_ MenuItem) Init() MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("init"))
	return rv
}

// The menu item’s unmodified key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514842-keyequivalent?language=objc
func (m_ MenuItem) KeyEquivalent() string {
	rv := objc.Call[string](m_, objc.Sel("keyEquivalent"))
	return rv
}

// The menu item’s unmodified key equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514842-keyequivalent?language=objc
func (m_ MenuItem) SetKeyEquivalent(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setKeyEquivalent:"), value)
}

// The image of the menu item that indicates an “off” state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514821-offstateimage?language=objc
func (m_ MenuItem) OffStateImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("offStateImage"))
	return rv
}

// The image of the menu item that indicates an “off” state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514821-offstateimage?language=objc
func (m_ MenuItem) SetOffStateImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setOffStateImage:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the menu item should be drawn highlighted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514856-highlighted?language=objc
func (m_ MenuItem) IsHighlighted() bool {
	rv := objc.Call[bool](m_, objc.Sel("isHighlighted"))
	return rv
}

// The user-assigned key equivalent for the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514850-userkeyequivalent?language=objc
func (m_ MenuItem) UserKeyEquivalent() string {
	rv := objc.Call[string](m_, objc.Sel("userKeyEquivalent"))
	return rv
}

// The menu item's target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc
func (m_ MenuItem) Target() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("target"))
	return rv
}

// The menu item's target. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514843-target?language=objc
func (m_ MenuItem) SetTarget(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setTarget:"), value)
}

// The state of the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc
func (m_ MenuItem) State() ControlStateValue {
	rv := objc.Call[ControlStateValue](m_, objc.Sel("state"))
	return rv
}

// The state of the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514804-state?language=objc
func (m_ MenuItem) SetState(value ControlStateValue) {
	objc.Call[objc.Void](m_, objc.Sel("setState:"), value)
}

// A Boolean value that indicates whether the menu item is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc
func (m_ MenuItem) IsHidden() bool {
	rv := objc.Call[bool](m_, objc.Sel("isHidden"))
	return rv
}

// A Boolean value that indicates whether the menu item is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514846-hidden?language=objc
func (m_ MenuItem) SetHidden(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setHidden:"), value)
}

// The menu item indentation level for the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514809-indentationlevel?language=objc
func (m_ MenuItem) IndentationLevel() int {
	rv := objc.Call[int](m_, objc.Sel("indentationLevel"))
	return rv
}

// The menu item indentation level for the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514809-indentationlevel?language=objc
func (m_ MenuItem) SetIndentationLevel(value int) {
	objc.Call[objc.Void](m_, objc.Sel("setIndentationLevel:"), value)
}

// The menu item's action-method selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc
func (m_ MenuItem) Action() objc.Selector {
	rv := objc.Call[objc.Selector](m_, objc.Sel("action"))
	return rv
}

// The menu item's action-method selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514825-action?language=objc
func (m_ MenuItem) SetAction(value objc.Selector) {
	objc.Call[objc.Void](m_, objc.Sel("setAction:"), value)
}

// A help tag for the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc
func (m_ MenuItem) ToolTip() string {
	rv := objc.Call[string](m_, objc.Sel("toolTip"))
	return rv
}

// A help tag for the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514848-tooltip?language=objc
func (m_ MenuItem) SetToolTip(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setToolTip:"), value)
}

// A Boolean value that marks the menu item as an alternate to the previous menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514823-alternate?language=objc
func (m_ MenuItem) IsAlternate() bool {
	rv := objc.Call[bool](m_, objc.Sel("isAlternate"))
	return rv
}

// A Boolean value that marks the menu item as an alternate to the previous menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514823-alternate?language=objc
func (m_ MenuItem) SetAlternate(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAlternate:"), value)
}

// The content view for the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514835-view?language=objc
func (m_ MenuItem) View() View {
	rv := objc.Call[View](m_, objc.Sel("view"))
	return rv
}

// The content view for the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514835-view?language=objc
func (m_ MenuItem) SetView(value IView) {
	objc.Call[objc.Void](m_, objc.Sel("setView:"), objc.Ptr(value))
}

// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/3787554-allowsautomatickeyequivalentloca?language=objc
func (m_ MenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	rv := objc.Call[bool](m_, objc.Sel("allowsAutomaticKeyEquivalentLocalization"))
	return rv
}

// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/3787554-allowsautomatickeyequivalentloca?language=objc
func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentLocalization(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsAutomaticKeyEquivalentLocalization:"), value)
}

// A custom string for a menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514860-attributedtitle?language=objc
func (m_ MenuItem) AttributedTitle() foundation.AttributedString {
	rv := objc.Call[foundation.AttributedString](m_, objc.Sel("attributedTitle"))
	return rv
}

// A custom string for a menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514860-attributedtitle?language=objc
func (m_ MenuItem) SetAttributedTitle(value foundation.IAttributedString) {
	objc.Call[objc.Void](m_, objc.Sel("setAttributedTitle:"), objc.Ptr(value))
}

// The image of the menu item that indicates an “on” state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514861-onstateimage?language=objc
func (m_ MenuItem) OnStateImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("onStateImage"))
	return rv
}

// The image of the menu item that indicates an “on” state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514861-onstateimage?language=objc
func (m_ MenuItem) SetOnStateImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setOnStateImage:"), objc.Ptr(value))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/2880316-allowskeyequivalentwhenhidden?language=objc
func (m_ MenuItem) AllowsKeyEquivalentWhenHidden() bool {
	rv := objc.Call[bool](m_, objc.Sel("allowsKeyEquivalentWhenHidden"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/2880316-allowskeyequivalentwhenhidden?language=objc
func (m_ MenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsKeyEquivalentWhenHidden:"), value)
}

// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/3787555-allowsautomatickeyequivalentmirr?language=objc
func (m_ MenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	rv := objc.Call[bool](m_, objc.Sel("allowsAutomaticKeyEquivalentMirroring"))
	return rv
}

// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/3787555-allowsautomatickeyequivalentmirr?language=objc
func (m_ MenuItem) SetAllowsAutomaticKeyEquivalentMirroring(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setAllowsAutomaticKeyEquivalentMirroring:"), value)
}

// The menu item’s menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514830-menu?language=objc
func (m_ MenuItem) Menu() Menu {
	rv := objc.Call[Menu](m_, objc.Sel("menu"))
	return rv
}

// The menu item’s menu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514830-menu?language=objc
func (m_ MenuItem) SetMenu(value IMenu) {
	objc.Call[objc.Void](m_, objc.Sel("setMenu:"), objc.Ptr(value))
}

// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514811-usesuserkeyequivalents?language=objc
func (mc _MenuItemClass) UsesUserKeyEquivalents() bool {
	rv := objc.Call[bool](mc, objc.Sel("usesUserKeyEquivalents"))
	return rv
}

// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514811-usesuserkeyequivalents?language=objc
func MenuItem_UsesUserKeyEquivalents() bool {
	return MenuItemClass.UsesUserKeyEquivalents()
}

// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514811-usesuserkeyequivalents?language=objc
func (mc _MenuItemClass) SetUsesUserKeyEquivalents(value bool) {
	objc.Call[objc.Void](mc, objc.Sel("setUsesUserKeyEquivalents:"), value)
}

// Returns a Boolean value that indicates whether menu items conform to user preferences for key equivalents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514811-usesuserkeyequivalents?language=objc
func MenuItem_SetUsesUserKeyEquivalents(value bool) {
	MenuItemClass.SetUsesUserKeyEquivalents(value)
}

// A Boolean value indicating whether the menu item is a separator item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514837-separatoritem?language=objc
func (m_ MenuItem) IsSeparatorItem() bool {
	rv := objc.Call[bool](m_, objc.Sel("isSeparatorItem"))
	return rv
}

// The object represented by the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514834-representedobject?language=objc
func (m_ MenuItem) RepresentedObject() objc.Object {
	rv := objc.Call[objc.Object](m_, objc.Sel("representedObject"))
	return rv
}

// The object represented by the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514834-representedobject?language=objc
func (m_ MenuItem) SetRepresentedObject(value objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("setRepresentedObject:"), value)
}

// The menu item's tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514840-tag?language=objc
func (m_ MenuItem) Tag() int {
	rv := objc.Call[int](m_, objc.Sel("tag"))
	return rv
}

// The menu item's tag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514840-tag?language=objc
func (m_ MenuItem) SetTag(value int) {
	objc.Call[objc.Void](m_, objc.Sel("setTag:"), value)
}

// The menu item's title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc
func (m_ MenuItem) Title() string {
	rv := objc.Call[string](m_, objc.Sel("title"))
	return rv
}

// The menu item's title. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514805-title?language=objc
func (m_ MenuItem) SetTitle(value string) {
	objc.Call[objc.Void](m_, objc.Sel("setTitle:"), value)
}

// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514827-mixedstateimage?language=objc
func (m_ MenuItem) MixedStateImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("mixedStateImage"))
	return rv
}

// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514827-mixedstateimage?language=objc
func (m_ MenuItem) SetMixedStateImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setMixedStateImage:"), objc.Ptr(value))
}

// The submenu of the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514845-submenu?language=objc
func (m_ MenuItem) Submenu() Menu {
	rv := objc.Call[Menu](m_, objc.Sel("submenu"))
	return rv
}

// The submenu of the menu item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514845-submenu?language=objc
func (m_ MenuItem) SetSubmenu(value IMenu) {
	objc.Call[objc.Void](m_, objc.Sel("setSubmenu:"), objc.Ptr(value))
}

// A Boolean value that indicates whether the menu item has a submenu. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514817-hassubmenu?language=objc
func (m_ MenuItem) HasSubmenu() bool {
	rv := objc.Call[bool](m_, objc.Sel("hasSubmenu"))
	return rv
}

// The menu item’s keyboard equivalent modifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514815-keyequivalentmodifiermask?language=objc
func (m_ MenuItem) KeyEquivalentModifierMask() EventModifierFlags {
	rv := objc.Call[EventModifierFlags](m_, objc.Sel("keyEquivalentModifierMask"))
	return rv
}

// The menu item’s keyboard equivalent modifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514815-keyequivalentmodifiermask?language=objc
func (m_ MenuItem) SetKeyEquivalentModifierMask(value EventModifierFlags) {
	objc.Call[objc.Void](m_, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// A Boolean value that indicates whether the menu item or any of its superitems is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514832-hiddenorhashiddenancestor?language=objc
func (m_ MenuItem) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.Call[bool](m_, objc.Sel("isHiddenOrHasHiddenAncestor"))
	return rv
}

// The menu item whose submenu contains the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514813-parentitem?language=objc
func (m_ MenuItem) ParentItem() MenuItem {
	rv := objc.Call[MenuItem](m_, objc.Sel("parentItem"))
	return rv
}

// A Boolean value that indicates whether the menu item is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc
func (m_ MenuItem) IsEnabled() bool {
	rv := objc.Call[bool](m_, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether the menu item is enabled. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514863-enabled?language=objc
func (m_ MenuItem) SetEnabled(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setEnabled:"), value)
}

// The menu item’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514819-image?language=objc
func (m_ MenuItem) Image() Image {
	rv := objc.Call[Image](m_, objc.Sel("image"))
	return rv
}

// The menu item’s image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsmenuitem/1514819-image?language=objc
func (m_ MenuItem) SetImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setImage:"), objc.Ptr(value))
}
