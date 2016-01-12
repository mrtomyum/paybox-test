package main

import (
	"fmt"
	m "github.com/mrtomyum/paybox-test/models"
	"log"
	"testing"
)

func init() {
	m.InitDB("./server.db")
	//	defer db.Close()
	fmt.Println("TEST START: Initiate Sqlite3 'paybox.db'")
	sites := m.LoadSites()
	fmt.Println("Load Site slice:=> ", sites)
	cards := m.LoadCards()
	fmt.Println("Load Card slice:=>", cards)
	fmt.Println()
}

func setup() (s m.Site, card *m.Card, box *m.Box, vendor *m.Vendor, trans *m.Trans) {
	// Site
	s = m.Site{
		Name: "บริษัท นพดลพานิช จำกัด",
	}
	s.SetBalance(0)

	// Card
	card = new(m.Card)
	card.Site = s
	card.Code = "123456"
	card.Group = "PURSE"
	card.Status = "OPEN"
	card.SetBalance(0)
	//	card := Card{
	//		Site: site,
	//		Code: "123456",
	//		Group: "PURSE",
	//		Status: "OPEN",
	//	}
	//	cards = LoadCards()

	//	paybox = NewDevice(BOX, "Paybox1", "P001")
	box = &m.Box{
		Device: m.Device{
			Name:   "Paybox1",
			Serial: "P001",
		},
		Cash:       0,
		IsLockOpen: false,
	}
	vendor = &m.Vendor{
		Device: m.Device{
			Name:   "V1",
			Serial: "V001",
		},
	}
	//	paybox = LoadDevice(BOX)
	//	vendor = LoadDevice(VENDOR)
	trans = new(m.Trans)
	//	v2  = NewDevice("V2", "VENDOR", "V002")
	//	v3  = NewDevice("V3", "VENDOR", "V003")
	//	sh1 = NewShop(s, "ร้านข้าวมันไก่", v1)
	return s, card, box, vendor, trans
}

//func setupTable() (sites []Site, cards []Card, payboxs, vendors []Device, tn *Trans) {
//	sites = LoadSites()
//	cards = LoadCards()
//	//	paybox = NewDevice("Paybox1", "BOX", "P001")
//	payboxs = LoadDevice(BOX)
//	//	vendor = NewDevice("V1", "VENDOR", "V001")
//	vendors = LoadDevice(VENDOR)
//	tn = new(Trans)
//	return sites, cards, payboxs, vendors, tn
//}

// เทสว่าการ์ดใหม่ต้องไม่มี code ซ้ำใน site เดียวกัน
//func Test_NewObject(t *testing.T) {
//	s, c, _, _, _ := setup()
//	// Test Card
//	if s.Name != "บริษัท ทดสอบ จำกัด" {
//		t.Error("Expected name = 'บริษัท ทดสอบ จำกัด'")
//	}
//	//	Test Card
//	if c.Code != "123456" {
//		t.Error("Expected code = '123456'")
//	}
//	fmt.Println("Cards=> ", c)
//}

// เทสคำนวณยอดคงเหลือบัตร
func TestCardBalance(t *testing.T) {
	_, c, _, _, _ := setup()
	//	c := cards[1]
	if c.SetCredit(100); c.Balance() != -100 {
		t.Errorf("Expected credit -100 but %v", c.Balance())
	}
	if c.SetDebit(30); c.Balance() != -70 {
		t.Errorf("Expected card balance = -70 but %v", c.Balance())
	}
}

// เทสคำนวณยอดคงเหลือใน Device ทั้ง Paybox และ Vendor
func TestDeviceBalance(t *testing.T) {
	_, _, b, v, _ := setup()
	if b.SetDebit(100); b.Balance() != 100 {
		t.Error("Expected debit = 100 box balance = 100 ")
	}
	if b.SetCredit(30); b.Balance() != 70 {
		t.Error("Expected credit 30 box balance = 70")
	}
	// ลองเอาบัตรซื้อของอาหารราคา 40 บาท
	if v.SetDebit(40); v.Balance() != 40 {
		t.Error("Expected debit  40 vendor balance = 40")
	}
	if v.SetCredit(5); v.Balance() != 35 {
		t.Error("Expected credit 5 vendor balance = 35")
	}
}

// เติมเงินเข้าบัตรใหม่
func Test_TransJob1_CardDeposit(t *testing.T) {
	_, c, b, _, tn := setup()
	//	_, cards, p, v, tn := setupTable()
	//	c = cards[1]
	value := 100
	cash := 100
	tn.Job1_CardDeposit( // เติมเงิน 50 ใส่เงิน 100 ทอน 50
		c,     // Card
		b,     // Host
		value, // Value
		cash,  // CashReceive
	)
	if c.Balance() != -100 ||
		b.Balance() != 100 {
		t.Errorf("Expected เติมเงิน 100 ต้องมีเงินในบัตรเพิ่ม c1.credit != 50 got %v|| c1.balance != 50 got %v||  p1.debit != 50 got %v|| p1.balance != 50 got %v", c.Credit(), c.Balance(), b.Debit(), b.Balance())
	}
	fmt.Println(
		"1.เติมเงินเข้าบัตร: ",
		"Value =", value,
		"c1.balance =", c.Balance(),
		"p1.balance =", b.Balance(),
		"tn.change =", tn.Change,
	)
}

// เทสคำนวณยอดคงเหลือในบัตร จากการเติมเงิน โดยใช้ข้อมูลจากตาราง Trans
func Test_Job12_CardDepositCalc(t *testing.T) {
	//	_, cards, boxs, vendors, tn := setupTable()
	//	transJ1 := LoadTrans(J1_CARD_DEPOSIT)
	card := m.LoadCard(1, "12345")
	fmt.Println("Load Card=> ", card)
	card.SetBalance(0)
	fmt.Println("Card.Balance=> ", card.Balance())
	card.Calc()
	fmt.Println("Card.Balance after Calc=> ", card.Balance())

	if card.Balance() != -490 {
		log.Fatalf("test 1.2 คำนวณไม่ตรง %d", card.Balance())
	}
}

func Test_TransJob11_CardDepositMustGreaterThan1(t *testing.T) {
	_, c, b, _, tn := setup()
	value := 0
	cash := 0
	err := tn.Job1_CardDeposit(
		c,     // Card
		b,     // Device
		value, // Value
		cash,  // CashReceive
	)
	if err == nil {
		t.Errorf("ยอดเงินเติมน้อยกว่าขั้นต่ำ %d แต่ไม่แจ้งเตือน err", value)
	}
	fmt.Println("1.1 Pass:เทสเติมเงินน้อยกว่า 1 บาท ต้องแสดง Error ดังนี้=>", err)
}

// เทสคืนเงินตามจำนวนที่กำหนด
func Test_TransJob2_CardWithdraw(t *testing.T) {
	value := 50
	_, c, b, _, tn := setup()
	c.SetBalance(-80)
	//	if err != nil {
	//		fmt.Printf("Check Card balance %v \n", c.Balance())
	//	}
	b.SetBalance(100)
	//	v.SetBalance(-20)
	//	if err != nil {
	//		fmt.Printf("Check Box balance %v\n", b.Balance())
	//	}
	err := tn.Job2_CardWithdraw(
		c,     // Card
		b,     // Box
		value, // Value 50
	)
	if c.Balance() != -30 ||
		b.Balance() != 50 ||
		err != nil {
		t.Errorf("Expected c.balance = -30/%d  p.balance = 50/%d", c.Balance(), b.Balance())
	}
	fmt.Println(
		"2.คืนเงินจากบัตร:  ",
		"Value =", value,
		"c1.balance=", c.Balance(),
		"p1.balance=", b.Balance(),
		"tn.change=", tn.Change,
	)
}

// เทสถอนเงินเกินจำนวนคงเหลือในบัตร ต้องแจ้งเตือน
func Test_TransJob21_CardOverWithdraw(t *testing.T) {
	_, c, b, _, tn := setup()
	value := 100
	err := tn.Job2_CardWithdraw(
		c,     // Card
		b,     // Box
		value, // Value
	)
	if err == nil {
		t.Error("ถอนเงินเกินจำนวนคงเหลือ...แต่ไม่แจ้งเตือน err =", err, c.Balance(), b.Balance())
	}
	fmt.Println(
		"21.ถอนเงินเกินจากบัตร:  ",
		"Value =", value,
		"c1.balance=", c.Balance(),
		"pb.balance=", b.Balance(),
		"tn.change=", tn.Change,
	)
}

// ชำระเงินจากบัตรให้ร้านค้า 20 บาท
func Test_TransJob3_ShopPayment(t *testing.T) {
	value := 20
	_, c, p, v, tn := setup()
	c.SetBalance(-100)
	p.SetBalance(100)
	v.SetBalance(0)
	tn.Job3_ShopPayment(
		c,     // Card
		v,     // Vendor
		value, // Value
	)
	if c.Balance() != -80 ||
		v.Balance() != -20 ||
		p.Balance() != 100 {
		t.Errorf("Expected ShopPayment 20 c1.balance = 80/%v  v1.balance = -20/%v", c.Balance(), v.Balance())
	}
	fmt.Println(
		"3.ชำระเงินร้านค้า: ",
		"Value =", value,
		"c1.balance=", c.Balance(),
		"v1.balance=", v.Balance(),
	)
}

// ยอดเงินสามารถคิดแบ่งสัดส่วน ตามอัตราที่กำหนด รองรับการแบ่งจ่าย เศษสตางค์ 2 หลัก

// เทสคำนวณเพื่อจ่ายเงินให้ร้านค้าหลังหักส่วนแบ่งสถานที่แล้ว 30% ของยอดคงเหลือในแต่ละ Device
func Test_ShopBalancer(t *testing.T) {
	s, _, _, v, _ := setup()
	sh := m.NewShop(s, "ร้านข้าวมันไก่โต้ง", v)
	v.SetBalance(5000)
	r, err := sh.RevenueCalc()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("3.1 Site: %v Balance: %v\nShop revenue: %v", s.Name, s.Balance(), r)
}

// Test Interface Balancer implement Method Debit(), Credit()
//func Test_Balancer(t *testing.T) {
//
//}
