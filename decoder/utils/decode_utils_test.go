package utils

import (
	"bytes"
	"testing"
)

func TestDecodeShort(t *testing.T) {
	buf := bytes.NewBuffer([]byte{0x00, 0x01})
	val, err := DecodeShort(buf)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Unexpected value. got: %v, want: %v", val, 1)
	}
}

func TestDecodeByte(t *testing.T) {
	buf := bytes.NewBuffer([]byte{0x01})
	val, err := DecodeByte(buf)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Unexpected value. got: %v, want: %v", val, 1)
	}
}

func TestDecodeChars(t *testing.T) {
	buf := bytes.NewBuffer([]byte("test"))
	val, err := DecodeChars(buf, 4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != "test" {
		t.Errorf("Unexpected value. got: %v, want: %v", val, "test")
	}
}

func TestDecodeLong(t *testing.T) {
	buf := bytes.NewBuffer([]byte{0x00, 0x00, 0x00, 0x01})
	val, err := DecodeLong(buf)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if val != 1 {
		t.Errorf("Unexpected value. got: %v, want: %v", val, 1)
	}
}
