package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/prantadas/qbit/api"
	tele "gopkg.in/telebot.v3"
)

var q api.QbitClient

func main() {
	q = api.New("admin", "adminadmin", "http://qbit:8090")
	_, err := q.Login(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	token := os.Getenv("QBIT_TOKEN")
	log.Println("Got Token:", token)

	pref := tele.Settings{
		Token:  token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", help)
	b.Handle("/help", help)

	b.Handle("/magnet", magnet)
	log.Println("Bot started")

	b.Start()
}

func help(c tele.Context) error {
	return c.Send(`
		/magnet - Add magnet link to download queue
	/help - Show this message
		`)
}

func magnet(c tele.Context) error {
	fmt.Println(c.Args())
	if len(c.Args()) == 0 {
		return c.Send("No magnet link provided")
	} else if len(c.Args()) > 1 {
		return c.Send("Too many magnet links provided")
	}
	res, err := q.Add(context.Background(), c.Args()[0])
	if err != nil {
		return c.Send("Cannot add magnet link provided")
	}
	defer res.Body.Close()
	return c.Send("Torrent added to download queue!")
}
