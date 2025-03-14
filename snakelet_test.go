package snakelet

import (
	"strconv"
	"testing"
)

// TestSuccessfulUnmarshal calls unmarshal on a testStruct to check if all fields are filled correctly.
func TestSuccessfulUnmarshal(t *testing.T) {
	type testStruct struct {
		TestBool   bool
		TestString string
		TestInt    int
	}

	const (
		testBoolValue   = true
		testStringValue = "Hello World!"
		testIntValue    = 1001
	)

	t.Setenv("TEST_BOOL", strconv.FormatBool(testBoolValue))
	t.Setenv("TEST_STRING", testStringValue)
	t.Setenv("TEST_INT", strconv.Itoa(testIntValue))

	test := testStruct{}
	if err := Unmarshal(&test); err != nil {
		t.Errorf("failed to unmarshal test struct (%v): %v", test, err)
	}

	if test.TestBool != testBoolValue {
		t.Errorf("test.TestBool(%v) does not equal testBoolValue (%v)", test.TestBool, testBoolValue)
	}

	if test.TestString != testStringValue {
		t.Errorf("test.TestString(%v) does not equal testStringValue (%v)", test.TestString, testStringValue)
	}

	if test.TestInt != testIntValue {
		t.Errorf("test.TestInt(%v) does not equal testIntValue (%v)", test.TestInt, testIntValue)
	}
}

func TestCustomName(t *testing.T) {
	type testStruct struct {
		TestName string `snakelet:"name=CUSTOM_TEST_NAME"`
	}

	testValue := "SUCCESS"
	t.Setenv("CUSTOM_TEST_NAME", testValue)

	testCustomName := testStruct{}
	if err := Unmarshal(&testCustomName); err != nil {
		t.Errorf("failed to unmarshal test struct (%v): %v", testCustomName, err)
	}

	if testCustomName.TestName != testValue {
		t.Errorf("testCustomName.TestName(%v) does not equal testValue (%v)", testCustomName.TestName, testValue)
	}
}

// Tests for:
//  - Env var key not been set
//  - Failed to parse env var value
