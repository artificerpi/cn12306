package main

import (
	"crypto/tls"
	"log"
	"net/http"
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

func q(rawurl string) [][]string {
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

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		// client.Timeout += 3 * time.Second
	}
	defer resp.Body.Close()

	return readData(resp.Body)
}

// // Query store the query string parameters
// type Query struct{
// 	TrainDate string 	`url:"leftTicketDTO.train_date"`
// 	FromStation string	`url:"leftTicketDTO.from_station"`
// 	ToStation string 	`url:"leftTicketDTO.to_station"`
// 	PurposeCode string 	`url:"purpose_codes"`
// }

// func (q *Query) Values()(values url.Values){
// 	values = url.Values{}
// 	iVal := reflect.ValueOf(q).Elem()
// 	typ := iVal.Type()
// 	for i := 0; i < iVal.NumField(); i++ {
// 		f := iVal.Field(i)
// 		// You ca use tags here...
// 		// tag := typ.Field(i).Tag.Get("tagname")
// 		// Convert each type into a string for the url.Values string map
// 		var v string
// 		switch f.Interface().(type) {
// 		case int, int8, int16, int32, int64:
// 			v = strconv.FormatInt(f.Int(), 10)
// 		case uint, uint8, uint16, uint32, uint64:
// 			v = strconv.FormatUint(f.Uint(), 10)
// 		case float32:
// 			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
// 		case float64:
// 			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
// 		case []byte:
// 			v = string(f.Bytes())
// 		case string:
// 			v = f.String()
// 		}
// 		values.Set(typ.Field(i).Name, v)
// 	}
// 	return
// }

// func (q *Query) Encode() string{
// 	v := q.Values()
// 	if v == nil {
// 		return ""
// 	}
// 	var buf bytes.Buffer
// 	keys := make([]string, 0, len(v))
// 	for k := range v {
// 		keys = append(keys, k)
// 	}

// 	for _, k := range keys {
// 		vs := v[k]
// 		prefix := url.QueryEscape(k) + "="
// 		for _, v := range vs {
// 			if buf.Len() > 0 {
// 				buf.WriteByte('&')
// 			}
// 			buf.WriteString(prefix)
// 			buf.WriteString(url.QueryEscape(v))
// 		}
// 	}
// 	return buf.String()
// }
