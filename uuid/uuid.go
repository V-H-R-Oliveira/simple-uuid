package uuid

import (
	"encoding/binary"
	"time"
)

func GenV1Timestamp() uint64 {
	now := time.Now()
	return UUID_TIMESTAMP + uint64(now.UTC().UnixNano()/100)
}

func genTimestamp() uint64 {
	timestamp := generateRandomBuffer(8)
	return binary.LittleEndian.Uint64(timestamp)
}

func genClockSequence() int64 {
	clockSequence := generateRandomBuffer(2)
	return int64(binary.LittleEndian.Uint16(clockSequence))
}

func getTimeLow(timestamp uint64) uint32 {
	return uint32(timestamp & 0xffffffff)
}

func getTimeMid(timestamp uint64) uint16 {
	return uint16((timestamp >> 32) & 0xffff)
}

func getTimeHighAndVersion(timestamp uint64, version int) uint16 {
	timeHigh := int64(timestamp >> 48)
	timeHigh = ^(^timeHigh & SET_4MSB)

	switch version {
	case 1:
		return uint16(timeHigh & V1)
	case 4:
		return uint16(timeHigh & V4)
	default:
		return 0
	}
}

func getClockSequenceAndVariant(variant string) uint16 {
	clock := genClockSequence()
	clock = ^(^clock & SET_3MSB)

	switch variant {
	case "dce":
		return uint16(clock & DCE)
	case "microsoft":
		return uint16(clock & MICROSOFT)
	default:
		return uint16(clock) // Future definition
	}
}

func getNode() []byte {
	node := generateRandomBuffer(8)
	return node[:6]
}
