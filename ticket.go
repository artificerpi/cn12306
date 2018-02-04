package main

type TicketData []string

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
