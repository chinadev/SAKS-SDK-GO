package entities

import "github.com/stianeikeland/go-rpio"

type IC_74HC595 struct {
	Pins map[string]rpio.Pin
	RealTrue rpio.State
	Data uint8
}

func (d *IC_74HC595) FlushSHCP() {
	d.Pins["shcp"].Write(d.RealTrue ^ 0x01)
	d.Pins["shcp"].Write(d.RealTrue)
}

func (d *IC_74HC595) FlushSTCP() {
	d.Pins["stcp"].Write(d.RealTrue ^ 0x01)
	d.Pins["stcp"].Write(d.RealTrue)
}

func (d *IC_74HC595) SetBit(bit rpio.State) {
	d.Pins["ds"].Write(bit)
	d.FlushSHCP()
}

func (d *IC_74HC595) SetDate(data uint8) {
	d.Data = data
	for i := 0; i < 8; i++ {
		d.SetBit((d.Data >> i) & 0x01)
	}
	d.FlushSTCP()
}

func (d *IC_74HC595) Clear(data uint8) {
	d.SetDate(0x00)
}