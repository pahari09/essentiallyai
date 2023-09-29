package _interface

type DecodedPacket interface {
	GetShort1() int16
	GetChars1() string
	GetByte1() byte
	GetChars2() string
	GetShort2() int16
	GetChars3() string
	GetLong1() int32
}
