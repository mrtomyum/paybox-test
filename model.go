package main

type Balancer interface {
	Debit(value int)
	Credit(value int)
}

type Shop struct {
	Site
	Name    string
	Vendor  *Device
	balance int
	debit   int
	credit  int
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

//func ShopRevenueCalc(s *Site, sh *Shop, v *Device) (revenue int, err error) {
//	sh.Credit(sh.Vendor.balance * .7)
//	s.Debit(sh.Vendor.balance * .3)
//	sh.Vendor.Debit(sh.Vendor.balance)
//	revenue = sh.credit //สมมุติเก็บส่วนแบ่ง 30%
//	return revenue, nil
//}
