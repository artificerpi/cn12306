package main

import (
	"log"
	"os"
)

const (
	queryURL string = "https://kyfw.12306.cn/otn/leftTicket/queryZ"
)

func init() {
	f, err := os.OpenFile("cn12306.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// defer f.Close()

	log.SetOutput(f)
}

func main() {
	// server health check
	// q("https://kyfw.12306.cn/otn/leftTicket/log")
	log.Println("Started")
	go pollTickets()
	startUI()
}
