package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

var keyLeftPress bool
var keyRightPress bool
var keyDownPress bool

var controlMutex = &sync.Mutex{}

func ProcessControls(s *Ship) {

	for {
		event := sdl.PollEvent()
		if event == nil {
			continue
		}
		switch t := event.(type) {
		case *sdl.KeyUpEvent:
			controlMutex.Lock()
			if t.Keysym.Sym == 1073741905 {
				keyDownPress = false
			} else if t.Keysym.Sym == 1073741903 {
				keyLeftPress = false
			} else if t.Keysym.Sym == 1073741904 {
				keyRightPress = false
			}
			controlMutex.Unlock()
		case *sdl.KeyDownEvent:
			controlMutex.Lock()
			if t.Keysym.Sym == 1073741905 {
				keyDownPress = true
			} else if t.Keysym.Sym == 1073741903 {
				keyLeftPress = true
			} else if t.Keysym.Sym == 1073741904 {
				keyRightPress = true
			}
			controlMutex.Unlock()
		}
	}
}
