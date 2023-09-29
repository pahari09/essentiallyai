package core

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecode(t *testing.T) {
	decoder := &BinaryDecoder{}

	packet := "\x04\xD2\x6B\x65\x65\x70\x64\x65\x63\x6F\x64\x69\x6E\x67\x38\x64\x6F\x6E\x74\x73\x74\x6F\x70\x03\x15\x63\x6F\x6E\x67\x72\x61\x74\x75\x6C\x61\x74\x69\x6F\x6E\x73\x07\x5B\xCD\x15"
	buffer := bytes.NewBufferString(packet)

	decodedStruct, err := decoder.Decode(buffer)

	assert.NoError(t, err)
	assert.Equal(t, int16(1234), decodedStruct.GetShort1())
	assert.Equal(t, "keepdecoding", decodedStruct.GetChars1())
	assert.Equal(t, int16(56), decodedStruct.GetByte1())
	assert.Equal(t, "dontstop", decodedStruct.GetChars2())
	assert.Equal(t, int16(789), decodedStruct.GetShort2())
	assert.Equal(t, "congratulations", decodedStruct.GetChars3())
	assert.Equal(t, int32(123456789), decodedStruct.GetLong1())
}
