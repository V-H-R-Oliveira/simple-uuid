package uuid

import (
	"crypto/rand"
	"log"
)

func generateRandomBuffer(size int) []byte {
	buffer := make([]byte, size)

	if _, err := rand.Read(buffer); err != nil {
		log.Fatal("Failed to read due error:", err)
	}

	return buffer
}
