package models

import (
	"errors"
	"math"
	"time"
)

type Job int

const (
	_ Job = iota
	J1_CARD_DEPOSIT
	J2_CARD_WITHDRAW
	J3_SHOP_PAYMENT
	J4_SHOP_REFUND
)

type Trans struct {
	Job
	*Card
	vendor    *Vendor
	host      *Box
	value     int
	cash      int
	Change    int
	timeStamp time.Time
}

//func LoadTrans(Job) []Trans{
//	trans := []Trans{}
//
//	return trans
//}

func (t *Trans) Job1_CardDeposit(card *Card, host *Box, value, cash int) error {
	if value < 1 {
		return errors.New("เงินไม่เพียงพอ ขั้นต่ำ 1 บาท")
	}
	host.SetDebit(value)
	card.SetCredit(value)

	t.Job = J1_CARD_DEPOSIT
	t.Card = card
	t.vendor = nil
	t.host = host
	t.value = value
	t.cash = cash
	t.Change = cash - value
	t.timeStamp = time.Now()

	return nil
}
func (t *Trans) Job2_CardWithdraw(card *Card, host *Box, value int) error {
	// ควรเช็คก่อนว่า card.balance พอหรือไม่?
	abs := int(math.Abs(float64(card.balance))) // math.Abs use float64
	if value > abs {
		return errors.New("ไม่สามารถถอนเงิน'เกิน'มูลค่าคงเหลือในบัตร")
	}
	// Balancer
	card.SetDebit(value)
	host.SetCredit(value)

	t.Job = J2_CARD_WITHDRAW
	t.Card = card
	t.vendor = nil
	t.host = host
	t.value = card.balance
	t.cash = 0
	t.Change = t.cash - t.value
	t.timeStamp = time.Now()
	return nil
}

func (t *Trans) Job3_ShopPayment(card *Card, vendor *Vendor, value int) *Trans {
	card.SetDebit(value)
	vendor.SetCredit(value)

	t.Job = J3_SHOP_PAYMENT
	t.Card = card
	t.vendor = vendor
	t.host = nil
	t.value = value
	t.cash = 0
	t.Change = 0
	t.timeStamp = time.Now()
	return t
}
