package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)


type DecodedStruct struct {
	Short1  int16
	Chars1  string
	Byte1   byte
	Chars2  string
	Short2  int16
	Chars3  string
	Long1   int32
}

func decodePacket(packet *bytes.Buffer) (*DecodedStruct, error) {
	var decoded DecodedStruct

	reader := bytes.NewBufferString(packet.String())

	// Read short1
	binary.Read(reader, binary.BigEndian, &decoded.Short1)

	// Read and convert chars1 to a string
	chars1 := make([]byte, 12)
	reader.Read(chars1)
	decoded.Chars1 = string(chars1)

	// Read byte1
	binary.Read(reader, binary.BigEndian, &decoded.Byte1)

	// Read and convert chars2 to a string
	chars2 := make([]byte, 8)
	reader.Read(chars2)
	decoded.Chars2 = string(chars2)

	// Read short2
	binary.Read(reader, binary.BigEndian, &decoded.Short2)

	// Read and convert chars3 to a string
	chars3 := make([]byte, 15)
	reader.Read(chars3)
	decoded.Chars3 = string(chars3)

	// Read long1
	binary.Read(reader, binary.BigEndian, &decoded.Long1)

	return &decoded, nil
}

func main() {
        
    packet := "\x04\xD2\x6B\x65\x65\x70\x64\x65\x63\x6F\x64\x69\x6E\x67\x38\x64\x6F\x6E\x74\x73\x74\x6F\x70\x03\x15\x63\x6F\x6E\x67\x72\x61\x74\x75\x6C\x61\x74\x69\x6F\x6E\x73\x07\x5B\xCD\x15"
    buffer := bytes.NewBufferString(packet)

    decodedStruct,err := decodePacket(buffer)

    if err != nil {
        fmt.Println("Error decoding packet:", err)
    } else {
        fmt.Printf("Decoded struct:{%d, %q, %d, %q, %d, %q, %d}\n", decodedStruct.Short1, decodedStruct.Chars1, decodedStruct.Byte1, decodedStruct.Chars2, decodedStruct.Short2, decodedStruct.Chars3, decodedStruct.Long1)
    }
}