package main

import (
	"testing"
)

var (
	s  = NewSite("SUS Lampoon")
	c1 = NewCard(s, "123456", "PURSE")

	p1  = NewDevice("Paybox1", "BOX", "P001")
	v1  = NewDevice("VENDOR1", "VENDOR", "V001")
	v2  = NewDevice("VENDOR2", "VENDOR", "V002")
	v3  = NewDevice("VENDOR3", "VENDOR", "V003")
	sh1 = NewShop(s, "ร้านข้าวมันไก่", v1)
)

func TestNewCard(t *testing.T) {
	if s.Name != "SUS Lampoon" {
		t.Error("Expected name = 'SUS Lampoon'")
	}
	if c1.Code != "123456" {
		t.Error("Expected code = '123456'")
	}
}

func TestCardBalance(t *testing.T) {
	c := NewCard(s, "123456", "PURSE")
	if c.Debit(100); c.balance != 100 {
		t.Error("Expected debit = 100 card balance = 100 ")
	}
	if c.Credit(30); c.balance != 70 {
		t.Error("Expected credit 30 should decrese balance = 70")
	}
}

func TestDeviceBalance(t *testing.T) {
	if p1.Debit(100); p1.balance != 100 {
		t.Error("Expected debit = 100 box balance = 100 ")
	}
	if p1.Credit(30); p1.balance != 70 {
		t.Error("Expected credit 30 box balance = 70")
	}
	// ลองเอาบัตรซื้อของอาหารราคา 40 บาท
	if v1.Debit(40); v1.balance != 40 {
		t.Error("Expected debit  40 vendor balance = 40")
	}
	if v1.Credit(5); v1.balance != 35 {
		t.Error("Expected credit 5 vendor balance = 35")
	}
}

func TestTransCardDeviceBalance(t *testing.T) {
	trans := new(Trans)

	trans.Job1CardDeposit(c1, p1, p1, 50, 100, 50) // เติมเงิน 50 ใส่เงิน 100 ทอน 50
	if c1.balance != 50 || c1.credit != 50 || p1.debit != 50 || p1.balance != 50 {
		t.Error("Expected เติมเงิน 100 ต้องมีเงินในบัตรเพิ่ม 100 เครดิต เงินในตู้เพิ่ม 100เป็น Debit")
	}

	trans.Job3ShopSales(c1, v1, p1, 20)
	if c1.debit != 20 || c1.balance != 30 || v1.credit != 20 || v1.balance != 20 {
		t.Error("Expected ซื้อของ 20 เดบิตเงินในบัตร 20 คงเหลือ 30 เครดิต เงิน Vendor เพิ่ม 20 คงเหลือ 20")
	}
}
