package sakshat

import (
	"testing"
	"time"
)

func TestBuzzer(t *testing.T) {
	Buzzer.BeepAction(time.Second / 5, time.Second / 5, 3)
}