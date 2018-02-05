package main

import (
	"fmt"

	"github.com/artificerpi/cn12306/query"
)

func printValues(values []string) {
	if len(values) != 36 {
		return
	}

	ticket := query.TicketInfo{
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

		ticket := query.TicketInfo{
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
