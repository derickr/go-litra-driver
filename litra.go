package litra

import (
	"bytes"
	"encoding/binary"
	"math"
	"time"

	"github.com/karalabe/usb"
)

const (
	vendor  = 0x046d
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

	return d, err
}

func getSwitchOn() []byte {
	return []byte{0x11, 0xff, 0x04, 0x1c, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func getSwitchOff() []byte {
	return []byte{0x11, 0xff, 0x04, 0x1c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func getSetBrightness(level int) []byte {
	minBrightness := float64(0x14)
	maxBrightness := float64(0xfa)

	if level < 0 {
		level = 0
	}
	if level > 100 {
		level = 100
	}

	value := minBrightness + ((float64(level) / 100) * (maxBrightness - minBrightness))
	adjusted_level := byte(math.Floor(float64(value)))

	return []byte{0x11, 0xff, 0x04, 0x4c, 0x00, adjusted_level, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func getSetTemperature(temp int16) []byte {
	if temp < 2700 {
		temp = 2700
	}
	if temp > 6500 {
		temp = 6500
	}

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, temp)
	byte0, _ := buf.ReadByte()
	byte1, _ := buf.ReadByte()

	return []byte{0x11, 0xff, 0x04, 0x9c, byte0, byte1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
}

func (d *LitraDevice) TurnOn() {
	var dummy []byte

	d.dev.Write(getSwitchOn())
	d.dev.Read(dummy)
	time.Sleep(30 * time.Millisecond)
}

func (d *LitraDevice) TurnOff() {
	var dummy []byte

	d.dev.Write(getSwitchOff())
	d.dev.Read(dummy)
	time.Sleep(30 * time.Millisecond)
}

func (d *LitraDevice) SetBrightness(level int) {
	var dummy []byte

	d.dev.Write(getSetBrightness(level))
	d.dev.Read(dummy)
	time.Sleep(30 * time.Millisecond)
}

func (d *LitraDevice) SetTemperature(temp int16) {
	var dummy []byte

	d.dev.Write(getSetTemperature(temp))
	d.dev.Read(dummy)
	time.Sleep(30 * time.Millisecond)
}

func (d *LitraDevice) Close() {
	d.dev.Close()
}
