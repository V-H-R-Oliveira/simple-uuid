package main

import (
	"fmt"
	"log"

	"github.com/V-H-R-Oliveira/simple-uuid/uuid"
)

func main() {
	const customNamespace = "2163d569-2c70-43d4-bb87-ff9c58814ade"

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
		log.Fatal("Failed to create an uuidv5 due error:", err)
	}

	args["namespace"] = customNamespace
	uuidv3CustomNamespace, err := uuid.NewUUID(3, args)

	if err != nil {
		log.Fatal("Failed to create an uuidv3 with a custom namespace due error:", err)
	}

	uuidv5CustomNamespace, err := uuid.NewUUID(5, args)

	if err != nil {
		log.Fatal("Failed to create an uuidv5 with a custom namespace due error:", err)
	}

	fmt.Printf(
		"V1: %s\nV3: %s\nV4: %s\nV5: %s\nV3 with a custom namespace: %s\nV5 with a custom namespace: %s\n",
		uuidv1.Stringify(),
		uuidv3.Stringify(),
		uuidv4.Stringify(),
		uuidv5.Stringify(),
		uuidv3CustomNamespace.Stringify(),
		uuidv5CustomNamespace.Stringify(),
	)
}
