package main

import (
	"testing"
	"time"
)

func TestQuery_q(t *testing.T) {
	// test with -v to enable verbose output
	// results := q(queryURL)
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			results := q("https://kyfw.12306.cn/otn/leftTicket/queryZ")
			t.Log("%q", parseResult(results)[:3])
		}
	}
}
