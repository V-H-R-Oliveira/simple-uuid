package uuid

import "testing"

func TestNameBasedUUID(t *testing.T) {
	const (
		dns             = "dns"
		url             = "url"
		oid             = "oid"
		x500            = "x500"
		dce             = "dce"
		dceValue        = 2
		customNamespace = "0c31efb8-d0c8-4b49-8962-0e2c51b65b96"
		v3              = 3
		v5              = 5
	)

	namespaces := []string{dns, url, oid, x500, customNamespace}

	testParams := &testParams{
		recorder: t,
	}

	t.Run("Test v3", func(t *testing.T) {
		args := make(map[string]string)
		testParams.version = v3
		testParams.args = args

		t.Run("It should have version 3", func(t *testing.T) {
			testParams.executor = testVersion

			for _, namespace := range namespaces {
				t.Logf("Testing with namespace %s\n", namespace)

				args["namespace"] = namespace
				args["name"] = "foo"
				got := basicUUIDTest(testParams).(uint16)

				if got != v3 {
					t.Errorf("Expected version %d, got %d\n", v3, got)
					break
				}
			}
		})

		t.Run("It should have the dce variant", func(t *testing.T) {
			testParams.executor = testDCEVariant

			for _, namespace := range namespaces {
				t.Logf("Testing with namespace %s\n", namespace)

				args["namespace"] = namespace
				args["name"] = "bar"
				got := basicUUIDTest(testParams).(uint16)

				if got != dceValue {
					t.Errorf("Expected variant %d (%s), got %d\n", dceValue, dce, got)
				}
			}
		})

		t.Run("It should have the format of 4hexOctet-2hexOctet-2hexOctet-2hexOctet-6hexOctet", func(t *testing.T) {
			testParams.executor = testStringFormat

			for _, namespace := range namespaces {
				t.Logf("Testing with namespace %s\n", namespace)

				args["namespace"] = namespace
				args["name"] = "go"
				got := basicUUIDTest(testParams).(string)

				if got != "" {
					t.Error(got)
					break
				}
			}
		})

		t.Run("Test uuid v3 constructor", func(t *testing.T) {
			t.Run("It should return valid uuids and a nil errors", func(t *testing.T) {
				args := make(map[string]string)

				args["name"] = "foo"
				args["namespace"] = dns

				if uuid, err := NewUUID(v3, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}

				args["name"] = "john doe"
				args["namespace"] = url

				if uuid, err := NewUUID(v3, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}

				args["name"] = "xyz"
				args["namespace"] = oid

				if uuid, err := NewUUID(v3, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}

				args["name"] = "test"
				args["namespace"] = x500

				if uuid, err := NewUUID(v3, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}

				args["name"] = "bar"
				args["namespace"] = customNamespace

				if uuid, err := NewUUID(v3, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}
			})

			t.Run("It should return nil uuids and non nil errors", func(t *testing.T) {
				if uuid, err := NewUUID(-1, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				delete(args, "name")

				if uuid, err := NewUUID(v3, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["name"] = "another one"
				delete(args, "namespace")

				if uuid, err := NewUUID(v3, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["name"] = ""

				if uuid, err := NewUUID(v3, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["namespace"] = "foo bar"
				args["name"] = "x"

				if uuid, err := NewUUID(v3, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["namespace"] = "2163d-2c70-43d4-bb87-ff9c58814ade"
				args["name"] = "mock"

				if uuid, err := NewUUID(v3, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["namespace"] = customNamespace + "-" + customNamespace
				args["name"] = "asdf"

				if uuid, err := NewUUID(v3, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}
			})
		})
	})

	t.Run("Test v5", func(t *testing.T) {
		args := make(map[string]string)
		testParams.version = v5
		testParams.args = args

		t.Run("It should have version 5", func(t *testing.T) {
			testParams.executor = testVersion

			for _, namespace := range namespaces {
				t.Logf("Testing with namespace %s\n", namespace)

				args["namespace"] = namespace
				args["name"] = "foo"
				got := basicUUIDTest(testParams).(uint16)

				if got != v5 {
					t.Errorf("Expected version %d, got %d\n", v5, got)
					break
				}
			}
		})

		t.Run("It should have the dce variant", func(t *testing.T) {
			testParams.executor = testDCEVariant

			for _, namespace := range namespaces {
				t.Logf("Testing with namespace %s\n", namespace)

				args["namespace"] = namespace
				args["name"] = "bar"
				got := basicUUIDTest(testParams).(uint16)

				if got != dceValue {
					t.Errorf("Expected variant %d (%s), got %d\n", dceValue, dce, got)
				}
			}
		})

		t.Run("It should have the format of 4hexOctet-2hexOctet-2hexOctet-2hexOctet-6hexOctet", func(t *testing.T) {
			testParams.executor = testStringFormat

			for _, namespace := range namespaces {
				t.Logf("Testing with namespace %s\n", namespace)

				args["namespace"] = namespace
				args["name"] = "go"
				got := basicUUIDTest(testParams).(string)

				if got != "" {
					t.Error(got)
					break
				}
			}
		})

		t.Run("Test uuid v5 constructor", func(t *testing.T) {
			t.Run("It should return valid uuids and a nil errors", func(t *testing.T) {
				args := make(map[string]string)

				args["name"] = "foo"
				args["namespace"] = dns

				if uuid, err := NewUUID(v5, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}

				args["name"] = "john doe"
				args["namespace"] = url

				if uuid, err := NewUUID(v5, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}

				args["name"] = "xyz"
				args["namespace"] = oid

				if uuid, err := NewUUID(v5, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}

				args["name"] = "test"
				args["namespace"] = x500

				if uuid, err := NewUUID(v5, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}

				args["name"] = "bar"
				args["namespace"] = customNamespace

				if uuid, err := NewUUID(v5, args); uuid == nil || err != nil {
					t.Errorf("Expected nil error, got %s\n", err.Error())
				}
			})

			t.Run("It should return nil uuids and non nil errors", func(t *testing.T) {
				if uuid, err := NewUUID(-1, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				delete(args, "name")

				if uuid, err := NewUUID(v5, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["name"] = "another one"
				delete(args, "namespace")

				if uuid, err := NewUUID(v5, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["name"] = ""

				if uuid, err := NewUUID(v5, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["namespace"] = "foo bar"
				args["name"] = "x"

				if uuid, err := NewUUID(v5, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["namespace"] = "2163d-2c70-43d4-bb87-ff9c58814ade"
				args["name"] = "mock"

				if uuid, err := NewUUID(v5, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}

				args["namespace"] = customNamespace + "-" + customNamespace
				args["name"] = "asdf"

				if uuid, err := NewUUID(v5, args); uuid != nil || err == nil {
					t.Error("Expected an error, got nil")
				}
			})
		})
	})
}
