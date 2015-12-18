package main

import (
	"testing"
)

var (
	s = NewSite("SUS Lampoon")
	c = NewCard(s, "123456", "PURSE")

	p1 = NewDevice("Paybox1", "BOX", "P001")
	v1 = NewDevice("VENDOR1", "VENDOR", "V001")
	v2 = NewDevice("VENDOR2", "VENDOR", "V002")
	v3 = NewDevice("VENDOR3", "VENDOR", "V003")

)
func TestNewCard(t *testing.T) {
	if s.Name != "SUS Lampoon" {
		t.Error("Expected name = 'SUS Lampoon'")
	}
	if c.Code != "123456" {
		t.Error("Expected code = '123456'")
	}
}


func TestCardBalance(t *testing.T) {
	if c.Debit(100); c.balance != 100 {
		t.Error("Expected debit = 100 card balance = 100 ")
	}
	if c.Credit(30); c.balance != 70 {
		t.Error("Expected credit 50 should decrese balance = 70")
	}
}

func TestDeviceBalance(t *testing.T) {
	if p1.Debit(100); p1.balance != 100 {
		t.Error("Expected debit = 100 card balance = 100 ")
	 }
	if p1.Credit(30); p1.balance != 70 {
		t.Error("Expected debit = 30 card balance = 70")
	}
}