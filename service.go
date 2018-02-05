package main

import (
	"log"
	"time"

	"github.com/artificerpi/cn12306/query"
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
	ticker := time.NewTicker(10 * time.Second)

	for {
		select {
		case <-ticker.C:
			log.Println("ticked")
			results := query.Q(queryURL)
			result := query.TicketResult{
				TicketList: results,
			}
			p := Data(result)
			cache.Store(TicketKey, &p)
			log.Println("Updated cache")
			// case <-p.exit:
			// 	ticker.Stop()

		}
	}
}

func getRows() (rows [][]string) {
	d, ok := cache.Load(TicketKey)
	o := *d
	if ok {
		p := o.(query.TicketResult)
		rows = parseResult(p.RawData())
	}

	return
}
