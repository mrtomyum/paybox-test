package main

type Balancer interface {
	Debit(value int)
	Credit(value int)
}

type Shop struct {
	*Site
	Name    string
	Vendor  *Device
	balance int
}

func NewShop(site *Site, name string, vendor *Device) *Shop {
	s := new(Shop)
	s.Site = site
	s.Name = name
	s.Vendor = vendor
	s.balance = 0
	return s
}
