package uuid

import (
	"encoding/binary"
	"errors"
	"fmt"
)

func newTimeBasedUUID(version int, variant string) (*UUID, error) {
	timestamp, err := getTimestampByVersion(version)

	if err != nil {
		return nil, err
	}

	node, err := getNode()

	if err != nil {
		return nil, err
	}

	clockAndVariant, err := getClockSequenceAndVariant(variant)

	if err != nil {
		return nil, err
	}

	return &UUID{
		TimeLow:            getTimeLow(timestamp),
		TimeMid:            getTimeMid(timestamp),
		TimeHighAndVersion: getTimeHighAndVersion(timestamp, version),
		ClockAndVariant:    clockAndVariant,
		Node:               node,
	}, nil
}

func newNamedBasedUUID(version int, namespace, name string) (*UUID, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid name length")
	}

 	availableNamespaces := getAvailableNamespaces()
	namespaceBuffer, hasNamespace := availableNamespaces[namespace]

	if !hasNamespace {
		var err error
		namespaceBuffer, err = uuidToBytes(namespace)

		if err != nil {
			return nil, err
		}
	}

	hashFunc, err := getHashFuncByVersion(version)

	if err != nil {
		return nil, err
	}

	namespaceNameHash, err := hashNamespaceName(hashFunc, namespaceBuffer, []byte(name))

	if err != nil {
		return nil, err
	}

	namespaceHashUint := binary.BigEndian.Uint64(namespaceNameHash[:8])
	nameHashUint := binary.BigEndian.Uint64(namespaceNameHash[8:])

	node, err := getNodeNamed(nameHashUint)

	if err != nil {
		return nil, err
	}

	return &UUID{
		TimeLow:            getTimeLowNamed(namespaceHashUint),
		TimeMid:            getTimeMidNamed(namespaceHashUint),
		TimeHighAndVersion: getTimeHighAndVersionNamed(nameHashUint, version),
		ClockAndVariant:    getClockSequenceAndVariantNamed(nameHashUint),
		Node:               node,
	}, nil
}

/*
	Creates a new UUID. Custom namespaces must be a valid string representation of an UUID.

	The accepted args keys/values are:
			[variant]:
				microsoft | dce | future
			[name]:
				valid string
			[namespace]:
				dns | url | oid | x500 | custom
*/
func NewUUID(version int, args map[string]string) (*UUID, error) {
	if version == 3 || version == 5 {
		if validateNamedUUIDArgs(args) {
			return newNamedBasedUUID(version, args["namespace"], args["name"])
		}

		return nil, fmt.Errorf("invalid arguments for version %d", version)
	}

	if version == 1 || version == 4 {
		if validateTimeUUIDArgs(args) {
			return newTimeBasedUUID(version, args["variant"])
		}

		return nil, fmt.Errorf("invalid arguments for version %d", version)
	}

	return nil, errors.New("invalid version")
}

/*
	Serialize an UUID to a hex string representation.

	Output example:
		1cedfbc2-676a-499f-8079-b5177528c26a
*/
func (uuid *UUID) Stringify() string {
	return fmt.Sprintf("%08x-%04x-%02x-%02x-%06x",
		uuid.TimeLow, uuid.TimeMid, uuid.TimeHighAndVersion, uuid.ClockAndVariant, uuid.Node)
}
