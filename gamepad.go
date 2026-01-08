package main

import (
	"github.com/holoplot/go-evdev"
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

// func c_pad() {
//
//	for {
//		C.btn_down(uidev, C.uint(evdev.BTN_A))
//		C.syn(uidev)
//		time.Sleep(5 * time.Second)
//		C.btn_up(uidev, C.uint(evdev.BTN_A))
//		C.syn(uidev)
//		time.Sleep(5 * time.Second)
//
//	}
//
// }
func Gamepad(c chan Data) {
	uidev := C.create_controller()
	defer C.destroy_device(uidev)
	for {
		// Works correct only with Steam Input
		data := <-c
		//gamepad.LeftStickMove((float32(data.leftX-127)/127)*1, (float32(data.leftY-127)/127)*1)
		//
		//gamepad.RightStickMoveX(float32(data.rightX-127) / 127)
		C.set_axis(uidev, C.uint(evdev.ABS_X), C.uint((data.leftX<<8)+data.leftX-32768))
		C.set_axis(uidev, C.uint(evdev.ABS_Y), C.uint((data.leftY<<8)+data.leftY-32768))
		C.set_axis(uidev, C.uint(evdev.ABS_RX), C.uint((data.rightX<<8)+data.rightX-32768))
		C.set_axis(uidev, C.uint(evdev.ABS_RY), C.uint((data.rightY<<8)+data.rightY-32768))
		//gamepad.RightStickMoveY(float32(data.rightY-127) / 127)

		if data.FirstButtons[0] == '1' {
			C.btn_down(uidev, C.uint(evdev.BTN_X))
		} else {
			C.btn_up(uidev, C.uint(evdev.BTN_X))
		}
		if data.FirstButtons[1] == '1' {
			C.btn_down(uidev, C.uint(evdev.BTN_A))
		} else {
			C.btn_up(uidev, C.uint(evdev.BTN_A))
		}
		if data.FirstButtons[2] == '1' {
			C.btn_down(uidev, C.uint(evdev.BTN_B))
		} else {
			C.btn_up(uidev, C.uint(evdev.BTN_B))
		}
		if data.FirstButtons[3] == '1' {
			C.btn_down(uidev, C.uint(evdev.BTN_Y))
		} else {
			C.btn_up(uidev, C.uint(evdev.BTN_Y))
		}

		if data.FirstButtons[4] == '1' {
			//C.set_axis(uidev, C.uint(evdev.ABS_HAT1X), C.uint(255))
			C.btn_down(uidev, C.uint(evdev.BTN_TR))

		} else {
			C.btn_up(uidev, C.uint(evdev.BTN_TR))
		}
		if data.FirstButtons[5] == '1' {
			C.btn_down(uidev, C.uint(evdev.BTN_TL))
		} else {
			C.btn_up(uidev, C.uint(evdev.BTN_TL))
		}

		if data.FirstButtons[6] == '1' {
			C.set_axis(uidev, C.uint(evdev.ABS_RZ), C.uint(255))
		} else {
			C.set_axis(uidev, C.uint(evdev.ABS_RZ), C.uint(0))
		}
		if data.FirstButtons[7] == '1' {
			C.set_axis(uidev, C.uint(evdev.ABS_Z), C.uint(255))
		} else {
			C.set_axis(uidev, C.uint(evdev.ABS_Z), C.uint(0))
		}
		//

		if data.SecondButtons[0] == '1' {
			C.btn_down(uidev, evdev.BTN_DPAD_LEFT)
		} else {
			C.btn_up(uidev, evdev.BTN_DPAD_LEFT)
		}
		if data.SecondButtons[1] == '1' {
			C.btn_down(uidev, evdev.BTN_DPAD_DOWN)
		} else {
			C.btn_up(uidev, evdev.BTN_DPAD_DOWN)
		}
		if data.SecondButtons[2] == '1' {
			C.btn_down(uidev, evdev.BTN_DPAD_RIGHT)
		} else {
			C.btn_up(uidev, evdev.BTN_DPAD_RIGHT)
		}
		if data.SecondButtons[3] == '1' {
			C.btn_down(uidev, evdev.BTN_DPAD_UP)
		} else {
			C.btn_up(uidev, evdev.BTN_DPAD_UP)
		}
		if data.SecondButtons[4] == '1' {
			C.btn_down(uidev, evdev.BTN_START)
		} else {
			C.btn_up(uidev, evdev.BTN_START)
		}
		if data.SecondButtons[5] == '1' {
			C.btn_down(uidev, evdev.BTN_THUMBL)
		} else {
			C.btn_up(uidev, evdev.BTN_THUMBL)
		}
		if data.SecondButtons[6] == '1' {
			C.btn_down(uidev, evdev.BTN_THUMBR)
		} else {
			C.btn_up(uidev, evdev.BTN_THUMBR)
		}
		if data.SecondButtons[7] == '1' {
			C.btn_down(uidev, evdev.BTN_SELECT)
		} else {
			C.btn_up(uidev, evdev.BTN_SELECT)
		}
		C.syn(uidev)
	}

}
func test(c chan Data) {
	<-c

}

type Data struct {
	leftX         int
	leftY         int
	rightX        int
	rightY        int
	FirstButtons  string
	SecondButtons string
}
