package uuid

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"log"
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

func newV3(namespace, name string) *UUID {
	availableNamespaces := getAvailableNamespaces()
	namespaceBuffer, hasNamespace := availableNamespaces[namespace]

	if !hasNamespace {
		log.Fatal("Invalid namespace")
	}

	namespaceNameHash := hashNamespaceName(md5.New(), namespaceBuffer, []byte(name))
	namespaceHashUint := binary.BigEndian.Uint64(namespaceNameHash[:8])
	nameHashUint := binary.BigEndian.Uint64(namespaceNameHash[8:])

	return &UUID{
		TimeLow:            uint32(namespaceHashUint >> 32),
		TimeMid:            uint16((namespaceHashUint >> 16) & 0xffff),
		TimeHighAndVersion: getTimeHighAndVersionNamed(int64(nameHashUint & 0xffff)),
		ClockAndVariant:    getClockSequenceAndVariantNamed(int64(nameHashUint >> 48)),
		Node:               getNodeNamed(nameHashUint & 0xffffffffffff),
	}
}

func NewUUID(version int, args map[string]string) (*UUID, error) {
	switch version {
	case 1:
		if validateTimeUUIDArgs(args) {
			return newV1(args["variant"]), nil
		}

		return nil, errors.New("invalid args for version 1")
	case 3:
		if validateNamedUUIDArgs(args) {
			return newV3(args["namespace"], args["name"]), nil
		}

		return nil, errors.New("invalid args for version 3")
	case 4:
		if validateTimeUUIDArgs(args) {
			return newV4(args["variant"]), nil
		}

		return nil, errors.New("invalid args for version 4")
	default:
		return nil, errors.New("invalid uuid version")
	}
}

func (uuid *UUID) Stringify() string {
	return fmt.Sprintf("%08x-%04x-%02x-%02x-%06x",
		uuid.TimeLow, uuid.TimeMid, uuid.TimeHighAndVersion, uuid.ClockAndVariant, uuid.Node)
}
