package snakelet

import (
	"strconv"
	"testing"
)

type testStruct struct {
	TestBool   bool
	TestString string
	TestInt    int
}

// TestSuccessfulUnmarshal calls unmarshal on a testStruct to check if all fields are filled correctly.
func TestSuccessfulUnmarshal(t *testing.T) {
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

// Tests for: 
//  - Env var key not been set
//  - Failed to parse env var value

