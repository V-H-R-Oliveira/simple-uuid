package uuid

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"strconv"
	"strings"
)

func generateRandomBuffer(size int) ([]byte, error) {
	buffer := make([]byte, size)
	_, err := rand.Read(buffer)
	return buffer, err
}

func hashNamespaceName(hashFunc hash.Hash, namespace, name []byte) ([]byte, error) {
	if err := binary.Write(hashFunc, binary.BigEndian, namespace); err != nil {
		return []byte{}, err
	}

	if err := binary.Write(hashFunc, binary.BigEndian, name); err != nil {
		return []byte{}, err
	}

	return hashFunc.Sum(nil), nil
}

func validateTimeUUIDArgs(args map[string]string) bool {
	_, hasVariant := args["variant"]
	return hasVariant
}

func validateNamedUUIDArgs(args map[string]string) bool {
	_, hasName := args["name"]
	_, hasNamespace := args["namespace"]
	return hasName && hasNamespace
}

func getAvailableNamespaces() map[string][]byte {
	namespaces := make(map[string][]byte)

	namespaces["dns"] = DNS_NAMESPACE
	namespaces["url"] = URL_NAMESPACE
	namespaces["oid"] = OID_NAMESPACE
	namespaces["x500"] = X500_NAMESPACE

	return namespaces
}

func validateStringifiedUUID(uuidComponents []string) error {
	componentsQtd := len(uuidComponents)

	if componentsQtd != 5 {
		return fmt.Errorf("mismatch between length 5 and %d", componentsQtd)
	}

	expectedHexOctects := 0

	if len(uuidComponents[0]) != 8 {
		expectedHexOctects = (len(uuidComponents[0]) / 2)
		return fmt.Errorf("mismatch between 4 hex numbers and %d", expectedHexOctects)
	}

	if len(uuidComponents[1]) != 4 {
		expectedHexOctects = (len(uuidComponents[1]) / 2)
		return fmt.Errorf("mismatch between 2 hex numbers and %d", expectedHexOctects)
	}

	if len(uuidComponents[2]) != 4 {
		expectedHexOctects = (len(uuidComponents[2]) / 2)
		return fmt.Errorf("mismatch between 2 hex numbers and %d", expectedHexOctects)
	}

	if len(uuidComponents[3]) != 4 {
		expectedHexOctects = (len(uuidComponents[3]) / 2)
		return fmt.Errorf("mismatch between 2 hex numbers and %d", expectedHexOctects)
	}

	if len(uuidComponents[4]) != 12 {
		expectedHexOctects = (len(uuidComponents[4]) / 2)
		return fmt.Errorf("mismatch between 6 hex numbers and %d", expectedHexOctects)
	}

	return nil
}

func uuidToBytes(uuid string) ([]byte, error) {
	splittedUUID := strings.Split(uuid, "-")

	if err := validateStringifiedUUID(splittedUUID); err != nil {
		return []byte{}, err
	}

	buffer := &bytes.Buffer{}

	for _, uuidPart := range splittedUUID {
		bitsize := (len(uuidPart) / 2) * 8
		uuidComponent, err := strconv.ParseUint(uuidPart, 16, bitsize)

		if err != nil {
			return []byte{}, err
		}

		if err := binary.Write(buffer, binary.BigEndian, uuidComponent); err != nil {
			return []byte{}, err
		}
	}

	if buffer.Len() == 0 {
		return buffer.Bytes(), errors.New("invalid namespace buffer length")
	}

	return buffer.Bytes(), nil
}
