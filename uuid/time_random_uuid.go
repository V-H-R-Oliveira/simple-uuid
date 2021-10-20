package uuid

import (
	"encoding/binary"
	"errors"
	"time"
)

func genV1Timestamp() uint64 {
	now := time.Now()
	return UUID_TIMESTAMP + uint64(now.UTC().UnixNano()/100)
}

func genRandomTimestamp() (uint64, error) {
	timestamp, err := generateRandomBuffer(8)

	if err != nil {
		return 0, nil
	}

	return binary.LittleEndian.Uint64(timestamp), nil
}

func genClockSequence() (int64, error) {
	clockSequence, err := generateRandomBuffer(2)

	if err != nil {
		return 0, err
	}

	return int64(binary.LittleEndian.Uint16(clockSequence)), nil
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

	if version == 1 {
		return uint16(timeHigh & V1)
	}

	return uint16(timeHigh & V4)
}

func getClockSequenceAndVariant(variant string) (uint16, error) {
	clock, err := genClockSequence()

	if err != nil {
		return 0, err
	}

	clock = ^(^clock & SET_3MSB)

	switch variant {
	case "dce":
		return uint16(clock & DCE_VARIANT), nil
	case "microsoft":
		return uint16(clock & MICROSOFT_VARIANT), nil
	case "future":
		return uint16(clock), nil
	default:
		return 0, errors.New("invalid variant")
	}
}

func getNode() ([]byte, error) {
	node, err := generateRandomBuffer(8)
	return node[:6], err
}

func getTimestampByVersion(version int) (uint64, error) {
	switch version {
	case 1:
		return genV1Timestamp(), nil
	case 4:
		return genRandomTimestamp()
	default:
		return 0, errors.New("invalid version")
	}
}
