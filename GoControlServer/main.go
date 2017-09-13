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
