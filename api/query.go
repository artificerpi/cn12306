package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	queryURL = "https://kyfw.12306.cn/otn/leftTicket/queryZ"
)

var (
	client *http.Client
	// cache  Cache
)

func init() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{
		Transport: tr,
		// Timeout:   10 * time.Second,
	}
}

type TicketResponse struct {
	Data       Payload     `json:"data"`
	HttpStatus int         `json:"httpstatus"`
	Messages   interface{} `json:"messages"`
	StatusOK   bool        `json:"status"`
}

type Payload struct {
	Flag        string      `json:"flag"`
	Map         interface{} `json:"map"`
	LeftTickets []string    `json:"result"`
}

// TicketQuery stores the query parameters
type TicketQuery struct {
	LeftTicketDTO LeftTicketDTO `json:"leftTicketDTO"`
	PassengerCode string        `json:"purpose_codes"`
}

// LeftTicketDTO is the dto obtained from 12306 query parameters
type LeftTicketDTO struct {
	TrainDate   string `json:"train_date"`
	FromStation string `json:"from_station"`
	ToStation   string `json:"to_station"`
}

// Encode returns the query string of TicketQuery
func (q *TicketQuery) Encode() string {
	values := q.Values()

	return unsortedEncode(values)
}

func unsortedEncode(values url.Values) string {
	if values == nil {
		return ""
	}

	var buf bytes.Buffer
	for k, vs := range values {
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

// Values returns the query parameters in map format
func (q *TicketQuery) Values() (values url.Values) {
	values = make(map[string][]string)
	m := structToMap(*q)
	for k, v := range m {
		values.Add(k, fmt.Sprint(v))
	}

	return
}

// RequestTicketInfo makes http get request to query ticket info
func RequestTicketInfo(q TicketQuery) []TicketData {
	req, err := http.NewRequest("GET", queryURL, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		// client.Timeout += 3 * time.Second
	}
	defer resp.Body.Close()

	return readData(resp.Body)
}

func readData(r io.Reader) []TicketData {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		log.Println(err)
	}

	var t TicketResponse
	err = json.Unmarshal(data, &t)
	if err != nil {
		log.Println(err)
	}

	records := make([]TicketData, 30)
	for _, v := range t.Data.LeftTickets {
		v = strings.TrimPrefix(v, "|")
		values := strings.Split(v, "|")
		// filter
		if strings.HasPrefix(values[2], "G") {
			records = append(records, values)
		}
	}

	log.Printf("Got %d line of records\n", len(records))

	return records
}
