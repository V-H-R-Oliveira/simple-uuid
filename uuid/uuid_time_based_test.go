package uuid

import (
	"testing"
)

func TestTimeBasedUUID(t *testing.T) {
	const (
		dce            = "dce"
		dceValue       = 2
		microsoft      = "microsoft"
		microsoftValue = 6
		future         = "future"
		futureValue    = 7
		v4             = 4
		v1             = 1
	)

	variants := []string{dce, microsoft, future}
	testParams := &testParams{
		recorder: t,
	}

	t.Run("Test v4", func(t *testing.T) {
		args := make(map[string]string)
		testParams.version = v4
		testParams.args = args

		t.Run("It should have version 4", func(t *testing.T) {
			testParams.executor = testVersion

			for _, variant := range variants {
				t.Logf("Testing with variant %s\n", variant)

				args["variant"] = variant
				got := basicUUIDTest(testParams).(uint16)

				if got != v4 {
					t.Errorf("Expected version %d, got %d\n", v4, got)
					break
				}
			}
		})

		t.Run("It should have the dce variant", func(t *testing.T) {
			testParams.executor = testDCEVariant
			args["variant"] = dce

			got := basicUUIDTest(testParams).(uint16)

			if got != dceValue {
				t.Errorf("Expected variant %d (%s), got %d\n", dceValue, dce, got)
			}
		})

		t.Run("It should have the microsoft variant", func(t *testing.T) {
			testParams.executor = testMicrosoftAndFutureVariant
			args["variant"] = microsoft

			got := basicUUIDTest(testParams).(uint16)

			if got != microsoftValue {
				t.Errorf("Expected variant %d (%s), got %d\n", microsoftValue, microsoft, got)
			}
		})

		t.Run("It should have the future variant", func(t *testing.T) {
			testParams.executor = testMicrosoftAndFutureVariant
			args["variant"] = future

			got := basicUUIDTest(testParams).(uint16)

			if got != futureValue {
				t.Errorf("Expected variant %d (%s), got %d\n", futureValue, future, got)
			}
		})

		t.Run("It should have the format of 4hexOctet-2hexOctet-2hexOctet-2hexOctet-6hexOctet", func(t *testing.T) {
			testParams.executor = testStringFormat

			for _, variant := range variants {
				t.Logf("Testing with variant %s\n", variant)

				args["variant"] = variant
				got := basicUUIDTest(testParams).(string)

				if got != "" {
					t.Error(got)
					break
				}
			}
		})

		t.Run("Test uuid v4 constructor", func(t *testing.T) {
			t.Run("It should return valid uuids and nil errors", func(t *testing.T) {
				args := make(map[string]string)
				args["variant"] = dce

				if uuid, err := NewUUID(v4, args); err != nil || uuid == nil {
					t.Errorf("Expected a nil error, got %s\n", err.Error())
				}

				args["variant"] = microsoft

				if uuid, err := NewUUID(v4, args); err != nil || uuid == nil {
					t.Errorf("Expected a nil error, got %s\n", err.Error())
				}

				args["variant"] = future

				if uuid, err := NewUUID(v4, args); err != nil || uuid == nil {
					t.Errorf("Expected a nil error, got %s\n", err.Error())
				}
			})

			t.Run("It should return nil uuids and non nil errors", func(t *testing.T) {
				args := make(map[string]string)
				args["variant"] = dce

				if uuid, err := NewUUID(-1, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				delete(args, "variant")

				if uuid, err := NewUUID(v4, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["variant"] = "xyz"

				if uuid, err := NewUUID(v4, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}
			})
		})
	})

	t.Run("Test v1", func(t *testing.T) {
		args := make(map[string]string)
		testParams.version = v1
		testParams.args = args

		t.Run("It should have version 1", func(t *testing.T) {
			testParams.executor = testVersion

			for _, variant := range variants {
				t.Logf("Testing with variant %s\n", variant)

				args["variant"] = variant
				got := basicUUIDTest(testParams).(uint16)

				if got != v1 {
					t.Errorf("Expected version %d, got %d\n", v1, got)
					break
				}
			}
		})

		t.Run("It should have the dce variant", func(t *testing.T) {
			testParams.executor = testDCEVariant
			args["variant"] = dce

			got := basicUUIDTest(testParams).(uint16)

			if got != dceValue {
				t.Errorf("Expected variant %d (%s), got %d\n", dceValue, dce, got)
			}
		})

		t.Run("It should have the microsoft variant", func(t *testing.T) {
			testParams.executor = testMicrosoftAndFutureVariant
			args["variant"] = microsoft

			got := basicUUIDTest(testParams).(uint16)

			if got != microsoftValue {
				t.Errorf("Expected variant %d (%s), got %d\n", microsoftValue, microsoft, got)
			}
		})

		t.Run("It should have the future variant", func(t *testing.T) {
			testParams.executor = testMicrosoftAndFutureVariant
			args["variant"] = future

			got := basicUUIDTest(testParams).(uint16)

			if got != futureValue {
				t.Errorf("Expected variant %d (%s), got %d\n", futureValue, future, got)
			}
		})

		t.Run("It should have the format of 4hexOctet-2hexOctet-2hexOctet-2hexOctet-6hexOctet", func(t *testing.T) {
			testParams.executor = testStringFormat

			for _, variant := range variants {
				t.Logf("Testing with variant %s\n", variant)

				args["variant"] = variant
				got := basicUUIDTest(testParams).(string)

				if got != "" {
					t.Error(got)
					break
				}
			}
		})

		t.Run("Test constructor", func(t *testing.T) {
			t.Run("It should return valid uuids and a nil errors", func(t *testing.T) {
				args := make(map[string]string)
				args["variant"] = dce

				if uuid, err := NewUUID(v1, args); err != nil || uuid == nil {
					t.Errorf("Expected a nil error, got %s\n", err.Error())
				}

				args["variant"] = microsoft

				if uuid, err := NewUUID(v1, args); err != nil || uuid == nil {
					t.Errorf("Expected a nil error, got %s\n", err.Error())
				}

				args["variant"] = future

				if uuid, err := NewUUID(v1, args); err != nil || uuid == nil {
					t.Errorf("Expected a nil error, got %s\n", err.Error())
				}
			})

			t.Run("It should return nil uuids and a non nil errors", func(t *testing.T) {
				args := make(map[string]string)
				args["variant"] = dce

				if uuid, err := NewUUID(-1, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				delete(args, "variant")

				if uuid, err := NewUUID(v1, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["variant"] = "xyz"

				if uuid, err := NewUUID(v1, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}
			})
		})
	})
}
