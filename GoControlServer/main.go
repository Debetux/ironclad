package main

import (
	"encoding/json"
	"fmt"

	"github.com/AllenDang/w32"
	"github.com/firstrow/tcp_server"
)

func main() {

	server := tcp_server.New("0.0.0.0:8080")

	server.OnNewClient(func(c *tcp_server.Client) {
		c.Send("Hello")
		fmt.Println("Hello")
	})

	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		// new message received
		c.Send("Hi")
		content := []byte(message)
		cursor := CursorPosition{}
		json.Unmarshal(content, &cursor)
		// fmt.Println(cursor)
		// fmt.Println(string(content))
		// fmt.Println(len(cursor.Keys))
		w32.SetCursorPos(cursor.PosX, cursor.PosY)

		if cursor.MouseLeftStatus == 1 {
			fmt.Println("Click")
			simulateMouse(w32.MOUSEEVENTF_LEFTDOWN)
		}

		if cursor.MouseLeftStatus == 2 {
			fmt.Println("up")
			simulateMouse(w32.MOUSEEVENTF_LEFTUP)
		}

		if len(cursor.Keys) > 0 {
			simulateKeys(cursor.Keys)
		}
	})

	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		// connection with client lost
		c.Send("Goodbye")
		fmt.Println("Goodbye")
	})

	server.Listen()
}

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
		var dwFlag uint32

		switch key.Status {
		case 1:
			dwFlag = w32.WM_KEYDOWN
		case 2:
			dwFlag = w32.WM_KEYUP
		default:
			dwFlag = w32.WM_KEYUP
		}

		inputs = append(inputs, w32.INPUT{
			Type: w32.INPUT_KEYBOARD,
			Ki: w32.KEYBDINPUT{
				WVk:         uint16(key.KeyCode),
				WScan:       0,
				DwFlags:     dwFlag,
				Time:        0,
				DwExtraInfo: 0,
			},
		})
	}

	w32.SendInput(inputs)
}
