// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods to adopt to respond to route picker view presentation events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerviewdelegate?language=objc
type PRoutePickerViewDelegate interface {
	// optional
	RoutePickerViewWillBeginPresentingRoutes(routePickerView RoutePickerView)
	HasRoutePickerViewWillBeginPresentingRoutes() bool

	// optional
	RoutePickerViewDidEndPresentingRoutes(routePickerView RoutePickerView)
	HasRoutePickerViewDidEndPresentingRoutes() bool
}

// A delegate implementation builder for the [PRoutePickerViewDelegate] protocol.
type RoutePickerViewDelegate struct {
	_RoutePickerViewWillBeginPresentingRoutes func(routePickerView RoutePickerView)
	_RoutePickerViewDidEndPresentingRoutes    func(routePickerView RoutePickerView)
}

func (di *RoutePickerViewDelegate) HasRoutePickerViewWillBeginPresentingRoutes() bool {
	return di._RoutePickerViewWillBeginPresentingRoutes != nil
}

// Tells the delegate that the route picker view is about to begin presenting routes to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerviewdelegate/2915788-routepickerviewwillbeginpresenti?language=objc
func (di *RoutePickerViewDelegate) SetRoutePickerViewWillBeginPresentingRoutes(f func(routePickerView RoutePickerView)) {
	di._RoutePickerViewWillBeginPresentingRoutes = f
}

// Tells the delegate that the route picker view is about to begin presenting routes to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerviewdelegate/2915788-routepickerviewwillbeginpresenti?language=objc
func (di *RoutePickerViewDelegate) RoutePickerViewWillBeginPresentingRoutes(routePickerView RoutePickerView) {
	di._RoutePickerViewWillBeginPresentingRoutes(routePickerView)
}
func (di *RoutePickerViewDelegate) HasRoutePickerViewDidEndPresentingRoutes() bool {
	return di._RoutePickerViewDidEndPresentingRoutes != nil
}

// Tells the delegate when the route picker view finishes presenting routes to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerviewdelegate/2915796-routepickerviewdidendpresentingr?language=objc
func (di *RoutePickerViewDelegate) SetRoutePickerViewDidEndPresentingRoutes(f func(routePickerView RoutePickerView)) {
	di._RoutePickerViewDidEndPresentingRoutes = f
}

// Tells the delegate when the route picker view finishes presenting routes to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerviewdelegate/2915796-routepickerviewdidendpresentingr?language=objc
func (di *RoutePickerViewDelegate) RoutePickerViewDidEndPresentingRoutes(routePickerView RoutePickerView) {
	di._RoutePickerViewDidEndPresentingRoutes(routePickerView)
}

// A concrete type wrapper for the [PRoutePickerViewDelegate] protocol.
type RoutePickerViewDelegateWrapper struct {
	objc.Object
}

func (r_ RoutePickerViewDelegateWrapper) HasRoutePickerViewWillBeginPresentingRoutes() bool {
	return r_.RespondsToSelector(objc.Sel("routePickerViewWillBeginPresentingRoutes:"))
}

// Tells the delegate that the route picker view is about to begin presenting routes to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerviewdelegate/2915788-routepickerviewwillbeginpresenti?language=objc
func (r_ RoutePickerViewDelegateWrapper) RoutePickerViewWillBeginPresentingRoutes(routePickerView IRoutePickerView) {
	objc.Call[objc.Void](r_, objc.Sel("routePickerViewWillBeginPresentingRoutes:"), objc.Ptr(routePickerView))
}

func (r_ RoutePickerViewDelegateWrapper) HasRoutePickerViewDidEndPresentingRoutes() bool {
	return r_.RespondsToSelector(objc.Sel("routePickerViewDidEndPresentingRoutes:"))
}

// Tells the delegate when the route picker view finishes presenting routes to the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avroutepickerviewdelegate/2915796-routepickerviewdidendpresentingr?language=objc
func (r_ RoutePickerViewDelegateWrapper) RoutePickerViewDidEndPresentingRoutes(routePickerView IRoutePickerView) {
	objc.Call[objc.Void](r_, objc.Sel("routePickerViewDidEndPresentingRoutes:"), objc.Ptr(routePickerView))
}
