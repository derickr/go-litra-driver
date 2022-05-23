package litra

import (
	"encoding/binary"
	"bytes"
	"math"

	"github.com/karalabe/usb"
)

const (
	vendor =  0x046d
	product = 0xc900
)

type LitraDevice struct {
	dev usb.Device
}

func New() (*LitraDevice, error) {
	d := &LitraDevice{}

	usbDevices, _ := usb.Enumerate(vendor, product)
	dev, err := usbDevices[0].Open()
	d.dev = dev
	
	d.SetBrightness(0)
	d.SetTemperature(4500)

	return d, err
}

func getSwitchOn() []byte {
	return []byte {0x11, 0xff, 0x04, 0x1c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func getSwitchOff() []byte {
	return []byte {0x11, 0xff, 0x04, 0x1c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func getSetBrightness(level int) []byte {
	minBrightness := 0x14
	maxBrightness := 0xfa

	value := minBrightness + ((level/100) * (maxBrightness - minBrightness));
	adjusted_level := byte(math.Floor(float64(value)))

	return []byte {0x11, 0xff, 0x04, 0x4c, 0x00, adjusted_level, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func getSetTemperature(temp int16) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, temp)
	byte0, _ := buf.ReadByte()
	byte1, _ := buf.ReadByte()

	return []byte {0x11, 0xff, 0x04, 0x9c, byte0, byte1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func (d *LitraDevice) TurnOn() {
	var dummy []byte

	d.dev.Write(getSwitchOn())
	d.dev.Read(dummy)
}

func (d *LitraDevice) TurnOff() {
	var dummy []byte

	d.dev.Write(getSwitchOff())
	d.dev.Read(dummy)
}

func (d *LitraDevice) SetBrightness(level int) {
	var dummy []byte

	d.dev.Write(getSetBrightness(level))
	d.dev.Read(dummy)
}

func (d *LitraDevice) SetTemperature(temp int16) {
	var dummy []byte

	d.dev.Write(getSetTemperature(temp))
	d.dev.Read(dummy)
}

func (d *LitraDevice) Close() {
	d.dev.Close()
}
