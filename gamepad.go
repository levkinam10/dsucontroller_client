package main

import (
	"github.com/bendahl/uinput"
	//"testing/quick"
)

func Gamepad(c chan Data) {
	gamepad, _ := uinput.CreateGamepad("/dev/uinput", []byte("DSU"), uint16(0x045e), uint16(0x028e))

	for {
		data := <-c
		gamepad.LeftStickMove(float32(data.leftX-127)/127, float32(data.leftY-127)/127)
		gamepad.RightStickMove(float32(data.rightX-127)/127, float32(data.rightY-127)/127)

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
