package main

import (
	"log"
	"os"

	"github.com/rogercoll/ipgeo"
	"github.com/rogercoll/ipgeo/db"
)

var (
	atlasAPI = os.Getenv("atlasAPI")
	token    = os.Getenv("ipToken")
)

func main() {
	dbClient, err := db.NewMongoClient(atlasAPI, "geoip", "ipscolletion", "ipsinformation")
	if err != nil {
		log.Fatal(err)
	}
	ipsToCheck, err := dbClient.Unseen()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", ipsToCheck)
	allInfo := make([]ipgeo.IPStack, len(*ipsToCheck))
	for i, ipToCheck := range *ipsToCheck {
		allInfo[i] = ipgeo.GetInfo(ipToCheck.Ip, token)
	}
	n, err := dbClient.Store(&allInfo)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d new IP information inserted to the %s database\n", n, db.GetDbType(dbClient))
}
