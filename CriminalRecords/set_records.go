package criminalrecords

import (
	"crypto/sha256"
	"encoding/json"
)

type RecordsSet struct {
	Records []Record `json:"records"`
}

func CreateRecordSet() *RecordsSet {
	return &RecordsSet{
		Records: []Record{},
	}
}

func (rs *RecordsSet) AddRecord(r *Record) {
	rs.Records = append(rs.Records, *r)
}

func (rs *RecordsSet) AddRecords(r ...*Record) {
	for _, rec := range r {
		rs.AddRecord(rec)
	}
}

func (rs *RecordsSet) Reset() {
	rs.Records = []Record{}
}

func (rs *RecordsSet) Hash() [32]byte {
	var jsonBlob, err = json.Marshal(rs)
	if err != nil {
		return [32]byte{}
	}
	return sha256.Sum256(jsonBlob)
}
