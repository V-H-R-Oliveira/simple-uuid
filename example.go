package main

import (
	"fmt"
	"log"

	"github.com/V-H-R-Oliveira/simple-uuid/uuid"
)

func main() {
	args := make(map[string]string)
	args["variant"] = "dce"

	uuidv4, err := uuid.NewUUID(4, args)

	if err != nil {
		log.Fatal("Failed to create an uuidv4 due error:", err)
	}

	uuidv1, err := uuid.NewUUID(1, args)

	if err != nil {
		log.Fatal("Failed to create an uuidv1 due error:", err)
	}

	args["name"] = "test"
	args["namespace"] = "dns"

	uuidv3, err := uuid.NewUUID(3, args)

	if err != nil {
		log.Fatal("Failed to create an uuidv3 due error:", err)
	}

	args["name"] = "foo"
	args["namespace"] = "url"

	uuidv5, err := uuid.NewUUID(5, args)

	if err != nil {
		log.Fatal("Failed to creatte an uuidv5 due error:", err)
	}

	fmt.Printf(
		"V1: %s\nV3: %s\nV4: %s\nV5: %s\n",
		uuidv1.Stringify(),
		uuidv3.Stringify(),
		uuidv4.Stringify(),
		uuidv5.Stringify(),
	)
}
