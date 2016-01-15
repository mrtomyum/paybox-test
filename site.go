package main

import (
	"fmt"
	"log"
)

type Site struct {
	ID      int
	Name    string
	debit   int
	credit  int
	balance int
}

//func NewSite(name string) *Site {
//	s := new(Site)
//	s.Name = name
//	return s
//}

func LoadSites() []Site {
	rs, err := db.Query("SELECT ID, Name FROM site ")
	if err != nil {
		log.Fatal(err)
	}
	defer rs.Close()
	sites := []Site{}
	var site Site
	for rs.Next() {
		err = rs.Scan(&site.ID, &site.Name)
		if err != nil {
			log.Fatal(err)
		}
		sites = append(sites, site)
	}
	fmt.Println("Sites = ", sites)
	return sites
}

func (s *Site) Debit(value int) {
	s.debit = value
	s.balance = +s.debit
}
