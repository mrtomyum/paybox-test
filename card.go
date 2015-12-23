package main

import "log"

type Card struct {
	ID int
	*Site
	Code    string
	Group   string
	Status  string
	debit   int
	credit  int
	balance int
}

func LoadCards() []Card {
	rs, _ := db.Query("SELECT ID, code, siteID FROM Card")
	cards := []Card{}
	var card Card
	var siteID int
	for rs.Next() {
		err := rs.Scan(&card.ID, &card.Code, &siteID)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
	}
	return cards
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
