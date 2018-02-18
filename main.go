package main

import (
	"flag"
	"log"
	"strings"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type ARGs struct {
	Token            *string
	LogFile          *string
	AllowedUserIDs   *string
	AllowedUserNames *string
	Debug            *bool
	Lang             *string
	DialogFlowToken  *string
}

func main() {
	args := ARGs{
		Token:            flag.String("tt", "", "telegram token"),
		LogFile:          flag.String("l", "", "log file (optional)"),
		AllowedUserIDs:   flag.String("uids", "", "allowed telegram user id`s (comma separated numbers)"),
		AllowedUserNames: flag.String("unames", "", "allowed telegram usernames (comma separated strings without (at) symbol (@))"),
		Debug:            flag.Bool("d", false, "debug enabled"),
		Lang:             flag.String("lang", "", "dialog flow lang"),
		DialogFlowToken:  flag.String("dft", "", "dialog flow token"),
	}
	flag.Parse()
	//validate input args
	if err := args.validate(); err != nil {
		log.Fatalln(err)
	}
	//set log output if l param is set
	setLogFiles(*args.LogFile)
	//starting bot
	log.Println("Starting telegram bot...")
	b, err := tb.NewBot(tb.Settings{
		Token:  *args.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Long polling started")

	allowedUserIDs := getUsers(*args.AllowedUserIDs)
	allowedUserNames := strings.Split(*args.AllowedUserNames, ",")

	b.Handle(tb.OnText, func(m *tb.Message) {
		if *args.Debug {
			log.Printf("%+v\n", m)
			log.Printf("%+v\n", m.Sender)
			log.Printf("%+v\n", m.Chat)
		}

		if isAllowedUsername(m.Sender.Username, allowedUserNames) || isAllowedID(m.Sender.ID, allowedUserIDs) {
			if *args.Debug {
				log.Println("user is allowed, sending request to dialogflow")
			}

			rsp, err := sendDialogFlow(m.Sender.ID, m.Text, *args.DialogFlowToken, *args.Lang)
			if err != nil {
				log.Println(err)
			}
			if *args.Debug {
				log.Printf("%+v\n", rsp)
			}
			b.Send(m.Sender, rsp.Result.Speech)
		}
	})

	b.Start()
}
