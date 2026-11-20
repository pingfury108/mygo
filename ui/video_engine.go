package ui

import (
	"github.com/egoist/mygo/internal/media"
	"github.com/egoist/mygo/internal/scene"
)

// A VideoEngine is what a VideoPlayer shows, for a video this app decodes
// itself: a decoder MyGo has no engine for, as HEVC where the system has no
// decoder for it and a GPU one is at hand. NewVideoPlayer takes a decoder
// that hands over pixels through memory; an engine hands over pictures that
// stay on the GPU, which a window's renderer draws from there.
//
//	mygo.NewVideoPlayerOf(engine)
//
// An engine calls the wake it was given, from any goroutine, as each
// picture comes due.
type VideoEngine = media.Engine

// A VideoPicture is a picture a VideoEngine hands over: its pixels, or a
// texture on the GPU that holds them.
type VideoPicture = media.Picture

// VideoState is how a VideoEngine loads and plays.
type VideoState = media.State

// The states of a VideoEngine.
const (
	VideoLoading   = media.Loading
	VideoReady     = media.Ready
	VideoBuffering = media.Buffering
	VideoEnded     = media.Ended
	VideoFailed    = media.Failed
)

// NewTexturePicture wraps a texture shared by another process or another
// part of this one as a picture: on Windows the NT handle of a Direct3D 11
// texture in BGRA with a keyed mutex, on macOS a CVPixelBufferRef, which
// the window's renderer draws from there, without a copy through memory.
func NewTexturePicture(native uintptr, id uint64, w, h int) VideoPicture {
	return media.Picture{
		Surface: &scene.Surface{Format: scene.SurfaceBGRA, Native: native, NativeID: id},
		W:       w,
		H:       h,
	}
}

// NewPixelsPicture wraps pixels in memory as a picture, which the renderer
// uploads: BGRA, premultiplied, w×h of them, in a slice that stays as it is
// while the picture shows.
func NewPixelsPicture(pix []byte, w, h int) VideoPicture {
	s := &scene.Surface{Format: scene.SurfaceBGRA}
	s.Planes[0] = scene.Plane{Pix: pix, Stride: w * 4}
	return media.Picture{Surface: s, W: w, H: h}
}

// NewVideoPlayerOf plays the pictures of an engine.
func NewVideoPlayerOf(e VideoEngine) *VideoPlayer {
	return NewEngineVideoPlayer(e)
}
