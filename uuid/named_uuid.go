package uuid

import (
	"bytes"
	"encoding/binary"
	"log"
)

func getTimeHighAndVersionNamed(hash int64) uint16 {
	hash = ^(^hash & SET_4MSB)
	return uint16(hash & V3)
}

func getClockSequenceAndVariantNamed(hash int64) uint16 {
	variant := ^(^hash & SET_3MSB)
	return uint16(variant & DCE)
}

func getNodeNamed(hash uint64) []byte {
	node := hash & 0xffffffffffff
	buffer := &bytes.Buffer{}

	if err := binary.Write(buffer, binary.BigEndian, node); err != nil {
		log.Fatal("Failed to get node named due error:", err)
	}

	return buffer.Bytes()[:6]
}