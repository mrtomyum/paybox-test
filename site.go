package main

import (
	"fmt"
	"log"
)

type Site struct {
	ID   int
	Name string
}

func NewSite(name string) *Site {
	s := new(Site)
	s.Name = name
	return s
}

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
