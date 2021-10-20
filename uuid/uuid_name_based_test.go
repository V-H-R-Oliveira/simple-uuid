package uuid

import (
	"strings"
	"testing"
)

func TestNameBasedUUID(t *testing.T) {
	t.Run("Test v3", func(t *testing.T) {
		t.Run("It should have version 3", func(t *testing.T) {
			namespaces := []string{"dns", "url", "oid", "x500"}

			for _, namespace := range namespaces {
				uuid := newNamedBasedUUID(3, namespace, "foo bar joe")
				version := uuid.TimeHighAndVersion >> 12

				if version != 3 {
					t.Errorf("Expect 3 got %d\n", version)
					break
				}
			}
		})

		t.Run("It should have the dce variant", func(t *testing.T) {
			namespaces := []string{"dns", "url", "oid", "x500"}

			for _, namespace := range namespaces {
				uuid := newNamedBasedUUID(3, namespace, "a random string")
				variant := uuid.ClockAndVariant >> 14

				if variant != 2 {
					t.Errorf("Expect 2 got %d\n", variant)
					break
				}
			}

		})

		t.Run("It should have the correct uuid format", func(t *testing.T) {
			namespaces := []string{"dns", "url", "oid", "x500"}

			for _, namespace := range namespaces {
				uuid := newNamedBasedUUID(3, namespace, "yet another string")
				uuidString := uuid.Stringify()
				splitted := strings.Split(uuidString, "-")

				if len(splitted[0]) != 8 {
					t.Errorf("Expect 8 got %d\n", len(splitted[0]))
					break
				}

				if len(splitted[1]) != 4 {
					t.Errorf("Expect 2 got %d\n", len(splitted[1]))
					break
				}

				if len(splitted[2]) != 4 {
					t.Errorf("Expect 2 got %d\n", len(splitted[2]))
					break
				}

				if len(splitted[3]) != 4 {
					t.Errorf("Expect 2 got %d\n", len(splitted[3]))
					break
				}

				if len(splitted[4]) != 12 {
					t.Errorf("Expect 12 got %d\n", len(splitted[4]))
					break
				}
			}
		})

		t.Run("Test constructor", func(t *testing.T) {
			args := make(map[string]string)
			args["name"] = "foo"
			args["namespace"] = "url"

			_, err := NewUUID(3, args)

			if err != nil {
				t.Errorf("Expected nil, got %s\n", err.Error())
			}

			uuid, err := NewUUID(-1, args)

			if uuid != nil || err == nil {
				t.Error("Expected err, got nil")
			}

			delete(args, "name")
			uuid, err = NewUUID(3, args)

			if uuid != nil || err == nil {
				t.Error("Expected err, got nil")
			}

			args["name"] = "another one"

			delete(args, "namespace")
			uuid, err = NewUUID(3, args)

			if uuid != nil || err == nil {
				t.Error("Expected err, got nil")
			}
		})
	})

	t.Run("Test v5", func(t *testing.T) {
		t.Run("It should have version 5", func(t *testing.T) {
			namespaces := []string{"dns", "url", "oid", "x500"}

			for _, namespace := range namespaces {
				uuid := newNamedBasedUUID(5, namespace, "foo bar joe")
				version := uuid.TimeHighAndVersion >> 12

				if version != 5 {
					t.Errorf("Expect 5 got %d\n", version)
					break
				}
			}
		})

		t.Run("It should have the dce variant", func(t *testing.T) {
			namespaces := []string{"dns", "url", "oid", "x500"}

			for _, namespace := range namespaces {
				uuid := newNamedBasedUUID(5, namespace, "a random string")
				variant := uuid.ClockAndVariant >> 14

				if variant != 2 {
					t.Errorf("Expect 2 got %d\n", variant)
					break
				}
			}

		})

		t.Run("It should have the correct uuid format", func(t *testing.T) {
			namespaces := []string{"dns", "url", "oid", "x500"}

			for _, namespace := range namespaces {
				uuid := newNamedBasedUUID(5, namespace, "yet another string")
				uuidString := uuid.Stringify()
				splitted := strings.Split(uuidString, "-")

				if len(splitted[0]) != 8 {
					t.Errorf("Expect 8 got %d\n", len(splitted[0]))
					break
				}

				if len(splitted[1]) != 4 {
					t.Errorf("Expect 2 got %d\n", len(splitted[1]))
					break
				}

				if len(splitted[2]) != 4 {
					t.Errorf("Expect 2 got %d\n", len(splitted[2]))
					break
				}

				if len(splitted[3]) != 4 {
					t.Errorf("Expect 2 got %d\n", len(splitted[3]))
					break
				}

				if len(splitted[4]) != 12 {
					t.Errorf("Expect 12 got %d\n", len(splitted[4]))
					break
				}
			}
		})

		t.Run("Test constructor", func(t *testing.T) {
			args := make(map[string]string)
			args["name"] = "foo"
			args["namespace"] = "url"

			_, err := NewUUID(5, args)

			if err != nil {
				t.Errorf("Expected nil, got %s\n", err.Error())
			}

			uuid, err := NewUUID(-1, args)

			if uuid != nil || err == nil {
				t.Error("Expected err, got nil")
			}

			delete(args, "name")
			uuid, err = NewUUID(5, args)

			if uuid != nil || err == nil {
				t.Error("Expected err, got nil")
			}

			args["name"] = "another one"

			delete(args, "namespace")
			uuid, err = NewUUID(5, args)

			if uuid != nil || err == nil {
				t.Error("Expected err, got nil")
			}
		})
	})
}