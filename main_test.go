package main

import (
	"fmt"
	"testing"
)

var (
	s  = NewSite("SUS Lampoon")
	c1 = NewCard(s, "123456", "PURSE")
	pb = NewDevice("Paybox1", "BOX", "P001")
	v1 = NewDevice("V1", "VENDOR", "V001")

//	v2  = NewDevice("V2", "VENDOR", "V002")
//	v3  = NewDevice("V3", "VENDOR", "V003")
//	sh1 = NewShop(s, "ร้านข้าวมันไก่", v1)
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
	if c.Credit(100); c.balance != -100 {
		t.Error("Expected credit -100")
	}
	if c.Debit(30); c.balance != -70 {
		t.Error("Expected card balance = -70 ")
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

func Test_TransJob1_CardDeposit(t *testing.T) {
	// เติมเงินเข้าบัตรใหม่
	value := 100
	cash := 100
	tn := new(Trans)
	tn.Job1_CardDeposit( // เติมเงิน 50 ใส่เงิน 100 ทอน 50
		c1,    // Card
		pb,    // Device Paybox1
		pb,    // Host Paybox1
		value, // Value 50
		cash,  // CashReceive 10
	)
	if c1.balance != -100 ||
		pb.balance != 100 {
		t.Errorf("Expected เติมเงิน 100 ต้องมีเงินในบัตรเพิ่ม c1.credit != 50 got %v|| c1.balance != 50 got %v||  p1.debit != 50 got %v|| p1.balance != 50 got %v", c1.credit, c1.balance, pb.debit, pb.balance)
	}
	fmt.Println(
		"1.เติมเงินเข้าบัตร: ",
		"Value =", value,
		"c1.balance =", c1.balance,
		"p1.balance =", pb.balance,
		"tn.change =", tn.change,
	)
}

func Test_TransJob3_ShopPayment(t *testing.T) {
	// ชำระเงินจากบัตรให้ร้านค้า 20 บาท
	value := 20
	tn := new(Trans)
	tn.Job3_ShopPayment(
		c1,    // Card
		v1,    // Device Vendor1
		pb,    // Host Paybox1
		value, // Value 20
	)
	if c1.balance != -80 ||
		v1.balance != -20 {
		t.Errorf("Expected ShopPayment 20 c1.balance = 80/%v  v1.balance = -20/%v", c1.balance, v1.balance)
	}
	fmt.Println(
		"3.ชำระเงินร้านค้า: ",
		"Value =", value,
		"c1.balance=", c1.balance,
		"v1.balance=", v1.balance,
	)
}

func Test_TransJob2_CardWithdraw(t *testing.T) {
	// คืนเงินตามจำนวนที่กำหนด แต่ไม่เกินมูลค่าคงเหลือ balance ในบัตร
	value := 50
	tn := new(Trans)
	tn.Job2_CardWithdraw(
		c1,    // Card
		pb,    // Device Paybox1
		pb,    // Host Paybox1
		value, // Value 30
	)
	if c1.balance != -30 ||
		pb.balance != 50 {
		t.Errorf("Expected c1.balance = -30/%d  pb.balance = 50/%d", c1.balance, pb.balance)
	}
	fmt.Println(
		"2.คืนเงินจากบัตร:  ",
		"Value =", value,
		"c1.balance=", c1.balance,
		"p1.balance=", pb.balance,
		"tn.change=", tn.change,
	)
}

func Test_TransJob21_CardOverWithdraw(t *testing.T) {
	// ถอนเงินเกินจำนวนคงเหลือในบัตรต้อง err != nil และแจ้งเตือน
	value := 100
	tn := new(Trans)
	err := tn.Job2_CardWithdraw(
		c1,    // Card
		pb,    // Device
		pb,    // Host
		value, // Value
	)
	if err == nil {
		t.Error("ถอนเงินเกินจำนวนคงเหลือ...แต่ไม่แจ้งเตือน err =", err, c1.balance, pb.balance)
	}
	fmt.Println(
		"21.ถอนเงินเกินจากบัตร:  ",
		"Value =", value,
		"c1.balance=", c1.balance,
		"pb.balance=", pb.balance,
		"tn.change=", tn.change,
	)
}

// Test Interface Balancer implement Method Debit(), Credit()
//func Test_Balancer(t *testing.T) {
//
//}
