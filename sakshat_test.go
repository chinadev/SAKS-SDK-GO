package sakshat

import (
	"testing"
	"time"
)

func TestLEDRow(t *testing.T) {
	for i := uint(0); i < 8; i++ {
		LEDRow.OnForIndex(i)
		time.Sleep(time.Second * 2)
		LEDRow.OffForIndex(i)
	}
}