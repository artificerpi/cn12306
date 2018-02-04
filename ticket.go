package main

type TicketData []string

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
