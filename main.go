package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	ui "github.com/gizak/termui"
)

func q(rawurl string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", rawurl, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Set("leftTicketDTO.train_date", "2018-02-11")
	q.Set("leftTicketDTO.from_station", "IZQ")
	q.Set("leftTicketDTO.to_station", "WHN")
	q.Set("purpose_codes", "ADULT")
	req.URL.RawQuery = unsortedEncode(q)

	log.Println(req.URL.String())

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("Data", string(data))

	var t TicketResponse
	err = json.Unmarshal(data, &t)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range t.Data.LeftTickets {
		v = strings.TrimPrefix(v, "|")
		values := strings.Split(v, "|")
		printValues(values)
	}
}

func printValues(values []string) {
	if len(values) != 36 {
		return
	}

	ticket := LeftTicket{
		StationTrainCode: values[2],
		FromStation:      values[5],
		ToStation:        values[6],
		StartTime:        values[7],
		ArriveTime:       values[8],
		BusinessSeat:     values[31],
		FirstClassSeat:   values[30],
		SecondClassSeat:  values[29],
		AdvancedSleeper:  values[20],
		SoftSleeper:      values[22],
		HardSleeper:      values[27],
		HardSeat:         values[28],
		StandUp:          values[25],
		Other:            values[32],
		Remark:           values[0],
	}

	fmt.Println(ticket.StationTrainCode, ticket.FromStation, ticket.ToStation, ticket.StartTime, ticket.ArriveTime, ticket.BusinessSeat, ticket.FirstClassSeat, ticket.SecondClassSeat,
		ticket.AdvancedSleeper, ticket.SoftSleeper, ticket.HardSleeper, ticket.HardSeat, ticket.StandUp, ticket.Other, ticket.Remark)
}

func startUI() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	rows1 := [][]string{
		[]string{"header1", "header2", "header3"},
		[]string{"你好块", "Go-lang is so cool", "Im working on Ruby"},
		[]string{"2016", "10", "11"},
	}
	table := ui.NewTable()
	table.Rows = rows1
	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(6, 3, table),
		),
	)

	// calculate layout
	ui.Body.Align()

	ui.Render(ui.Body)

	// handle key q pressing
	ui.Handle("/sys/kbd/q", func(ui.Event) {
		// press q to quit
		ui.StopLoop()
	})

	ui.Handle("/sys/kbd/C-x", func(ui.Event) {
		// handle Ctrl + x combination
	})

	ui.Handle("/sys/kbd", func(ui.Event) {
		// handle all other key pressing
	})

	// handle a 1s timer
	ui.Handle("/timer/1s", func(e ui.Event) {
		t := e.Data.(ui.EvtTimer)
		// t is a EvtTimer
		if t.Count%2 == 0 {
			// do something
			q("https://kyfw.12306.cn/otn/leftTicket/queryZ")
			ui.Render(ui.Body)
		}
	})

	ui.Loop() // block until StopLoop is called
}

func main() {
	// server health check
	// q("https://kyfw.12306.cn/otn/leftTicket/log")

	// query every 5 seconds
	// ticker := time.NewTicker(5 * time.Second)
	// for {
	// 	select {
	// 	case <-ticker.C:
	// 		q("https://kyfw.12306.cn/otn/leftTicket/queryZ")
	// 	}
	// }
	startUI()
}
