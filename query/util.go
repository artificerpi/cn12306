package query

import (
	"bytes"
	"net/url"
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
