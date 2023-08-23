// AUTO-GENERATED CODE, DO NOT MODIFY

package avkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol that defines the methods to implement to participate in the player view’s full-screen presentation life cycle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate?language=objc
type PPlayerViewDelegate interface {
	// optional
	PlayerViewWillEnterFullScreen(playerView PlayerView)
	HasPlayerViewWillEnterFullScreen() bool

	// optional
	PlayerViewDidEnterFullScreen(playerView PlayerView)
	HasPlayerViewDidEnterFullScreen() bool

	// optional
	PlayerViewWillExitFullScreen(playerView PlayerView)
	HasPlayerViewWillExitFullScreen() bool

	// optional
	PlayerViewDidExitFullScreen(playerView PlayerView)
	HasPlayerViewDidExitFullScreen() bool

	// optional
	PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerView PlayerView, completionHandler func(restored bool))
	HasPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler() bool
}

// A delegate implementation builder for the [PPlayerViewDelegate] protocol.
type PlayerViewDelegate struct {
	_PlayerViewWillEnterFullScreen                                        func(playerView PlayerView)
	_PlayerViewDidEnterFullScreen                                         func(playerView PlayerView)
	_PlayerViewWillExitFullScreen                                         func(playerView PlayerView)
	_PlayerViewDidExitFullScreen                                          func(playerView PlayerView)
	_PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler func(playerView PlayerView, completionHandler func(restored bool))
}

func (di *PlayerViewDelegate) HasPlayerViewWillEnterFullScreen() bool {
	return di._PlayerViewWillEnterFullScreen != nil
}

// Tells the delegate that the player view is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752989-playerviewwillenterfullscreen?language=objc
func (di *PlayerViewDelegate) SetPlayerViewWillEnterFullScreen(f func(playerView PlayerView)) {
	di._PlayerViewWillEnterFullScreen = f
}

// Tells the delegate that the player view is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752989-playerviewwillenterfullscreen?language=objc
func (di *PlayerViewDelegate) PlayerViewWillEnterFullScreen(playerView PlayerView) {
	di._PlayerViewWillEnterFullScreen(playerView)
}
func (di *PlayerViewDelegate) HasPlayerViewDidEnterFullScreen() bool {
	return di._PlayerViewDidEnterFullScreen != nil
}

// Tells the delegate that the player view entered full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752987-playerviewdidenterfullscreen?language=objc
func (di *PlayerViewDelegate) SetPlayerViewDidEnterFullScreen(f func(playerView PlayerView)) {
	di._PlayerViewDidEnterFullScreen = f
}

// Tells the delegate that the player view entered full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752987-playerviewdidenterfullscreen?language=objc
func (di *PlayerViewDelegate) PlayerViewDidEnterFullScreen(playerView PlayerView) {
	di._PlayerViewDidEnterFullScreen(playerView)
}
func (di *PlayerViewDelegate) HasPlayerViewWillExitFullScreen() bool {
	return di._PlayerViewWillExitFullScreen != nil
}

// Tells the delegate that the player view is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752990-playerviewwillexitfullscreen?language=objc
func (di *PlayerViewDelegate) SetPlayerViewWillExitFullScreen(f func(playerView PlayerView)) {
	di._PlayerViewWillExitFullScreen = f
}

// Tells the delegate that the player view is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752990-playerviewwillexitfullscreen?language=objc
func (di *PlayerViewDelegate) PlayerViewWillExitFullScreen(playerView PlayerView) {
	di._PlayerViewWillExitFullScreen(playerView)
}
func (di *PlayerViewDelegate) HasPlayerViewDidExitFullScreen() bool {
	return di._PlayerViewDidExitFullScreen != nil
}

// Tells the delegate that the player view exited full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752988-playerviewdidexitfullscreen?language=objc
func (di *PlayerViewDelegate) SetPlayerViewDidExitFullScreen(f func(playerView PlayerView)) {
	di._PlayerViewDidExitFullScreen = f
}

// Tells the delegate that the player view exited full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752988-playerviewdidexitfullscreen?language=objc
func (di *PlayerViewDelegate) PlayerViewDidExitFullScreen(playerView PlayerView) {
	di._PlayerViewDidExitFullScreen(playerView)
}
func (di *PlayerViewDelegate) HasPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler() bool {
	return di._PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler != nil
}

// Tells the delegate to restore the app’s user interface when exiting full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752986-playerview?language=objc
func (di *PlayerViewDelegate) SetPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(f func(playerView PlayerView, completionHandler func(restored bool))) {
	di._PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler = f
}

// Tells the delegate to restore the app’s user interface when exiting full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752986-playerview?language=objc
func (di *PlayerViewDelegate) PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerView PlayerView, completionHandler func(restored bool)) {
	di._PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerView, completionHandler)
}

// A concrete type wrapper for the [PPlayerViewDelegate] protocol.
type PlayerViewDelegateWrapper struct {
	objc.Object
}

func (p_ PlayerViewDelegateWrapper) HasPlayerViewWillEnterFullScreen() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewWillEnterFullScreen:"))
}

// Tells the delegate that the player view is about to enter full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752989-playerviewwillenterfullscreen?language=objc
func (p_ PlayerViewDelegateWrapper) PlayerViewWillEnterFullScreen(playerView IPlayerView) {
	objc.Call[objc.Void](p_, objc.Sel("playerViewWillEnterFullScreen:"), objc.Ptr(playerView))
}

func (p_ PlayerViewDelegateWrapper) HasPlayerViewDidEnterFullScreen() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewDidEnterFullScreen:"))
}

// Tells the delegate that the player view entered full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752987-playerviewdidenterfullscreen?language=objc
func (p_ PlayerViewDelegateWrapper) PlayerViewDidEnterFullScreen(playerView IPlayerView) {
	objc.Call[objc.Void](p_, objc.Sel("playerViewDidEnterFullScreen:"), objc.Ptr(playerView))
}

func (p_ PlayerViewDelegateWrapper) HasPlayerViewWillExitFullScreen() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewWillExitFullScreen:"))
}

// Tells the delegate that the player view is about to exit full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752990-playerviewwillexitfullscreen?language=objc
func (p_ PlayerViewDelegateWrapper) PlayerViewWillExitFullScreen(playerView IPlayerView) {
	objc.Call[objc.Void](p_, objc.Sel("playerViewWillExitFullScreen:"), objc.Ptr(playerView))
}

func (p_ PlayerViewDelegateWrapper) HasPlayerViewDidExitFullScreen() bool {
	return p_.RespondsToSelector(objc.Sel("playerViewDidExitFullScreen:"))
}

// Tells the delegate that the player view exited full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752988-playerviewdidexitfullscreen?language=objc
func (p_ PlayerViewDelegateWrapper) PlayerViewDidExitFullScreen(playerView IPlayerView) {
	objc.Call[objc.Void](p_, objc.Sel("playerViewDidExitFullScreen:"), objc.Ptr(playerView))
}

func (p_ PlayerViewDelegateWrapper) HasPlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler() bool {
	return p_.RespondsToSelector(objc.Sel("playerView:restoreUserInterfaceForFullScreenExitWithCompletionHandler:"))
}

// Tells the delegate to restore the app’s user interface when exiting full-screen mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avkit/avplayerviewdelegate/3752986-playerview?language=objc
func (p_ PlayerViewDelegateWrapper) PlayerViewRestoreUserInterfaceForFullScreenExitWithCompletionHandler(playerView IPlayerView, completionHandler func(restored bool)) {
	objc.Call[objc.Void](p_, objc.Sel("playerView:restoreUserInterfaceForFullScreenExitWithCompletionHandler:"), objc.Ptr(playerView), completionHandler)
}
