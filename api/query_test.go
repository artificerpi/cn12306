package api

import "testing"

func TestTicketQuery_Encode(t *testing.T) {
	type fields struct {
		LeftTicketDTO LeftTicketDTO
		PassengerCode string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"good_query", fields{LeftTicketDTO{"2019-01-29", "IZQ", "WHN"}, "ADULT"},
			"leftTicketDTO.train_date=2019-01-29&leftTicketDTO.from_station=IZQ&leftTicketDTO.to_station=WHN&purpose_codes=ADULT"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &TicketQuery{
				LeftTicketDTO: tt.fields.LeftTicketDTO,
				PassengerCode: tt.fields.PassengerCode,
			}
			if got := q.Encode(); got != tt.want {
				t.Errorf("TicketQuery.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
