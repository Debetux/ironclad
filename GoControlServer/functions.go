package main

import "github.com/AllenDang/w32"

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

		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         0,
				WScan:       uint16(key.KeyCode),
				DwFlags:     0x0004,
				Time:        0,
				DwExtraInfo: 0,
			},
		})
	}

	w32.SendInput(inputs)
	return
}
