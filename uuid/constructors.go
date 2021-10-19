package uuid

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
)

func newTimeBasedUUID(version int, variant string) *UUID {
	timestamp := getTimestampByVersion(version)

	if timestamp == 0 {
		log.Fatal("invalid timestamp")
	}

	return &UUID{
		TimeLow:            getTimeLow(timestamp),
		TimeMid:            getTimeMid(timestamp),
		TimeHighAndVersion: getTimeHighAndVersion(timestamp, version),
		ClockAndVariant:    getClockSequenceAndVariant(variant),
		Node:               getNode(),
	}
}

func newNamedBasedUUID(version int, namespace, name string) *UUID {
	availableNamespaces := getAvailableNamespaces()
	namespaceBuffer, hasNamespace := availableNamespaces[namespace]

	if !hasNamespace {
		log.Fatal("Invalid namespace")
	}

	hashFunc := getHashFuncByVersion(version)

	if hashFunc == nil {
		log.Fatal("Invalid version")
	}

	namespaceNameHash := hashNamespaceName(hashFunc, namespaceBuffer, []byte(name))
	namespaceHashUint := binary.BigEndian.Uint64(namespaceNameHash[:8])
	nameHashUint := binary.BigEndian.Uint64(namespaceNameHash[8:])

	return &UUID{
		TimeLow:            getTimeLowNamed(namespaceHashUint),
		TimeMid:            getTimeMidNamed(namespaceHashUint),
		TimeHighAndVersion: getTimeHighAndVersionNamed(nameHashUint, version),
		ClockAndVariant:    getClockSequenceAndVariantNamed(nameHashUint),
		Node:               getNodeNamed(nameHashUint),
	}
}

/*
	Creates a new UUID.

	The accepted args keys/values are:
			[variant]:
				microsoft | dce | ""
			[name]:
				valid string
			[namespace]:
				dns | url | oid | x500
*/
func NewUUID(version int, args map[string]string) (*UUID, error) {
	if version == 3 || version == 5 {
		if validateNamedUUIDArgs(args) {
			return newNamedBasedUUID(version, args["namespace"], args["name"]), nil
		}

		return nil, fmt.Errorf("invalid arguments for version %d", version)
	}

	if version == 1 || version == 4 {
		if validateTimeUUIDArgs(args) {
			return newTimeBasedUUID(version, args["variant"]), nil
		}

		return nil, fmt.Errorf("invalid arguments for version %d", version)
	}

	return nil, errors.New("invalid version")
}

/*
	Serialize an UUID to a hex string representation

	Output example:
		1cedfbc2-676a-499f-8079-b5177528c26a
*/
func (uuid *UUID) Stringify() string {
	return fmt.Sprintf("%08x-%04x-%02x-%02x-%06x",
		uuid.TimeLow, uuid.TimeMid, uuid.TimeHighAndVersion, uuid.ClockAndVariant, uuid.Node)
}
