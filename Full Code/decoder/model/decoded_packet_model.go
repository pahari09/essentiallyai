package model

import (
	"rssentiallyai/decoder/interface"
)

type decodedPacket struct {
	short1 int16
	chars1 string
	byte1  byte
	chars2 string
	short2 int16
	chars3 string
	long1  int32
}

func NewDecodedPacket(short1 int16, chars1 string, byte1 byte, chars2 string, short2 int16, chars3 string, long1 int32) _interface.DecodedPacket {
	return &decodedPacket{
		short1: short1,
		chars1: chars1,
		byte1:  byte1,
		chars2: chars2,
		short2: short2,
		chars3: chars3,
		long1:  long1,
	}
}

func (d *decodedPacket) GetShort1() int16 {
	return d.short1
}

func (d *decodedPacket) GetChars1() string {
	return d.chars1
}

func (d *decodedPacket) GetByte1() byte {
	return d.byte1
}

func (d *decodedPacket) GetChars2() string {
	return d.chars2
}

func (d *decodedPacket) GetShort2() int16 {
	return d.short2
}

func (d *decodedPacket) GetChars3() string {
	return d.chars3
}

func (d *decodedPacket) GetLong1() int32 {
	return d.long1
}
