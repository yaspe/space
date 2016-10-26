package main

import (
	"./core"
	
	"github.com/veandco/go-sdl2/sdl"
)

func ProcessControls(s *core.Ship) {
	
	var keyLeftPress bool
	var keyRightPress bool
	var keyDownPress bool
	
	for {
		event := sdl.PollEvent()
		switch t := event.(type) {
		case *sdl.KeyUpEvent:
			if t.Keysym.Sym == 1073741905 {
				keyDownPress = false
			} else if t.Keysym.Sym == 1073741903 {
				keyLeftPress = false
			} else if t.Keysym.Sym == 1073741904 {
				keyRightPress = false
			}
		case *sdl.KeyDownEvent:
			sdl.FlushEvent(sdl.KEYDOWN)
			if t.Keysym.Sym == 1073741905 {
				keyDownPress = true
			} else if t.Keysym.Sym == 1073741903 {
				keyLeftPress = true
			} else if t.Keysym.Sym == 1073741904 {
				keyRightPress = true
			}
		}
		
		if keyDownPress {
			s.EngineMain()
		} else {
			s.EngineMainDesable()
		}
		
		if keyLeftPress && !keyRightPress {
			s.EngineLeft()
		} else if !keyLeftPress && keyRightPress {
			s.EngineRight()
		} else {
			s.EngineRightDesable()
		}
	}
}