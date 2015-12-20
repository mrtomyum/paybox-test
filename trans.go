package main

import (
	"time"
	//	"log"
	//	"math"
)

type Job int

const (
	J1_CARD_DEPOSIT Job = iota
	J2_CARD_WITHDRAW
	J3_SHOP_PAYMENT
	J4_SHOP_REFUND
)

type Trans struct {
	Job
	*Card
	device    *Device
	host      *Device
	value     int
	cash      int
	change    int
	timeStamp time.Time
}

func (t *Trans) Job1_CardDeposit(card *Card, device *Device, host *Device, value, cash int) *Trans {
	device.Debit(value)
	card.Credit(value)

	t.Job = J1_CARD_DEPOSIT
	t.Card = card
	t.device = device
	t.host = host
	t.value = value
	t.cash = cash
	t.change = cash - value
	t.timeStamp = time.Now()
	return t
}
func (t *Trans) Job2_CardWithdraw(card *Card, device *Device, host *Device, value int) *Trans {
	//ยังติดปัญหา ABS ค่า value ไม่ได้
	//	if value > card.balance {
	//		log.Fatalf("มูลค่าคงเหลือในบัตร %d ไม่พอจ่าย %d", card.balance, value)
	//	}
	// Balancer
	card.Debit(value)
	device.Credit(value)

	t.Job = J2_CARD_WITHDRAW
	t.Card = card
	t.device = device
	t.host = host
	t.value = card.balance
	t.cash = 0
	t.change = t.cash - t.value
	t.timeStamp = time.Now()
	return t
}

func (t *Trans) Job3_ShopPayment(card *Card, device *Device, host *Device, value int) *Trans {
	card.Debit(value)
	device.Credit(value)

	t.Job = J3_SHOP_PAYMENT
	t.Card = card
	t.device = device
	t.host = host
	t.value = value
	t.cash = 0
	t.change = 0
	t.timeStamp = time.Now()
	return t
}
