package main

import (
	"testing"
)

var (
	s = NewSite("SUS Lampoon")
	c1 = NewCard(s, "123456", "PURSE")

	p1 = NewDevice("Paybox1", "BOX", "P001")
	v1 = NewDevice("VENDOR1", "VENDOR", "V001")
	v2 = NewDevice("VENDOR2", "VENDOR", "V002")
	v3 = NewDevice("VENDOR3", "VENDOR", "V003")
	t1 = Job1CardDeposit(c1, p1, p1, 50, 100, 50) // เติมเงิน 50 ใส่เงิน 100 ทอน 50
	t2 = Job3ShopSales(c1, v1, p1, 20,)
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
	if c1.Debit(100); c1.balance != 100 {
		t.Error("Expected debit = 100 card balance = 100 ")
	}
	if c1.Credit(30); c1.balance != 70 {
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
	t1.JobCardDeposit(c1, p1, 100)
	if c1.balance && c1.credit && p1.balance != 100 {
		t.Error("Expected เติมเงิน 100 ต้องมีเงินในบัตรเพิ่ม 100 เครดิต เงินในตู้เพิ่ม 100เป็น Debit")
	}
}
