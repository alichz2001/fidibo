package model

import (
	"encoding/json"
)

type Book struct {
	ID    uint64 `gorm:"primaryKey"`
	Title string
	Cover string
	Type  uint8 `gorm:"type:tinyint;unsigned"`
}

func (i *Book) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Book) UnmarshalBinary(data []byte) error {

	return nil
}
