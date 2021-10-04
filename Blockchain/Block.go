package blockchain

import (
	"crypto/sha256"
	"encoding/json"
	"time"
)

type SHAable interface {
	Hash() [32]byte
}

type block_header struct {
	Index     int       `json:"index"`
	PrevHash  [32]byte  `json:"prev_hash"`
	Timestamp time.Time `json:"timestamp"`
}

type block struct {
	Header    block_header `json:"header"`
	DataHash  [32]byte     `json:"data_hash"`
	PrevBlock *block       `json:"-"`
	Data      SHAable      `json:"-"`
}

func CreateGenesisBlock() *block {
	return &block{
		Header: block_header{
			Index:     0,
			PrevHash:  [32]byte{},
			Timestamp: time.Now(),
		},
		PrevBlock: nil,
		DataHash:  [32]byte{},
		Data:      nil,
	}
}

func CreateBlock(index int, prev_hash [32]byte, data SHAable, prev_block *block) *block {
	return &block{
		Header: block_header{
			Index:     index,
			PrevHash:  prev_hash,
			Timestamp: time.Now(),
		},
		PrevBlock: prev_block,
		DataHash:  data.Hash(),
		Data:      data,
	}
}

func (b *block) Hash() ([32]byte, error) {
	//Recalculate Data Hash
	var jsonBlob, err = json.Marshal(b)
	if err != nil {
		return [32]byte{}, err
	}
	return sha256.Sum256(jsonBlob), nil
}

func (curr *block) Validate(prev *block) bool {
	var prevHash, err = prev.Hash()
	if err != nil {
		return false
	}
	if curr.Header.PrevHash != prevHash {
		return false
	} else if curr.Header.Timestamp.Before(prev.Header.Timestamp) {
		return false
	} else if curr.Header.Index != prev.Header.Index+1 {
		return false
	} else if curr.Data.Hash() != curr.DataHash {
		return false
	}
	return true
}
