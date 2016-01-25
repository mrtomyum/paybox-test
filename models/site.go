package models

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
	rs, err := db.Query("SELECT ID, Name FROM Site ")
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

func (s *Site) SetDebit(value int) {
	s.debit = +value
	s.balance = +value
}

func (s *Site) SetBalance(value int) {
	s.balance = value
}

func (s *Site) Balance() int {
	return s.balance
}
