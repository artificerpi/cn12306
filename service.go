package main

import (
	"log"
	"time"

	"github.com/artificerpi/cn12306/api"
)

const (
	TicketKey string = "ticket_key"
)

var (
	cache Cache
)

func init() {
	cache = &SimpleCache{
		data: make(map[string]Data),
	}
}

func pollTickets() {
	ticker := time.NewTicker(5 * time.Second)

	for {
		select {
		case <-ticker.C:
			q := api.TicketQuery{
				LeftTicketDTO: api.LeftTicketDTO{TrainDate: "2019-01-29", FromStation: "IZQ", ToStation: "WHN"},
				PassengerCode: "ADULT"}
			results := api.RequestTicketInfo(q)
			result := api.TicketResult{TicketList: results}
			p := Data(result)
			cache.Store(TicketKey, &p)

			log.Println("Updating")
			// case <-p.exit:
			// 	ticker.Stop()

		}
	}
}

func getRows() (rows [][]string) {
	d, ok := cache.Load(TicketKey)
	o := *d
	if ok {
		p := o.(api.TicketResult)
		rows = parseResult(p.RawData())
	}

	return
}
