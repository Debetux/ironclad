package main

// KeyboardEvent
// Status: 0 none, 1 down, 2 up
type KeyboardEvent struct {
	KeyCode int
	Status  int
}

// CursorPostion's struct received when new events are sent by the client
type CursorPosition struct {
	PosX            int
	PosY            int
	MouseLeftStatus int
	Keys            []KeyboardEvent
}
