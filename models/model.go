package models

type Balancer interface {
	Debit(value int)
	Credit(value int)
}
