package main

type TicketResponse struct{
	Data Payload `json:"data"`
	HttpStatus int  `json:"httpstatus"`
	Messages interface{} `json:"messages"`
	StatusOK	bool `json:"status"`
}

type Payload struct{
	Flag string  `json:"flag"`
	Map interface{} `json:"map"`
	LeftTickets []string `json:"result"`
}

type LeftTicket struct{
	TrainNo string
	FromStation string
	ToStation string
	StartTime string
	ArriveTime string
	BusinessSeat int
	FirstClassSeat int
	SecondClassSeat int
	AdvancedSleeper int
	SoftSleeper int
	FastSleeper int
	HardSleeper int
	SoftSeat int
	HardSeat int
	StandUp int
	Other int
	Remark string
}