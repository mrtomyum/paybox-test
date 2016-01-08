package main

type Balancer interface {
	Debit(value int)
	Credit(value int)
}