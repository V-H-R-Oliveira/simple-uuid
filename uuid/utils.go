package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"hash"
	"log"
)

func generateRandomBuffer(size int) []byte {
	buffer := make([]byte, size)

	if _, err := rand.Read(buffer); err != nil {
		log.Fatal("Failed to read due error:", err)
	}

	return buffer
}

func hashNamespaceName(hashFunc hash.Hash, namespace, name []byte) []byte {
	if err := binary.Write(hashFunc, binary.BigEndian, namespace); err != nil {
		log.Fatal("Failed to write the namespace due error:", err)
	}

	if err := binary.Write(hashFunc, binary.BigEndian, name); err != nil {
		log.Fatal("Failed to write the name due error:", err)
	}

	return hashFunc.Sum(nil)
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
