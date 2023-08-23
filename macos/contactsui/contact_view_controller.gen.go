// AUTO-GENERATED CODE, DO NOT MODIFY

package contactsui

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/contacts"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ContactViewController] class.
var ContactViewControllerClass = _ContactViewControllerClass{objc.GetClass("CNContactViewController")}

type _ContactViewControllerClass struct {
	objc.Class
}

// An interface definition for the [ContactViewController] class.
type IContactViewController interface {
	appkit.IViewController
	Contact() contacts.Contact
}

// A view controller that displays a new, unknown, or existing contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactviewcontroller?language=objc
type ContactViewController struct {
	appkit.ViewController
}

func ContactViewControllerFrom(ptr unsafe.Pointer) ContactViewController {
	return ContactViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

func (cc _ContactViewControllerClass) Alloc() ContactViewController {
	rv := objc.Call[ContactViewController](cc, objc.Sel("alloc"))
	return rv
}

func ContactViewController_Alloc() ContactViewController {
	return ContactViewControllerClass.Alloc()
}

func (cc _ContactViewControllerClass) New() ContactViewController {
	rv := objc.Call[ContactViewController](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContactViewController() ContactViewController {
	return ContactViewControllerClass.New()
}

func (c_ ContactViewController) Init() ContactViewController {
	rv := objc.Call[ContactViewController](c_, objc.Sel("init"))
	return rv
}

func (c_ ContactViewController) InitWithNibNameBundle(nibNameOrNil appkit.NibName, nibBundleOrNil foundation.IBundle) ContactViewController {
	rv := objc.Call[ContactViewController](c_, objc.Sel("initWithNibName:bundle:"), nibNameOrNil, objc.Ptr(nibBundleOrNil))
	return rv
}

// Returns a view controller object initialized to the nib file in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsviewcontroller/1434481-initwithnibname?language=objc
func NewContactViewControllerWithNibNameBundle(nibNameOrNil appkit.NibName, nibBundleOrNil foundation.IBundle) ContactViewController {
	instance := ContactViewControllerClass.Alloc().InitWithNibNameBundle(nibNameOrNil, nibBundleOrNil)
	instance.Autorelease()
	return instance
}

// Returns the descriptor for all the keys that must be fetched on the contact before setting it on the view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactviewcontroller/1550990-descriptorforrequiredkeys?language=objc
func (cc _ContactViewControllerClass) DescriptorForRequiredKeys() objc.Object {
	rv := objc.Call[objc.Object](cc, objc.Sel("descriptorForRequiredKeys"))
	return rv
}

// Returns the descriptor for all the keys that must be fetched on the contact before setting it on the view controller. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactviewcontroller/1550990-descriptorforrequiredkeys?language=objc
func ContactViewController_DescriptorForRequiredKeys() objc.Object {
	return ContactViewControllerClass.DescriptorForRequiredKeys()
}

// The contact being displayed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contactsui/cncontactviewcontroller/1522596-contact?language=objc
func (c_ ContactViewController) Contact() contacts.Contact {
	rv := objc.Call[contacts.Contact](c_, objc.Sel("contact"))
	return rv
}
