package sakshat

import (
	"testing"
	"time"
)

//func TestBuzzer(t *testing.T) {
//	Buzzer.BeepAction(time.Second / 5, time.Second / 5, 2)
//}

func TestLEDRow(t *testing.T) {
	// Test for Index
	for i := uint(0); i < 8; i++ {
		LEDRow.OnForIndex(i)
		time.Sleep(time.Second)
		LEDRow.OffForIndex(i)
	}
	time.Sleep(time.Second)

	// Test for SetRow
	LED1 := [8]bool{false, true, false, true, false, true, false, true}
	LED2 := [8]bool{true, false, true, false, true, false, true, false}
	LEDRow.SetRow(LED1)
	time.Sleep(time.Second)
	LEDRow.SetRow(LED2)
	time.Sleep(time.Second)

	// Test for status acquiring
	LED3 := LEDRow.RowStatus()
	t.Log(LED2 == LED3)

	// Test for All
	LEDRow.On()
	time.Sleep(time.Second)
	LEDRow.Off()
}