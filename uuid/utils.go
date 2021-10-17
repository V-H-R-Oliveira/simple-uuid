package uuid

import (
	"crypto/rand"
	"log"
)

func generateRandomBuffer(size int) []byte {
	buffer := make([]byte, 8)

	if _, err := rand.Read(buffer); err != nil {
		log.Fatal("Failed to read due error:", err)
	}

	return buffer
}
