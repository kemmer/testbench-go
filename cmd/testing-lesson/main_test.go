package main

import (
	"reflect"
	"testing"
)

func TestInt64ToInt(t *testing.T) {
	expectedValue := 837733

	result, err := Int64ToInt(int64(expectedValue))

	if err != nil {
		t.Errorf("error when casting to int64: %s", err)
	}

	if reflect.TypeOf(result).Name() != "int" {
		t.Error("returned type is not int")
	}

	if result != expectedValue {
		t.Errorf("expected value differ, want = %v, got = %v", expectedValue, result)
	}
}
