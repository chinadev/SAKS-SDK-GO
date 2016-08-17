package entities

type Led74HC595 struct {
	IC *IC_74HC595
}

func (d *Led74HC595) IsOn(index uint) bool {
	if index > 7 {
		return false
	}
	return (d.IC.Data >> index) & 0x01 != 0
}

func (d *Led74HC595) RowStatus() []bool {
	var r []bool
	for i := uint(0); i < 8; i++ {
		r = append(r, d.IsOn(i))
	}
	return r
}

func (d *Led74HC595) On() {
	d.IC.SetData(0xff)
}

func (d *Led74HC595) Off() {
	d.IC.Clear()
}

func (d *Led74HC595) OnForIndex(index uint) {
	d.IC.SetData(d.IC.Data | (0x01 << index))
}

func (d *Led74HC595) OffForIndex(index uint) {
	arr := []uint8{0xfe, 0xfd, 0xfb, 0xf7, 0xef, 0xdf, 0xbf, 0x7f}
	d.IC.SetData(d.IC.Data & arr[index])
}

func (d *Led74HC595) SetRow(status [8]bool) {
	for i := range(status) {
		if status[i] {
			d.OnForIndex(uint(i))
		} else {
			d.OffForIndex(uint(i))
		}
	}
}