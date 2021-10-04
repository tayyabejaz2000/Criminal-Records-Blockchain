package main

import (
	blockchain "crime_records/Blockchain"
	criminalrecords "crime_records/CriminalRecords"
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	var bchain = blockchain.CreateBlockChain()

	var records = criminalrecords.CreateRecordSet()
	records.AddRecords(
		criminalrecords.CreateRecord("name1", "12345-1234567-5", "Street1", "302", time.Now()),
	)
	bchain.AddBlock(records)

	records = criminalrecords.CreateRecordSet()
	records.AddRecords(
		criminalrecords.CreateRecord("name2", "54321-7654321-0", "Street2", "402", time.Now()),
		criminalrecords.CreateRecord("name3", "52345-1234567-5", "Street3", "602", time.Now()),
		criminalrecords.CreateRecord("name4", "14321-7654321-0", "Street4", "502", time.Now()),
	)
	bchain.AddBlock(records)

	fmt.Printf("Vaidation before tampering: %v\n", bchain.Validate())

	//Tampering
	bchain.LastBlock().PrevBlock.Data.(*criminalrecords.RecordsSet).Records[0].Name = "1"
	bchain.LastBlock().PrevBlock.DataHash = bchain.LastBlock().Data.Hash()

	fmt.Printf("Vaidation after tampering: %v\n", bchain.Validate())

	var json, _ = json.Marshal(bchain.LastBlock())
	fmt.Printf("Last Block: %v\n", string(json))
}
