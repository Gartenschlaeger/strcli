package strutilities_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/pkg/strutilities"
	"github.com/Gartenschlaeger/strcli/pkg/testasserts"
)

func TestFirstFieldWithSpaceSeparator(t *testing.T) {
	input := "Das ist ein Test"

	field, err := strutilities.GetField(input, 1, " ")

	testasserts.NoError(t, err)
	testasserts.StringEquals(t, field, "Das")
}
