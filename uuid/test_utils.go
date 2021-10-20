package uuid

import (
	"errors"
	"math"
	"strings"
	"testing"
)

type testParams struct {
	recorder testing.TB
	version  int
	executor func(*UUID) interface{}
	args     map[string]string
}

func newUUID(recorder testing.TB, version int, args map[string]string) (*UUID, error) {
	if version == 1 || version == 4 {
		if !validateTimeUUIDArgs(args) {
			recorder.Errorf("invalid arguments for version %d.\nExpected \"variant\"\n", version)
		}

		return newTimeBasedUUID(version, args["variant"])
	}

	if version == 3 || version == 5 {
		if !validateNamedUUIDArgs(args) {
			recorder.Errorf("invalid arguments for version %d.\nExpected \"namespace\" and \"name\"\n", version)
		}

		return newNamedBasedUUID(version, args["namespace"], args["name"])
	}

	recorder.Errorf("version %d is not a valid one.\n", version)
	return nil, errors.New("invalid version")
}

func basicUUIDTest(params *testParams) interface{} {
	uuid, err := newUUID(params.recorder, params.version, params.args)

	if err != nil {
		return uint16(math.MaxUint16)
	}

	return params.executor(uuid)
}

func testVersion(uuid *UUID) interface{} {
	return uuid.TimeHighAndVersion >> 12
}

func testDCEVariant(uuid *UUID) interface{} {
	return uuid.ClockAndVariant >> 14
}

func testMicrosoftAndFutureVariant(uuid *UUID) interface{} {
	return uuid.ClockAndVariant >> 13
}

func testStringFormat(uuid *UUID) interface{} {
	uuidString := uuid.Stringify()
	uuidComponents := strings.Split(uuidString, "-")

	if err := validateStringifiedUUID(uuidComponents); err != nil {
		return err.Error()
	}

	return ""
}
