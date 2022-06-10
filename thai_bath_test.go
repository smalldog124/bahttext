package bahttext_test

import (
	"bahttext"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_0_should_be_zero_bath(t *testing.T) {
	expected := "ศูนย์บาท"
	baht := 0.0

	actual := bahttext.THBText(baht)

	assert.Equal(t, expected, actual)
}

func Test_0_dot_25_should_be_zero_dot_twenty_five_satang_bath(t *testing.T) {
	expected := "ศูนย์บาทยี้สิบห้าสตางค์"
	baht := 0.25

	actual := bahttext.THBText(baht)

	assert.Equal(t, expected, actual)
}
