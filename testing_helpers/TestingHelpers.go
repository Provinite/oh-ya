package testing_helpers

import "testing"

func ExpectToHaveLen(t *testing.T, actual int, expectedLength int) {
	if actual != expectedLength {
		t.Fatalf("Expected length %d, got %d", expectedLength, actual)
	}
}

func ExpectToEqual(t *testing.T, actual interface{}, expected interface{}) {
	if actual != expected {
		t.Fatalf("Expected %s to equal %s", actual, expected)
	}
}
