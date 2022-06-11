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
func Test_20_should_be_twenty_bath(t *testing.T) {
	expected := "ศูนย์บาท"
	baht := 20.0

	actual := bahttext.THBText(baht)

	assert.Equal(t, expected, actual)
}
func Test_0_dot_25_should_be_zero_dot_twenty_five_satang_bath(t *testing.T) {
	expected := "ศูนย์บาทยี้สิบห้าสตางค์"
	baht := 0.25

	actual := bahttext.THBText(baht)

	assert.Equal(t, expected, actual)
}

func Test_2000000000000_should_be_tow_hundred_million_bath(t *testing.T) {
	expected := "สองล้านล้านบาทถ้วน"
	baht := 2000000000000.0

	actual := bahttext.THBText(baht)

	assert.Equal(t, expected, actual)
}

func Test_2345_dot_01_should_be_two_thousand_three_hundred_forty_five_bath(t *testing.T) {
	expected := "สองพันสามร้อยสี่สิบห้าบาทหนึ่งสตางค์"
	baht := 2345.01

	actual := bahttext.THBText(baht)

	assert.Equal(t, expected, actual)
}
