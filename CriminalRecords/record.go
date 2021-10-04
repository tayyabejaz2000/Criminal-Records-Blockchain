package criminalrecords

import (
	"time"
)

type Record struct {
	Name    string    `json:"name"`
	CNIC    string    `json:"cnic"`
	Address string    `json:"address"`
	Crime   string    `json:"crime"`
	Time    time.Time `json:"time"`
}

func CreateRecord(name string, cnic string, address string, crime string, timestamp time.Time) *Record {
	return &Record{
		Name:    name,
		CNIC:    cnic,
		Address: address,
		Crime:   crime,
		Time:    timestamp,
	}
}
