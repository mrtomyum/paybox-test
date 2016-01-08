package main

import (
	"fmt"
	"log"
)

type Card struct {
	ID int
	Site
	Code    string
	Group   string
	Status  string
	debit   int
	credit  int
	balance int
}

func LoadCard(siteID int, code string) Card {
	stmt, err := db.Prepare("SELECT ID, siteID, code FROM Card WHERE siteID = ? AND code = ?")
	if err != nil {
		log.Fatal(err)
	}
	var card Card
	err = stmt.QueryRow(siteID, code).Scan(&card.ID, &card.Site.ID, &card.Code)
	if err != nil {
		log.Fatalf("บัตรมากกว่า 1 มั้ง: %s ", err)
	}
	return card
}

func LoadCards() []Card {
	rs, _ := db.Query("SELECT ID, siteID, code FROM Card")
	cards := []Card{}
	var card Card
	//	var siteID int
	for rs.Next() {
		err := rs.Scan(&card.ID, &card.Site.ID, &card.Code)
		if err != nil {
			log.Fatal(err)
		}
		cards = append(cards, card)
	}
	return cards
}

func (c *Card) New(site Site, code string, group string) {
//	c := new(Card)
	c.Site = site
	c.Code = code
	c.Group = group
	c.Status = "OPEN"
	c.balance = 0
//	return c
}

func (c *Card) Debit(value int) {
	c.debit = value
	c.balance = c.balance + c.debit
}

func (c *Card) Credit(value int) {
	c.credit = value
	c.balance = c.balance - c.credit
}

func (c *Card) Calc() {
	var j Job
	var v int
	r, err := db.Query("SELECT jobID, value FROM Trans WHERE cardID = ?", c.ID)
	if err != nil {
		log.Fatal("db.Exec Error: ", err)
	}
	//	fmt.Println("Rows Affected: ",  )
	//	fmt.Println("r.Next(): ", r.Next())
	for r.Next() {
		fmt.Printf("Begin-> %d ", c.balance)
		_ = r.Scan(&j, &v)
		switch j {
		case J1_CARD_DEPOSIT:
			c.balance = c.balance - v
		case J2_CARD_WITHDRAW:
			c.balance = c.balance + v
		case J3_SHOP_PAYMENT:
			c.balance = c.balance + v
		default:
		}
		fmt.Printf("Job: %v Value: %d Balance: %d\n", j, v, c.balance)
	}
}
