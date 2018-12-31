package api

import "time"

type TicketData []string

type TicketResult struct {
	TicketList []TicketData
	QueryTime  time.Time
}

func (d *TicketResult) RawData() [][]string {
	list := make([][]string, len(d.TicketList))
	for _, e := range d.TicketList {
		list = append(list, e)
	}

	return list
}

type TicketInfo struct {
	StationTrainCode string
	FromStation      string
	ToStation        string
	StartTime        string
	ArriveTime       string
	BusinessSeat     string
	FirstClassSeat   string
	SecondClassSeat  string
	AdvancedSleeper  string
	SoftSleeper      string
	HardSleeper      string
	HardSeat         string
	StandUp          string
	Other            string
	Remark           string
}

func (t *TicketInfo) Format() []string {
	return nil
}

func parseTicket(data TicketData) *TicketInfo {
	return nil
}
