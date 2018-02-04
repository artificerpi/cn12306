package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"strings"
)

func unsortedEncode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := []string{"leftTicketDTO.train_date",
		"leftTicketDTO.from_station",
		"leftTicketDTO.to_station",
		"purpose_codes"}

	for _, k := range keys {
		vs := v[k]
		prefix := url.QueryEscape(k) + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(url.QueryEscape(v))
		}
	}
	return buf.String()
}

func readData(r io.Reader) [][]string {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(err)
	}
	// log.Println("Data", string(data))

	var t TicketResponse
	err = json.Unmarshal(data, &t)
	if err != nil {
		log.Println(err)
	}

	results := make([][]string, 30)
	for _, v := range t.Data.LeftTickets {
		v = strings.TrimPrefix(v, "|")
		values := strings.Split(v, "|")
		// filter
		if strings.HasPrefix(values[2], "G") {
			results = append(results, values)
		}
	}

	return results
}

func printValues(values []string) {
	if len(values) != 36 {
		return
	}

	ticket := TicketInfo{
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

func parseResult(results [][]string) [][]string {
	rows := make([][]string, 0)
	header := []string{"车次", "起始", "--->", "终止", "发车时间", "--", "到达时间", "历时", "座位信息", "", "", ""}
	rows = append(rows, header)

	for _, v := range results {

		if len(v) < 32 {
			continue
		}
		// length 36
		// fmt.Println(v)

		ticket := TicketInfo{
			StationTrainCode: v[2],
			FromStation:      v[5],
			ToStation:        v[6],
			StartTime:        v[7],
			ArriveTime:       v[8],
			BusinessSeat:     v[31],
			FirstClassSeat:   v[30],
			SecondClassSeat:  v[29],
			AdvancedSleeper:  v[20],
			SoftSleeper:      v[22],
			HardSleeper:      v[27],
			HardSeat:         v[28],
			StandUp:          v[25],
			Other:            v[32],
			Remark:           v[0],
		}

		data := []string{ticket.StationTrainCode, ticket.FromStation, "", ticket.ToStation, ticket.StartTime, "", ticket.ArriveTime, "", ticket.BusinessSeat, ticket.FirstClassSeat, ticket.SecondClassSeat,
			ticket.Other}
		rows = append(rows, data)
		// fmt.Println(data)

	}

	// fmt.Println(rows)
	return rows
}
