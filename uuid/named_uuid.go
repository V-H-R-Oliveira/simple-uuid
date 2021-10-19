package uuid

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"hash"
	"log"
)

func getTimeLowNamed(hash uint64) uint32 {
	return uint32(hash >> 32)
}

func getTimeMidNamed(hash uint64) uint16 {
	return uint16((hash >> 16) & 0xffff)
}

func getTimeHighAndVersionNamed(hash uint64, version int) uint16 {
	timeHigh := int64(hash & 0xffff)
	timeHigh = ^(^timeHigh & SET_4MSB)

	if version == 3 {
		return uint16(timeHigh & V3)
	}

	return uint16(timeHigh & V5)
}

func getClockSequenceAndVariantNamed(hash uint64) uint16 {
	variant := ^(^int64(hash>>48) & SET_3MSB)
	return uint16(variant & DCE_VARIANT)
}

func getNodeNamed(hash uint64) []byte {
	node := hash & 0xffffffffffff
	buffer := &bytes.Buffer{}

	if err := binary.Write(buffer, binary.BigEndian, node); err != nil {
		log.Fatal("Failed to get node named due error:", err)
	}

	return bytes.Trim(buffer.Bytes(), "\x00")
}

func getHashFuncByVersion(version int) hash.Hash {
	switch version {
	case 3:
		return md5.New()
	case 5:
		return sha1.New()
	default:
		return nil
	}
}
