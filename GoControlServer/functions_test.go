package main

import (
	"testing"
	"time"

	"github.com/AllenDang/w32"
)

func TestSimulateMouse(t *testing.T) {
	simulateMouse(w32.MOUSEEVENTF_LEFTDOWN)
	simulateMouse(w32.MOUSEEVENTF_LEFTUP)
}
func TestSimulateKeys(t *testing.T) {
	time.Sleep(250 * time.Millisecond)

	var keyCode = 0x52

	var keys = []KeyboardEvent{
		KeyboardEvent{
			KeyCode: keyCode,
			Status:  0x0004 | 0x0003,
		},
	}
	simulateKeys(keys)
}
