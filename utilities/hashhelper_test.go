package utilities_test

import (
	"testing"

	"github.com/Gartenschlaeger/strcli/utilities"
	"github.com/stretchr/testify/assert"
)

type hashTestCase struct {
	s        string
	expected string
}

func TestMD5(t *testing.T) {
	testCases := []hashTestCase{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"1234", "81dc9bdb52d04dc20036dbd8313ed055"},
		{"abcd", "e2fc714c4727ee9395f324cd2e7f331f"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.GetMD5Hash(tc.s)

			assert.Equal(t, tc.expected, r)
		})
	}
}

func TestSHA1(t *testing.T) {
	testCases := []hashTestCase{
		{"", "da39a3ee5e6b4b0d3255bfef95601890afd80709"},
		{"1234", "7110eda4d09e062aa5e4a390b0a572ac0d2c0220"},
		{"abcd", "81fe8bfe87576c3ecb22426f8e57847382917acf"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.GetSHA1Hash(tc.s)

			assert.Equal(t, tc.expected, r)
		})
	}
}

func TestSHA256(t *testing.T) {
	testCases := []hashTestCase{
		{"", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"1234", "03ac674216f3e15c761ee1a5e255f067953623c8b388b4459e13f978d7c846f4"},
		{"abcd", "88d4266fd4e6338d13b845fcf289579d209c897823b9217da3e161936f031589"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.GetSHA256Hash(tc.s)

			assert.Equal(t, tc.expected, r)
		})
	}
}

func TestSHA512(t *testing.T) {
	testCases := []hashTestCase{
		{"", "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		{"1234", "d404559f602eab6fd602ac7680dacbfaadd13630335e951f097af3900e9de176b6db28512f2e000b9d04fba5133e8b1c6e8df59db3a8ab9d60be4b97cc9e81db"},
		{"abcd", "d8022f2060ad6efd297ab73dcc5355c9b214054b0d1776a136a669d26a7d3b14f73aa0d0ebff19ee333368f0164b6419a96da49e3e481753e7e96b716bdccb6f"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.GetSHA512Hash(tc.s)

			assert.Equal(t, tc.expected, r)
		})
	}
}

func TestBase64Encode(t *testing.T) {
	testCases := []hashTestCase{
		{"", ""},
		{"1234", "MTIzNA=="},
		{"abcd", "YWJjZA=="},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r := utilities.Base64Encode(tc.s)

			assert.Equal(t, tc.expected, r)
		})
	}
}

func TestBase64Decode(t *testing.T) {
	testCases := []hashTestCase{
		{"", ""},
		{"MTIzNA==", "1234"},
		{"YWJjZA==", "abcd"},
	}

	for i, tc := range testCases {
		t.Run(utilities.PadInt(i, 2), func(t *testing.T) {
			r, err := utilities.Base64Decode(tc.s)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, r)
		})
	}
}
