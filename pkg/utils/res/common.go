package res

import "time"

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
	Errors     interface{}
}

type SampleCardStatus struct {
	ID          uint      `json:"id"`
	CardID      string    `json:"card_id"`
	UserContact string    `json:"user_contact"`
	TimeStamp   time.Time `json:"time_stamp"`
	Comment     string    `json:"comment"`
}
