package uuid

import (
	"errors"
	"fmt"
	"strings"
)

func newV4(variant string) *UUID {
	timestamp := genTimestamp()

	return &UUID{
		TimeLow:            getTimeLow(timestamp),
		TimeMid:            getTimeMid(timestamp),
		TimeHighAndVersion: getTimeHighAndVersion(timestamp, 4),
		ClockAndVariant:    getClockSequenceAndVariant(variant),
		Node:               getNode(),
	}
}

func newV1(variant string) *UUID {
	timestamp := GenV1Timestamp()

	return &UUID{
		TimeLow:            getTimeLow(timestamp),
		TimeMid:            getTimeMid(timestamp),
		TimeHighAndVersion: getTimeHighAndVersion(timestamp, 1),
		ClockAndVariant:    getClockSequenceAndVariant(variant),
		Node:               getNode(),
	}
}

func NewUUID(version int, variant string) (*UUID, error) {
	variant = strings.ToLower(variant)

	switch version {
	case 1:
		return newV1(variant), nil
	case 4:
		return newV4(variant), nil
	default:
		return nil, errors.New("invalid uuid version")
	}
}

func (uuid *UUID) Stringify() string {
	return fmt.Sprintf("%08x-%04x-%02x-%02x-%06x",
		uuid.TimeLow, uuid.TimeMid, uuid.TimeHighAndVersion, uuid.ClockAndVariant, uuid.Node)
}
