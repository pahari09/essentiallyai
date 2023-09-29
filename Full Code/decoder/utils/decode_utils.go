package utils

import (
	"bytes"
	"encoding/binary"
)

func DecodeShort(reader *bytes.Buffer) (int16, error) {
	var val int16
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func DecodeByte(reader *bytes.Buffer) (byte, error) {
	var val byte
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}

func DecodeChars(reader *bytes.Buffer, length int) (string, error) {
	val := make([]byte, length)
	if _, err := reader.Read(val); err != nil {
		return "", err
	}
	return string(val), nil
}

func DecodeLong(reader *bytes.Buffer) (int32, error) {
	var val int32
	if err := binary.Read(reader, binary.BigEndian, &val); err != nil {
		return 0, err
	}
	return val, nil
}
