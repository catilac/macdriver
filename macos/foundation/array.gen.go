// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Array] class.
var ArrayClass = _ArrayClass{objc.GetClass("NSArray")}

type _ArrayClass struct {
	objc.Class
}

// An interface definition for the [Array] class.
type IArray interface {
	objc.IObject
	ObjectEnumerator() Enumerator
	SortedArrayUsingFunctionContextHint(comparator func(arg0 objc.Object, arg1 objc.Object, arg2 unsafe.Pointer) int, context unsafe.Pointer, hint []byte) []objc.Object
	IsEqualToArray(otherArray []objc.IObject) bool
	SortedArrayWithOptionsUsingComparator(opts SortOptions, cmptr Comparator) []objc.Object
	ObjectsAtIndexes(indexes IIndexSet) []objc.Object
	ShuffledArray() []objc.Object
	EnumerateObjectsUsingBlock(block func(obj objc.Object, idx uint, stop *bool))
	SortedArrayUsingDescriptors(sortDescriptors []ISortDescriptor) []objc.Object
	ArrayByAddingObjectsFromArray(otherArray []objc.IObject) []objc.Object
	IndexesOfObjectsWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet
	IndexOfObjectPassingTest(predicate func(obj objc.Object, idx uint, stop *bool) bool) uint
	SortedArrayUsingSelector(comparator objc.Selector) []objc.Object
	IndexOfObjectIdenticalToInRange(anObject objc.IObject, range_ Range) uint
	ObjectAtIndex(index uint) objc.Object
	FilteredArrayUsingPredicate(predicate IPredicate) []objc.Object
	ObjectAtIndexedSubscript(idx uint) objc.Object
	ArrayByApplyingDifference(difference IOrderedCollectionDifference) []objc.Object
	EnumerateObjectsWithOptionsUsingBlock(opts EnumerationOptions, block func(obj objc.Object, idx uint, stop *bool))
	GetObjectsRange(objects objc.IObject, range_ Range)
	InitWithContentsOfURLError(url IURL, error IError) []objc.Object
	ContainsObject(anObject objc.IObject) bool
	MakeObjectsPerformSelector(aSelector objc.Selector)
	ReverseObjectEnumerator() Enumerator
	WriteToURLError(url IURL, error IError) bool
	IndexesOfObjectsAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet
	RemoveObserverForKeyPath(observer objc.IObject, keyPath string)
	DescriptionWithLocale(locale objc.IObject) string
	SubarrayWithRange(range_ Range) []objc.Object
	DifferenceFromArray(other []objc.IObject) OrderedCollectionDifference
	SetValueForKey(value objc.IObject, key string)
	IndexesOfObjectsPassingTest(predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet
	ArrayByAddingObject(anObject objc.IObject) []objc.Object
	IndexOfObject(anObject objc.IObject) uint
	AddObserverForKeyPathOptionsContext(observer objc.IObject, keyPath string, options KeyValueObservingOptions, context unsafe.Pointer)
	IndexOfObjectWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) uint
	EnumerateObjectsAtIndexesOptionsUsingBlock(s IIndexSet, opts EnumerationOptions, block func(obj objc.Object, idx uint, stop *bool))
	ShuffledArrayWithRandomSource(randomSource objc.IObject) []objc.Object
	PathsMatchingExtensions(filterTypes []string) []string
	ComponentsJoinedByString(separator string) string
	ValueForKey(key string) objc.Object
	IndexOfObjectAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) uint
	FirstObjectCommonWithArray(otherArray []objc.IObject) objc.Object
	SortedArrayUsingComparator(cmptr Comparator) []objc.Object
	SortedArrayHint() []byte
	LastObject() objc.Object
	Description() string
	Count() uint
	FirstObject() objc.Object
}

// A static ordered collection of objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray?language=objc
type Array struct {
	objc.Object
}

func ArrayFrom(ptr unsafe.Pointer) Array {
	return Array{
		Object: objc.ObjectFrom(ptr),
	}
}

func (a_ Array) InitWithObjects(firstObj objc.IObject, args ...any) Array {
	rv := objc.Call[Array](a_, objc.Sel("initWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Initializes a newly allocated array by placing in it the objects in the argument list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460068-initwithobjects?language=objc
func NewArrayWithObjects(firstObj objc.IObject, args ...any) Array {
	instance := ArrayClass.Alloc().InitWithObjects(firstObj, args...)
	instance.Autorelease()
	return instance
}

func (ac _ArrayClass) Array() Array {
	rv := objc.Call[Array](ac, objc.Sel("array"))
	return rv
}

// Creates and returns an empty array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460120-array?language=objc
func Array_Array() Array {
	return ArrayClass.Array()
}

func (ac _ArrayClass) ArrayWithObjects(firstObj objc.IObject, args ...any) Array {
	rv := objc.Call[Array](ac, objc.Sel("arrayWithObjects:"), append([]any{objc.Ptr(firstObj)}, args...)...)
	return rv
}

// Creates and returns an array containing the objects in the argument list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460145-arraywithobjects?language=objc
func Array_ArrayWithObjects(firstObj objc.IObject, args ...any) Array {
	return ArrayClass.ArrayWithObjects(firstObj, args...)
}

func (a_ Array) InitWithArray(array []objc.IObject) Array {
	rv := objc.Call[Array](a_, objc.Sel("initWithArray:"), array)
	return rv
}

// Initializes a newly allocated array by placing in it the objects contained in a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1412169-initwitharray?language=objc
func NewArrayWithArray(array []objc.IObject) Array {
	instance := ArrayClass.Alloc().InitWithArray(array)
	instance.Autorelease()
	return instance
}

func (ac _ArrayClass) ArrayWithArray(array []objc.IObject) Array {
	rv := objc.Call[Array](ac, objc.Sel("arrayWithArray:"), array)
	return rv
}

// Creates and returns an array containing the objects in another given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460122-arraywitharray?language=objc
func Array_ArrayWithArray(array []objc.IObject) Array {
	return ArrayClass.ArrayWithArray(array)
}

func (ac _ArrayClass) ArrayWithObject(anObject objc.IObject) Array {
	rv := objc.Call[Array](ac, objc.Sel("arrayWithObject:"), objc.Ptr(anObject))
	return rv
}

// Creates and returns an array containing a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1411981-arraywithobject?language=objc
func Array_ArrayWithObject(anObject objc.IObject) Array {
	return ArrayClass.ArrayWithObject(anObject)
}

func (a_ Array) Init() Array {
	rv := objc.Call[Array](a_, objc.Sel("init"))
	return rv
}

func (ac _ArrayClass) Alloc() Array {
	rv := objc.Call[Array](ac, objc.Sel("alloc"))
	return rv
}

func Array_Alloc() Array {
	return ArrayClass.Alloc()
}

func (ac _ArrayClass) New() Array {
	rv := objc.Call[Array](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewArray() Array {
	return ArrayClass.New()
}

// Returns an enumerator object that lets you access each object in the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1416048-objectenumerator?language=objc
func (a_ Array) ObjectEnumerator() Enumerator {
	rv := objc.Call[Enumerator](a_, objc.Sel("objectEnumerator"))
	return rv
}

// Returns a new array that lists the receiving array’s elements in ascending order as defined by the comparison function comparator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1414839-sortedarrayusingfunction?language=objc
func (a_ Array) SortedArrayUsingFunctionContextHint(comparator func(arg0 objc.Object, arg1 objc.Object, arg2 unsafe.Pointer) int, context unsafe.Pointer, hint []byte) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("sortedArrayUsingFunction:context:hint:"), comparator, context, hint)
	return rv
}

// Compares the receiving array to another array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1411770-isequaltoarray?language=objc
func (a_ Array) IsEqualToArray(otherArray []objc.IObject) bool {
	rv := objc.Call[bool](a_, objc.Sel("isEqualToArray:"), otherArray)
	return rv
}

// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given NSComparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1417804-sortedarraywithoptions?language=objc
func (a_ Array) SortedArrayWithOptionsUsingComparator(opts SortOptions, cmptr Comparator) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("sortedArrayWithOptions:usingComparator:"), opts, cmptr)
	return rv
}

// Returns an array containing the objects in the array at the indexes specified by a given index set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1411296-objectsatindexes?language=objc
func (a_ Array) ObjectsAtIndexes(indexes IIndexSet) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("objectsAtIndexes:"), objc.Ptr(indexes))
	return rv
}

// Returns a new array that lists this array’s elements in a random order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1640855-shuffledarray?language=objc
func (a_ Array) ShuffledArray() []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("shuffledArray"))
	return rv
}

// Executes a given block using each object in the array, starting with the first object and continuing through the array to the last object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1415846-enumerateobjectsusingblock?language=objc
func (a_ Array) EnumerateObjectsUsingBlock(block func(obj objc.Object, idx uint, stop *bool)) {
	objc.Call[objc.Void](a_, objc.Sel("enumerateObjectsUsingBlock:"), block)
}

// Returns a copy of the receiving array sorted as specified by a given array of sort descriptors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1415069-sortedarrayusingdescriptors?language=objc
func (a_ Array) SortedArrayUsingDescriptors(sortDescriptors []ISortDescriptor) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("sortedArrayUsingDescriptors:"), sortDescriptors)
	return rv
}

// Returns a new array that is a copy of the receiving array with the objects contained in another array added to the end. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1412087-arraybyaddingobjectsfromarray?language=objc
func (a_ Array) ArrayByAddingObjectsFromArray(otherArray []objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("arrayByAddingObjectsFromArray:"), otherArray)
	return rv
}

// Returns the indexes of objects in the array that pass a test in a given block for a given set of enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1415087-indexesofobjectswithoptions?language=objc
func (a_ Array) IndexesOfObjectsWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](a_, objc.Sel("indexesOfObjectsWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Returns the index of the first object in the array that passes a test in a given block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1408043-indexofobjectpassingtest?language=objc
func (a_ Array) IndexOfObjectPassingTest(predicate func(obj objc.Object, idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](a_, objc.Sel("indexOfObjectPassingTest:"), predicate)
	return rv
}

// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given selector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1410025-sortedarrayusingselector?language=objc
func (a_ Array) SortedArrayUsingSelector(comparator objc.Selector) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("sortedArrayUsingSelector:"), comparator)
	return rv
}

// Returns the lowest index within a specified range whose corresponding array value is equal to a given object . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1415805-indexofobjectidenticalto?language=objc
func (a_ Array) IndexOfObjectIdenticalToInRange(anObject objc.IObject, range_ Range) uint {
	rv := objc.Call[uint](a_, objc.Sel("indexOfObjectIdenticalTo:inRange:"), objc.Ptr(anObject), range_)
	return rv
}

// Returns the object located at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1417555-objectatindex?language=objc
func (a_ Array) ObjectAtIndex(index uint) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("objectAtIndex:"), index)
	return rv
}

// Evaluates a given predicate against each object in the receiving array and returns a new array containing the objects for which the predicate returns true. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1411033-filteredarrayusingpredicate?language=objc
func (a_ Array) FilteredArrayUsingPredicate(predicate IPredicate) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("filteredArrayUsingPredicate:"), objc.Ptr(predicate))
	return rv
}

// Returns the object at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1414084-objectatindexedsubscript?language=objc
func (a_ Array) ObjectAtIndexedSubscript(idx uint) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("objectAtIndexedSubscript:"), idx)
	return rv
}

// Creates a new array by applying a difference object to an existing array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/3152165-arraybyapplyingdifference?language=objc
func (a_ Array) ArrayByApplyingDifference(difference IOrderedCollectionDifference) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("arrayByApplyingDifference:"), objc.Ptr(difference))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/2879153-arraywithcontentsofurl?language=objc
func (ac _ArrayClass) ArrayWithContentsOfURLError(url IURL, error IError) []objc.Object {
	rv := objc.Call[[]objc.Object](ac, objc.Sel("arrayWithContentsOfURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/2879153-arraywithcontentsofurl?language=objc
func Array_ArrayWithContentsOfURLError(url IURL, error IError) []objc.Object {
	return ArrayClass.ArrayWithContentsOfURLError(url, error)
}

// Executes a given block using each object in the array with the specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1413010-enumerateobjectswithoptions?language=objc
func (a_ Array) EnumerateObjectsWithOptionsUsingBlock(opts EnumerationOptions, block func(obj objc.Object, idx uint, stop *bool)) {
	objc.Call[objc.Void](a_, objc.Sel("enumerateObjectsWithOptions:usingBlock:"), opts, block)
}

// Copies references to objects contained in the array that fall within the specified range to aBuffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1414725-getobjects?language=objc
func (a_ Array) GetObjectsRange(objects objc.IObject, range_ Range) {
	objc.Call[objc.Void](a_, objc.Sel("getObjects:range:"), objc.Ptr(objects), range_)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/2879134-initwithcontentsofurl?language=objc
func (a_ Array) InitWithContentsOfURLError(url IURL, error IError) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("initWithContentsOfURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Returns a Boolean value that indicates whether a given object is present in the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1407477-containsobject?language=objc
func (a_ Array) ContainsObject(anObject objc.IObject) bool {
	rv := objc.Call[bool](a_, objc.Sel("containsObject:"), objc.Ptr(anObject))
	return rv
}

// Sends to each object in the array the message identified by a given selector, starting with the first object and continuing through the array to the last object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1460115-makeobjectsperformselector?language=objc
func (a_ Array) MakeObjectsPerformSelector(aSelector objc.Selector) {
	objc.Call[objc.Void](a_, objc.Sel("makeObjectsPerformSelector:"), aSelector)
}

// Returns an enumerator object that lets you access each object in the array, in reverse order. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1415734-reverseobjectenumerator?language=objc
func (a_ Array) ReverseObjectEnumerator() Enumerator {
	rv := objc.Call[Enumerator](a_, objc.Sel("reverseObjectEnumerator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/2879138-writetourl?language=objc
func (a_ Array) WriteToURLError(url IURL, error IError) bool {
	rv := objc.Call[bool](a_, objc.Sel("writeToURL:error:"), objc.Ptr(url), objc.Ptr(error))
	return rv
}

// Returns the indexes, from a given set of indexes, of objects in the array that pass a test in a given block for a given set of enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1413512-indexesofobjectsatindexes?language=objc
func (a_ Array) IndexesOfObjectsAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](a_, objc.Sel("indexesOfObjectsAtIndexes:options:passingTest:"), objc.Ptr(s), opts, predicate)
	return rv
}

// Raises an exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1414976-removeobserver?language=objc
func (a_ Array) RemoveObserverForKeyPath(observer objc.IObject, keyPath string) {
	objc.Call[objc.Void](a_, objc.Sel("removeObserver:forKeyPath:"), objc.Ptr(observer), keyPath)
}

// Returns a string that represents the contents of the array, formatted as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1412374-descriptionwithlocale?language=objc
func (a_ Array) DescriptionWithLocale(locale objc.IObject) string {
	rv := objc.Call[string](a_, objc.Sel("descriptionWithLocale:"), locale)
	return rv
}

// Returns a new array containing the receiving array’s elements that fall within the limits specified by a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1415157-subarraywithrange?language=objc
func (a_ Array) SubarrayWithRange(range_ Range) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("subarrayWithRange:"), range_)
	return rv
}

// Compares two arrays to create a difference object that represents the changes between them. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/3152166-differencefromarray?language=objc
func (a_ Array) DifferenceFromArray(other []objc.IObject) OrderedCollectionDifference {
	rv := objc.Call[OrderedCollectionDifference](a_, objc.Sel("differenceFromArray:"), other)
	return rv
}

// Invokes setValue:forKey: on each of the array's items using the specified value and key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1408301-setvalue?language=objc
func (a_ Array) SetValueForKey(value objc.IObject, key string) {
	objc.Call[objc.Void](a_, objc.Sel("setValue:forKey:"), value, key)
}

// Returns the indexes of objects in the array that pass a test in a given block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1417603-indexesofobjectspassingtest?language=objc
func (a_ Array) IndexesOfObjectsPassingTest(predicate func(obj objc.Object, idx uint, stop *bool) bool) IndexSet {
	rv := objc.Call[IndexSet](a_, objc.Sel("indexesOfObjectsPassingTest:"), predicate)
	return rv
}

// Returns a new array that is a copy of the receiving array with a given object added to the end. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1408534-arraybyaddingobject?language=objc
func (a_ Array) ArrayByAddingObject(anObject objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("arrayByAddingObject:"), objc.Ptr(anObject))
	return rv
}

// Returns the lowest index whose corresponding array value is equal to a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1417076-indexofobject?language=objc
func (a_ Array) IndexOfObject(anObject objc.IObject) uint {
	rv := objc.Call[uint](a_, objc.Sel("indexOfObject:"), objc.Ptr(anObject))
	return rv
}

// Raises an exception. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1409775-addobserver?language=objc
func (a_ Array) AddObserverForKeyPathOptionsContext(observer objc.IObject, keyPath string, options KeyValueObservingOptions, context unsafe.Pointer) {
	objc.Call[objc.Void](a_, objc.Sel("addObserver:forKeyPath:options:context:"), objc.Ptr(observer), keyPath, options, context)
}

// Returns the index of an object in the array that passes a test in a given block for a given set of enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1417053-indexofobjectwithoptions?language=objc
func (a_ Array) IndexOfObjectWithOptionsPassingTest(opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](a_, objc.Sel("indexOfObjectWithOptions:passingTest:"), opts, predicate)
	return rv
}

// Executes a given block using the objects in the array at the specified indexes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1417577-enumerateobjectsatindexes?language=objc
func (a_ Array) EnumerateObjectsAtIndexesOptionsUsingBlock(s IIndexSet, opts EnumerationOptions, block func(obj objc.Object, idx uint, stop *bool)) {
	objc.Call[objc.Void](a_, objc.Sel("enumerateObjectsAtIndexes:options:usingBlock:"), objc.Ptr(s), opts, block)
}

// Returns a new array that lists this array’s elements in a random order, using the specified random source. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1640687-shuffledarraywithrandomsource?language=objc
func (a_ Array) ShuffledArrayWithRandomSource(randomSource objc.IObject) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("shuffledArrayWithRandomSource:"), objc.Ptr(randomSource))
	return rv
}

// Returns an array containing all the pathname elements in the receiving array that have filename extensions from a given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1418275-pathsmatchingextensions?language=objc
func (a_ Array) PathsMatchingExtensions(filterTypes []string) []string {
	rv := objc.Call[[]string](a_, objc.Sel("pathsMatchingExtensions:"), filterTypes)
	return rv
}

// Constructs and returns an NSString object that is the result of interposing a given separator between the elements of the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1412075-componentsjoinedbystring?language=objc
func (a_ Array) ComponentsJoinedByString(separator string) string {
	rv := objc.Call[string](a_, objc.Sel("componentsJoinedByString:"), separator)
	return rv
}

// Returns an array containing the results of invoking valueForKey: using key on each of the array's objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1412219-valueforkey?language=objc
func (a_ Array) ValueForKey(key string) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("valueForKey:"), key)
	return rv
}

// Returns the index, from a given set of indexes, of the first object in the array that passes a test in a given block for a given set of enumeration options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1407652-indexofobjectatindexes?language=objc
func (a_ Array) IndexOfObjectAtIndexesOptionsPassingTest(s IIndexSet, opts EnumerationOptions, predicate func(obj objc.Object, idx uint, stop *bool) bool) uint {
	rv := objc.Call[uint](a_, objc.Sel("indexOfObjectAtIndexes:options:passingTest:"), objc.Ptr(s), opts, predicate)
	return rv
}

// Returns the first object contained in the receiving array that’s equal to an object in another given array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1408825-firstobjectcommonwitharray?language=objc
func (a_ Array) FirstObjectCommonWithArray(otherArray []objc.IObject) objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("firstObjectCommonWithArray:"), otherArray)
	return rv
}

// Returns an array that lists the receiving array’s elements in ascending order, as determined by the comparison method specified by a given NSComparator block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1411195-sortedarrayusingcomparator?language=objc
func (a_ Array) SortedArrayUsingComparator(cmptr Comparator) []objc.Object {
	rv := objc.Call[[]objc.Object](a_, objc.Sel("sortedArrayUsingComparator:"), cmptr)
	return rv
}

// Analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to sortedArrayUsingFunction:context:hint:. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1413063-sortedarrayhint?language=objc
func (a_ Array) SortedArrayHint() []byte {
	rv := objc.Call[[]byte](a_, objc.Sel("sortedArrayHint"))
	return rv
}

// The last object in the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1408316-lastobject?language=objc
func (a_ Array) LastObject() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("lastObject"))
	return rv
}

// A string that represents the contents of the array, formatted as a property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1413042-description?language=objc
func (a_ Array) Description() string {
	rv := objc.Call[string](a_, objc.Sel("description"))
	return rv
}

// The number of objects in the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1409982-count?language=objc
func (a_ Array) Count() uint {
	rv := objc.Call[uint](a_, objc.Sel("count"))
	return rv
}

// The first object in the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/1412852-firstobject?language=objc
func (a_ Array) FirstObject() objc.Object {
	rv := objc.Call[objc.Object](a_, objc.Sel("firstObject"))
	return rv
}
