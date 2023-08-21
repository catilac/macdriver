// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactRelation] class.
var ContactRelationClass = _ContactRelationClass{objc.GetClass("CNContactRelation")}

type _ContactRelationClass struct {
	objc.Class
}

// An interface definition for the [ContactRelation] class.
type IContactRelation interface {
	objc.IObject
	Name() string
}

// An immutable object that represents the relationship between one contact to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactrelation?language=objc
type ContactRelation struct {
	objc.Object
}

func ContactRelationFrom(ptr unsafe.Pointer) ContactRelation {
	return ContactRelation{
		Object: objc.ObjectFrom(ptr),
	}
}

func (c_ ContactRelation) InitWithName(name string) ContactRelation {
	rv := objc.Call[ContactRelation](c_, objc.Sel("initWithName:"), name)
	return rv
}

// Creates an object with the name of the related contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactrelation/1403385-initwithname?language=objc
func NewContactRelationWithName(name string) ContactRelation {
	instance := ContactRelationClass.Alloc().InitWithName(name)
	instance.Autorelease()
	return instance
}

func (cc _ContactRelationClass) ContactRelationWithName(name string) ContactRelation {
	rv := objc.Call[ContactRelation](cc, objc.Sel("contactRelationWithName:"), name)
	return rv
}

// Instantiate a class instance with the name of the related contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactrelation/1416519-contactrelationwithname?language=objc
func ContactRelation_ContactRelationWithName(name string) ContactRelation {
	return ContactRelationClass.ContactRelationWithName(name)
}

func (cc _ContactRelationClass) Alloc() ContactRelation {
	rv := objc.Call[ContactRelation](cc, objc.Sel("alloc"))
	return rv
}

func ContactRelation_Alloc() ContactRelation {
	return ContactRelationClass.Alloc()
}

func (cc _ContactRelationClass) New() ContactRelation {
	rv := objc.Call[ContactRelation](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactRelation() ContactRelation {
	return ContactRelationClass.New()
}

func (c_ ContactRelation) Init() ContactRelation {
	rv := objc.Call[ContactRelation](c_, objc.Sel("init"))
	return rv
}

// The name of the related contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactrelation/1403399-name?language=objc
func (c_ ContactRelation) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}
