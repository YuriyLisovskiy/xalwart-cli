package core

import "testing"

func TestRandomString_CheckLength(t *testing.T) {
	expectedLength := 256
	randString, _ := RandomString(uint(expectedLength))
	actualLength := len(randString)
	if actualLength != expectedLength {
		t.Errorf("Expected %d, received %d", expectedLength, actualLength)
	}
}
