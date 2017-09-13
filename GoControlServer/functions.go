package main

import (
	"fmt"
	"strconv"

	"github.com/AllenDang/w32"
)

func simulateMouse(mouseEventType int) {
	var inputs []w32.INPUT

	inputs = append(inputs, w32.INPUT{
		Type: w32.INPUT_MOUSE,
		Mi: w32.MOUSEINPUT{
			Dx:          0,
			Dy:          0,
			MouseData:   0,
			DwFlags:     uint32(mouseEventType),
			Time:        0,
			DwExtraInfo: 0,
		},
	})

	w32.SendInput(inputs)
	return
}

func simulateKeys(keys []KeyboardEvent) {
	var inputs []w32.INPUT

	for _, key := range keys {
		keyCode, _ := strconv.ParseInt(key.KeyCode, 0, 64)

		if keyCode == 0x7f {
			key.KeyCode = "0x0e"
			inputs = append(inputs, generateSpecialKey(key))
			continue
		}

		if key.Status == 2 {
			continue
		}

		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         0,
				WScan:       uint16(keyCode),
				DwFlags:     0x0004,
				Time:        0,
				DwExtraInfo: 0,
			},
		})
	}

	w32.SendInput(inputs)
	return
}

func generateSpecialKey(key KeyboardEvent) w32.INPUT {
	var dwFlag uint32
	var input w32.INPUT
	var keyCode int64

	keyCode, _ = strconv.ParseInt(key.KeyCode, 0, 64)
	fmt.Println("SPECIAL KEY")

	switch key.Status {
	case 1:
		dwFlag = w32.WM_KEYDOWN
	case 2:
		dwFlag = w32.WM_KEYUP
	default:
		dwFlag = w32.WM_KEYUP
	}

	input = w32.INPUT{
		Type: w32.INPUT_KEYBOARD,
		Ki: w32.KEYBDINPUT{
			WVk:         uint16(keyCode),
			WScan:       0,
			DwFlags:     dwFlag,
			Time:        0,
			DwExtraInfo: 0,
		},
	}

	return input
}
