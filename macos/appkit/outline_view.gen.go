// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [OutlineView] class.
var OutlineViewClass = _OutlineViewClass{objc.GetClass("NSOutlineView")}

type _OutlineViewClass struct {
	objc.Class
}

// An interface definition for the [OutlineView] class.
type IOutlineView interface {
	ITableView
	SetDropItemDropChildIndex(item objc.IObject, index int)
	MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject)
	IsItemExpanded(item objc.IObject) bool
	RowForItem(item objc.IObject) int
	ExpandItem(item objc.IObject)
	LevelForItem(item objc.IObject) int
	ReloadItem(item objc.IObject)
	RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	ParentForItem(item objc.IObject) objc.Object
	CollapseItem(item objc.IObject)
	LevelForRow(row int) int
	NumberOfChildrenOfItem(item objc.IObject) int
	InsertItemsAtIndexesInParentWithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions)
	FrameOfOutlineCellAtRow(row int) foundation.Rect
	IsExpandable(item objc.IObject) bool
	ItemAtRow(row int) objc.Object
	ChildOfItem(index int, item objc.IObject) objc.Object
	ChildIndexForItem(item objc.IObject) int
	ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool
	AutosaveExpandedItems() bool
	SetAutosaveExpandedItems(value bool)
	AutoresizesOutlineColumn() bool
	SetAutoresizesOutlineColumn(value bool)
	StronglyReferencesItems() bool
	SetStronglyReferencesItems(value bool)
	IndentationMarkerFollowsCell() bool
	SetIndentationMarkerFollowsCell(value bool)
	OutlineTableColumn() TableColumn
	SetOutlineTableColumn(value ITableColumn)
	IndentationPerLevel() float64
	SetIndentationPerLevel(value float64)
}

// A view that uses a row-and-column format to display hierarchical data like directories and files that can be expanded and collapsed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview?language=objc
type OutlineView struct {
	TableView
}

func OutlineViewFrom(ptr unsafe.Pointer) OutlineView {
	return OutlineView{
		TableView: TableViewFrom(ptr),
	}
}

func (oc _OutlineViewClass) Alloc() OutlineView {
	rv := objc.Call[OutlineView](oc, objc.Sel("alloc"))
	return rv
}

func OutlineView_Alloc() OutlineView {
	return OutlineViewClass.Alloc()
}

func (oc _OutlineViewClass) New() OutlineView {
	rv := objc.Call[OutlineView](oc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewOutlineView() OutlineView {
	return OutlineViewClass.New()
}

func (o_ OutlineView) Init() OutlineView {
	rv := objc.Call[OutlineView](o_, objc.Sel("init"))
	return rv
}

func (o_ OutlineView) InitWithFrame(frameRect foundation.Rect) OutlineView {
	rv := objc.Call[OutlineView](o_, objc.Sel("initWithFrame:"), frameRect)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstableview/1525511-initwithframe?language=objc
func NewOutlineViewWithFrame(frameRect foundation.Rect) OutlineView {
	instance := OutlineViewClass.Alloc().InitWithFrame(frameRect)
	instance.Autorelease()
	return instance
}

// Used to “retarget” a proposed drop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1525351-setdropitem?language=objc
func (o_ OutlineView) SetDropItemDropChildIndex(item objc.IObject, index int) {
	objc.Call[objc.Void](o_, objc.Sel("setDropItem:dropChildIndex:"), item, index)
}

// Moves an item at a given index in the given parent to a new index in a new parent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1530467-moveitematindex?language=objc
func (o_ OutlineView) MoveItemAtIndexInParentToIndexInParent(fromIndex int, oldParent objc.IObject, toIndex int, newParent objc.IObject) {
	objc.Call[objc.Void](o_, objc.Sel("moveItemAtIndex:inParent:toIndex:inParent:"), fromIndex, oldParent, toIndex, newParent)
}

// Returns a Boolean value that indicates whether a given item is expanded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1525355-isitemexpanded?language=objc
func (o_ OutlineView) IsItemExpanded(item objc.IObject) bool {
	rv := objc.Call[bool](o_, objc.Sel("isItemExpanded:"), item)
	return rv
}

// Returns the row associated with a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1533885-rowforitem?language=objc
func (o_ OutlineView) RowForItem(item objc.IObject) int {
	rv := objc.Call[int](o_, objc.Sel("rowForItem:"), item)
	return rv
}

// Expands a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1535323-expanditem?language=objc
func (o_ OutlineView) ExpandItem(item objc.IObject) {
	objc.Call[objc.Void](o_, objc.Sel("expandItem:"), item)
}

// Returns the indentation level for a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1535891-levelforitem?language=objc
func (o_ OutlineView) LevelForItem(item objc.IObject) int {
	rv := objc.Call[int](o_, objc.Sel("levelForItem:"), item)
	return rv
}

// Reloads and redisplays the data for the given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1531263-reloaditem?language=objc
func (o_ OutlineView) ReloadItem(item objc.IObject) {
	objc.Call[objc.Void](o_, objc.Sel("reloadItem:"), item)
}

// Removes items at the given indexes in the given parent with the specified optional animations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1527168-removeitemsatindexes?language=objc
func (o_ OutlineView) RemoveItemsAtIndexesInParentWithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	objc.Call[objc.Void](o_, objc.Sel("removeItemsAtIndexes:inParent:withAnimation:"), objc.Ptr(indexes), parent, animationOptions)
}

// Returns the parent for a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1535031-parentforitem?language=objc
func (o_ OutlineView) ParentForItem(item objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("parentForItem:"), item)
	return rv
}

// Collapses a given item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1532295-collapseitem?language=objc
func (o_ OutlineView) CollapseItem(item objc.IObject) {
	objc.Call[objc.Void](o_, objc.Sel("collapseItem:"), item)
}

// Returns the indentation level for a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1531449-levelforrow?language=objc
func (o_ OutlineView) LevelForRow(row int) int {
	rv := objc.Call[int](o_, objc.Sel("levelForRow:"), row)
	return rv
}

// Returns the number of children for the specified parent item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1534304-numberofchildrenofitem?language=objc
func (o_ OutlineView) NumberOfChildrenOfItem(item objc.IObject) int {
	rv := objc.Call[int](o_, objc.Sel("numberOfChildrenOfItem:"), item)
	return rv
}

// Inserts new items at the given indexes in the given parent with the specified optional animations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1528656-insertitemsatindexes?language=objc
func (o_ OutlineView) InsertItemsAtIndexesInParentWithAnimation(indexes foundation.IIndexSet, parent objc.IObject, animationOptions TableViewAnimationOptions) {
	objc.Call[objc.Void](o_, objc.Sel("insertItemsAtIndexes:inParent:withAnimation:"), objc.Ptr(indexes), parent, animationOptions)
}

// Returns the frame of the outline cell for a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1532314-frameofoutlinecellatrow?language=objc
func (o_ OutlineView) FrameOfOutlineCellAtRow(row int) foundation.Rect {
	rv := objc.Call[foundation.Rect](o_, objc.Sel("frameOfOutlineCellAtRow:"), row)
	return rv
}

// Returns a Boolean value that indicates whether a given item is expandable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1533676-isexpandable?language=objc
func (o_ OutlineView) IsExpandable(item objc.IObject) bool {
	rv := objc.Call[bool](o_, objc.Sel("isExpandable:"), item)
	return rv
}

// Returns the item associated with a given row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1534846-itematrow?language=objc
func (o_ OutlineView) ItemAtRow(row int) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("itemAtRow:"), row)
	return rv
}

// Returns the specified child of an item. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1527501-child?language=objc
func (o_ OutlineView) ChildOfItem(index int, item objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](o_, objc.Sel("child:ofItem:"), index, item)
	return rv
}

// Returns the child index of the specified item within its parent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1529954-childindexforitem?language=objc
func (o_ OutlineView) ChildIndexForItem(item objc.IObject) int {
	rv := objc.Call[int](o_, objc.Sel("childIndexForItem:"), item)
	return rv
}

// Returns a Boolean value that indicates whether auto-expanded items should return to their original collapsed state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1535362-shouldcollapseautoexpandeditemsf?language=objc
func (o_ OutlineView) ShouldCollapseAutoExpandedItemsForDeposited(deposited bool) bool {
	rv := objc.Call[bool](o_, objc.Sel("shouldCollapseAutoExpandedItemsForDeposited:"), deposited)
	return rv
}

// A Boolean value indicating whether the expanded items are automatically saved across launches of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1530638-autosaveexpandeditems?language=objc
func (o_ OutlineView) AutosaveExpandedItems() bool {
	rv := objc.Call[bool](o_, objc.Sel("autosaveExpandedItems"))
	return rv
}

// A Boolean value indicating whether the expanded items are automatically saved across launches of the app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1530638-autosaveexpandeditems?language=objc
func (o_ OutlineView) SetAutosaveExpandedItems(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setAutosaveExpandedItems:"), value)
}

// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1532304-autoresizesoutlinecolumn?language=objc
func (o_ OutlineView) AutoresizesOutlineColumn() bool {
	rv := objc.Call[bool](o_, objc.Sel("autoresizesOutlineColumn"))
	return rv
}

// A Boolean value that indicates whether the outline view resizes its outline column when the user expands or collapses items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1532304-autoresizesoutlinecolumn?language=objc
func (o_ OutlineView) SetAutoresizesOutlineColumn(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setAutoresizesOutlineColumn:"), value)
}

// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1644626-stronglyreferencesitems?language=objc
func (o_ OutlineView) StronglyReferencesItems() bool {
	rv := objc.Call[bool](o_, objc.Sel("stronglyReferencesItems"))
	return rv
}

// A Boolean value that indicates whether the outline view retains and releases the objects returned from its data source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1644626-stronglyreferencesitems?language=objc
func (o_ OutlineView) SetStronglyReferencesItems(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setStronglyReferencesItems:"), value)
}

// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1531707-indentationmarkerfollowscell?language=objc
func (o_ OutlineView) IndentationMarkerFollowsCell() bool {
	rv := objc.Call[bool](o_, objc.Sel("indentationMarkerFollowsCell"))
	return rv
}

// A Boolean value indicating whether the indentation marker symbol displayed in the outline column should be indented along with the cell contents. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1531707-indentationmarkerfollowscell?language=objc
func (o_ OutlineView) SetIndentationMarkerFollowsCell(value bool) {
	objc.Call[objc.Void](o_, objc.Sel("setIndentationMarkerFollowsCell:"), value)
}

// The table column in which hierarchical data is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1533581-outlinetablecolumn?language=objc
func (o_ OutlineView) OutlineTableColumn() TableColumn {
	rv := objc.Call[TableColumn](o_, objc.Sel("outlineTableColumn"))
	return rv
}

// The table column in which hierarchical data is displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1533581-outlinetablecolumn?language=objc
func (o_ OutlineView) SetOutlineTableColumn(value ITableColumn) {
	objc.Call[objc.Void](o_, objc.Sel("setOutlineTableColumn:"), objc.Ptr(value))
}

// The per-level indentation, measured in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1535181-indentationperlevel?language=objc
func (o_ OutlineView) IndentationPerLevel() float64 {
	rv := objc.Call[float64](o_, objc.Sel("indentationPerLevel"))
	return rv
}

// The per-level indentation, measured in points. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsoutlineview/1535181-indentationperlevel?language=objc
func (o_ OutlineView) SetIndentationPerLevel(value float64) {
	objc.Call[objc.Void](o_, objc.Sel("setIndentationPerLevel:"), value)
}
