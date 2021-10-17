package uuid

import (
	"errors"
	"fmt"
)

func newV4() *UUID {
	timestamp := genTimestamp()

	return &UUID{
		TimeLow:            getTimeLow(timestamp),
		TimeMid:            getTimeMid(timestamp),
		TimeHighAndVersion: getTimeHighAndVersion(timestamp, 4),
		ClockAndVariant:    getClockSequenceAndVariant(),
		Node:               getNode(),
	}
}

func newV1() *UUID {
	timestamp := GenV1Timestamp()

	return &UUID{
		TimeLow:            getTimeLow(timestamp),
		TimeMid:            getTimeMid(timestamp),
		TimeHighAndVersion: getTimeHighAndVersion(timestamp, 1),
		ClockAndVariant:    getClockSequenceAndVariant(),
		Node:               getNode(),
	}
}

func NewUUID(version int) (*UUID, error) {
	switch version {
	case 1:
		return newV1(), nil
	case 4:
		return newV4(), nil
	default:
		return nil, errors.New("invalid uuid version")
	}
}

func (uuid *UUID) Stringify() string {
	return fmt.Sprintf("%08x-%04x-%02x-%02x-%06x",
		uuid.TimeLow, uuid.TimeMid, uuid.TimeHighAndVersion, uuid.ClockAndVariant, uuid.Node)
}
