package main

import (
	"errors"
	"math"
	"time"
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
func (t *Trans) Job2_CardWithdraw(card *Card, device *Device, host *Device, value int) (err error) {
	// ควรเช็คก่อนว่า card.balance พอหรือไม่?
	abs := int(math.Abs(float64(card.balance))) // math.Abs use float64
	if value > abs {
		return errors.New("ไม่สามารถถอนเงิน'เกิน'มูลค่าคงเหลือในบัตร")
	}
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

	err = nil
	return err
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
