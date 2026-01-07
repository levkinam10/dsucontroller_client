package main

import (
	"github.com/bendahl/uinput"
	"github.com/holoplot/go-evdev"
	"time"
	//"testing/quick"
)

/*
#cgo pkg-config: libevdev
#include <libevdev/libevdev-uinput.h>
#include <libevdev/libevdev.h>
#include <linux/input-event-codes.h>
#include <linux/input.h>
#include <unistd.h>
#include<gamepad.h>
*/
import "C"

func c_pad() {
	uidev := C.create_controller()
	defer C.destroy_device(uidev)
	for {
		C.btn_down(uidev, C.uint(evdev.BTN_A))
		C.syn(uidev)
		time.Sleep(5 * time.Second)
		C.btn_up(uidev, C.uint(evdev.BTN_A))
		C.syn(uidev)
		time.Sleep(5 * time.Second)

	}

}
func Gamepad(c chan Data) {
	gamepad, _ := uinput.CreateGamepad("/dev/uinput", []byte("DSU"), uint16(0x1234), uint16(0x5678))

	for {
		// Works correct only with Steam Input
		data := <-c
		gamepad.LeftStickMove((float32(data.leftX-127)/127)*1, (float32(data.leftY-127)/127)*1)

		gamepad.RightStickMoveX(float32(data.rightX-127) / 127)
		gamepad.RightStickMoveY(float32(data.rightY-127) / 127)

		if data.FirstButtons[0] == '1' {
			gamepad.ButtonDown(uinput.ButtonNorth)
		} else {
			gamepad.ButtonUp(uinput.ButtonNorth)
		}
		if data.FirstButtons[1] == '1' {
			gamepad.ButtonDown(uinput.ButtonSouth)
		} else {
			gamepad.ButtonUp(uinput.ButtonSouth)
		}
		if data.FirstButtons[2] == '1' {
			gamepad.ButtonDown(uinput.ButtonEast)
		} else {
			gamepad.ButtonUp(uinput.ButtonEast)
		}
		if data.FirstButtons[3] == '1' {
			gamepad.ButtonDown(uinput.ButtonWest)
		} else {
			gamepad.ButtonUp(uinput.ButtonWest)
		}

		//
		if data.FirstButtons[4] == '1' {
			gamepad.ButtonDown(uinput.ButtonBumperRight)
		} else {
			gamepad.ButtonUp(uinput.ButtonBumperRight)
		}
		if data.FirstButtons[5] == '1' {
			gamepad.ButtonDown(uinput.ButtonBumperLeft)
		} else {
			gamepad.ButtonUp(uinput.ButtonBumperLeft)
		}
		if data.FirstButtons[6] == '1' {
			gamepad.ButtonDown(uinput.ButtonTriggerRight)

		} else {
			gamepad.ButtonUp(uinput.ButtonTriggerRight)
		}
		if data.FirstButtons[7] == '1' {
			gamepad.ButtonDown(uinput.ButtonTriggerLeft)
		} else {
			gamepad.ButtonUp(uinput.ButtonTriggerLeft)
		}
		if data.SecondButtons[0] == '1' {
			gamepad.ButtonDown(uinput.ButtonDpadLeft)
		} else {
			gamepad.ButtonUp(uinput.ButtonDpadLeft)
		}
		if data.SecondButtons[1] == '1' {
			gamepad.ButtonDown(uinput.ButtonDpadDown)
		} else {
			gamepad.ButtonUp(uinput.ButtonDpadDown)
		}
		if data.SecondButtons[2] == '1' {
			gamepad.ButtonDown(uinput.ButtonDpadRight)
		} else {
			gamepad.ButtonUp(uinput.ButtonDpadRight)
		}
		if data.SecondButtons[3] == '1' {
			gamepad.ButtonDown(uinput.ButtonDpadUp)
		} else {
			gamepad.ButtonUp(uinput.ButtonDpadUp)
		}
		if data.SecondButtons[4] == '1' {
			gamepad.ButtonDown(uinput.ButtonStart)
		} else {
			gamepad.ButtonUp(uinput.ButtonStart)
		}
		if data.SecondButtons[5] == '1' {
			gamepad.ButtonDown(uinput.ButtonThumbRight)
		} else {
			gamepad.ButtonUp(uinput.ButtonThumbRight)
		}
		if data.SecondButtons[6] == '1' {
			gamepad.ButtonDown(uinput.ButtonThumbLeft)
		} else {
			gamepad.ButtonUp(uinput.ButtonThumbLeft)
		}
		if data.SecondButtons[7] == '1' {
			gamepad.ButtonDown(uinput.ButtonSelect)
		} else {
			gamepad.ButtonUp(uinput.ButtonSelect)
		}

	}

}

type Data struct {
	leftX         int
	leftY         int
	rightX        int
	rightY        int
	FirstButtons  string
	SecondButtons string
}
