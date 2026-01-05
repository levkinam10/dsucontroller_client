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

	}

}

type Data struct {
	leftX        int
	leftY        int
	rightX       int
	rightY       int
	FirstButtons string
}
