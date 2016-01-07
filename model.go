package main

type Balancer interface {
	Debit(value int)
	Credit(value int)
}

type Shop struct {
	Site
	Name    string
	Vendor  *Device
	debit   int
	credit  int
	balance int
}

func (sh *Shop) Debit(value int) {
	sh.debit = value
	sh.balance = sh.balance - sh.debit
}

func (sh *Shop) Credit(value int) {
	sh.credit = value
	sh.balance = sh.balance + sh.credit
}

func NewShop(site Site, name string, vendor *Device) Shop {
	s := Shop{}
	s.Site = site
	s.Name = name
	s.Vendor = vendor
	s.balance = 0
	return s
}

func (s *Shop) RevenueCalc() (revenue int, err error) {
	s.Credit(s.Vendor.balance * 70 / 100)
	s.Debit(s.Vendor.balance * 30 / 100)
	s.Vendor.Debit(s.Vendor.balance)
	revenue = s.credit //สมมุติเก็บส่วนแบ่ง 30%
	return revenue, nil
}
