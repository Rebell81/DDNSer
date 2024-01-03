package main

import (
	"DDNSer/flare"
	"DDNSer/keen"
	"DDNSer/telegram"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	telegramBot    *telegram.Telegram
	keeneticClient *keen.Keenetic
	err            error
)

func main() {
	ctx := context.Background()

	cloudFlareToken := os.Getenv("CLOUDFLARE_TOKEN")
	cloudFlareZoneId := os.Getenv("CLOUDFLARE_ZON_ID")
	keeneticHost := os.Getenv("KEENETIC_HOST")
	keeneticLogin := os.Getenv("LOGIN")
	keeneticPassword := os.Getenv("PASSWORD")
	telegramBotToken := os.Getenv("TELEGRAM_TOKEN")
	telegramChatId, _ := strconv.ParseInt(os.Getenv("TELEGRAM_CHAT_ID"), 10, 64)

	if telegramBotToken != "" {
		telegramBot, err = telegram.InitBot(telegramBotToken)
	}

	keeneticClient, err = keen.New(keeneticHost, keeneticLogin, keeneticPassword)
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

	fmt.Println("start cycle....")
	for {
		interfaces, err := keeneticClient.GetSingle("PPPoE1")
		if err != nil {
			logErr(err.Error(), telegramChatId)
		}

		if len(list) == 1 && len(interfaces) == 1 {
			cloudFlareRecord := list[0]
			byflyInterface := &interfaces[0]

			//log.Println(fmt.Sprintf("Cloud ip: '%s'		| Keen ip: '%s'", cloudFlareRecord.Content, byflyInterface.Address))

			if cloudFlareRecord.Content != byflyInterface.Address && byflyInterface.Address != "" {
				log.Println(fmt.Sprintf("'%s'			==============>			'%s'", cloudFlareRecord.Content, byflyInterface.Address))
				err = flareClient.UpdateDnsRecord(ctx, cloudFlareZoneId, cloudFlareRecord.ID, byflyInterface.Address)
				if err != nil {
					logErr(err.Error(), telegramChatId)
				}

				list, err = flareClient.GetRecords(ctx, cloudFlareZoneId)
				if err != nil {
					logErr(err.Error(), telegramChatId)
				}

			}
		}

		time.Sleep(time.Second)
	}

}

func logErr(msg string, chatId int64) {
	if telegramBot != nil {
		errs := telegramBot.SendMsg(msg, chatId)
		if errs != nil {
			fmt.Println(errs)
		}
	}

	fmt.Println(msg)
}
