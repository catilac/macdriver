// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coredata"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PersistentDocument] class.
var PersistentDocumentClass = _PersistentDocumentClass{objc.GetClass("NSPersistentDocument")}

type _PersistentDocumentClass struct {
	objc.Class
}

// An interface definition for the [PersistentDocument] class.
type IPersistentDocument interface {
	IDocument
	PersistentStoreTypeForFileType(fileType string) string
	ManagedObjectModel() coredata.ManagedObjectModel
	ManagedObjectContext() coredata.ManagedObjectContext
	SetManagedObjectContext(value coredata.IManagedObjectContext)
}

// A document object that can integrate with Core Data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument?language=objc
type PersistentDocument struct {
	Document
}

func PersistentDocumentFrom(ptr unsafe.Pointer) PersistentDocument {
	return PersistentDocument{
		Document: DocumentFrom(ptr),
	}
}

func (pc _PersistentDocumentClass) Alloc() PersistentDocument {
	rv := objc.Call[PersistentDocument](pc, objc.Sel("alloc"))
	return rv
}

func PersistentDocument_Alloc() PersistentDocument {
	return PersistentDocumentClass.Alloc()
}

func (pc _PersistentDocumentClass) New() PersistentDocument {
	rv := objc.Call[PersistentDocument](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPersistentDocument() PersistentDocument {
	return PersistentDocumentClass.New()
}

func (p_ PersistentDocument) Init() PersistentDocument {
	rv := objc.Call[PersistentDocument](p_, objc.Sel("init"))
	return rv
}

func (p_ PersistentDocument) InitWithTypeError(typeName string, outError foundation.IError) PersistentDocument {
	rv := objc.Call[PersistentDocument](p_, objc.Sel("initWithType:error:"), typeName, objc.Ptr(outError))
	return rv
}

// Initializes a document of a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515159-initwithtype?language=objc
func NewPersistentDocumentWithTypeError(typeName string, outError foundation.IError) PersistentDocument {
	instance := PersistentDocumentClass.Alloc().InitWithTypeError(typeName, outError)
	instance.Autorelease()
	return instance
}

func (p_ PersistentDocument) InitForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError foundation.IError) PersistentDocument {
	rv := objc.Call[PersistentDocument](p_, objc.Sel("initForURL:withContentsOfURL:ofType:error:"), objc.Ptr(urlOrNil), objc.Ptr(contentsURL), typeName, objc.Ptr(outError))
	return rv
}

// Initializes a document with the specified contents, and places the resulting document's file at the designated location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/1515041-initforurl?language=objc
func NewPersistentDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError foundation.IError) PersistentDocument {
	instance := PersistentDocumentClass.Alloc().InitForURLWithContentsOfURLOfTypeError(urlOrNil, contentsURL, typeName, outError)
	instance.Autorelease()
	return instance
}

// Returns the type of persistent store associated with the specified file type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument/1396168-persistentstoretypeforfiletype?language=objc
func (p_ PersistentDocument) PersistentStoreTypeForFileType(fileType string) string {
	rv := objc.Call[string](p_, objc.Sel("persistentStoreTypeForFileType:"), fileType)
	return rv
}

// The managed object model of the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument/1396152-managedobjectmodel?language=objc
func (p_ PersistentDocument) ManagedObjectModel() coredata.ManagedObjectModel {
	rv := objc.Call[coredata.ManagedObjectModel](p_, objc.Sel("managedObjectModel"))
	return rv
}

// The managed object context for the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument/1396162-managedobjectcontext?language=objc
func (p_ PersistentDocument) ManagedObjectContext() coredata.ManagedObjectContext {
	rv := objc.Call[coredata.ManagedObjectContext](p_, objc.Sel("managedObjectContext"))
	return rv
}

// The managed object context for the document. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument/1396162-managedobjectcontext?language=objc
func (p_ PersistentDocument) SetManagedObjectContext(value coredata.IManagedObjectContext) {
	objc.Call[objc.Void](p_, objc.Sel("setManagedObjectContext:"), objc.Ptr(value))
}
