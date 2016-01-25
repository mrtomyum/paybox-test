package models

type Shop struct {
	Site
	*Vendor
	Name    string
	debit   int
	credit  int
	balance int
}

func (s *Shop) SetDebit(value int) {
	s.debit = +value
	s.balance = +value
}

func (s *Shop) SetCredit(value int) {
	s.credit = -value
	s.balance = -value
}

func NewShop(site Site, name string, vendor *Vendor) Shop {
	s := Shop{}
	s.Site = site
	s.Name = name
	s.Vendor = vendor
	s.balance = 0
	return s
}

func (s *Shop) RevenueCalc() (revenue int, err error) {
	s.SetCredit(s.Vendor.Balance() * 70 / 100)
	s.Site.SetDebit(s.Vendor.Balance() * 30 / 100)
	s.Vendor.SetDebit(s.Vendor.Balance())
	revenue = s.credit //สมมุติเก็บส่วนแบ่ง 30%
	return revenue, nil
}
