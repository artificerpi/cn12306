package main

import (
	"fmt"
	"time"

	"github.com/artificerpi/cn12306/query"
)

const (
	queryURL string = "https://kyfw.12306.cn/otn/leftTicket/queryZ"
)

func runQuery() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			results := query.Q("https://kyfw.12306.cn/otn/leftTicket/queryZ")
			// data := parseResult(results)
			// if len(data) > 3 {
			// 	data = data[:3]
			// }

			// fmt.Println(data)
			fmt.Println(len(results))
		}
	}
}

func main() {
	runQuery()
}
