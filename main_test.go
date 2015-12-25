package main

import (
	"fmt"
	"log"
	"testing"
)

func init() {
	InitDB("./server.db")
	//	defer db.Close()
	fmt.Println("TEST START: Initiate Sqlite3 'paybox.db'")
	sites := LoadSites()
	fmt.Println("Load Site slice:=> ", sites)
	cards := LoadCards()
	fmt.Println("Load Card slice:=>", cards)
	fmt.Println()
}

func setup() (sites []Site, cards *Card, paybox, vendor *Device, trans *Trans) {
	//	site := sites[0]
	sites = LoadSites()
	card := NewCard(sites[0], "123456", "PURSE")
	//	cards = LoadCards()
	paybox = NewDevice(BOX, "Paybox1", "P001")
	vendor = NewDevice(VENDOR, "V1", "V001")
	//	paybox = LoadDevice(BOX)
	//	vendor = LoadDevice(VENDOR)
	trans = new(Trans)
	//	v2  = NewDevice("V2", "VENDOR", "V002")
	//	v3  = NewDevice("V3", "VENDOR", "V003")
	//	sh1 = NewShop(s, "ร้านข้าวมันไก่", v1)
	return sites, card, paybox, vendor, trans
}

func setupTable() (sites []Site, cards []Card, payboxs, vendors []Device, tn *Trans) {
	sites = LoadSites()
	cards = LoadCards()
	//	paybox = NewDevice("Paybox1", "BOX", "P001")
	payboxs = LoadDevice(BOX)
	//	vendor = NewDevice("V1", "VENDOR", "V001")
	vendors = LoadDevice(VENDOR)
	tn = new(Trans)
	return sites, cards, payboxs, vendors, tn
}

// เทสว่าการ์ดใหม่ต้องไม่มี code ซ้ำใน site เดียวกัน
func Test_NewObject(t *testing.T) {
	sites, c, _, _, _ := setup()
	s := sites[0]
	if s.Name != "บริษัท ทดสอบ จำกัด" {
		t.Error("Expected name = 'บริษัท ทดสอบ จำกัด'")
	}
	//	c := cards[0]
	if c.Code != "123456" {
		t.Error("Expected code = '123456'")
	}
	fmt.Println("Cards=> ", c)
}

// เทสคำนวณยอดคงเหลือบัตร
func TestCardBalance(t *testing.T) {
	_, c, _, _, _ := setup()
	//	c := cards[1]
	if c.Credit(100); c.balance != -100 {
		t.Error("Expected credit -100")
	}
	if c.Debit(30); c.balance != -70 {
		t.Error("Expected card balance = -70 ")
	}
}

// เทสคำนวณยอดคงเหลือใน Device ทั้ง Paybox และ Vendor
func TestDeviceBalance(t *testing.T) {
	_, _, p, v, _ := setup()
	if p.Debit(100); p.balance != 100 {
		t.Error("Expected debit = 100 box balance = 100 ")
	}
	if p.Credit(30); p.balance != 70 {
		t.Error("Expected credit 30 box balance = 70")
	}
	// ลองเอาบัตรซื้อของอาหารราคา 40 บาท
	if v.Debit(40); v.balance != 40 {
		t.Error("Expected debit  40 vendor balance = 40")
	}
	if v.Credit(5); v.balance != 35 {
		t.Error("Expected credit 5 vendor balance = 35")
	}
}

// เติมเงินเข้าบัตรใหม่
func Test_TransJob1_CardDeposit(t *testing.T) {
	_, c, p, _, tn := setup()
	//	_, cards, p, v, tn := setupTable()
	//	c = cards[1]
	value := 100
	cash := 100
	tn.Job1_CardDeposit( // เติมเงิน 50 ใส่เงิน 100 ทอน 50
		c,     // Card
		p,     // Device
		p,     // Host
		value, // Value
		cash,  // CashReceive
	)
	if c.balance != -100 ||
		p.balance != 100 {
		t.Errorf("Expected เติมเงิน 100 ต้องมีเงินในบัตรเพิ่ม c1.credit != 50 got %v|| c1.balance != 50 got %v||  p1.debit != 50 got %v|| p1.balance != 50 got %v", c.credit, c.balance, p.debit, p.balance)
	}
	fmt.Println(
		"1.เติมเงินเข้าบัตร: ",
		"Value =", value,
		"c1.balance =", c.balance,
		"p1.balance =", p.balance,
		"tn.change =", tn.change,
	)
}

// เทสคำนวณยอดคงเหลือในบัตร จากการเติมเงิน โดยใช้ข้อมูลจากตาราง Trans
func Test_Job12_CardDepositCalc(t *testing.T) {
	//	_, cards, boxs, vendors, tn := setupTable()
	//	transJ1 := LoadTrans(J1_CARD_DEPOSIT)
	card := LoadCard(1, "12345")
	fmt.Println("Load Card=> ", card)
	card.balance = 0
	fmt.Println("Card.Balance=> ", card.balance)
	card.Calc()
	fmt.Println("Card.Balance after Calc=> ", card.balance)

	if card.balance != -490 {
		log.Fatalf("test 1.2 คำนวณไม่ตรง %d", card.balance)
	}
}

func Test_TransJob11_CardDepositMustGreaterThan1(t *testing.T) {
	_, c, p, _, tn := setup()
	value := 0
	cash := 0
	err := tn.Job1_CardDeposit(
		c,     // Card
		p,     // Device
		p,     // Host
		value, // Value
		cash,  // CashReceive
	)
	if err == nil {
		t.Errorf("ยอดเงินเติมน้อยกว่าขั้นต่ำ %d แต่ไม่แจ้งเตือน err", value)
	}
	fmt.Println("1.1 เทสเติมเงินน้อยกว่า 1 บาท ต้องแสดง Error=>", err)
}

// ชำระเงินจากบัตรให้ร้านค้า 20 บาท
func Test_TransJob3_ShopPayment(t *testing.T) {
	value := 20
	_, c, p, v, tn := setup()
	c.balance = -100
	p.balance = 100
	v.balance = 0
	tn.Job3_ShopPayment(
		c,     // Card
		v,     // Device
		p,     // Host
		value, // Value
	)
	if c.balance != -80 ||
		v.balance != -20 ||
		p.balance != 100 {
		t.Errorf("Expected ShopPayment 20 c1.balance = 80/%v  v1.balance = -20/%v", c.balance, v.balance)
	}
	fmt.Println(
		"3.ชำระเงินร้านค้า: ",
		"Value =", value,
		"c1.balance=", c.balance,
		"v1.balance=", v.balance,
	)
}

// เทสคืนเงินตามจำนวนที่กำหนด
func Test_TransJob2_CardWithdraw(t *testing.T) {
	value := 50
	_, c, p, v, tn := setup()
	c.balance = -80
	p.balance = 100
	v.balance = -20
	err := tn.Job2_CardWithdraw(
		c,     // Card
		p,     // Device Paybox1
		p,     // Host Paybox1
		value, // Value 30
	)
	if c.balance != -30 ||
		p.balance != 50 ||
		err != nil {
		t.Errorf("Expected c.balance = -30/%d  p.balance = 50/%d", c.balance, p.balance)
	}
	fmt.Println(
		"2.คืนเงินจากบัตร:  ",
		"Value =", value,
		"c1.balance=", c.balance,
		"p1.balance=", p.balance,
		"tn.change=", tn.change,
	)
}

// เทสถอนเงินเกินจำนวนคงเหลือในบัตร ต้องแจ้งเตือน
func Test_TransJob21_CardOverWithdraw(t *testing.T) {
	_, c, p, _, tn := setup()
	value := 100
	err := tn.Job2_CardWithdraw(
		c,     // Card
		p,     // Device
		p,     // Host
		value, // Value
	)
	if err == nil {
		t.Error("ถอนเงินเกินจำนวนคงเหลือ...แต่ไม่แจ้งเตือน err =", err, c.balance, p.balance)
	}
	fmt.Println(
		"21.ถอนเงินเกินจากบัตร:  ",
		"Value =", value,
		"c1.balance=", c.balance,
		"pb.balance=", p.balance,
		"tn.change=", tn.change,
	)
}

// เทสคำนวณเพื่อจ่ายเงินให้ร้านค้าหลังหักส่วนแบ่งสถานที่แล้ว 30%
//func Test_ShopBalancer(t *testing.T)  {
//	s, c, p, v, tn := setup()
//	shop := NewShop(s, "ร้านข้าวมันไก่โต้ง", v)
//
//}
// Test Interface Balancer implement Method Debit(), Credit()
//func Test_Balancer(t *testing.T) {
//
//}
