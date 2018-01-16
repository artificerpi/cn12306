package main

import(
	"net/http"
	"log"
	"io/ioutil"
	"crypto/tls"
	"encoding/json"
	"strings"
	"fmt"
)


func q(rawurl string){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err:= http.NewRequest("GET", rawurl, nil)
	if err!=nil{
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Set("leftTicketDTO.train_date","2018-02-11")
	q.Set("leftTicketDTO.from_station","IZQ")
	q.Set("leftTicketDTO.to_station","WHN")
	q.Set("purpose_codes","ADULT")
	req.URL.RawQuery = unsortedEncode(q)

	log.Println(req.URL.String())

	resp, err:= client.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		log.Fatal(err)
	}
	// log.Println("Data", string(data))

	var t TicketResponse
	err = json.Unmarshal(data, &t)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range t.Data.LeftTickets {
		v = strings.TrimPrefix(v,"|")
		values := strings.Split(v,"|")
		printValues(values)
	}
}

func printValues(values []string){
	if len(values) != 36 {
		return
	}
	
	ticket := LeftTicket{
		StationTrainCode: values[2],
		FromStation: values[5],
		ToStation: values[6],
		StartTime: values[7],
		ArriveTime: values[8],
		BusinessSeat: values[31],
		FirstClassSeat: values[30],
		SecondClassSeat: values[29],
		AdvancedSleeper: values[20],
		SoftSleeper: values[22],
		HardSleeper: values[27],
		HardSeat: values[28],
		StandUp: values[25],
		Other: values[32],
		Remark: values[0],
	}

	fmt.Println(ticket.StationTrainCode, ticket.FromStation, ticket.ToStation,ticket.StartTime, ticket.ArriveTime, ticket.BusinessSeat, ticket.FirstClassSeat, ticket.SecondClassSeat,
	ticket.AdvancedSleeper, ticket.SoftSleeper, ticket.HardSleeper, ticket.HardSeat, ticket.StandUp, ticket.Other, ticket.Remark)
}

func main(){
	// server health check
	// q("https://kyfw.12306.cn/otn/leftTicket/log")
	
	q("https://kyfw.12306.cn/otn/leftTicket/queryZ")
}