package uuid

import (
	"strings"
	"testing"
)

const (
	dce            = "dce"
	microsoft      = "microsoft"
	defaultVariant = ""
)

func TestUUID(t *testing.T) {
	t.Run("Test v4", func(t *testing.T) {
		t.Run("It should have version 4", func(t *testing.T) {
			uuid := newV4(dce)
			version := uuid.TimeHighAndVersion >> 12

			if version != 4 {
				t.Errorf("Expect 4 got %d\n", version)
			}
		})

		t.Run("It should have the dce variant", func(t *testing.T) {
			uuid := newV4(dce)
			variant := uuid.ClockAndVariant >> 14

			if variant != 2 {
				t.Errorf("Expect 2 got %d\n", variant)
			}
		})

		t.Run("It should have the microsoft variant", func(t *testing.T) {
			uuid := newV4(microsoft)
			variant := uuid.ClockAndVariant >> 13

			if variant != 6 {
				t.Errorf("Expect 6 got %d\n", variant)
			}
		})

		t.Run("It should have the future variant", func(t *testing.T) {
			uuid := newV4(defaultVariant)
			variant := uuid.ClockAndVariant >> 13

			if variant != 7 {
				t.Errorf("Expect 7 got %d\n", variant)
			}
		})

		t.Run("It should have the correct uuid format", func(t *testing.T) {
			uuid := newV4(dce)
			uuidString := uuid.Stringify()
			splitted := strings.Split(uuidString, "-")

			if len(splitted[0]) != 8 {
				t.Errorf("Expect 8 got %d\n", len(splitted[0]))
			}

			if len(splitted[1]) != 4 {
				t.Errorf("Expect 2 got %d\n", len(splitted[1]))
			}

			if len(splitted[2]) != 4 {
				t.Errorf("Expect 2 got %d\n", len(splitted[2]))
			}

			if len(splitted[3]) != 4 {
				t.Errorf("Expect 2 got %d\n", len(splitted[3]))
			}

			if len(splitted[4]) != 12 {
				t.Errorf("Expect 12 got %d\n", len(splitted[4]))
			}
		})
	})

	t.Run("Test v1", func(t *testing.T) {
		t.Run("It should have version 1", func(t *testing.T) {
			uuid := newV1(dce)
			version := uuid.TimeHighAndVersion >> 12

			if version != 1 {
				t.Errorf("Expect 1 got %d\n", version)
			}
		})

		t.Run("It should have the dce variant", func(t *testing.T) {
			uuid := newV1(dce)
			variant := uuid.ClockAndVariant >> 14

			if variant != 2 {
				t.Errorf("Expect 2 got %d\n", variant)
			}
		})

		t.Run("It should have the microsoft variant", func(t *testing.T) {
			uuid := newV1(microsoft)
			variant := uuid.ClockAndVariant >> 13

			if variant != 6 {
				t.Errorf("Expect 6 got %d\n", variant)
			}
		})

		t.Run("It should have the future variant", func(t *testing.T) {
			uuid := newV1(defaultVariant)
			variant := uuid.ClockAndVariant >> 13

			if variant != 7 {
				t.Errorf("Expect 7 got %d\n", variant)
			}
		})

		t.Run("It should have the correct uuid format", func(t *testing.T) {
			uuid := newV1(dce)
			uuidString := uuid.Stringify()
			splitted := strings.Split(uuidString, "-")

			if len(splitted[0]) != 8 {
				t.Errorf("Expect 8 got %d\n", len(splitted[0]))
			}

			if len(splitted[1]) != 4 {
				t.Errorf("Expect 2 got %d\n", len(splitted[1]))
			}

			if len(splitted[2]) != 4 {
				t.Errorf("Expect 2 got %d\n", len(splitted[2]))
			}

			if len(splitted[3]) != 4 {
				t.Errorf("Expect 2 got %d\n", len(splitted[3]))
			}

			if len(splitted[4]) != 12 {
				t.Errorf("Expect 12 got %d\n", len(splitted[4]))
			}
		})
	})
}
