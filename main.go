package main

import (
	"fmt"
	"time"
)

const (
	queryURL string = "https://kyfw.12306.cn/otn/leftTicket/queryZ"
)

func runQuery() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			results := q("https://kyfw.12306.cn/otn/leftTicket/queryZ")
			data := parseResult(results)
			if len(data) > 3 {
				data = data[:3]
			}

			fmt.Println(data)
		}
	}
}

func main() {
	// server health check
	// q("https://kyfw.12306.cn/otn/leftTicket/log")

	// startUI()
	runQuery()
}
