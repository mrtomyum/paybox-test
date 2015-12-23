package main

// Site
type Site struct {
	Name string
}

func NewSite(name string) *Site {
	s := new(Site)
	s.Name = name
	return s
}

type Balancer interface {
	Debit(value int)
	Credit(value int)
}

type Card struct {
	*Site
	Code    string
	Group   string
	Status  string
	debit   int
	credit  int
	balance int
}

func NewCard(site *Site, code string, group string) *Card {
	c := new(Card)
	c.Site = site
	c.Code = code
	c.Group = group
	c.Status = "OPEN"
	c.balance = 0
	return c
}

func (c *Card) Debit(value int) {
	c.debit = value
	c.balance = c.balance + c.debit
}

func (c *Card) Credit(value int) {
	c.credit = value
	c.balance = c.balance - c.credit
}

type Device struct {
	*Site
	Host     *Device
	Name     string
	Group    string
	Serial   string
	debit    int
	credit   int
	balance  int
	isHost   bool
	isOnline bool
}

func NewDevice(name, group, serial string) *Device {
	d := new(Device)
	d.Name = name
	d.Group = group
	d.Serial = serial
	d.debit = 0
	d.credit = 0
	d.balance = 0
	d.isHost = false
	d.isOnline = false
	return d
}

func (d *Device) Debit(value int) {
	d.debit = value
	d.balance = d.balance + d.debit
}

func (d *Device) Credit(value int) {
	d.credit = value
	d.balance = d.balance - d.credit
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
