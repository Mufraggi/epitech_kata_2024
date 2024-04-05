package muf

import "testing"

func TestIsLeapYearDivide400(t *testing.T) {
	isLeapYear := IsLeapYear(2000)
	if isLeapYear != true {
		t.Errorf("isLeapYear returned false for 2000")
	}
}
func TestIsLeapYearDivide400ButNotBy100(t *testing.T) {
	isLeapYear := IsLeapYear(1900)
	if isLeapYear != true {
		t.Errorf("isLeapYear returned false for 1900")
	}
}

func TestIsLeapYearDivide4ButNotBy100(t *testing.T) {
	isLeapYear := IsLeapYear(2008)
	if isLeapYear != true {
		t.Errorf("isLeapYear returned false for 2008")
	}
}

func TestIsNotLeapYear(t *testing.T) {
	isLeapYear := IsLeapYear(2017)
	if isLeapYear != false {
		t.Errorf("isLeapYear returned true for 2017")
	}
}
