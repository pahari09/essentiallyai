package core

import (
	"bytes"
	"fmt"
	"rssentiallyai/decoder/interface"
	"rssentiallyai/decoder/model"
	"rssentiallyai/decoder/utils"
)

type Decoder interface {
	Decode(*bytes.Buffer) (_interface.DecodedPacket, error)
}

type BinaryDecoder struct{}

func (bd *BinaryDecoder) Decode(packet *bytes.Buffer) (_interface.DecodedPacket, error) {
	reader := bytes.NewBufferString(packet.String())

	short1, err := utils.DecodeShort(reader)
	if err != nil {
		return nil, fmt.Errorf("error decoding short1: %w", err)
	}

	chars1, err := utils.DecodeChars(reader, 12)
	if err != nil {
		return nil, fmt.Errorf("error decoding chars1: %w", err)
	}

	byte1, err := utils.DecodeByte(reader)
	if err != nil {
		return nil, fmt.Errorf("error decoding byte1: %w", err)
	}

	chars2, err := utils.DecodeChars(reader, 8)
	if err != nil {
		return nil, fmt.Errorf("error decoding chars2: %w", err)
	}

	short2, err := utils.DecodeShort(reader)
	if err != nil {
		return nil, fmt.Errorf("error decoding short2: %w", err)
	}

	chars3, err := utils.DecodeChars(reader, 15)
	if err != nil {
		return nil, fmt.Errorf("error decoding chars3: %w", err)
	}

	long1, err := utils.DecodeLong(reader)
	if err != nil {
		return nil, fmt.Errorf("error decoding long1: %w", err)
	}

	return model.NewDecodedPacket(short1, chars1, byte1, chars2, short2, chars3, long1), nil
}
