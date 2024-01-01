package main

import (
	"DDNSer/flare"
	"DDNSer/keen"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	cloudFlareToken := os.Getenv("CLOUDFLARE_TOKEN")
	cloudFlareZoneId := os.Getenv("CLOUDFLARE_ZON_ID")
	keeneticHost := os.Getenv("KEENETIC_HOST")
	keeneticLogin := os.Getenv("LOGIN")
	keeneticPassword := os.Getenv("PASSWORD")

	keenetiClient, err := keen.New(keeneticHost, keeneticLogin, keeneticPassword)
	if err != nil {
		log.Fatal(err)
	}

	flareClient, err := flare.InitApi(ctx, cloudFlareToken)
	if err != nil {
		log.Fatal(err)
	}

	list, err := flareClient.GetRecords(ctx, cloudFlareZoneId)
	if err != nil {
		fmt.Println(err)
	}

	for {
		interfaces, err := keenetiClient.GetSingle("PPPoE1")
		if err != nil {
			fmt.Println(err)
		}

		if len(list) == 1 && len(interfaces) == 1 {
			cloudFlareRecord := list[0]
			byflyInterface := &interfaces[0]

			log.Println(fmt.Sprintf("Cloud ip: '%s'		| Keen ip: '%s'", cloudFlareRecord.Content, byflyInterface.Address))

			if cloudFlareRecord.Content != byflyInterface.Address && byflyInterface.Address != "" {
				log.Println(fmt.Sprintf("'%s'			==============>			'%s'", cloudFlareRecord.Content, byflyInterface.Address))
				err = flareClient.UpdateDnsRecord(ctx, cloudFlareZoneId, cloudFlareRecord.ID, byflyInterface.Address)
				if err != nil {
					fmt.Println(err)
				}

				list, err = flareClient.GetRecords(ctx, cloudFlareZoneId)
				if err != nil {
					fmt.Println(err)
				}
			}
		}

		time.Sleep(time.Second)
	}

}
