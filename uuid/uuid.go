package uuid

import "encoding/binary"

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

func getTimeHighAndVersion(timestamp uint64) uint16 {
	timeHigh := int64(timestamp >> 48)
	timeHigh = ^(^timeHigh & SET_4MSB)
	return uint16(timeHigh & V4)
}

func getClockSequenceAndVariant() uint16 {
	clock := genClockSequence()
	clock = ^(^clock & SET_3MSB)
	return uint16(clock & ^(1 << 14))
}

func getNode() []byte {
	node := generateRandomBuffer(8)
	return node[:6]
}
