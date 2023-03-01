package testasserts

import "testing"

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf(err.Error())
	}
}

func StringEquals(t *testing.T, actual string, expected string) {
	if actual != expected {
		t.Errorf("Expected \"%s\" but got \"%s\"", expected, actual)
	}
}
