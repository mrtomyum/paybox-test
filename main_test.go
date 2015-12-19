package main

import (
	"testing"
)

var (
	s  = NewSite("SUS Lampoon")
	c1 = NewCard(s, "123456", "PURSE")

	p1  = NewDevice("Paybox1", "BOX", "P001")
	v1  = NewDevice("V1", "VENDOR", "V001")
	v2  = NewDevice("V2", "VENDOR", "V002")
	v3  = NewDevice("V3", "VENDOR", "V003")
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
	if c.Credit(100); c.balance != 100 {
		t.Error("Expected credit 30 should decrese balance = 70")
	}
	if c.Debit(30); c.balance != 70 {
		t.Error("Expected debit = 100 card balance = 100 ")
	}
}

func TestDeviceBalance(t *testing.T) {
	p := NewDevice("Paybox1", "BOX", "P001")
	if p.Debit(100); p.balance != 100 {
		t.Error("Expected debit = 100 box balance = 100 ")
	}
	if p.Credit(30); p.balance != 70 {
		t.Error("Expected credit 30 box balance = 70")
	}
	// ลองเอาบัตรซื้อของอาหารราคา 40 บาท
	v := NewDevice("V1", "VENDOR", "V001")
	if v.Debit(40); v.balance != 40 {
		t.Error("Expected debit  40 vendor balance = 40")
	}
	if v.Credit(5); v.balance != 35 {
		t.Error("Expected credit 5 vendor balance = 35")
	}
}

func TestTransCalCardDeviceBalance(t *testing.T) {
	trans := new(Trans)

	trans.Job1CardDeposit( // เติมเงิน 50 ใส่เงิน 100 ทอน 50
		c1,  // Card
		p1,  // Device Paybox1
		p1,  // Host Paybox1
		50,  // Value 50
		100, // CashReceive 10
		50,  // Change
	)
	if c1.credit != 50 ||
		c1.balance != 50 ||
		p1.debit != 50 ||
		p1.balance != 50 {
		t.Errorf("Expected เติมเงิน 100 ต้องมีเงินในบัตรเพิ่ม c1.credit != 50 got %v|| c1.balance != 50 got %v||  p1.debit != 50 got %v|| p1.balance != 50 got %v", c1.credit, c1.balance, p1.debit, p1.balance)
	}

	trans.Job3ShopSales(
		c1, // Card
		v1, // Device Vendor1
		p1, // Host Paybox1
		20, // Value 20
	)
	if c1.debit != 20 ||
		c1.balance != 30 ||
		v1.credit != 20 ||
		v1.balance != -20 {
		t.Errorf("Expected ซื้อของ 20 เดบิตเงินในบัตร 20 got %v บัตรคงเหลือ 30 got %v เครดิตเงิน Vendor เพิ่ม 20 got %v คงเหลือ -20 got %v", c1.debit, c1.balance, v1.credit, v1.balance)
	}
}

// Test Interface Balancer implement Method Debit(), Credit()
func TestBalancer(t *testing.T) {

}
